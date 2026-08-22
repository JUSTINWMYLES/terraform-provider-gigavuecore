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
  end_time = null
  map = null
  max_points = null
  max_series = null
  metric = null
  rate = null
  since = null
  sort_by = null
  start_time = null
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
* `rate` (Bool, optional) - Return results as rate (true) or absolute values (false)
* `since` (String, optional) - Relative time window (e.g., 1h, 24h) to query from
* `sort_by` (String, optional) - Field to sort the results by
* `start_time` (String, optional) - Start Time in ISO 8601 format

### Attributes

In addition to all arguments above, the following attributes are exported:

* `end_time` (String, computed) - End Time in ISO 8601 format
* `metric` (String, computed) - Time Series metric type
* `series` (List(Object({context, data_points})), computed) - time series
* `start_time` (String, computed) - Start Time in ISO 8601 format
* `type` (String, computed) - Time Series domain type

