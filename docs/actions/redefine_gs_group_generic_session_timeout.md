---
page_title: "gigavuecore_redefine_gs_group_generic_session_timeout Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Session Timeout
---

# gigavuecore_redefine_gs_group_generic_session_timeout Action

Redefine GS Group's Session Timeout

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_generic_session_timeout" "example" {
  config {
    alias = "example"
    time  = 5
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `time` (Number, optional) - Maximum timeout for session entry


