#!/bin/bash

set -euo pipefail

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
REPO_DIR="$( cd "$SCRIPT_DIR"/../../ && pwd )"
BUNDLE_DIR="$( mktemp -d -p /tmp )"
REPORTS_DIR="$REPO_DIR/test_output"
CACHE_DIR="$HOME/.cache/dependency-check"
IMAGE_NAME="${1:-"quay.io/signalfx/signalfx-agent-dev:$( "$SCRIPT_DIR"/../current-version )"}"

mkdir -p "$REPORTS_DIR"

cid=$( docker create $IMAGE_NAME true )
docker export $cid | tar -C "$BUNDLE_DIR" -xf -
docker rm -fv $cid

docker build -t dependency-check "$SCRIPT_DIR"

set +e

docker run --rm \
    -v "$BUNDLE_DIR":/bundle \
    -v "$CACHE_DIR":/usr/share/dependency-check/data \
    -v "$REPORTS_DIR":/reports \
    dependency-check \
        --scan /bundle \
        --exclude 'lib/python2.7/site-packages/pip*/' --exclude 'lib/python2.7/ensurepip/' --exclude 'lib/python2.7/site-packages/setuptools*/' \
        --enableExperimental \
        --propertyfile /usr/share/dependency-check/go.properties \
        --suppression /usr/share/dependency-check/suppression.xml \
        --project signalfx-agent -o /reports -f JUNIT -f HTML \
        --junitFailOnCVSS 9 --failOnCVSS 9
rc=$?

rm -rf "$BUNDLE_DIR"

exit $rc
