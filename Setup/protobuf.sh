
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip

go get -u github.com/golang/protobuf/protoc-gen-go

protoc  --go_out=plugins=grpc:$GOPATH/src *.proto 

protoc  --go_out=. *.proto
