---
page_title: "gigavuecore_load_battery_optimization Data Source - gigavuecore"
subcategory: ""
description: |-
  load all battery optimization profile
---

# gigavuecore_load_battery_optimization Data Source

load all battery optimization profile

## Example Usage

```terraform
data "gigavuecore_load_battery_optimization" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `cpu_hibernation` (Object({enabled, sleep_in_mins, threshold_level}), computed)
  * `enabled` (Bool, computed) - Enable/disable cpu hibernation battery optimization
  * `sleep_in_mins` (Number, computed) - Apply cpu hibernation battery optimization for this configured sleep time
  * `threshold_level` (Number, computed) - Apply cpu hibernation battery optimization at this battery level
* `monitor_port` (List(Object({enabled, level, port_group_id})), computed) - Monitor port off battery optimization
* `unused_port` (List(Object({enabled, level, port_group_id})), computed) - Unused port off battery optimization

