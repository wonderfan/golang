./peer chaincode deploy -p github.com/hyperledger/fabric/examples/chaincode/go/eventsender -c '{"Function":"init", "Args":[]}'

CORE_PEER_ADDRESS=172.17.0.2:7051 peer chaincode deploy -p github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02 -c '{"Function":"init", "Args": ["a","100", "b", "200"]}'

CORE_PEER_ADDRESS=172.17.0.2:7051 peer chaincode invoke -n 1edd7021ab71b766f4928a9ef91182c018dffb86fef7a4b5a5516ac590a87957e21a62d939df817f5105f524abddcddfc7b1a60d780f02d8235bd7af9db81b66 -c '{"Function":"invoke", "Args": ["a","b","10"]}'
