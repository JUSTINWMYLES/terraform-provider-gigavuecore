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
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `cpu_hibernation` (Attributes, computed) (see [below for nested schema](#nestedatt--cpu_hibernation))
* `monitor_port` (Attributes List, computed) - Monitor port off battery optimization (see [below for nested schema](#nestedatt--monitor_port))
* `unused_port` (Attributes List, computed) - Unused port off battery optimization (see [below for nested schema](#nestedatt--unused_port))

<a id="nestedatt--cpu_hibernation"></a>
### Nested Schema for `cpu_hibernation`

Read-Only:

* `enabled` (Boolean) - Enable/disable cpu hibernation battery optimization
* `sleep_in_mins` (Number) - Apply cpu hibernation battery optimization for this configured sleep time
* `threshold_level` (Number) - Apply cpu hibernation battery optimization at this battery level
<a id="nestedatt--monitor_port"></a>
### Nested Schema for `monitor_port`

Read-Only:

* `enabled` (Boolean) - Enable/disable port battery optimization
* `level` (String) - Apply port battery optimization at this battery level
* `port_group_id` (String) - Port group Id
<a id="nestedatt--unused_port"></a>
### Nested Schema for `unused_port`

Read-Only:

* `enabled` (Boolean) - Enable/disable port battery optimization
* `level` (String) - Apply port battery optimization at this battery level
* `port_group_id` (String) - Port group Id

