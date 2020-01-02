#!/usr/bin/env bash
set -e
arg=${1:-"-dev"}
DIR=build

[[ "$arg" == "-prod" ]] && DIR=bin
[ ! -d ${DIR} ] && mkdir DIR

FILENAME=${DIR}/blockchain
rm -rf ${FILENAME}
go build -o ${FILENAME} ./src
