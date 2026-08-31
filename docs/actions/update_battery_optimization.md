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
    cluster_id = "example"
    cpu_hibernation = {
      enabled         = true
      sleep_in_mins   = 6
      threshold_level = 25
    }
    monitor_port = [{
      enabled       = true
      level         = "alarm"
      port_group_id = "example"
    }]
    unused_port = [{
      enabled       = true
      level         = "alarm"
      port_group_id = "example"
    }]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `cpu_hibernation` (Attributes, optional) - CPU hibernation battery optimization (see [below for nested schema](#nestedatt--cpu_hibernation))
* `monitor_port` (Attributes List, optional) - Monitor port off battery optimization (see [below for nested schema](#nestedatt--monitor_port))
* `unused_port` (Attributes List, optional) - Unused port off battery optimization (see [below for nested schema](#nestedatt--unused_port))

<a id="nestedatt--cpu_hibernation"></a>
### Nested Schema for `cpu_hibernation`

Optional:

* `enabled` (Boolean) - Enable/disable cpu hibernation battery optimization
* `sleep_in_mins` (Number) - Apply cpu hibernation battery optimization for this configured sleep time
* `threshold_level` (Number) - Apply cpu hibernation battery optimization at this battery level

<a id="nestedatt--monitor_port"></a>
### Nested Schema for `monitor_port`

Required:

* `port_group_id` (String) - Port group Id

Optional:

* `enabled` (Boolean) - Enable/disable port battery optimization
* `level` (String) - Apply port battery optimization at this battery level

<a id="nestedatt--unused_port"></a>
### Nested Schema for `unused_port`

Required:

* `port_group_id` (String) - Port group Id

Optional:

* `enabled` (Boolean) - Enable/disable port battery optimization
* `level` (String) - Apply port battery optimization at this battery level

