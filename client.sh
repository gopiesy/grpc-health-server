#!/usr/bin/env bash

set -x

REQ='{"service": "health"}'
URL='localhost:9111'
RPC='grpc.health.v1.Health/Check'
ROOT_CA='./certs/cacert.pem'
CLIENT_CRT='./certs/client.pem'
CLIENT_KEY='./certs/client.key'

#~/go/bin/grpcurl --plaintext --import-path ./proto --proto health.proto \
#    -d "$REQ" \
#    $URL \
#    $RPC

~/go/bin/grpcurl --insecure --import-path ./proto --proto health.proto \
    -cacert "$ROOT_CA"\
    -key "$CLIENT_KEY"\
    -cert "$CLIENT_CRT"\
    -d "$REQ" \
    $URL \
    $RPC
