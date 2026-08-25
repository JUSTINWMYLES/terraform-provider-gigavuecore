---
page_title: "gigavuecore_list_bulk_replicate_config_restore_log Data Source - gigavuecore"
subcategory: ""
description: |-
  Find bulk replicate config restore log file
---

# gigavuecore_list_bulk_replicate_config_restore_log Data Source

Find bulk replicate config restore log file

## Example Usage

```terraform
data "gigavuecore_list_bulk_replicate_config_restore_log" "example" {
  filename = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `filename` (String, required) - Bulk Replicate restore log file name

### Attributes

In addition to all arguments above, the following attributes are exported:

* `created_ts` (Number, computed) - File creation timestamp in UTC milliseconds
* `created_ts_utc` (String, computed) - File creation timestamp in ISO 8601 format


