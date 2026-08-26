---
page_title: "gigavuecore_update_battery_optimization Action - gigavuecore"
subcategory: ""
description: |-
  update battery optimization
---

# gigavuecore_update_battery_optimization Action

update battery optimization

## Example Usage

```terraform
action "gigavuecore_update_battery_optimization" "example" {
  config {
    cluster_id      = "example"
    cpu_hibernation = null
    monitor_port    = null
    unused_port     = null
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `cpu_hibernation` (Dynamic, optional)
* `monitor_port` (List of Dynamic, optional) - Monitor port off battery optimization
* `unused_port` (List of Dynamic, optional) - Unused port off battery optimization


