#!/bin/ash

# create go.mod file
cd src/simple-api/ && go mod init

# get client-go module
go get k8s.io/client-go@v12.0.0
