#!/bin/bash

export GOPATH=/Users/sam/code/GitHub/dptool

# Clean up to begin with
go clean
rm -rf out
# rm -rf pkg

# Optioally get our deps. Whatever version happens to be current today - yeah I know, good luck :-)
# go get github.com/inconshreveable/mousetrap
# go get github.com/spf13/cobra
# go get github.com/fsnotify/fsnotify

cd src

# Mac build
mkdir -p out/darwin
GOOS=darwin GOARCH=amd64 go build -o ../out/darwin/dptool github.com/samdjones/dptool

# Linux build
mkdir -p out/linux
GOOS=linux GOARCH=amd64 go build -o ../out/linux/dptool github.com/samdjones/dptool

# Windows build
mkdir -p out/windows
GOOS=windows GOARCH=amd64 go build -o ../out/windows/dptool.exe github.com/samdjones/dptool

# Now have some fun. Or something.
