#! /usr/bin/env bash

wget https://golang.google.cn/dl/go1.16.linux-arm64.tar.gz
tar -xvf go1.16.linux-arm64.tar.gz
mv go /usr/local

export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
