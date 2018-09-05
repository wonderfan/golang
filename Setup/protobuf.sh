
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip

protoc  --go_out=plugins=grpc:$GOPATH/src *.proto 

protoc  --go_out=. *.proto
