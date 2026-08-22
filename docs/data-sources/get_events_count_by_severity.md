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
  end_time = null
  group_by = null
  start_time = null
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

* `agg_groups` (List(Object({count_, group_name})), computed) - list of aggregate counts per groupBy type

