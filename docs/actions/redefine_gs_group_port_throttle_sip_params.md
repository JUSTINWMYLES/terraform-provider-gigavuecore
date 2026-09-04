---
page_title: "gigavuecore_redefine_gs_group_port_throttle_sip_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Port Throttle Sip Params
---

# gigavuecore_redefine_gs_group_port_throttle_sip_params Action

Redefine GS Group's Port Throttle Sip Params

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_port_throttle_sip_params" "example" {
  config {
    alias         = "example"
    port_throttle = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `port_throttle` (String, optional) - Alias of referenced Port Throttle


