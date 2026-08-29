---
page_title: "gigavuecore_redefine_gs_group_gs_group_flow_mask_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Flow Mask Params
---

# gigavuecore_redefine_gs_group_gs_group_flow_mask_params Action

Redefine GS Group's Flow Mask Params

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_gs_group_flow_mask_params" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    enabled    = true
    length     = 0
    offset     = 0
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `enabled` (Boolean, optional)
* `length` (Number, optional)
* `offset` (Number, optional)


