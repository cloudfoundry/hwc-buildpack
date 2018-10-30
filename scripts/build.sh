#!/usr/bin/env bash
set -exuo pipefail

cd "$( dirname "${BASH_SOURCE[0]}" )/.."
source .envrc



GOOS=windows go build -o bin/compile.exe ./src/hwc/compile/cli
GOOS=windows go build -o bin/supply.exe ./src/hwc/supply/cli
GOOS=windows go build -o bin/finalize.exe ./src/hwc/finalize/cli
