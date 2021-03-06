# Reminder-gRPC
Reminder service takes only one argument : timestamp. ‘When’ value here specify the desired date/time of when a user should get reminder notification.

## prerequisite
##### install protocol buffer compiler from [here](https://developers.google.com/protocol-buffers/) or use [snap](https://snapcraft.io/protobuf) for linux
##### check if protoc is successfully installed 
    protoc --version

#### gRPC package for golang 
    go get -u google.golang.org/grpc
#### protobuf packages for golang 
    go get -u github.com/golang/protobuf/protoc-gen-go
#### install protoc grpc gateway for rest api
    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
## Generate Proto files : 
    protoc --proto_path=proto --go_out=plugins=grpc:proto reminder.proto
## Generate Proto with rest api
    protoc \
    --proto_path=proto \
    --proto_path $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway*/third_party/googleapis/ \
    --go_out=plugins=grpc:proto \
    --grpc-gateway_out=logtostderr=true:proto \
    reminder.proto
    
## running server and client and testing the reminder:
open a first terminal and run : 
 
    go run cmd/main.go
open a second terminal and run :
    
    go run cmd/client/client.go