#!/bin/bash
protoc -I ../../../../../github.com/vitalksocialnetwork/ProtoApi/auth/proto/ *.proto --go_out=../../../../../ --go-grpc_out=../../../../../
