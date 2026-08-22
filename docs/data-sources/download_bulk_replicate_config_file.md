---
page_title: "gigavuecore_download_bulk_replicate_config_file Data Source - gigavuecore"
subcategory: ""
description: |-
  Download bulk replicate config file
---

# gigavuecore_download_bulk_replicate_config_file Data Source

Download bulk replicate config file

## Example Usage

```terraform
data "gigavuecore_download_bulk_replicate_config_file" "example" {
  filename = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `filename` (String, required) - Target config file name

### Attributes

In addition to all arguments above, the following attributes are exported:


