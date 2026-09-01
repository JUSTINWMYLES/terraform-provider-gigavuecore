---
page_title: "gigavuecore_redefine_gs_group_engine_watchdog_timer Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Engine Watchdog Timer
---

# gigavuecore_redefine_gs_group_engine_watchdog_timer Action

Redefine GS Group's Engine Watchdog Timer

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_engine_watchdog_timer" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    time       = 0
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `time` (Number, optional) - Engine Watchdog Timer value, time to restart in seconds .Valid values: 60-600. 0 is disable


