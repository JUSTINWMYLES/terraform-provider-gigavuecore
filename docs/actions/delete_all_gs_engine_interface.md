---
page_title: "gigavuecore_delete_all_gs_engine_interface Action - gigavuecore"
subcategory: ""
description: |-
  Delete All GigaSMART port interface
---

# gigavuecore_delete_all_gs_engine_interface Action

Delete All GigaSMART port interface

## Example Usage

```terraform
action "gigavuecore_delete_all_gs_engine_interface" "example" {
  config {
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
