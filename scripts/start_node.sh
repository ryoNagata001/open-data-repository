#!/bin/zsh

home=$GOPATH/src/open-data-repository

tendermint init --home ${home}
tendermint node --home ${home}