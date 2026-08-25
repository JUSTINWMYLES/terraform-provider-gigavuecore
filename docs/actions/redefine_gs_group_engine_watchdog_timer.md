---
page_title: "gigavuecore_redefine_gs_group_engine_watchdog_timer Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Engine Watchdog Timer
---

# gigavuecore_redefine_gs_group_engine_watchdog_timer Action

Redefine GS Group's Engine Watchdog Timer

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_engine_watchdog_timer" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    time       = 1
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `time` (Number, optional) - Engine Watchdog Timer value, time to restart in seconds .Valid values: 60-600. 0 is disable


