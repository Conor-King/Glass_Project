/** helper function allows to load the configuration */

const path = require("path");
const fs = require("fs");

//Configuration for Organisation 1
exports.buildCCPOrg1 = function () {
  const ccpPath = path.resolve(
    __dirname,
    "..",
    "ledger",
    "test-network",
    "organizations",
    "peerOrganizations",
    "org1.example.com",
    "connection-org1.json"
  );

  const fileExists = fs.existsSync(ccpPath);
  if (!fileExists) {
    throw new Error("Cannot find: ${ccpPath}");
  }

  const contents = fs.readFileSync(ccpPath, "utf8");

  const ccp = JSON.parse(contents);

  console.log("Loaded network config from: ${ccpPath}");

  return ccp;
};

// Permission enforcement from the administrative side user.
// Creates the cedentials and stores; the identity and the certificates
exports.buildWallet = async function (Wallets, walletPath) {
  let wallet;

  if (walletPath) {
    wallet = await Wallets.newFileSystemWallet(walletPath);
    console.log("Built wallet from ${walletPath}");
  } else {
    wallet = await Wallets.newInMemoryWallet();
    console.log("Build in-memory wallet");
  }

  return wallet;
};

exports.prettyJSONString = function (inputString) {
  return JSON.stringify(JSON.parse(inputString), null, 2);
};
