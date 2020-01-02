#!/usr/bin/env bash
set -e
arg=${1:-"-dev"}
FILENAME=build/blockchain
if [[ "$arg" == "-prod" ]]; then
    FILENAME=bin/blockchain
fi
rm -rf ${FILENAME}
go build -o ${FILENAME} ./src

