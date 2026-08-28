---
page_title: "gigavuecore_redefine_gs_group_hsm_group_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's HSM Group Params
---

# gigavuecore_redefine_gs_group_hsm_group_params Action

Redefine GS Group's HSM Group Params

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_hsm_group_params" "example" {
  config {
    alias     = "example"
    hsm_group = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `hsm_group` (String, optional) - alias of Hsm Group


