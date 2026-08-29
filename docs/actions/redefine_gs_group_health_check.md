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
    action          = "example"
    alias           = "example"
    dst_port        = 0
    enabled         = true
    interval        = 0
    protocol        = "example"
    rcv_port        = 0
    retries         = 0
    round_trip_time = 0
    src_port        = 0
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


