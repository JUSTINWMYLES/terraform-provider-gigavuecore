---
page_title: "gigavuecore_query_gs_engine_time_series Data Source - gigavuecore"
subcategory: ""
description: |-
  GS Engine Time Series Query interface
---

# gigavuecore_query_gs_engine_time_series Data Source

GS Engine Time Series Query interface

## Example Usage

```terraform
data "gigavuecore_query_gs_engine_time_series" "example" {
  box_port    = "example"
  cluster     = "example"
  end_time    = "example"
  max_points  = 0
  max_series  = 0
  metric      = "example"
  node        = "example"
  port        = "example"
  rate        = true
  since       = "example"
  slot        = "example"
  sort_by     = "example"
  start_time  = "example"
  traffic_dir = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_port` (String, optional) - BoxPort id to filter by. This is a {slotId}/{portId} part of the full cluster port id {boxId}/{slotId} (excludes the {boxId} prefix). Value of '\*' indicates a 'group-by-port' request. Multiple values separated by '\|' can be specified, in which case time series are filtered and grouped by the requested ports
* `cluster` (String, optional) - Cluster id to filter by. Value of '\*' indicates a 'group-by-clusterId' request. Multiple values separated by '\|' can be specified, in which case time series are filtered and grouped by the requested clusters
* `end_time` (String, optional) - End Time in ISO 8601 format
* `max_points` (Number, optional) - Number of the data points to return for the requested time period. Useful for building charts. Default value of 0 will return every datapoint, which could be costly for extensive time periods and/or high data sampling rates
* `max_series` (Number, optional) - In Conjunction with the 'sortBy', returns the first N elements of the sorted time series list. Useful for the 'top-N' queries
* `metric` (String, required) - Time Series metric type
* `node` (String, optional) - Node id to filter by. Value of '\*' indicates a 'group-by-nodeId' request. Multiple values separated by '\|' can be specified, in which case time series are filtered and grouped by the requested node
* `port` (String, optional) - Port id to filter by. This is a bare {portId} part of the full cluster port id {boxId}/{slotId}/{portId}. Value of '\*' indicates a 'group-by-port' request. Multiple values separated by '\|' can be specified, in which case time series are filtered and grouped by the requested ports
* `rate` (Boolean, optional) - Controls whether rate calculation should be done on the stored raw data before results are returned
* `since` (String, optional) - Time Series start time reference in the past relative to 'now'. In the {number}-{timeUnit} format, where {timeUnit} is one of: \['minute', 'hour', 'day', 'week', 'month'\]. (Ex: '3-hour'). Mutually exclusive with and takes Precedence over the 'startTime' and 'endTime' attributes
* `slot` (String, optional) - Slot id to filter by. This is a bare {slotId} part of the full cluster slot id {boxId}/{slotId}. Value of '\*' indicates a 'group-by-slot' request. Multiple values separated by '\|' can be specified, in which case time series are filtered and grouped by the requested slots
* `sort_by` (String, optional) - If specified, each resulted timeseries over the specified time period is averaged to a single value, and then sorted in the requested order. Useful for the 'top-N' queries
* `start_time` (String, optional) - Start Time in ISO 8601 format
* `traffic_dir` (String, optional) - Traffic direction to filter by. Value of '\*' indicates a 'group-by-traffic-direction' request

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

