---
page_title: "gigavuecore_query_raw_data Function - gigavuecore"
subcategory: ""
description: |-
  Query Raw Data
---

# gigavuecore_query_raw_data Function

Query Raw Data

## Example Usage

```terraform
# Example: provider::gigavuecore::gigavuecore_query_raw_data("<since>", "<start_time>", "<end_time>", "<index_pattern>", "<fm_tags>", "<group_by>", "<tags>", "<size>", true, "<response_scroll_timeout>", "<sort_by>", "<query_objects>", "<agg_field>", "<bucket_interval>", true, true, "<date_time_stamp_field>", true, "<exclude_fields>", "<sort_by_field>", true, "<aggregator>", "<metrics>", true, true, true, "<result_aggregator>", "<exclude_agg_from_response>")
output "example" {
  value = provider::gigavuecore::gigavuecore_query_raw_data("<since>", "<start_time>", "<end_time>", "<index_pattern>", "<fm_tags>", "<group_by>", "<tags>", "<size>", true, "<response_scroll_timeout>", "<sort_by>", "<query_objects>", "<agg_field>", "<bucket_interval>", true, true, "<date_time_stamp_field>", true, "<exclude_fields>", "<sort_by_field>", true, "<aggregator>", "<metrics>", true, true, true, "<result_aggregator>", "<exclude_agg_from_response>")
}
```

## Signature

```text
gigavuecore_query_raw_data(since: String, start_time: String, end_time: String, index_pattern: String, fm_tags: String, group_by: String, tags: String, size: String, scroll: Boolean, response_scroll_timeout: String, sort_by: String, query_objects: String, agg_field: String, bucket_interval: String, trending: Boolean, rate: Boolean, date_time_stamp_field: String, exclude: Boolean, exclude_fields: String, sort_by_field: String, inventory_index: Boolean, aggregator: String, metrics: String, skip_normalized: Boolean, skip_zero_values: Boolean, is_source_req_in_aggregation: Boolean, result_aggregator: String, exclude_agg_from_response: String) -> String
```

## Arguments

The following arguments are supported:

* `since` (String) - Start time reference in the past relative to 'now'. In the {number}-{timeUnit} format, where {timeUnit} is one of: \['minute', 'hour', 'day', 'week', 'month'\]. (Ex: '3-hour'). Mutually exclusive with and takes Precedence over the 'startTime' and 'endTime' attributes
* `start_time` (String) - Start time. In ISO 8601 format. Mutually exclusive with and overridden by the 'since' attribute
* `end_time` (String) - End time. In ISO 8601 format. Mutually exclusive with and overridden by the 'since' attribute
* `index_pattern` (String) - Index name of the module
* `fm_tags` (String) - Used to filter based on given input , example : {resourceTypeName=Port, type=TrapPacketDrop}. can have multiple values with comma separated and specified within {}
* `group_by` (String) - To Terms aggregate the data based on the fields we can use group by field. eg : tag.NodeType>portIdToClusterId. can have multiple values with greater than(>) separated
* `tags` (String) - Tags to filter the results. can have multiple values with comma separated and specified within {}. eg :{"Region":"USA","State":"NC","Site":"DC1"}
* `size` (String) - The maximum number of hits to be returned with each batch of results. Default is 1
* `scroll` (Boolean) - Scroll is required or not (true/false → default is false) i.e. to keep the search context open
* `response_scroll_timeout` (String) - Scroll Time out to mention how long it should keep the "search context" alive
* `sort_by` (String) - If specified, results sorted in the requested order.
* `query_objects` (String) - if index pattern is 'fmstats' then objects that can be queried based on specified values. Can have multiple values with comma separated
* `agg_field` (String) - On which field metric aggreagation(Min, Max, Sum, Avg and Count) to be applied. Mutually exclusive with metrics param
* `bucket_interval` (String) - When date histogram is applied what is the bucket interval to be chosen.
* `trending` (Boolean) - If enabled aggregation can be applied based on metric aggreagation(Min, Max, Sum, Avg and Count) input param "aggregator"
* `rate` (Boolean) - Whether to return derivative data points or actual data points if enabled. Controls whether rate calculation should be done on the stored raw data before results are returned
* `date_time_stamp_field` (String) - TimeStamp field name for the OpenSearch index
* `exclude` (Boolean) - Exclude the fields specified in param 'excludeFields' if enabled
* `exclude_fields` (String) - If the param 'exclude' is true, we can specify the fields in this param to be excluded from result.
* `sort_by_field` (String) - If specified, then result will be sorted by the field name supplied
* `inventory_index` (Boolean) - This flag will be true for a non time-series indices
* `aggregator` (String) - Aggregator is mandatory when trending is true
* `metrics` (String) - Applicable for stats index pattern (eg : port.octets,port.packets) can have multiple values with comma separated
* `skip_normalized` (Boolean) - Whether to consider the derivative value with unit (eg: per second rate value) for further aggregations or result.
* `skip_zero_values` (Boolean) - When metric aggregations applied whether to consider resultant zero values for further aggregations.
* `is_source_req_in_aggregation` (Boolean) - Whether topHits sub aggregation data required in the response. When doing multi level terms aggregations, set this to false as the topHits may not be useful.
* `result_aggregator` (String) - When multiple terms aggregation applied what metric aggregation to be used at top level, for example when ports data grouped by a tag, port type and portId what is the result aggregation to be done at tag and type groupby level. Multiple values can be given with comma separated.
* `exclude_agg_from_response` (String) - When multiple terms aggregations (groupBy) applied, the response will have the nested levels. If any one value from groupBy is provided, the result will not have any nested aggregation data from the given term. Say for example, If groupBy param is given as tag.Site>PortType>PortIdToClusterId and excludeAggFromResponse is given as PortIdToClusterId then aggregation data of PortIdToClusterId and below level will not come in response. If trending is true, by providing dateHistogramAgg excludes date histogram buckets from response.


## Return

The function returns a value of type `String`.
