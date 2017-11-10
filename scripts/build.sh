#!/usr/bin/env bash
set -exuo pipefail

cd "$( dirname "${BASH_SOURCE[0]}" )/.."
source .envrc

GOOS=windows go build -o $BINDIR/compile.exe hwc/compile/cli
