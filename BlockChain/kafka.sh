#!/usr/bin/env bash

docker run -d --name zookeeper0 --add-host zookeeper0.example.com:39.104.145.229 \
--add-host zookeeper1.exampe.com:39.104.189.169 --network host -e "ZOO_MY_ID=1" \
-e "ZOO_SERVERS=server.1=zookeeper0.example.com:2888:3888 server.2=zookeeper1.example.com:2888:3888" \
-p 2181:2181 -p 2888:2888 -p 3888:3888 hyperledger/fabric-zookeeper

docker run -d --name zookeeper1 --add-host zookeeper0.example.com:39.104.145.229 \
--add-host zookeeper1.exampe.com:39.104.189.169 --network host -e "ZOO_MY_ID=2" \
-e "ZOO_SERVERS=server.1=zookeeper0.example.com:2888:3888 server.2=zookeeper1.example.com:2888:3888" \
-p 2181:2181 -p 2888:2888 -p 3888:3888 hyperledger/fabric-zookeeper

39.104.189.169  zookeeper1.exampe.com
39.104.145.229  zookeeper0.example.com
