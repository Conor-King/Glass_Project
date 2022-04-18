### Feel free to suggest any ideas, improvements and changes to the approach!


## How to

### Work Enviroment
- my go version - 1.17.8
- ```curl -o minifab -sL https://tinyurl.com/yxa2q6yr && chmod +x minifab ``` - download latest minifab
- ```PATH="$(pwd):$PATH"``` - add minifab path to your path tempoarily (execute each time you develop)

### Steps
1. ```minifab up``` - bring up the network (first time it takes around 3-4 mins)
2. copy the "strings" chaincode from ```/chaincodes``` folder to ```/vars/chaincode```
3. ```minifab install -v 1.01 -n strings``` - install chaincode named stings, specify higher version each time you want to update the chaincode
4. ```minifab approve```
5. ```minifab commit```
6. ```minifab initialize -p '"init","a","uniqueid1:cid1 uniqueid2:cid2","b","uniqueid3:cid3 uniqueid4:cid4"'``` - create 2 entities and associate them with two examlpe uniqueids:cids
7. ```minifab discover```
8. ```minifab invoke -p '"query", "a"'``` //expected output payload: uniqueid1:cid1 uniqueid2:cid2
9. ```minifab invoke -p '"invoke","a","b","uniqueid2"'``` - transfer asset with uniqueid2 from a to b
10. ```minifab invoke -p '"invoke","add","a","exampleid:examplecid"'``` - add asset to "a" entity with exampleid:examplecid

## Exlanation


The Steps above are to get familiar with the workflow with minifab and the chaincode. The values are for example puprose, as the conncetion between HLF and IPFS is still under development. 

After many different approaches, the only way to add asset  pair to an entity was by using the invoke method in the chaincode provided, with some flow control,  as if you try to make a new method, for eaxmple "add" medthod the transaction is invalid, as the chaincode which this chaincode is based on accepts only invoke as a valid way to modify the ledger. Not the most elegant way to implement an add functionality, but it works. 

Delete functionality to be added soon, possibly by similar way that add is implemented. See line 89 in chaincodes/strings/go/main.go

After this the only thing remianing is to populate the entities with actual unique ids and cids using the ipfs code. The method do do this is with a bash script which being developed. Thiss will make the populating simpler and more user-friendly. The script already have a method which gives the user the option to input and upload an asset to the ipfs. The generation of unique ids is with "github.com/google/uuid" see line 50 in ipfs/main.go. 

To develop simply edit the main.go file in chaicodes/sample/go/main.go and copy it to /vars/chaincode/sample. You can edit directly in vars/chaincode but there is a chance to lose your code, for ex - if u execute ```minifab cleanup```. 
```cp chaincodes/strings/go/main.go vars/chaincode/strings/go/``` - copy and replace the main.go file 

## To Do:
- Query the ledger on uuid
- Delete an asset pair from an entity
- Populate an entity with real uuid:cid pairs
- Add an entity
- Delete an entity

This should be completed untill Tuesday.



## Relevant Links
https://github.com/hyperledger-labs/minifabric/blob/main/docs/README.md - Minifab documentation

https://github.com/google/uuid - Unique Id generation
