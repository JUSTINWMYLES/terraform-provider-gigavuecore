---
page_title: "gigavuecore_redefine_gs_group_diameter_s6_a_session Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Diameter s6a session
---

# gigavuecore_redefine_gs_group_diameter_s6_a_session Action

Redefine GS Group's Diameter s6a session

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_diameter_s6_a_session" "example" {
  config {
    alias   = "example"
    limit   = 1
    timeout = 30
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `limit` (Number, optional) - Number of sessions to allocate for Diameter S6A
* `timeout` (Number, optional) - timeout in seconds used to clean inactive sessions


