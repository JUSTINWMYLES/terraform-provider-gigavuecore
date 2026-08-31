---
page_title: "gigavuecore_query_port_group_lb_time_series Data Source - gigavuecore"
subcategory: ""
description: |-
  PortGroup Load Balancing Time Series Query interface
---

# gigavuecore_query_port_group_lb_time_series Data Source

PortGroup Load Balancing Time Series Query interface

## Example Usage

```terraform
data "gigavuecore_query_port_group_lb_time_series" "example" {
  cluster    = "example"
  end_time   = "example"
  gs_group   = "example"
  max_points = 0
  max_series = 0
  metric     = "pgLb.octets"
  port_group = "example"
  since      = "example"
  sort_by    = "example"
  start_time = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster` (String, optional) - Cluster id to filter by. Value of '\*' indicates a 'group-by-clusterId' request. Multiple values separated by '\|' can be specified, in which case time series are filtered and grouped by the requested clusters
* `end_time` (String, optional) - End Time in ISO 8601 format
* `gs_group` (String, optional) - GsGroup alias to filter by. Value of '\*' indicates a 'group-by-GsGroup' request. Multiple values separated by '\|' can be specified, in which case time series are filtered and grouped by the requested GsGroup
* `max_points` (Number, optional) - Number of the data points to return for the requested time period. Useful for building charts. Default value of 0 will return every datapoint, which could be costly for extensive time periods and/or high data sampling rates
* `max_series` (Number, optional) - In conjunction with the 'sortBy', returns the first N elements of the sorted time series list. Useful for the 'top-N' queries
* `metric` (String, required) - Time Series metric type
* `port_group` (String, optional) - PortGroup alias to filter by. Value of '\*' indicates a 'group-by-PortGroup' request. Multiple values separated by '\|' can be specified, in which case time series are filtered and grouped by the requested PortGroup
* `since` (String, optional) - Time Series start time reference in the past relative to 'now'. In the {number}-{timeUnit} format, where {timeUnit} is one of: \['minute', 'hour', 'day', 'week', 'month'\]. (Ex: '3-hour'). Mutually exclusive with and takes Precedence over the 'startTime' and 'endTime' attributes
* `sort_by` (String, optional) - If specified, each resulted timeseries over the specified time period is averaged to a single value, and then sorted in the requested order. Useful for the 'top-N' queries
* `start_time` (String, optional) - Start Time in ISO 8601 format

### Attributes

In addition to all arguments above, the following attributes are exported:

* `series` (Attributes List, computed) - time series (see [below for nested schema](#nestedatt--series))
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

