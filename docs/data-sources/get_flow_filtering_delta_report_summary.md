---
page_title: "gigavuecore_get_flow_filtering_delta_report_summary Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Flow Filtering Delta Report Summary
---

# gigavuecore_get_flow_filtering_delta_report_summary Data Source

Load Flow Filtering Delta Report Summary

## Example Usage

```terraform
data "gigavuecore_get_flow_filtering_delta_report_summary" "example" {
  cluster_id = null
  gs_group_alias = null
  since = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `gs_group_alias` (String, required) - alias of the target GS Group
* `since` (String, required) - Start time reference in the past relative to 'now'. In the {number}-{timeUnit} format, where {timeUnit} is one of: \['minute', 'hour', 'day', 'week', 'month'\]. (Ex: '3-hour').

### Attributes

In addition to all arguments above, the following attributes are exported:

* `control_only_session` (Object({max, min}), computed)
  * `max` (Number, computed) - This is the latest record of data available
  * `min` (Number, computed) - This is the earliest record of data available
* `control_tunnels` (Object({max, min}), computed)
  * `max` (Number, computed) - This is the latest record of data available
  * `min` (Number, computed) - This is the earliest record of data available
* `control_user_tunnels` (Object({max, min}), computed)
  * `max` (Number, computed) - This is the latest record of data available
  * `min` (Number, computed) - This is the earliest record of data available
* `end_time` (String, computed) - Timestamp of the latest record of data
* `error_msg` (String, computed) - Description of error regarding reset/ Maintenance
* `gsgroup` (String, computed) - alias of gsgroup
* `gtp_correlation_statistics` (List(Object({gtp_control_message_stats, gtp_user_message_stats})), computed) - control messages and user data message counters
* `gtp_interface_statistics` (List(Object({dropped_bytes, dropped_pkts, gtp_c_packets, interface_type, rx_bytes, rx_pkts, tx_bytes, tx_pkts})), computed) - number of sessions and tunnels by interface
* `gtp_session_statistics` (List(Object({interface_type, sessions, tunnels})), computed)
* `pending_session` (Object({max, min}), computed)
  * `max` (Number, computed) - This is the latest record of data available
  * `min` (Number, computed) - This is the earliest record of data available
* `reset_happened` (Bool, computed) - flag to notify reset/ Maintenance happened if true
* `start_time` (String, computed) - Timestamp of the earliest record of data

