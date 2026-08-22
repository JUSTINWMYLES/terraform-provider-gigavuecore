---
page_title: "gigavuecore_load_cluster_port_filter_counters Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Cluster Port Filter Counters
---

# gigavuecore_load_cluster_port_filter_counters Data Source

Load Cluster Port Filter Counters

## Example Usage

```terraform
data "gigavuecore_load_cluster_port_filter_counters" "example" {
  cluster_id = null
  port_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `port_id` (String, required) - port id to filter by

### Attributes

In addition to all arguments above, the following attributes are exported:

* `port_filter_stats` (Object({port, rules}), computed) - Egress Port Filter counters
  * `port` (String, computed)
  * `rules` (Object({drop_rules, pass_rules}), computed) - Port Filter Rules Container
    * `drop_rules` (List(Object({bytes, filters, pkts, rule_id})), computed)
    * `pass_rules` (List(Object({bytes, filters, pkts, rule_id})), computed)
* `sys_up_time` (Number, computed) - Device sysUpTime. The time (in hundredths of a second) since the network management portion of the system was last re-initialized

