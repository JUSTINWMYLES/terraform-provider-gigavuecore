---
page_title: "gigavuecore_download_backup_config_file_of_format_text Data Source - gigavuecore"
subcategory: ""
description: |-
  Download a Backup config file of format text
---

# gigavuecore_download_backup_config_file_of_format_text Data Source

Download a Backup config file of format text

~> **Note:** This data source is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
data "gigavuecore_download_backup_config_file_of_format_text" "example" {
  backup_id  = "example"
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `backup_id` (String, required) - id of the config backup snapshot to download
* `cluster_id` (String, required) - ID of the target device


