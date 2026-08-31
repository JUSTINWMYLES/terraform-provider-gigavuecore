---
page_title: "gigavuecore_redefine_gs_group_sip_session Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Sip Session
---

# gigavuecore_redefine_gs_group_sip_session Action

Redefine GS Group's Sip Session

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_sip_session" "example" {
  config {
    alias   = "example"
    timeout = 30
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `timeout` (Number, optional) - Sip session inactivity timer, value in seconds .Valid values: 30-300.


