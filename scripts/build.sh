#!/usr/bin/env bash

BINDIR="$( dirname "$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )" )/bin"

set -ex

GOOS=windows go build -o $BINDIR/hwc.exe github.com/greenhouse-org/hwc-buildpack/hwc
GOOS=windows go build -o $BINDIR/detect.exe github.com/greenhouse-org/hwc-buildpack/detect
GOOS=windows go build -o $BINDIR/compile.exe github.com/greenhouse-org/hwc-buildpack/compile
GOOS=windows go build -o $BINDIR/release.exe github.com/greenhouse-org/hwc-buildpack/release

zip -r hwc_buildpack.zip $BINDIR
