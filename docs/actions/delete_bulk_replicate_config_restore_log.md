---
page_title: "gigavuecore_delete_bulk_replicate_config_restore_log Action - gigavuecore"
subcategory: ""
description: |-
  Delete bulk replicate config restore log
---

# gigavuecore_delete_bulk_replicate_config_restore_log Action

Delete bulk replicate config restore log

## Example Usage

```terraform
action "gigavuecore_delete_bulk_replicate_config_restore_log" "example" {
  config {
    filename = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `filename` (String, required) - Target restore log file name


