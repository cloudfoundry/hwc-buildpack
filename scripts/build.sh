#!/usr/bin/env bash

ROOTDIR="$( dirname "$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )" )"
BINDIR=$ROOTDIR/bin

set -ex

GOOS=windows go build -o $BINDIR/hwc.exe github.com/greenhouse-org/hwc-buildpack/hwc
GOOS=windows go build -o $BINDIR/detect.exe github.com/greenhouse-org/hwc-buildpack/detect
GOOS=windows go build -o $BINDIR/compile.exe github.com/greenhouse-org/hwc-buildpack/compile
GOOS=windows go build -o $BINDIR/release.exe github.com/greenhouse-org/hwc-buildpack/release

pushd $ROOTDIR &> /dev/null
zip -r hwc_buildpack.zip bin
popd &> /dev/null
