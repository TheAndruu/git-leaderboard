#!/bin/bash

# Everything is 64 bit
GOARCH=amd64

# Build windows
GOOS=windows
go build -v -o build/show-commits-x64.exe

GOARCH=386
go build -v -o build/show-commits-x32.exe


# Build mac
GOOS=darwin
go build -v -o build/show-commits-mac

# Build windows
GOOS=linux
go build -v -o build/show-commits
