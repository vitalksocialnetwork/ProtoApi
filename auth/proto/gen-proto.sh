#!/bin/bash
protoc -I ../proto/ ../proto/*.proto --go_out=../gen-go --go-grpc_out=../gen-go
