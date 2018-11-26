#!/usr/bin/env bash

docker run -d --name zookeeper1 --add-host zookeeper0.example.com:39.104.145.229 --add-host zookeeper1.exampe.com:39.104.189.169 -e "ZOO_SERVERS=server.1=zookeeper0.example.com:2888:3888 server.2=zookeeper1.example.com:2888:3888" -p 2181:2181 -p 2888:2888 -p 3888:3888 hyperledger/fabric-zookeeper
