---
page_title: "gigavuecore_redefine_gs_group_hsm_group_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's HSM Group Params
---

# gigavuecore_redefine_gs_group_hsm_group_params Action

Redefine GS Group's HSM Group Params

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

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


