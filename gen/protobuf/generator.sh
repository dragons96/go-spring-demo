#!/bin/bash
pathprefix=$1

protoc -I %pathprefix% --go_out=%pathprefix%/internal/proto/pb --go-grpc_out=%pathprefix%/internal/proto/pb %pathprefix%/internal/proto/*.proto
