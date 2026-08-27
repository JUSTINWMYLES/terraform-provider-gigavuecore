---
page_title: "gigavuecore_redefine_gs_group_tcp_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's TCP Params
---

# gigavuecore_redefine_gs_group_tcp_params Action

Redefine GS Group's TCP Params

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_tcp_params" "example" {
  config {
    alias        = "example"
    application  = "example"
    load_balance = true
    tcp_control  = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `application` (String, optional) - To choose the action on Unknown Application Data
* `load_balance` (Boolean, optional) - Enables TCP loadbalancing for the Tool Ports. Supports SIP & N11 only
* `tcp_control` (String, optional) - To choose the action on TCP Control messages


