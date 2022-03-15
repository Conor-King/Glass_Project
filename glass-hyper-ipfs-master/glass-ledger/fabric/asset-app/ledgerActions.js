const helper = require("./helper");
const path = require("path");
const { Gateway, Wallets } = require("fabric-network");

const walletPath = path.join(__dirname, "wallet");

const channelName = "channel1";
const chaincodeName = "basic";

const org1UserId = "org1UserId";

async function main() {
  try {
    // build an in memory object with the network configuration (also known as a connection profile)
    const ccp = helper.buildCCPOrg1();

    // setup the wallet to hold the credentials of the application user
    const wallet = await helper.buildWallet(Wallets, walletPath);

    const gateway = new Gateway();

    // setup the gateway instance
    // The user will now be able to create connections to the fabric network and be able to
    // submit transactions and query. All transactions submitted by this gateway will be
    // signed by this user using the credentials stored in the wallet.
    await gateway.connect(ccp, {
      wallet,
      identity: org1UserId,
      discovery: { enabled: true, asLocalhost: true }, // using asLocalhost as this gateway is using a fabric network deployed locally
    });

    // Build a network instance based on the channel where the smart contract is deployed
    const network = await gateway.getNetwork(channelName);

    // Get the contract from the network.
    const contract = network.getContract(chaincodeName);

    // Initialize a set of asset data on the channel using the chaincode 'InitLedger' function.
    // This type of transaction would only be run once by an application the first time it was started after it
    // deployed the first time. Any updates to the chaincode deployed later would likely not need to run
    // an "init" type function.
    console.log(
      "\n--> Submit Transaction: InitLedger, function creates the initial set of assets on the ledger"
    );
    // await contract.submitTransaction('InitLedger');
    console.log("*** Result: committed");

    let args = process.argv;
    if (args[2] === "GetAllAssets") {
      let result = await contract.evaluateTransaction("GetAllAssets");
      console.log(helper.prettyJSONString(result.toString()));

      // evaluateTransaction only sends the transaction to the peer you are connected to
    } else if (args[2] === "ReadAsset") {
      let asset = args[3];
      let result = await contract.evaluateTransaction("ReadAsset", asset);
      console.log(helper.prettyJSONString(result.toString()));
    } else if (args[2] === "CreateAsset") {
      let r = await contract.submitTransaction(
        "CreateAsset",
        "asset101",
        "violet",
        "5",
        "Snnorre 3",
        "1300"
      );
      console.log(" -> Committed: ", r.toString());

      // submitTransaction sends the transaction to the orderer, which makes sense here
      // you can do submitTransaction for ReadAsset too but it will be slow as it will go to different peers
    } else {
      console.error("Bad command: ", args[2]);
    }

    gateway.disconnect();
  } catch (error) {
    console.log(error);
  }
}

main();
