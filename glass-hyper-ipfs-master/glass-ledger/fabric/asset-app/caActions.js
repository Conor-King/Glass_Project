// It performs the administrative actions; 1. creating an admin user, & 2. a normal user

const path = require("path");
const helper = require("./helper");

const FabricCAServices = require("fabric-ca-client");
const { Wallets } = require("fabric-network");

const adminUserId = "admin";
const adminUserPasswd = "adminpw";

const walletPath = path.join(__dirname, "wallet");

// creates a ca-client that communicates with the certification authority.
// The certification authority will return a certificate which will be saved on the wallet.
function buildCAClient(FabricCAServices, ccp, caHostName) {
  const caInfo = ccp.certificateAuthorities[caHostName];
  const caTLSCACerts = caInfo.tlsCACerts.pem;
  const caClient = new FabricCAServices(
    caInfo.url,
    {
      trustedRoots: caTLSCACerts,
      verify: false,
    },
    caInfo.caName
  );

  console.log("Built a CA client named: ${caInfo.caName}");
  return caClient;
}

async function enrollAdmin(caClient, wallet, orgMspId) {
  try {
    const identity = await wallet.get(adminUserId);
    if (identity) {
      console.log("An identity for admin already exists");
      return;
    }

    const enrollment = await caClient.enroll({
      enrollmentID: adminUserId,
      enrollmentSecret: adminUserPasswd,
    });

    const x509Identity = {
      credentials: {
        certificate: enrollment.certificate,
        privateKey: enrollment.key.toBytes(),
      },
      mspId: orgMspId,
      type: "X.509",
    };

    await wallet.put(adminUserId, x509Identity);
  } catch (error) {
    console.log("Failed to enroll admin: ${error}");
  }
}

async function registerAndEnrollUser(
  caClient,
  wallet,
  orgMspId,
  userId,
  affiliation
) {
  try {
    // Check to see if we've already enrolled the user
    const userIdentity = await wallet.get(userId);
    if (userIdentity) {
      console.log(
        `An identity for the user ${userId} already exists in the wallet`
      );
      return;
    }

    // Must use an admin to register a new user
    const adminIdentity = await wallet.get(adminUserId);
    if (!adminIdentity) {
      console.log(
        "An identity for the admin user does not exist in the wallet"
      );
      console.log("Enroll the admin user before retrying");
      return;
    }

    // build a user object for authenticating with the CA
    const provider = wallet
      .getProviderRegistry()
      .getProvider(adminIdentity.type);
    const adminUser = await provider.getUserContext(adminIdentity, adminUserId);

    // Register the user, enroll the user, and import the new identity into the wallet.
    // if affiliation is specified by client, the affiliation value must be configured in CA
    const secret = await caClient.register(
      {
        enrollmentID: userId,
        role: "client",
      },
      adminUser
    );
    const enrollment = await caClient.enroll({
      enrollmentID: userId,
      enrollmentSecret: secret,
    });
    const x509Identity = {
      credentials: {
        certificate: enrollment.certificate,
        privateKey: enrollment.key.toBytes(),
      },
      mspId: orgMspId,
      type: "X.509",
    };
    await wallet.put(userId, x509Identity);
    console.log(
      `Successfully registered and enrolled user ${userId} and imported it into the wallet`
    );
  } catch (error) {
    console.error(`Failed to register user : ${error}`);
  }
}

async function getAdmin() {
  let ccp = helper.buildCCPOrg1();

  // get fabric services instance based on network configuration
  const caClient = buildCAClient(FabricCAServices, ccp, "ca.org1.example.com");

  // save creds in the wallet
  const wallet = await helper.buildWallet(Wallets, walletPath);

  // should be done only once
  await enrollAdmin(caClient, wallet, "Org1Msp");
}

async function getUser(org1UserId) {
  let ccp = helper.buildCCPOrg1();

  const caClient = buildCAClient(FabricCAServices, ccp, "ca.org1.example.com");

  // save creds in the wallet
  const wallet = await helper.buildWallet(Wallets, walletPath);

  await registerAndEnrollUser(
    caClient,
    wallet,
    "Org1MSP",
    org1UserId,
    "org1.department"
  );
}

let args = process.argv;

if (args[2] === "admin") {
  getAdmin();
} else if (args[2] === "user") {
  let org1UserId = args[3];
  getUser("org1UserId");
} else {
  console.log("Invalid command");
}
