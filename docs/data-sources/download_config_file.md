---
page_title: "gigavuecore_download_config_file Data Source - gigavuecore"
subcategory: ""
description: |-
  Download a config file
---

# gigavuecore_download_config_file Data Source

Download a config file

## Example Usage

```terraform
data "gigavuecore_download_config_file" "example" {
  backup_id  = null
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `backup_id` (String, required) - id of the config backup snapshot to download
* `cluster_id` (String, required) - ID of the target device


