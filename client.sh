#!/usr/bin/env bash

set -x

REQ='{"service": "health"}'
URL='localhost:9111'
RPC='grpc.health.v1.Health/Check'

~/go/bin/grpcurl --plaintext --import-path ./proto --proto health.proto \
    -d "$REQ" \
    $URL \
    $RPC
