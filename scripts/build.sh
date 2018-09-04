#!/usr/bin/env bash
set -exuo pipefail

cd "$( dirname "${BASH_SOURCE[0]}" )/.."
source .envrc

GOOS=windows go build -o bin/compile.exe hwc/compile/cli
GOOS=windows go build -o bin/supply.exe hwc/supply/cli
GOOS=windows go build -o bin/finalize.exe hwc/finalize/cli
