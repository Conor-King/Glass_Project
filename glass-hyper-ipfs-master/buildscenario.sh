#!/bin/bash

# Make the ./vars/chaincode directory and copy over chaincode with name 'glassipfs'
mkdir $(pwd)/vars
mkdir $(pwd)/vars/chaincode
cp -R $(pwd)/glass-ipfs/chaincode_glassipfs $(pwd)/vars/chaincode/glassipfs
cp -R $(pwd)/glass-ipfs/glassipfs_collection_config.json $(pwd)/vars/glassipfs_collection_config.json

minifab="$(pwd)/minifab"

$minifab up -o org1.org -i 2.3.0 -r true -s couchdb -n glassipfs #Startup minifab and install glassipfs chaincode (from /vars/chaincode/ directory)

# Below command should no longer be necessary as approve, commit and initialize occurs automatically when
# the chaincode is installed via '-n glassipfs'
#$minifab approve,commit,initialize -p ''


#Below example is outdated. Run glass-portal.py instead to create entries.

# ========================================================================================================
# Create 3 IPFS_RESOURCE, namely "resource1", "resource2", and "resource3" with some dummy data in the fields:
# ========================================================================================================

#IPFS_RESOURCE=$( echo '{"CID":"QmPEedw2im9ui56RMTSQKWgz8evqgXrXaU1hVeW1HJn48m","URI":"ipfs://QmPEedw2im9ui56RMTSQKWgz8evqgXrXaU1hVeW1HJn48m","key":"password123"}' | base64 | tr -d \\n )
#$minifab invoke -p '"createGlassResource"' -t '{"IPFSResource":"'$IPFS_RESOURCE'"}'

#IPFS_RESOURCE=$( echo '{"CID":"resource2","CID":"QmPXtnomzYtjLWJwxCfG7ujhVh3aw2PdMGjgbSUvjQESV1","key":"anotherSecretPassword"}' | base64 | tr -d \\n )
#$minifab invoke -p '"createGlassResource"' -t '{"IPFSResource":"'$IPFS_RESOURCE'"}'

#IPFS_RESOURCE=$( echo '{"CID":"resource3","CID":"QmZonswuZB5tXPtgqGLRwEK1f4Zddu3U1HFS15Vs7DDTu9","key":"Computer123Password"}' | base64 | tr -d \\n )
#$minifab invoke -p '"createGlassResource"' -t '{"IPFSResource":"'$IPFS_RESOURCE'"}'

#Note that the data field CID is public while the (decryption) key is private.
