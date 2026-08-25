---
page_title: "gigavuecore_redefine_gs_group_node_role_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Node Role Params
---

# gigavuecore_redefine_gs_group_node_role_params Action

Redefine GS Group's Node Role Params

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_node_role_params" "example" {
  config {
    alias            = "example"
    mob5_g_limit     = 1
    mob_lte_limit    = 1
    stand_alone_mode = true
    type             = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `mob5_g_limit` (Number, optional) - Number of sessions to allocate for Control 5G Node
* `mob_lte_limit` (Number, optional) - Number of sessions to allocate for LTE CPN / UPN Node
* `stand_alone_mode` (Boolean, optional) - Enables UPN Stand-alone mode
* `type` (String, optional)


