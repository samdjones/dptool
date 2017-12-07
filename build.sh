#!/bin/bash

export GOPATH=`pwd`

# Clean up to begin with
go clean
rm -rf out
rm -rf pkg

# Optionally, get our deps. Whatever version happens to be current today - yeah I know, good luck :-)
# go get github.com/inconshreveable/mousetrap
# go get github.com/spf13/cobra
# go get github.com/fsnotify/fsnotify

# Mac build
mkdir -p bin/darwin
GOOS=darwin GOARCH=amd64 go build -o bin/darwin/dptool github.com/samdjones/dptool

# Linux build
mkdir -p bin/linux
GOOS=linux GOARCH=amd64 go build -o bin/linux/dptool github.com/samdjones/dptool

# Windows 64 build
mkdir -p bin/windows64
GOOS=windows GOARCH=amd64 go build -o bin/windows64/dptool.exe github.com/samdjones/dptool

# Windows 32 build
mkdir -p bin/windows32
GOOS=windows GOARCH=386 go build -o bin/windows32/dptool.exe github.com/samdjones/dptool

# Now have some fun. Or something.
