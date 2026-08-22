---
page_title: "gigavuecore_query_for_applications Data Source - gigavuecore"
subcategory: ""
description: |-
  Query application visibility time series statistics
---

# gigavuecore_query_for_applications Data Source

Query application visibility time series statistics

## Example Usage

```terraform
data "gigavuecore_query_for_applications" "example" {
  end_time = null
  filter_field = null
  group_by_clause = null
  metric = null
  monitor_tags = null
  order_by_desc = null
  rate = null
  solution_created_timestamp = null
  solution_monitor_interval_in_mins = null
  start_time = null
  tags = null
  time = null
  time_unit = null
  top = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `end_time` (String, optional) - End time for the query (ISO 8601 or epoch)
* `filter_field` (String, optional) - Field to filter results by
* `group_by_clause` (String, optional) - Group by clause for aggregation
* `metric` (String, optional) - Metric to query (e.g., application bytes, flows)
* `monitor_tags` (String, optional) - Monitor tags for filtering
* `order_by_desc` (Bool, optional) - Order results in descending order
* `rate` (Bool, optional) - Return results as rate (true) or absolute values (false)
* `solution_created_timestamp` (String, optional) - Solution creation timestamp
* `solution_monitor_interval_in_mins` (String, optional) - Solution monitor interval in minutes
* `start_time` (String, optional) - Start time for the query (ISO 8601 or epoch)
* `tags` (String, optional) - Tags to filter the query
* `time` (String, optional) - Time window for the query
* `time_unit` (String, optional) - Time unit for the query (e.g., minutes, hours)
* `top` (String, optional) - Number of top results to return

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({time_series})), computed)

