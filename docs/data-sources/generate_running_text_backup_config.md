---
page_title: "gigavuecore_generate_running_text_backup_config Data Source - gigavuecore"
subcategory: ""
description: |-
  Generate running text backup config
---

# gigavuecore_generate_running_text_backup_config Data Source

Generate running text backup config

~> **Note:** This data source is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
data "gigavuecore_generate_running_text_backup_config" "example" {
  backup_id  = "example"
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `backup_id` (String, required) - id of the config backup snapshot to download
* `cluster_id` (String, required) - ID of the target device


