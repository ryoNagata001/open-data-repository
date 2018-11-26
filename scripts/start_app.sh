#!/bin/zsh

home=$GOPATH/src/open-data-repository

cd ${home}/src/open-data-repository-app
go build
./open-data-repository-app