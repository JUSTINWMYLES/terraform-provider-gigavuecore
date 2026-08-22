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
  cluster = null
  end_time = null
  gs_group = null
  max_points = null
  max_series = null
  metric = null
  port_group = null
  since = null
  sort_by = null
  start_time = null
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

* `end_time` (String, computed) - End Time in ISO 8601 format
* `series` (List(Object({context, data_points})), computed) - time series
* `start_time` (String, computed) - Start Time in ISO 8601 format
* `type` (String, computed) - Time Series domain type

