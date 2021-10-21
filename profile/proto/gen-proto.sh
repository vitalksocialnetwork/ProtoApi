#!/bin/bash
protoc -I ../../../../../github.com/vitalksocialnetwork/ProtoApi/profile/proto/ *.proto --go_out=../../../../../ --go-grpc_out=../../../../../
