export BASE_DIR=$GOPATH/src/github.com/stratumn/sdk/fabricstore/integration

# Stop docker containers
docker rm -f $(docker ps -aq)
docker rmi $(docker images | grep dev-peer0.org1.example.com-pop-1.0 | awk "{print \$3}")

# Remove working directories
rm -R $BASE_DIR/../chaincode/hyperledger
rm -R $BASE_DIR/../chaincode/stratumn
rm -R $BASE_DIR/../keystore
rm -R $BASE_DIR/../msp

# remove the local state
rm -f ~/.hfc-key-store/*