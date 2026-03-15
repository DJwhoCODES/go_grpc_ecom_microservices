#!/bin/bash

PROTO_DIR=./proto

echo "Generating protobuf files..."

for dir in $PROTO_DIR/*; do
  protoc \
    --go_out=. \
    --go-grpc_out=. \
    $dir/*.proto
done

echo "Proto generation complete"