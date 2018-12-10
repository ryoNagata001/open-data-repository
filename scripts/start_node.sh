#!/bin/zsh

home=$GOPATH/src/open-data-repository

tendermint init --home ${home}
tendermint node --home ${home} --consensus.create_empty_blocks=false #ここに他のパラメーターも加えれば、常時追加可能だろう