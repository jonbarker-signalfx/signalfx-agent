#!/bin/bash

set -euo pipefail

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
REPO_DIR="$( cd "$SCRIPT_DIR"/../../ && pwd )"
REPORTS_DIR="$REPO_DIR/test_output"
CACHE_DIR="$HOME/.cache/dependency-check"

mkdir -p "$REPORTS_DIR"

docker build -t dependency-check "$SCRIPT_DIR"

docker run --rm \
    -v "$REPO_DIR":/src \
    -v "$CACHE_DIR":/usr/share/dependency-check/data \
    -v "$REPORTS_DIR":/reports \
    dependency-check \
        --scan /src \
        --exclude 'scripts/' --exclude 'tests/' --exclude 'test-services/' --exclude 'vendor/' \
        --enableExperimental \
        --propertyfile /usr/share/dependency-check/go.properties \
        --suppression /usr/share/dependency-check/suppression.xml \
        --project signalfx-agent -o /reports -f JUNIT -f HTML \
        --junitFailOnCVSS 9 --failOnCVSS 9
