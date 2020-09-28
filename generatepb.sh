#!/bin/bash

protoc --proto_path=rpc --go_out=rpc --go-grpc_out=rpc --go_opt=paths=source_relative rpc/calculator.proto