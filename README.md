Group project: Glass research project, a blockchain network that shares EU citizen details with other member states.
# What is GLASS
EU research project which aims to establish a common infrastructure at the European level to provide shared public storage for documentation, files and data, and hosting services to cross-sector organizations.

GLASS aims to design and deploy a blockchain-based distributed environment and deliver the operational framework for sharing common services including

-the definition of the resources requirements
-the incorporation of the Interplanetary File System (IPFS) in the GLASS architecture
-the design, development and deployment of the distributed ledger
-seamless identity management

# What is IPFS
A peer-to-peer hypermedia protocol designed to preserve and grow humanity's knowledge by making the web upgradeable, resilient, and more open. IPFS is a file sharing system that can be leveraged to more efficiently store and share large files. It relies on cryptographic hashes that can easily be stored on a blockchain Learn more here: https://ipfs.io/#how


# Distributed Ledger
At the heart of a blockchain network is a distributed ledger that records all the transactions that take place on the network. the information recorded to a blockchain is append-only, using cryptographic techniques that guarantee that once a transaction has been added to the ledger it cannot be modified. It’s why blockchains are sometimes described as systems of proof

# Smart Contracts
To support the consistent update of information — and to enable a whole host of ledger functions (transacting, querying, etc) — a blockchain network uses smart contracts  to provide controlled access to the ledger. smart contracts  to interact with the channel ledger. Smart contracts contain the business logic that governs assets on the blockchain ledger. Applications run by members of the network can invoke smart contracts to create assets on the ledger, as well as change and transfer those assets. Applications also query smart contracts to read data on the ledger. Hyperledger Fabric users often use the terms smart contract and chaincode interchangeably.  Chaincode can be implemented in several programming languages. Currently, Go, Node.js, and Java chaincode are supported. . In general, a smart contract defines the transaction logic that controls the lifecycle of a business object contained in the world state. It is then packaged into a chaincode which is then deployed to a blockchain network. Think of smart contracts as governing transactions, whereas chaincode governs how smart contracts are packaged for deployment.

# Consensus
The process of keeping the ledger transactions synchronized across the network — to ensure that ledgers update only when transactions are approved by the appropriate participants, and that when ledgers do update, they update with the same transactions in the same order — is called consensus Transactions must be written to the ledger in the order in which they occur. For this to happen, the order of transactions must be established and a method for rejecting bad transactions that have been inserted into the ledger in error (or maliciously) must be put into place. Hyperledger Fabric has been designed to allow network starters to choose a consensus mechanism that best represents the relationships that exist between participants.

# Hyperledger Fabric
Like other blockchain technologies, it has a ledger, uses smart contracts, and is a system by which participants manage their transactions  . Where Hyperledger Fabric breaks from some other blockchain systems is that it is private and permissioned. 
Design Features:

Assets
exchange tangible and intangible assets

Assets are represented as key-value pairs and state changes recorded on a channel ledger.

Chaincode
Software defining the assets and instructions to modify them. Rules for reading and altering the key-value pairs. Execution - write a set of key-value that is submitted and validated by all peers

Ledger
Sequenced, temper-resistant, immutable

Store the chain of sequenced records in blocks + the current state. One ledger per channel. Each peer have a copy of the ledger for the channels which he participates. Read about features of the ledger here:

Hyperledger Fabric Model - hyperledger-fabricdocs master documentation

Privacy
Ledger per channel and chaincode which modifies the state of assets. Scope - the channel. It can be shared across the network or be private including only a set of members.

In later stages of development:

participants can create separate channels and isolate their transactions. Solving the transparency privacy issue. Chaincode (the ability to read and write) can be installed only on selected peers. Other ones would not have access to the transaction history and etc. This would keep the transaction data confidential, accessible only by

authorized organizations. Data kept away from the broader network but shared across subsets of organizations on the channel.

Additionally, the values can be AES encrypted and can be only decrypted by the peer with a key.

Security & Membership Services
All members have known identities. Public keys to generate certificates tied to organizations and users.

Consensus
Stages: Proposal, endorsement, ordering, validation, and commitment. FInal check protection against double-spend operations plus Identity verification.

# What is Self-sovereign identity SSI 🆔

Self-sovereign identity (SSI) is a term used to describe the digital movement that recognizes an individual should own and control their identity without the intervening administrative authorities. SSI allows people to interact in the digital world with the same freedom and capacity for trust as they do in the offline world.. With SSI, the power to control personal data resides with the individual, and not an administrative third party granting or tracking access to these credentials. The SSI identity system gives you the ability to use your digital wallet and authenticate your own identity using the credentials you have been issued. You no longer have to give up control of personal information to dozens of databases each time you want to access new goods and services, with the risk of your identity being stolen by hackers. This is called “self-sovereign” identity because each person is now in control of their own identity—they are their own sovereign nation. People can control their own information and relationships. A person’s digital existence is now independent of any organization: no-one can take their identity away.

# What is Distributed Ledger Technologies (DLT)

Digital database which every member can supplement the data stored there. The data is stored locally on each machine not on a centralized cloud ( decentral peer-to-peer network). Mining verifies the data and makes DLT transparent, safe and decentral. Every blockchain is a form of DLT but not every DLT is blockchain. Nodes of the DLT are located on different locations.

## Characteristics of DLT
-Immutable
-Transparent
-Anonymous
-Single Source of truth


## System architecture

![visual representation of our architecture](https://github.com/Conor-King/Glass_Project/blob/main/Teams_Files/system_architecture.png)

The above diagram is a visual representation of our architecture. 

### Limitation of the architecture:
The system is not truly a distributed system because if our main IPFS or our mani HLF goes down the whole system is down. Moreover, the system has a careerist in subordinate network design. This can be improved in a further version. A possible improvement to create the system fully decentralized could be based on current DEFI technology. Secondly, an alternative solution could be using a redundant connection from our application to the network. 

## Acknowledgement
We would like to express our sincere gratitude to several individuals and organizations for supporting us throughout our project. First, we wish to express our sincere gratitude to our supervisor, Professor Collins, for his enthusiasm, patience, insightful comments, helpful information, and practical advice. Moreover, we would like to express our gratitude to Dr. Sarwar Sayeed and to Dr. Owen Lo. Last but not least, we would like to thank you for Dr Naghmeh Moradpoor Sheykhkanloo for sponsorship during the project.

