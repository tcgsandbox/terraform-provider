#!/bin/bash
# Sync the OpenAPI spec to a local copy for client generation.
#
# If a local spec file exists at the given path (or the default), it is copied.
# Otherwise the spec is fetched from the public API.
#
# Usage: ./scripts/sync-api-docs.sh [OPENAPI_SPEC_PATH]

set -e

PROJECT_DIR="$(git rev-parse --show-toplevel)"
cd "$PROJECT_DIR"

OPENAPI_SPEC="${1:-$HOME/projects/tcg-sandbox/openapi/api.yaml}"
DEST="docs/api/api-spec-copy.yaml"

mkdir -p "$(dirname "$DEST")"

if [ -f "$OPENAPI_SPEC" ]; then
    echo "Copying local spec from $OPENAPI_SPEC..."
    cp "$OPENAPI_SPEC" "$DEST"
else
    echo "Local spec not found at $OPENAPI_SPEC, fetching from API..."
    curl -fSL https://api.tcg-sandbox.com/docs/api.yaml -o "$DEST"
fi

if [ ! -s "$DEST" ]; then
    echo "Error: $DEST is empty or missing after sync"
    exit 1
fi

echo "✓ API spec synced to $DEST"
