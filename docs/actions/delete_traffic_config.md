---
page_title: "gigavuecore_delete_traffic_config Action - gigavuecore"
subcategory: ""
description: |-
  Delete all traffic configurations
---

# gigavuecore_delete_traffic_config Action

Delete all traffic configurations

## Example Usage

```terraform
action "gigavuecore_delete_traffic_config" "example" {
  config {
    keep_stack = true
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `keep_stack` (Boolean, optional) - If true, delete all traffic but keep stack configurations. Default: false


