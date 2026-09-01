---
page_title: "gigavuecore_redefine_gs_group_gs_group_flow_mask_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Flow Mask Params
---

# gigavuecore_redefine_gs_group_gs_group_flow_mask_params Action

Redefine GS Group's Flow Mask Params

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_gs_group_flow_mask_params" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    enabled    = true
    length     = 1
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


