## How to
- my go version - 1.17.8
- ```curl -o minifab -sL https://tinyurl.com/yxa2q6yr && chmod +x minifab ```
- ```PATH="$(pwd):$PATH" ``` - to add minifab path to your path tempoarily

1. ```minifab up``` - bring up the network (first time it takes around 3-4 mins)
2. copy the "strings" chaincode from minifabric/chaincodes folder to minifabric/vars/chaincode
3. minifab install -v 1.01 -n strings - install chaincode named stings, specify higher version each time you want to update the chaincode
4. minifab approve
5. minifab commit
6. minifab initialize -p '"init","a","uniqueid1:cid1 uniqueid2:cid2","b","uniqueid3:cid3 uniqueid4:cid4"' - crate 2 entities and associate them with two examlpe uniqueids:cids
7. minifab discover
8. minifab invoke -p '"query", "a"' //expected output payload: uniqueid1:cid1 uniqueid2:cid2
9. minifab invoke -p '"invoke","a","b","uniqueid2"' - transfer asset with uniqueid2 from a to b
