#!/bin/bash
# Generate the API client from the OpenAPI spec using oapi-codegen
#
# This script generates Go code from the TCG Sandbox OpenAPI spec.
# The generated client is stored in internal/provider/client_generated.go
#
# Usage: ./scripts/generate-client.sh [OPENAPI_SPEC_PATH]
#
# Examples:
#   ./scripts/generate-client.sh
#   ./scripts/generate-client.sh /path/to/api.yaml

set -e

# Default to ~/projects/tcg-sandbox/openapi/api.yaml
OPENAPI_SPEC="${1:-$HOME/projects/tcg-sandbox/openapi/api.yaml}"

# Validate the OpenAPI spec file exists
if [ ! -f "$OPENAPI_SPEC" ]; then
    echo "Error: OpenAPI spec not found at $OPENAPI_SPEC"
    echo "Usage: ./scripts/generate-client.sh [OPENAPI_SPEC_PATH]"
    exit 1
fi

# Check if oapi-codegen is installed, install if not
if ! command -v oapi-codegen &> /dev/null; then
    echo "Installing oapi-codegen..."
    go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
fi

echo "Generating client from OpenAPI spec at $OPENAPI_SPEC..."
oapi-codegen -package provider -generate types,client \
    "$OPENAPI_SPEC" > internal/provider/client_generated.go

echo "✓ Client generated successfully at internal/provider/client_generated.go"
