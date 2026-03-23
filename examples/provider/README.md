# Provider Example

This directory contains a working example of the tcg-sandbox provider. Outside of acting as an example for reference, it can also be used for local development and manual testing of the `tcg-sandbox` provider.

## Using the Local Dev Build

A `.terraformrc-dev` file is included that points Terraform at your locally built binary instead of the published registry version. It is scoped to this directory only — your other Terraform projects are unaffected.

> **Setup required:** `.terraformrc-dev` contains placeholder values. Copy it to a local file and update it for your system:
> ```bash
> cp .terraformrc-dev .terraformrc.local
> ```
> Then edit `.terraformrc.local` and replace `your-org` with your registry namespace and the binary path with the directory where `go install` places your binary (e.g. `~/go/bin` on most systems). Point `TF_CLI_CONFIG_FILE` at this local copy instead.

### Activate config

```bash
export TF_CLI_CONFIG_FILE="$(pwd)/.terraformrc.local"
```

## Running the Example

Make sure the API is running locally, then build and install the provider:

```bash
# From the repo root
go install .
```

Then apply the configuration:

```bash
terraform apply
```

> **Note:** `terraform init` is not required (and will warn) when using `dev_overrides` — Terraform skips version resolution for overridden providers.
