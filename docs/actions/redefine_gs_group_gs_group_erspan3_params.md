---
page_title: "gigavuecore_redefine_gs_group_gs_group_erspan3_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's ERSPAN III Params
---

# gigavuecore_redefine_gs_group_gs_group_erspan3_params Action

Redefine GS Group's ERSPAN III Params

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_gs_group_erspan3_params" "example" {
  config {
    alias = "example"
    cluster_id = "example"
    timestamp_format = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `timestamp_format` (String, optional) - erspan III tunnelDecap timestamp format
