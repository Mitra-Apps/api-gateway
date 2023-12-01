# API-GATEWAY

## Installation
Install go version 1.21.3
Install protobuf : https://github.com/protocolbuffers/protobuf/releases/tag/v25.1
Execute command line in terminal :
go get -u google.golang.org/protobuf
go get -u google.golang.org/grpc       
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go mod tidy
go mod vendor

## How to run
run the apps using command : go run main.go

## Generate pb file from proto file
This is example to create protobuf of user entity and service :
protoc --go_out=. domain/user/proto/*.proto
protoc --go-grpc_out=require_unimplemented_servers=false:. ./domain/user/proto/*.proto