---
page_title: "gigavuecore_redefine_gs_group_port_throttle_sip_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Port Throttle Sip Params
---

# gigavuecore_redefine_gs_group_port_throttle_sip_params Action

Redefine GS Group's Port Throttle Sip Params

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


