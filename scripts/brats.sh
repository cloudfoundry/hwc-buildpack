#!/usr/bin/env bash
# Test that the compiled binaries of the buildpacks are working as expected

set -euo pipefail

cd "$( dirname "${BASH_SOURCE[0]}" )/.."
source .envrc
./scripts/install_tools.sh

GINKGO_NODES=${GINKGO_NODES:-3}
GINKGO_ATTEMPTS=${GINKGO_ATTEMPTS:-1}
export CF_STACK=${CF_STACK:-windows}

cd src/*/brats

echo "Run Buildpack Runtime Acceptance Tests"
ginkgo -mod vendor -r --flakeAttempts=$GINKGO_ATTEMPTS -nodes $GINKGO_NODES
