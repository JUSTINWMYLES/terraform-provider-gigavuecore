---
page_title: "gigavuecore_get_correlated_alarms_count_by_parameter Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Counts of Correlated Alarms Grouped By The Requested Alarm Parameter
---

# gigavuecore_get_correlated_alarms_count_by_parameter Data Source

Get Counts of Correlated Alarms Grouped By The Requested Alarm Parameter

## Example Usage

```terraform
data "gigavuecore_get_correlated_alarms_count_by_parameter" "example" {
  acknowledged     = "example"
  acknowledgedby   = "example"
  alias            = "example"
  cluster_id       = "example"
  device_ip        = "example"
  end_time         = "example"
  group_by         = "example"
  hostname         = "example"
  resource_id      = "example"
  resource_type    = "example"
  severity         = "example"
  start_time       = "example"
  suppressed       = "example"
  type             = "example"
  unacknowledgedby = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `acknowledged` (String, optional) - filter by acknowledged alarms
* `acknowledgedby` (String, optional) - filter alarms that are acknowledged by a specific user
* `alias` (String, optional) - resource alias name to filter by
* `cluster_id` (String, optional) - Cluster ID to filter by
* `device_ip` (String, optional) - IP address of Device to filter by
* `end_time` (String, optional) - Optionally specify the end time to filter the records by. In ISO 8601 format 'yyyy-MM-ddTHH:mm:ssZ'
* `group_by` (String, required) - aggregation grouping category
* `hostname` (String, optional) - hostname to filter by
* `resource_id` (String, optional) - Affected entity to filter by
* `resource_type` (String, optional) - Affected entity type to filter by
* `severity` (String, optional) - Alarm severity to filter by
* `start_time` (String, optional) - Optionally specify the start time to filter the records by. In ISO 8601 format 'yyyy-MM-ddTHH:mm:ssZ'
* `suppressed` (String, optional) - filter by suppressed alarms
* `type` (String, optional) - Alarm Type to filter by
* `unacknowledgedby` (String, optional) - filter alarms that are unacknowledged by a specific user

### Attributes

In addition to all arguments above, the following attributes are exported:

* `agg_groups` (Attributes List, computed) - list of aggregate counts per groupBy type (see [below for nested schema](#nestedatt--agg_groups))

<a id="nestedatt--agg_groups"></a>
### Nested Schema for `agg_groups`

Read-Only:

* `count_` (Number) - the number of alarms in the aggregation group
* `group_name` (String) - Name of the group. For grouping based on severity, this field will have a value from the set \['Clear', 'Warn', 'Minor', 'Major', 'Critical'\]

