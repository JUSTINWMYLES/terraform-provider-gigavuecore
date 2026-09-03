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
  cluster_id     = "example"
  gs_group_alias = "example"
  since          = "example"
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

* `control_only_session` (Attributes, computed) - number of sessions without user bearers (see [below for nested schema](#nestedatt--control_only_session))
* `control_tunnels` (Attributes, computed) - total number of control tunnels (see [below for nested schema](#nestedatt--control_tunnels))
* `control_user_tunnels` (Attributes, computed) - total number of control and user tunnels (see [below for nested schema](#nestedatt--control_user_tunnels))
* `end_time` (String, computed) - Timestamp of the latest record of data
* `error_msg` (String, computed) - Description of error regarding reset/ Maintenance
* `gsgroup` (String, computed) - alias of gsgroup
* `gtp_correlation_statistics` (Attributes List, computed) - control messages and user data message counters (see [below for nested schema](#nestedatt--gtp_correlation_statistics))
* `gtp_interface_statistics` (Attributes List, computed) - number of sessions and tunnels by interface (see [below for nested schema](#nestedatt--gtp_interface_statistics))
* `gtp_session_statistics` (Attributes List, computed) (see [below for nested schema](#nestedatt--gtp_session_statistics))
* `pending_session` (Attributes, computed) - number of sessions waiting for control message response (see [below for nested schema](#nestedatt--pending_session))
* `reset_happened` (Boolean, computed) - flag to notify reset/ Maintenance happened if true
* `start_time` (String, computed) - Timestamp of the earliest record of data

<a id="nestedatt--control_only_session"></a>
### Nested Schema for `control_only_session`

Read-Only:

* `max` (Number) - This is the latest record of data available
* `min` (Number) - This is the earliest record of data available

<a id="nestedatt--control_tunnels"></a>
### Nested Schema for `control_tunnels`

Read-Only:

* `max` (Number) - This is the latest record of data available
* `min` (Number) - This is the earliest record of data available

<a id="nestedatt--control_user_tunnels"></a>
### Nested Schema for `control_user_tunnels`

Read-Only:

* `max` (Number) - This is the latest record of data available
* `min` (Number) - This is the earliest record of data available

<a id="nestedatt--gtp_correlation_statistics"></a>
### Nested Schema for `gtp_correlation_statistics`

Read-Only:

* `gtp_control_message_stats` (Attributes) - Statistics for Control Messages GTP-C (see [below for nested schema](#nestedatt--gtp_correlation_statistics--gtp_control_message_stats))
* `gtp_user_message_stats` (Attributes List) (see [below for nested schema](#nestedatt--gtp_correlation_statistics--gtp_user_message_stats))

<a id="nestedatt--gtp_correlation_statistics--gtp_control_message_stats"></a>
### Nested Schema for `gtp_correlation_statistics.gtp_control_message_stats`

Read-Only:

* `gtp_c_stats` (Attributes List) (see [below for nested schema](#nestedatt--gtp_correlation_statistics--gtp_control_message_stats--gtp_c_stats))
* `gtp_version` (String)

<a id="nestedatt--gtp_correlation_statistics--gtp_control_message_stats--gtp_c_stats"></a>
### Nested Schema for `gtp_correlation_statistics.gtp_control_message_stats.gtp_c_stats`

Read-Only:

* `col_no_rule` (Attributes) (see [below for nested schema](#nestedatt--gtp_correlation_statistics--gtp_control_message_stats--gtp_c_stats--col_no_rule))
* `col_no_session` (Attributes) (see [below for nested schema](#nestedatt--gtp_correlation_statistics--gtp_control_message_stats--gtp_c_stats--col_no_session))
* `col_no_tnlx` (Attributes) (see [below for nested schema](#nestedatt--gtp_correlation_statistics--gtp_control_message_stats--gtp_c_stats--col_no_tnlx))
* `col_other` (Attributes) (see [below for nested schema](#nestedatt--gtp_correlation_statistics--gtp_control_message_stats--gtp_c_stats--col_other))
* `col_parse_er` (Attributes) (see [below for nested schema](#nestedatt--gtp_correlation_statistics--gtp_control_message_stats--gtp_c_stats--col_parse_er))
* `control_message` (String)
* `tool_pass` (Attributes) (see [below for nested schema](#nestedatt--gtp_correlation_statistics--gtp_control_message_stats--gtp_c_stats--tool_pass))

<a id="nestedatt--gtp_correlation_statistics--gtp_control_message_stats--gtp_c_stats--col_no_rule"></a>
### Nested Schema for `gtp_correlation_statistics.gtp_control_message_stats.gtp_c_stats.col_no_rule`

Read-Only:

* `delta` (Number) - This is the differential data calculated for the selected duration
* `max` (Number) - This is the latest record of data available
* `min` (Number) - This is the earliest record of data available

<a id="nestedatt--gtp_correlation_statistics--gtp_control_message_stats--gtp_c_stats--col_no_session"></a>
### Nested Schema for `gtp_correlation_statistics.gtp_control_message_stats.gtp_c_stats.col_no_session`

Read-Only:

* `delta` (Number) - This is the differential data calculated for the selected duration
* `max` (Number) - This is the latest record of data available
* `min` (Number) - This is the earliest record of data available

<a id="nestedatt--gtp_correlation_statistics--gtp_control_message_stats--gtp_c_stats--col_no_tnlx"></a>
### Nested Schema for `gtp_correlation_statistics.gtp_control_message_stats.gtp_c_stats.col_no_tnlx`

Read-Only:

* `delta` (Number) - This is the differential data calculated for the selected duration
* `max` (Number) - This is the latest record of data available
* `min` (Number) - This is the earliest record of data available

<a id="nestedatt--gtp_correlation_statistics--gtp_control_message_stats--gtp_c_stats--col_other"></a>
### Nested Schema for `gtp_correlation_statistics.gtp_control_message_stats.gtp_c_stats.col_other`

Read-Only:

* `delta` (Number) - This is the differential data calculated for the selected duration
* `max` (Number) - This is the latest record of data available
* `min` (Number) - This is the earliest record of data available

<a id="nestedatt--gtp_correlation_statistics--gtp_control_message_stats--gtp_c_stats--col_parse_er"></a>
### Nested Schema for `gtp_correlation_statistics.gtp_control_message_stats.gtp_c_stats.col_parse_er`

Read-Only:

* `delta` (Number) - This is the differential data calculated for the selected duration
* `max` (Number) - This is the latest record of data available
* `min` (Number) - This is the earliest record of data available

<a id="nestedatt--gtp_correlation_statistics--gtp_control_message_stats--gtp_c_stats--tool_pass"></a>
### Nested Schema for `gtp_correlation_statistics.gtp_control_message_stats.gtp_c_stats.tool_pass`

Read-Only:

* `delta` (Number) - This is the differential data calculated for the selected duration
* `max` (Number) - This is the latest record of data available
* `min` (Number) - This is the earliest record of data available

<a id="nestedatt--gtp_correlation_statistics--gtp_user_message_stats"></a>
### Nested Schema for `gtp_correlation_statistics.gtp_user_message_stats`

Read-Only:

* `collector` (Attributes) (see [below for nested schema](#nestedatt--gtp_correlation_statistics--gtp_user_message_stats--collector))
* `drop` (Attributes) (see [below for nested schema](#nestedatt--gtp_correlation_statistics--gtp_user_message_stats--drop))
* `tool_pass` (Attributes) (see [below for nested schema](#nestedatt--gtp_correlation_statistics--gtp_user_message_stats--tool_pass))

<a id="nestedatt--gtp_correlation_statistics--gtp_user_message_stats--collector"></a>
### Nested Schema for `gtp_correlation_statistics.gtp_user_message_stats.collector`

Read-Only:

* `delta` (Number) - This is the differential data calculated for the selected duration
* `max` (Number) - This is the latest record of data available
* `min` (Number) - This is the earliest record of data available

<a id="nestedatt--gtp_correlation_statistics--gtp_user_message_stats--drop"></a>
### Nested Schema for `gtp_correlation_statistics.gtp_user_message_stats.drop`

Read-Only:

* `delta` (Number) - This is the differential data calculated for the selected duration
* `max` (Number) - This is the latest record of data available
* `min` (Number) - This is the earliest record of data available

<a id="nestedatt--gtp_correlation_statistics--gtp_user_message_stats--tool_pass"></a>
### Nested Schema for `gtp_correlation_statistics.gtp_user_message_stats.tool_pass`

Read-Only:

* `delta` (Number) - This is the differential data calculated for the selected duration
* `max` (Number) - This is the latest record of data available
* `min` (Number) - This is the earliest record of data available

<a id="nestedatt--gtp_interface_statistics"></a>
### Nested Schema for `gtp_interface_statistics`

Read-Only:

* `dropped_bytes` (Attributes) (see [below for nested schema](#nestedatt--gtp_interface_statistics--dropped_bytes))
* `dropped_pkts` (Attributes) (see [below for nested schema](#nestedatt--gtp_interface_statistics--dropped_pkts))
* `gtp_c_packets` (Attributes) (see [below for nested schema](#nestedatt--gtp_interface_statistics--gtp_c_packets))
* `interface_type` (String) - interface type
* `rx_bytes` (Attributes) (see [below for nested schema](#nestedatt--gtp_interface_statistics--rx_bytes))
* `rx_pkts` (Attributes) (see [below for nested schema](#nestedatt--gtp_interface_statistics--rx_pkts))
* `tx_bytes` (Attributes) (see [below for nested schema](#nestedatt--gtp_interface_statistics--tx_bytes))
* `tx_pkts` (Attributes) (see [below for nested schema](#nestedatt--gtp_interface_statistics--tx_pkts))

<a id="nestedatt--gtp_interface_statistics--dropped_bytes"></a>
### Nested Schema for `gtp_interface_statistics.dropped_bytes`

Read-Only:

* `delta` (Number) - This is the differential data calculated for the selected duration
* `max` (Number) - This is the latest record of data available
* `min` (Number) - This is the earliest record of data available

<a id="nestedatt--gtp_interface_statistics--dropped_pkts"></a>
### Nested Schema for `gtp_interface_statistics.dropped_pkts`

Read-Only:

* `delta` (Number) - This is the differential data calculated for the selected duration
* `max` (Number) - This is the latest record of data available
* `min` (Number) - This is the earliest record of data available

<a id="nestedatt--gtp_interface_statistics--gtp_c_packets"></a>
### Nested Schema for `gtp_interface_statistics.gtp_c_packets`

Read-Only:

* `delta` (Number) - This is the differential data calculated for the selected duration
* `max` (Number) - This is the latest record of data available
* `min` (Number) - This is the earliest record of data available

<a id="nestedatt--gtp_interface_statistics--rx_bytes"></a>
### Nested Schema for `gtp_interface_statistics.rx_bytes`

Read-Only:

* `delta` (Number) - This is the differential data calculated for the selected duration
* `max` (Number) - This is the latest record of data available
* `min` (Number) - This is the earliest record of data available

<a id="nestedatt--gtp_interface_statistics--rx_pkts"></a>
### Nested Schema for `gtp_interface_statistics.rx_pkts`

Read-Only:

* `delta` (Number) - This is the differential data calculated for the selected duration
* `max` (Number) - This is the latest record of data available
* `min` (Number) - This is the earliest record of data available

<a id="nestedatt--gtp_interface_statistics--tx_bytes"></a>
### Nested Schema for `gtp_interface_statistics.tx_bytes`

Read-Only:

* `delta` (Number) - This is the differential data calculated for the selected duration
* `max` (Number) - This is the latest record of data available
* `min` (Number) - This is the earliest record of data available

<a id="nestedatt--gtp_interface_statistics--tx_pkts"></a>
### Nested Schema for `gtp_interface_statistics.tx_pkts`

Read-Only:

* `delta` (Number) - This is the differential data calculated for the selected duration
* `max` (Number) - This is the latest record of data available
* `min` (Number) - This is the earliest record of data available

<a id="nestedatt--gtp_session_statistics"></a>
### Nested Schema for `gtp_session_statistics`

Read-Only:

* `interface_type` (String) - interface type
* `sessions` (Attributes) (see [below for nested schema](#nestedatt--gtp_session_statistics--sessions))
* `tunnels` (Attributes) (see [below for nested schema](#nestedatt--gtp_session_statistics--tunnels))

<a id="nestedatt--gtp_session_statistics--sessions"></a>
### Nested Schema for `gtp_session_statistics.sessions`

Read-Only:

* `max` (Number) - This is the latest record of data available
* `min` (Number) - This is the earliest record of data available

<a id="nestedatt--gtp_session_statistics--tunnels"></a>
### Nested Schema for `gtp_session_statistics.tunnels`

Read-Only:

* `max` (Number) - This is the latest record of data available
* `min` (Number) - This is the earliest record of data available

<a id="nestedatt--pending_session"></a>
### Nested Schema for `pending_session`

Read-Only:

* `max` (Number) - This is the latest record of data available
* `min` (Number) - This is the earliest record of data available

