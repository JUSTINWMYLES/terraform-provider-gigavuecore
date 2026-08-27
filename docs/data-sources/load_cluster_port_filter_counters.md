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
  port_id    = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `port_id` (String, required) - port id to filter by

### Attributes

In addition to all arguments above, the following attributes are exported:

* `port_filter_stats` (Attributes, computed) - Egress Port Filter counters (see [below for nested schema](#nestedatt--port_filter_stats))
* `sys_up_time` (Number, computed) - Device sysUpTime. The time (in hundredths of a second) since the network management portion of the system was last re-initialized

<a id="nestedatt--port_filter_stats"></a>
### Nested Schema for `port_filter_stats`

Read-Only:

* `port` (String)
* `rules` (Attributes) - Port Filter Rules Container (see [below for nested schema](#nestedatt--port_filter_stats--rules))
<a id="nestedatt--port_filter_stats--rules"></a>
### Nested Schema for `port_filter_stats.rules`

Read-Only:

* `drop_rules` (Attributes List) (see [below for nested schema](#nestedatt--port_filter_stats--rules--drop_rules))
* `pass_rules` (Attributes List) (see [below for nested schema](#nestedatt--port_filter_stats--rules--pass_rules))
<a id="nestedatt--port_filter_stats--rules--drop_rules"></a>
### Nested Schema for `port_filter_stats.rules.drop_rules`

Read-Only:

* `bytes` (Number)
* `filters` (Attributes List) - list of the filters configured for rule (see [below for nested schema](#nestedatt--port_filter_stats--rules--drop_rules--filters))
* `pkts` (Number)
* `rule_id` (Number)
<a id="nestedatt--port_filter_stats--rules--drop_rules--filters"></a>
### Nested Schema for `port_filter_stats.rules.drop_rules.filters`

Read-Only:

* `filter_type` (String)
* `filter_value` (String)
<a id="nestedatt--port_filter_stats--rules--pass_rules"></a>
### Nested Schema for `port_filter_stats.rules.pass_rules`

Read-Only:

* `bytes` (Number)
* `filters` (Attributes List) - list of the filters configured for rule (see [below for nested schema](#nestedatt--port_filter_stats--rules--pass_rules--filters))
* `pkts` (Number)
* `rule_id` (Number)
<a id="nestedatt--port_filter_stats--rules--pass_rules--filters"></a>
### Nested Schema for `port_filter_stats.rules.pass_rules.filters`

Read-Only:

* `filter_type` (String)
* `filter_value` (String)

