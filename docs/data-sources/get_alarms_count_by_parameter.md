---
page_title: "gigavuecore_get_alarms_count_by_parameter Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Counts of Alarms Grouped By The Requested Alarm Parameter
---

# gigavuecore_get_alarms_count_by_parameter Data Source

Get Counts of Alarms Grouped By The Requested Alarm Parameter

## Example Usage

```terraform
data "gigavuecore_get_alarms_count_by_parameter" "example" {
  acknowledged = null
  acknowledgedby = null
  alias = null
  cluster_id = null
  device_ip = null
  end_time = null
  group_by = null
  hostname = null
  resource_id = null
  resource_type = null
  severity = null
  start_time = null
  suppressed = null
  type = null
  unacknowledgedby = null
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

* `agg_groups` (List(Object({count_, group_name})), computed) - list of aggregate counts per groupBy type

