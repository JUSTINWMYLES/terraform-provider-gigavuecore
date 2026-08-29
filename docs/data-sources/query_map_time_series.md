---
page_title: "gigavuecore_query_map_time_series Data Source - gigavuecore"
subcategory: ""
description: |-
  Physical Map Time Series Query interface
---

# gigavuecore_query_map_time_series Data Source

Physical Map Time Series Query interface

## Example Usage

```terraform
data "gigavuecore_query_map_time_series" "example" {
  cluster    = "example"
  end_time   = "example"
  map        = "example"
  max_points = 0
  max_series = 0
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

* `cluster` (String, required) - Cluster id to filter by. Value of '\*' indicates a 'group-by-clusterId' request. Multiple values separated by '\|' can be specified, in which case time series are filtered and grouped by the requested clusters
* `end_time` (String, optional) - End Time in ISO 8601 format
* `map` (String, required) - Map alias to filter by. Value of '\*' indicates a 'group-by-map' request. Multiple values separated by '\|' can be specified, in which case time series are filtered and grouped by the requested map
* `max_points` (Number, optional) - Number of the data points to return for the requested time period. Useful for building charts. Default value of 0 will return every datapoint, which could be costly for extensive time periods and/or high data sampling rates
* `max_series` (Number, optional) - In Conjunction with the 'sortBy', returns the first N elements of the sorted time series list. Useful for the 'top-N' queries
* `metric` (String, required) - Time Series metric type
* `rate` (Boolean, optional) - Controls whether rate calculation should be done on the stored raw data before results are returned
* `since` (String, optional) - Time Series start time reference in the past relative to 'now'. In the {number}-{timeUnit} format, where {timeUnit} is one of: \['minute', 'hour', 'day', 'week', 'month'\]. (Ex: '3-hour'). Mutually exclusive with and takes Precedence over the 'startTime' and 'endTime' attributes
* `sort_by` (String, optional) - If specified, each resulted timeseries over the specified time period is averaged to a single value, and then sorted in the requested order. Useful for the 'top-N' queries
* `start_time` (String, optional) - Start Time in ISO 8601 format

### Attributes

In addition to all arguments above, the following attributes are exported:

* `end_time` (String, computed) - End Time in ISO 8601 format
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

