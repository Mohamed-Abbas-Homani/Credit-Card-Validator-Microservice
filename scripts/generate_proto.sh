#!/bin/sh

set -e

PROTO_PATH=./pkg/proto
PROTO_FILE=cardvalidator.proto
OUT_DIR=.

protoc --go_out=$OUT_DIR --go_opt=paths=source_relative \
    --go-grpc_out=$OUT_DIR --go-grpc_opt=paths=source_relative \
    $PROTO_PATH/$PROTO_FILE