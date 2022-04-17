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

### Exlanation
The Steps above are to get familiar with the workflow with minifab and the chaincode. The values are for example puprose, as the conncetion between HLF and IPFS is still in development.

To develop simply edit the main.go file in chaicodes/sample/go/main.go and copy it to /vars/chaincode/sample. You can edit directly in vars/chaincode but there is a chance to lose your code, for ex - if u execute ```minifab cleanup```. 


### Relevant Links
https://github.com/hyperledger-labs/minifabric/blob/main/docs/README.md - Minifab documentation
