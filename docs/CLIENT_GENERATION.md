# API Client Generation

This document explains how the Terraform provider generates its API client from the TCG Sandbox OpenAPI specification.

## Overview

Rather than manually writing the API client, we auto-generate it from the OpenAPI spec using `oapi-codegen`. This ensures:

- **Type safety**: All request/response types are automatically generated
- **Consistency**: Client stays in sync with API changes
- **Maintainability**: No manual boilerplate to maintain
- **Coverage**: All API endpoints are automatically included

## Generated Files

- `internal/provider/client_generated.go` - Auto-generated API client (do not edit)

## How to Regenerate the Client

Run the generation script:

```bash
./scripts/generate-client.sh
```

By default, it looks for the OpenAPI spec at `$HOME/projects/tcg-sandbox/openapi/api.yaml`. You can specify a custom path:

```bash
./scripts/generate-client.sh /path/to/api.yaml
```

## What Gets Generated

The script generates:

1. **Types** - All request/response models from the OpenAPI schemas
2. **Client** - HTTP client with methods for each API endpoint
3. **Auto-generated code** - Handles serialization, validation, etc.

## Customizing Generation

To change what gets generated, edit `scripts/generate-client.sh` and modify the `oapi-codegen` flags:

```bash
# Current: generates types and client
oapi-codegen -package client -generate types,client "$OPENAPI_SPEC"

# Other options:
# -generate types,client,server   # Also generate server interfaces
# -generate types                 # Only types
```

See [oapi-codegen documentation](https://github.com/oapi-codegen/oapi-codegen) for more options.

## When to Regenerate

Regenerate the client after:

- API changes in the OpenAPI spec
- New endpoints are added
- Request/response schemas change
- Adding new features that need new endpoints
