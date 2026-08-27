---
page_title: "gigavuecore_delete_fabric_map_with_last_update_timestamp_validation Action - gigavuecore"
subcategory: ""
description: |-
  since FM 6.8.00
---

# gigavuecore_delete_fabric_map_with_last_update_timestamp_validation Action

since FM 6.8.00

## Example Usage

```terraform
action "gigavuecore_delete_fabric_map_with_last_update_timestamp_validation" "example" {
  config {
    alias        = "example"
    updated_time = 1.0
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the fabric map to be deleted
* `updated_time` (Number, required) - last updated timestamp of the fabric map to be deleted


