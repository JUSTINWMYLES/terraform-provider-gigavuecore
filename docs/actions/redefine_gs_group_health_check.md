---
page_title: "gigavuecore_redefine_gs_group_health_check Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's health check
---

# gigavuecore_redefine_gs_group_health_check Action

Redefine GS Group's health check

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_health_check" "example" {
  config {
    action          = "pass"
    alias           = "example"
    dst_port        = 1
    enabled         = true
    interval        = 5
    protocol        = "icmp"
    rcv_port        = 1
    retries         = 1
    round_trip_time = 1
    src_port        = 1
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `action` (String, optional)
* `alias` (String, required) - alias of the target GS Group
* `dst_port` (Number, optional)
* `enabled` (Boolean, optional)
* `interval` (Number, optional)
* `protocol` (String, optional)
* `rcv_port` (Number, optional)
* `retries` (Number, optional)
* `round_trip_time` (Number, optional)
* `src_port` (Number, optional)


