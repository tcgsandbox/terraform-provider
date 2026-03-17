# TCG Sandbox Terraform Provider

## Build & Test
- `make generate` — sync API spec, regenerate client, format examples, generate docs
- `make test` — run unit tests
- `make testacc` — run acceptance tests (requires `TF_ACC=1` and live API credentials)
- `make fmt` — format Go code
- `make lint` — run golangci-lint

## Project Structure
- `internal/provider/client_generated.go` — **Auto-generated, DO NOT EDIT.** Regenerate with `make generate`.
- `internal/provider/{entity}_resource.go` — Terraform resource (CRUD)
- `internal/provider/{entity}_data_source.go` — Terraform data source (read-only)
- `internal/provider/{entity}_models.go` — Shared model types and API<->Terraform mappers
- `internal/provider/provider.go` — Provider config and resource/data source registration
- `internal/provider/provider_test.go` — Provider test factories (update when adding new resources)
- `docs/api/api-spec-copy.yaml` — Local copy of the OpenAPI spec (synced by `scripts/sync-api-docs.sh`)
- `scripts/sync-api-docs.sh` — Syncs OpenAPI spec from local path or public API
- `scripts/generate-client.sh` — Generates Go client from the synced spec

## Naming Conventions
- Resource file: `{entity}_resource.go`, struct: `{entity}Resource`, constructor: `New{Entity}Resource()`
- Data source file: `{entity}_data_source.go`, struct: `{entity}DataSource`, constructor: `New{Entity}DataSource()`
- Models file: `{entity}_models.go` for types shared between resource and data source
- Register constructors in `provider.go` → `Resources()` or `DataSources()`
- Terraform type name: `req.ProviderTypeName + "_{entity}"` in `Metadata()`

## Provider Patterns

### API Call Pattern
1. Call client method (e.g., `r.client.CreateGame(ctx, body)`)
2. Parse response: `ParseCreateGameResponse(httpResp)`
3. Check the typed JSON field (e.g., `gameResp.JSON201`); treat `nil` as an error
4. Always re-read the resource after Create/Update to get the server's canonical state

### Image Handling
Local file path → read file → base64 data URL for the API. Track SHA-256 hash in resource private state (`resp.Private.SetKey`) to detect file content changes even when the path hasn't changed. See `game_resource.go` for the full pattern with `readImageAsDataURL()` and `hashImageFile()`.

### Plan Modifiers
- `stringplanmodifier.UseStateForUnknown()` — for Computed+immutable fields (id, owner)
- `boolplanmodifier.UseStateForUnknown()` — for Computed booleans (playable)
- `mapplanmodifier.RequiresReplace()` — for fields that force resource replacement
- `useStateForUnknownOrNullObject()` — custom modifier for Optional+Computed nested objects; resolves to null on create, prior state on update

### Read Preserves Local-Only Fields
When the API doesn't return a field (e.g., `banner_image_path`, `attributes`, `rules`), save the value from prior state before calling `mapGameToResourceState()`, then restore it after.

### Sub-Resources
Resources scoped under a parent (e.g., Cards under Game+Set) need parent IDs as Required string attributes in the schema.

## Coding Practices
- Keep solutions simple and self-contained — one file per resource/data source
- Share model types in `{entity}_models.go` only when both resource and data source exist for the same entity
- Use plan mode for complex provider changes to review approach before implementation
- Prefer straightforward code over abstractions; three similar lines > premature helper function

## Agent Tooling
- `.claude/references/analyze-api-changes.sh` — Syncs spec, regenerates client, produces a structured report of changes, existing implementations, and missing implementations. Report is saved to `/tmp/terraform-provider-api-analysis.md`.
- `/sync-provider` — Orchestration skill that runs the analysis and dispatches parallel subagents to implement missing or changed resources/data sources.
