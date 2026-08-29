---
page_title: "gigavuecore_redefine_gs_group_sip_tcp_idle_timeout Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Sip Tcp Idle Timeout
---

# gigavuecore_redefine_gs_group_sip_tcp_idle_timeout Action

Redefine GS Group's Sip Tcp Idle Timeout

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_sip_tcp_idle_timeout" "example" {
  config {
    alias = "example"
    time  = 0
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `time` (Number, optional) - Sip tcp idle timeout value in seconds .Valid values: 20-600.


