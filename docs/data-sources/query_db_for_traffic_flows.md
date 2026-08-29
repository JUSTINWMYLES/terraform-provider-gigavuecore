---
page_title: "gigavuecore_query_db_for_traffic_flows Data Source - gigavuecore"
subcategory: ""
description: |-
  Query time series traffic flow statistics
---

# gigavuecore_query_db_for_traffic_flows Data Source

Query time series traffic flow statistics

## Example Usage

```terraform
data "gigavuecore_query_db_for_traffic_flows" "example" {
  end_time   = "example"
  map        = "example"
  max_points = "example"
  max_series = "example"
  metric     = "example"
  rate       = true
  since      = "example"
  sort_by    = "example"
  start_time = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `end_time` (String, optional) - End Time in ISO 8601 format
* `map` (String, required) - Traffic flow alias (required)
* `max_points` (String, optional) - Maximum number of data points to return
* `max_series` (String, optional) - Maximum number of series to return
* `metric` (String, optional) - Time Series metric type
* `rate` (Boolean, optional) - Return results as rate (true) or absolute values (false)
* `since` (String, optional) - Relative time window (e.g., 1h, 24h) to query from
* `sort_by` (String, optional) - Field to sort the results by
* `start_time` (String, optional) - Start Time in ISO 8601 format

### Attributes

In addition to all arguments above, the following attributes are exported:

* `end_time` (String, computed) - End Time in ISO 8601 format
* `metric` (String, computed) - Time Series metric type
* `series` (Attributes List, computed) - time series (see [below for nested schema](#nestedatt--series))
* `start_time` (String, computed) - Start Time in ISO 8601 format
* `type` (String, computed) - Time Series domain type

<a id="nestedatt--series"></a>
### Nested Schema for `series`

Read-Only:

* `context` (Attributes List) - Time Series query filtering values (see [below for nested schema](#nestedatt--series--context))
* `data_points` (Attributes List) - Time Series timestamped data points (see [below for nested schema](#nestedatt--series--data_points))
<a id="nestedatt--series--context"></a>
### Nested Schema for `series.context`

Read-Only:

* `name` (String) - Time-series filtering/grouping tag
* `value` (String) - tag value
<a id="nestedatt--series--data_points"></a>
### Nested Schema for `series.data_points`

Read-Only:

* `ts_utc` (Number) - Timestamp in UTC milliseconds
* `value` (Number) - time series value

