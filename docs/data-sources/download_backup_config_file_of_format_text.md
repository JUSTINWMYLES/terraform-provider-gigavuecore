---
page_title: "gigavuecore_download_backup_config_file_of_format_text Data Source - gigavuecore"
subcategory: ""
description: |-
  Download a Backup config file of format text
---

# gigavuecore_download_backup_config_file_of_format_text Data Source

Download a Backup config file of format text

## Example Usage

```terraform
data "gigavuecore_download_backup_config_file_of_format_text" "example" {
  backup_id = null
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `backup_id` (String, required) - id of the config backup snapshot to download
* `cluster_id` (String, required) - ID of the target device

### Attributes

In addition to all arguments above, the following attributes are exported:


