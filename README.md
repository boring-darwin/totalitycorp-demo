# GRPC DEMO
Basic implentation of grpc protocal in golang


## Server
Server directory have all the server side implementation.

Run server cmd: `-> go run server.go`


## Client
Client directory have code related to get user by id and get list od user by publishing the list of id.

Run client cmd: `-> go run client.go`

## Proto
Proto directory contains the protobuff file, in which the message contract has been defined.
To create golang implementation out of the .proto file use the below cmd.

-> `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/user.proto`