#!/bin/bash
cd $(cd `dirname $0`;pwd)

protoc -I ../api/proto ../api/proto/*.proto -I. \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--go_out=plugins=grpc:../api/grpcapi
