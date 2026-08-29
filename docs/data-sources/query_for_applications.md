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
  end_time                          = "example"
  filter_field                      = "example"
  group_by_clause                   = "example"
  metric                            = "example"
  monitor_tags                      = "example"
  order_by_desc                     = true
  rate                              = true
  solution_created_timestamp        = "example"
  solution_monitor_interval_in_mins = "example"
  start_time                        = "example"
  tags                              = "example"
  time                              = "example"
  time_unit                         = "example"
  top                               = "example"
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
* `order_by_desc` (Boolean, optional) - Order results in descending order
* `rate` (Boolean, optional) - Return results as rate (true) or absolute values (false)
* `solution_created_timestamp` (String, optional) - Solution creation timestamp
* `solution_monitor_interval_in_mins` (String, optional) - Solution monitor interval in minutes
* `start_time` (String, optional) - Start time for the query (ISO 8601 or epoch)
* `tags` (String, optional) - Tags to filter the query
* `time` (String, optional) - Time window for the query
* `time_unit` (String, optional) - Time unit for the query (e.g., minutes, hours)
* `top` (String, optional) - Number of top results to return

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `time_series` (Attributes) - Fabric Map Time Series Group (see [below for nested schema](#nestedatt--items--time_series))
<a id="nestedatt--items--time_series"></a>
### Nested Schema for `items.time_series`

Read-Only:

* `end_time` (String) - End Time in ISO 8601 format
* `metric` (String) - Time Series metric type
* `series` (Attributes List) - time series (see [below for nested schema](#nestedatt--items--time_series--series))
* `start_time` (String) - Start Time in ISO 8601 format
* `type` (String) - Time Series domain type
<a id="nestedatt--items--time_series--series"></a>
### Nested Schema for `items.time_series.series`

Read-Only:

* `context` (Attributes List) - Time Series query filtering values (see [below for nested schema](#nestedatt--items--time_series--series--context))
* `data_points` (Attributes List) - Time Series timestamped data points (see [below for nested schema](#nestedatt--items--time_series--series--data_points))
<a id="nestedatt--items--time_series--series--context"></a>
### Nested Schema for `items.time_series.series.context`

Read-Only:

* `name` (String) - Time-series filtering/grouping tag
* `value` (String) - tag value
<a id="nestedatt--items--time_series--series--data_points"></a>
### Nested Schema for `items.time_series.series.data_points`

Read-Only:

* `ts_utc` (Number) - Timestamp in UTC milliseconds
* `value` (Number) - time series value

