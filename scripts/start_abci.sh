#!/bin/zsh

home=$GOPATH/src/open-data-repository

cd ${home}/src/open-data-repository-abci
go build
./open-data-repository-abci