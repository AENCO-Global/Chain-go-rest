#!/usr/bin/env bash

# Set the GoPath must be set (assume the deployed folder is the path)
export GOPATH=~/go
export GOBIN=$GOPATH/pkg
export GOEXE=$GOPATH/bin
if [[ $PATH != *"go"* ]]; then
    export PATH=$PATH:$GOPATH/bin
fi

# Install the Libraries for Go to build
go get -v -u github.com/gorilla/mux
go get -v -u github.com/AENCO-Global/Chain-go-sdk
go get -v -u github.com/AENCO-Global/Chain-go-rest
