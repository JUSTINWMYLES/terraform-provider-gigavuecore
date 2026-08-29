---
page_title: "gigavuecore_get_events_count_by_severity Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Counts of Events Grouped By Severity
---

# gigavuecore_get_events_count_by_severity Data Source

Get Counts of Events Grouped By Severity

## Example Usage

```terraform
data "gigavuecore_get_events_count_by_severity" "example" {
  end_time   = "example"
  group_by   = "example"
  start_time = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `end_time` (String, optional) - Optionally specify the end time to filter the records by. In ISO 8601 format
* `group_by` (String, required) - aggregation grouping category
* `start_time` (String, optional) - Optionally specify the start time to filter the records by. In ISO 8601 format

### Attributes

In addition to all arguments above, the following attributes are exported:

* `agg_groups` (Attributes List, computed) - list of aggregate counts per groupBy type (see [below for nested schema](#nestedatt--agg_groups))

<a id="nestedatt--agg_groups"></a>
### Nested Schema for `agg_groups`

Read-Only:

* `count_` (Number) - the number of events in the aggregation group
* `group_name` (String) - Name of the group. For grouping based on severity, this field will have a value from the set \['info', 'minor', 'major', 'critical'\]

