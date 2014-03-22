#!/bin/bash

pkill -f mhd2014

cd "$GOPATH/src/github.com/advincze/mhd2014"
go get -t ./...
go install ./...
"$GOPATH/bin/mhd2014" -port=8080 2>&1> /dev/null &

