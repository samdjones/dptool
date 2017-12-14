#!/bin/bash

#
# You really don't need to use this for normal dev builds, just us gb directly, e.g.:
#
#   $ gb build
#
# This script only exists to build a pile of platform binaries for a release.
# I know, it's kinda dirty and I should make a Makefile, in my Infinite Free Time (TM)...
#

rm -r bin

GOOS=darwin GOARCH=amd64 gb build
GOOS=linux GOARCH=amd64 gb build
GOOS=linux GOARCH=386 gb build
GOOS=windows GOARCH=amd64 gb build
GOOS=windows GOARCH=386 gb build

mkdir bin/dptoolbinaries
cp README.md bin/dptoolbinaries
cd bin
mv dptool-* dptoolbinaries
zip -r dptoolbinaries.zip dptoolbinaries/*
cd ..
