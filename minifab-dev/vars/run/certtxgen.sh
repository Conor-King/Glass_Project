#!/bin/bash
cd $FABRIC_CFG_PATH
# cryptogen generate --config crypto-config.yaml --output keyfiles
configtxgen -profile OrdererGenesis -outputBlock genesis.block -channelID systemchannel

configtxgen -printOrg org0-example-com > JoinRequest_org0-example-com.json
configtxgen -printOrg org1-example-com > JoinRequest_org1-example-com.json
