### Feel free to suggest any ideas, improvements and changes to the approach!

## How to
### Work Enviroment
- my go version - 1.17.8
- ```curl -o minifab -sL https://tinyurl.com/yxa2q6yr && chmod +x minifab ``` - download latest minifab
- ```PATH="$(pwd):$PATH"``` - add minifab path to your path tempoarily (execute each time you develop)

### Steps
1. ```minifab up -o france.eu.com``` - bring up the network (first time it takes around 3-4 mins)
2. copy the "strings" chaincode from ```/chaincodes``` folder to ```/vars/chaincode``` ```cp -r chaincodes/strings vars/chaincode```
3. Make sure the spec.yaml file is at the root folder for the minifab tool
4. ```minifab install -v 1.01 -n strings``` - install chaincode named stings, specify higher version each time you want to update the chaincode
5. ```minifab approve```
6. ```minifab commit```
7. ```minifab initialize -p '"init","a","uniqueid1:cid1 uniqueid2:cid2","b","uniqueid3:cid3 uniqueid4:cid4"'``` - create 2 entities and associate them with two examlpe uniqueids:cids
8. ```minifab discover```
9. ```minifab invoke -p '"query", "a"'``` //expected output payload: uniqueid1:cid1 uniqueid2:cid2
10. ```minifab invoke -p '"invoke","a","b","uniqueid2"'``` - **transfer** asset with uniqueid2 from a to b
11. ```minifab invoke -p '"invoke","add","a","uniqueid:examplecid"'``` - **add** asset to "a" entity with exampleid:examplecid
12. ```minifab invoke -p '"invoke","delete","a","uniqueid"'``` - **delete** an asset from "a" entity with uniqueid:examplecid

## Explanation

The Steps above are to get familiar with the workflow with minifab and the chaincode. The values are for example puprose, as the conncetion between HLF and IPFS is still under development. 

After many different approaches, the only way to add asset  pair to an entity was by using the invoke method in the chaincode provided, with some flow control,  as if you try to make a new method, for eaxmple "add" medthod the transaction is invalid, as the chaincode which this chaincode is based on accepts only invoke as a valid way to modify the ledger. Not the most elegant way to implement an add functionality, but it works. 

Delete functionality is based on the way that add is implemented. See line 89 in chaincodes/strings/go/main.go

After this the only thing remianing is to populate the entities with actual unique ids and cids using the ipfs code. The method do do this is with a bash script which being developed. Thiss will make the populating simpler and more user-friendly. The script already have a method which gives the user the option to input and upload an asset to the ipfs. The generation of unique ids is with "github.com/google/uuid" see line 50 in ipfs/main.go. 

## Connect Minifab with IPFS
1. Make sure the network is up and running - make sure you define the starting org with -o
2. Open a new terminal and start IPFS ```ipfs daemon```
3. ```minifab install -v 1.02 -n strings``` - install the chaincode, but make sure that you **specify higher version** 
4. ```minifab approve```
5. ```minifab commit```
6. ```./start.sh``` - Start a script which manages IPFS and writes to Minifab. Follow the instructions. Start from step 1.
8. ```minifab discover```

You should now have a ledger with uuid:cid pairs. You can add, delete and trasfer between the entities with the commands explained above.


## Encryption
An encryption function has been implemented in the program. The main goal of the function is to store encrypted data on the IPFS network to protect PPI and meet standards. The encryption function uses the asymmetric encryption method, AES. Moreover, a decryption function is implemented as well. Symmetric key-based encryption might not be the best for the project, however, in a future version, asymmetric key encryption can be implemented with a rolling key.

## System Flow Chart 

![visual representation of our System Flow](https://github.com/Conor-King/Glass_Project/blob/main/Teams_Files/flow_chart.png)
The flow chart represent the how the program works.
## Script
The scrip is written in bash, there is a lot to improve there. Possible feauteres to be added :
- Chaincode install, approve and commit in one command. 
- Add the newly added IPFS asset to to the Ledger
- Write down all the newly added asssets to a file

### Develop suggestion
Simply edit the main.go file in chaicodes/sample/go/main.go and copy it to /vars/chaincode/sample. You can edit directly in vars/chaincode but there is a chance to lose your code, for ex - if u execute ```minifab cleanup```. 

```cp chaincodes/strings/go/main.go vars/chaincode/strings/go/``` - copy and replace the main.go file 

## Relevant Links
https://github.com/hyperledger-labs/minifabric/blob/main/docs/README.md - Minifab documentation

https://github.com/google/uuid - Unique Id generation
