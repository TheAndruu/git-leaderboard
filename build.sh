#!/bin/bash

# Everything is 64 bit
GOARCH=amd64

# Build windows
GOOS=windows
go build -v -o build/show-commits.exe

# Build mac
GOOS=darwin
go build -v -o build/show-commits-mac

# Build windows
GOOS=linux
go build -v -o build/show-commits
