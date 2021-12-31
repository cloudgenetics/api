#!/usr/bin/env bash 
set -xe
# install packages and dependencies
go get -d
# build command
GOOS=linux GOARCH=amd64 go build -o bin/application application.go
