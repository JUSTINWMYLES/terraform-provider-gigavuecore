---
page_title: "gigavuecore_config Resource - gigavuecore"
subcategory: ""
description: |-
  Find bulk replicate config file by file name
---

# gigavuecore_config Resource

Find bulk replicate config file by file name

## Example Usage

```terraform
resource "gigavuecore_config" "example" {
  filename = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `filename` (String, required) - Bulk Replicate config file name

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `comment` (String, computed) - User comments for the config file
* `created_ts` (Number, computed) - File creation timestamp in UTC milliseconds
* `created_ts_utc` (String, computed) - File creation timestamp in ISO 8601 format
* `family` (String, computed) - \['H' or 'G'\]: identifies whether this is an H-series or a G-series Config file


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_config.example {filename}
```
