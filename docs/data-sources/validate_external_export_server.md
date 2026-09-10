---
page_title: "gigavuecore_validate_external_export_server Data Source - gigavuecore"
subcategory: ""
description: |-
  Validate External Export Server
---

# gigavuecore_validate_external_export_server Data Source

Validate External Export Server

~> **Note:** This data source is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
data "gigavuecore_validate_external_export_server" "example" {
  export_target_alias = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `export_target_alias` (String, required) - Alias for external export target target server


