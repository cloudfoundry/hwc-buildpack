#!/bin/bash
set -euo pipefail

cd "$( dirname "${BASH_SOURCE[0]}" )/.."
source .envrc

if [ ! -f .bin/ginkgo ]; then
  go install github.com/onsi/ginkgo
# (cd src/*/vendor/github.com/onsi/ginkgo/ginkgo/ && go install)
fi
if [ ! -f .bin/buildpack-packager ]; then
  go install github.com/cloudfoundry/libbuildpack/packager/buildpack-packager
# (cd src/*/vendor/github.com/cloudfoundry/libbuildpack/packager/buildpack-packager && go install)
fi
