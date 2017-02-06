#!/usr/bin/env bash

ROOTDIR="$( dirname "$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )" )"
BINDIR=$ROOTDIR/bin
export GOPATH=$ROOTDIR

set -ex

GOOS=windows go build -o $BINDIR/compile.exe compile
