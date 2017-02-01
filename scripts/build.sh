#!/usr/bin/env bash

ROOTDIR="$( dirname "$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )" )"
BINDIR=$ROOTDIR/bin

set -ex

GOOS=windows go build -o $BINDIR/hwc.exe ./src/hwc
