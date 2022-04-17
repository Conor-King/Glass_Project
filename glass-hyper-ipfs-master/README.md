# Glass-Hyper-IPFS
Proof-of-concept chain code which aims to demonstrate how one may integrate IPFS with Hyperledger Fabric. This code has been based on [fish-farm](https://github.com/blockpass-identity-lab/fish-farm) repo. This README focuses specifically on the Glass-IPFS chaincode. The [fish-farm README](https://github.com/blockpass-identity-lab/fish-farm#readme) provides further technical details and some troubleshooting tips which may also be applicable to this repo since they share the same workflow.

This branch of code (glass-portal) specifically demonstrates the integration of Hyperledger Fabric with IPFS using a user front-end we name the 'glass-portal'. The workflow implemented demonstrates how we may encrypt and store GLASS resources (e.g. passports, drivers license, medical records and other citizen centric resources) on IPFS in a secure manner. All resources stored on IPFS are first encrypted by the glass-portal. A triplet, in the form of `CID`, `Encryption Key` and `URI` are then recorded on Hyperledger Fabric. Only permissioned individuals may read the Hyperledger Fabric record, obtain the encryption key and use it to decrypt a given resource stored on IPFS. 

# Table of Contents
1. [Components Overview](#overview)
2. [Workflow Description](#workflow)
3. [Installation and Usage](#installation)

# 1. Components Overview<a name="overview"></a>

There are three main components in the proof-of-concept implemented: Hyperledger Fabric, IPFS and Glass Portal as shown in the figure below:

![Components](/figures/fig1.png)

## 1.1. Hyperledger Fabric
Hyperledger Fabric is used as our permissioned blockchain for the storage of the _triplets_ (`CID`, `Key` and `URI`) for each resource encrypted and distributed on IPFS. The triplets are stored in a data collection while the ledger itself is used to record read/write transactions (for audit and record keeping purposes). See Section 1.1.1 below for more details on the design concept of this triplet.

In Hyperledger Fabric, we implement a [glassipfs chaincode](https://github.com/blockpass-identity-lab/glass-hyper-ipfs/blob/glass-portal/glass-ipfs/chaincode_glassipfs/go/main.go) to allow users to `createGlassResource()` and `readGlassResourceKey()`. The former function allows a user to insert a new triplet while the latter function allows a user to read an existing triplet to obtain encryption key of a resource. 

Given the sensitive nature of encryption keys, and to demonstrate access policies capabilities in Hyperledger Fabric, two generic organisations have been configured in this workflow: `org1.org` and `org2.org`. By default, `org1.org` has full permission to create new triplets and read existing triplet values. On the other hand, `org2.org` only has permission to create new triplets and cannot read any encryption key. The access policy is defined in [glassipfs_collection_config.json](https://github.com/blockpass-identity-lab/glass-hyper-ipfs/edit/glass-portal/glass-ipfs/glassipfs_collection_config.json) as shown below:

```
[
   {
      "name": "collectionGlassResources",
      "policy": "OR( 'org1-org.member', 'org2-org.member' )",
      "requiredPeerCount": 0,
      "maxPeerCount": 3,
      "blockToLive":30,
      "memberOnlyRead": true
   },
   {
      "name": "collectionGlassResourcesKeys",
      "policy": "OR( 'org1-org.member')",
      "requiredPeerCount": 0,
      "maxPeerCount": 3,
      "blockToLive":30,
      "memberOnlyRead": true
   }
  ]
```

As shown in the above policy, the Hyperledger Fabric environment is configured to use two data collections: `collectionGlassResources` and `collectionGlassResourcesKeys`. The `collectionGlassResources` is a public data collection (readable by both `org1.org` and `org2.org`) while `collectionGlassResourcesKeys` is private data collection (readable by org1.org only). The former collection stores the CID and URI of our Glass resources while latter collection stores the encryption key. These two data collections, in combintation, form our _triplet_ concept. The fields for each resource and read permissions for each of the two orgs are shown in the tables below:

| collectionGlassResources (public data) Fields| collectionGlassResourcesKeys (private data) Fields|
| ------------- | ------------- |
| CID  | Secret Key |
| URI  |  |

| org1.org | org2.org |
| ------------- | ------------- |
| CID| CID|
| URI| URI|
| Encryption Key| |

### 1.1.1. Triplets Design Overview
The triplets stored on Hyperledger Fabric data collections form the core metadata which allows us to decrypt data distributed on IPFS (or other distribution mechanisms in future work). The triplets we stored consist of (`CID`, `Encryption Key` and `URI`):

- **CID** : Content Identifier, using the content addressing algorithm as defined by [IPFS](https://docs.ipfs.io/concepts/content-addressing). The CID, in short, is a uses the SHA-256 algorithm to hash the contents of a file and allows us to generate a unique ID to **identify** the file.
- **URI** : Unique Resource Identifier allows to **locate** the resource. In other words, it defines the actual location whereby the resource is stored.
- **Key** : The encryption key which can be used to decrypt the resource (which is associated with the CID and hosted in a location as defined by the URI). We assume, and require, all resources to be encrypted before they are distributed or stored given the sensitive nature of the data we deal with in GLASS.

Since the focus of this work is integration of Hyperledger Fabric with IPFS, it may have been noted that the CID and URI will store the same content. Since the IPFS protocol uses the CID to both **identify** and **locate** resources, this is as expected for IPFS. However, by using the property of URI as a seperate field, we may choose to distribute our content in other mechanisms in future work such as Dropbox and Sharepoint (as examples). In such scenarios, the CID of a resource will remain the same but the URI will differ depending on where the resource is located. Naturally, the encryption key will also remain the same. See table below for illustrative example of how the URIs may differ depending on where the resource is hosted:

| Distribution Mechanism | IPFS                                                 | Dropbox                                        | Sharepoint                                                        |
|------------------------|------------------------------------------------------|------------------------------------------------|-------------------------------------------------------------------|
| CID                    | QmVy3VhfsENEvLsetgren2Y2jrnj227GZFuVQHqyKyBQYC       | QmVy3VhfsENEvLsetgren2Y2jrnj227GZFuVQHqyKyBQYC | QmVy3VhfsENEvLsetgren2Y2jrnj227GZFuVQHqyKyBQYC                    |
| URI                    | /ipfs/QmVy3VhfsENEvLsetgren2Y2jrnj227GZFuVQHqyKyBQYC | https://www.dropbox.com/sh/5nhrjsksha          | https://glassproject.sharepoint.com/sites/share/fa112fba-c3de5f8a |


## 1.2. IPFS
The interplanetary file system is a peer-to-peer content sharing protocol widely used on the internet. This protocol is used as our primarily resource distribution mechanism in the developed prototype. All resources distributed on the IPFS network, within our scenario, are encrypted. However, for greater security, a private instance of IPFS was used as the testing ground in the scope of this codebase. A private IPFS functions the same as a the public instance of IPFS. However, only nodes which posses a shared private key (referred to as the swarm key) can participate in the private IPFS network. The use of a private IPFS network helps prevent accidental disclosure of confidential or sensitive information.


## 1.3. Glass Portal
The glass portal is built on Python Flask (backend) and HTML (frontend). It acts as the interface between Hyperledger Fabric and Private IPFS. The glass portal serves as the main component users interact with when uploading GLASS resources (to IPFS) or retrieving GLASS resource metadata (i.e. encryption keys) from Hyperledger Fabric. The user may use the Glass Portal to safely encrypt and distribute resources on the private IPFS network and query the distributed ledger to obtain encryption keys for resources which are already on IPFS (assuming they have the correct permissions). The three main aspects of glass-portal are:

1) Allow a user to upload a GLASS resource to the portal (passport, medical data etc). The portal will automatically encrypt and distribute the GLASS resources onto the (private) IPFS network. AES-256 CBC is used as our encryption function in this proof-of-concept. 
2) In conjunction with distributing the encrypted resource onto IPFS, the glass portal will automatically record the triplet metadata (`CID`, `Encryption Key` and `URI`) generated in Step 1 and store this in the data collections of Hyperledger Fabric. Naturally, the permissioned blockchain will record this event in the ledger for audit purposes.
3) Read and retrieve encryption keys from Hyperledger Fabric data collection and perform decryption of GLASS resources, assuming the user possesses the correct permissions.

The screenshot below shows the front-end and the three sections which allow the above-described aspects to take place: 

![Frontend](/figures/fig2.png)

Emphasis should be made that the main contribution to this work is the backend code. Improvements to UI will be addressed in future work.

# 2. Workflow Description<a name="workflow"></a>

## 2.1. Upload, Encrypt, Distribute and Record Resource

A step-by-step visualisation on how the components allow a user to 1) upload, encrypt and distribute a resource on IPFS; and record the metadata (i.e. the _triplet_ `CID`, `Encryption Key` and `URI`) on Hyperledger Fabric is shown in the figure below:

![Frontend](/figures/fig3.png)

1) A user uploads their sensitive resource (e.g. a photocopy of their passport in this example) to the GLASS portal. The portal will first automatically encrypt the resource using AES-256 CBC. The original resource (plaintext) is deleted from the server once encryption is successful. 
2) Steps 2a + 2b occur at the same time. The GLASS portal will distribute the encrypted resource on the (private) IPFS network. In doing so, a CID will be generated for the encrypted resource. This CID also acts as our URI in this example since the IPFS uses the CID to both identify (hash) and locate resources. In conjunction with distributing the encypted resource on IPFS, the GLASS portal also interacts with Hyperledger Fabric to record the triplet of CID, URI and encryption key used to encrypt the resource.
3) The status result is returned to the user. If the actions were successful, the user is provided with a CID and URI of their encrypted resource. This allows them or other permissioned individuals to locate it in the future.

## 2.2. Obtain Encryption Key and Decrypt Resource from IPFS

The below figure shows how a user can query Hyperledger Fabric and obtain the encryption key for a resource. The resource can then be retrieved from IPFS and decrypted by the glass-portal by inputting the correct key and IV:

![Frontend](/figures/fig4.png)

1) A user provides a CID for the GLASS portal to query against Hyperledger Fabric. 
2) If the user has satisfies the correct policy (i.e. they belong to org1.org in this example), permission will be granted to allow the GLASS portal to read the encryption key which is mapped to the given CID.
3) The encryption key is returned to the user. In the above example, the encryption key and IV are simply separated by a colon for ease of interpretation.
4) Finally, using the known CID, the user can retrieve the encrypted resource from IPFS. Since they now have the encryption key, this encrypted resource and key can be provided to the GLASS portal for decryption. If decryption is successful, the user is able to download the decrypted content from the GLASS portal and obtain the original photocopy of their passport.

# 3. Installation & Usage Instructions<a name="installation"></a>
Note that this code was tested on Mac OS/Linux environment. Additional configuration and modification of code may be required to run this on Windows. To run this implementation, first clone this branch of code:

`git clone -b glass-portal https://github.com/blockpass-identity-lab/glass-hyper-ipfs.git`

Instructions for setting up each of the three components follow. 

## 3.1. Setup Hyperledger Fabric
The software tool [minifabric](https://github.com/hyperledger-labs/minifabric) has been used to ease and speed up deployment of Hyperledger Fabric. See minifabric documentation for full details on usage. To quickly deploy and start a preconfigured instance of Hyperledger Fabric (catered for this codebase) follow the instructions:

1. First, make sure to install [docker](https://www.docker.com/) (18.03 or newer) in your development environment
2. Navigate to the root directory of this repo: ```cd glass-hyper-ipfs```
3. Run ```chmod +x buildscenario.sh``` and ```./buildscenario.sh``` respectively to build Hyperledger Fabric environment with glass-ipfs chaincode.

If ```./buildscenario.sh``` finishes running successful, our Hyperledger Fabric environment should be deployed and ready to be used. The commands ```./minifab down``` and ```./minifab cleanup``` can be used to bring down and reset your Hyperledger Fabric environment if necessary.

## 3.2. Setup (Private) IPFS
Instructions for installing and setting up a private IPFS instance follows:

1. Download and install the command line instance of the IPFS protocol here: https://ipfs.io/#install
2. Follow the instructions in this overleaf document to setup a private instance of IPFS: https://www.overleaf.com/read/ksqbjtcgjpys 

**Important Note:** In the current scope of work, the private IPFS instance must be configured on the same environment as where Hyperledger Fabric and GLASS portal is running. Running a private instance of IPFS as a docker container (as detailed in the overleaf notes) is not supported for now.


## 3.3. Setup GLASS Portal
The GLASS portal is implemented in Python. Ensure Python 3.7 (minimum) is installed on your machine then perform the following steps:

1. Navigate to the root directory of this repo: ```cd glass-hyper-ipfs```
2. Install the dependencies required by glass portal: ```pip install requirements.txt```
3. Start an instance of the server by running ```pyton glass-portal.py```
4. The portal can then be accessed via the URL: http://127.0.0.1:5000/ by default.

If setup of Hyperledger Fabric and Private IPFS was successful, the GLASS portal can now be used to upload, encrypt, and distribute resources on (private) IPFS and read encryption keys from Hyperledger Fabric (ensure ```org1.org``` is selected when performing read actions to Hyperledger else a permission error will be raised as per policy).
