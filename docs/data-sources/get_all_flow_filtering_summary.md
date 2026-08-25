---
page_title: "gigavuecore_get_all_flow_filtering_summary Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Flow Filtering Report Summary
---

# gigavuecore_get_all_flow_filtering_summary Data Source

Load all Flow Filtering Report Summary

## Example Usage

```terraform
data "gigavuecore_get_all_flow_filtering_summary" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `control_only_session` (Number) - number of sessions without user bearers
* `control_tunnels` (Number) - total number of control tunnels
* `control_user_tunnels` (Number) - total number of control and user tunnels
* `gsgroup` (String) - alias of gsgroup
* `gtp_corelation_statistics` (Attributes List) - control messages and user data message counters (see [below for nested schema](#nestedatt--items--gtp_corelation_statistics))
* `gtp_interface_statistics` (Attributes List) - number of sessions and tunnels by interface (see [below for nested schema](#nestedatt--items--gtp_interface_statistics))
* `gtp_pfcp_statistics` (Attributes) - PFCP message stats (see [below for nested schema](#nestedatt--items--gtp_pfcp_statistics))
* `gtp_session_statistics` (Attributes List) (see [below for nested schema](#nestedatt--items--gtp_session_statistics))
* `pending_session` (Number) - number of sessions waiting for control message response
<a id="nestedatt--items--gtp_corelation_statistics"></a>
### Nested Schema for `items.gtp_corelation_statistics`

Read-Only:

* `gtp_control_message_stats` (Attributes) - Statistics for Control Messages GTP-C (see [below for nested schema](#nestedatt--items--gtp_corelation_statistics--gtp_control_message_stats))
* `gtp_user_message_stats` (Attributes List) (see [below for nested schema](#nestedatt--items--gtp_corelation_statistics--gtp_user_message_stats))
<a id="nestedatt--items--gtp_corelation_statistics--gtp_control_message_stats"></a>
### Nested Schema for `items.gtp_corelation_statistics.gtp_control_message_stats`

Read-Only:

* `gtp_c_stats` (Attributes List) (see [below for nested schema](#nestedatt--items--gtp_corelation_statistics--gtp_control_message_stats--gtp_c_stats))
* `gtp_version` (String)
<a id="nestedatt--items--gtp_corelation_statistics--gtp_control_message_stats--gtp_c_stats"></a>
### Nested Schema for `items.gtp_corelation_statistics.gtp_control_message_stats.gtp_c_stats`

Read-Only:

* `col_no_rule` (Number)
* `col_no_session` (Number)
* `col_no_tnlx` (Number)
* `col_other` (Number)
* `col_parse_er` (Number)
* `control_message` (String)
* `tool_pass` (Number)
<a id="nestedatt--items--gtp_corelation_statistics--gtp_user_message_stats"></a>
### Nested Schema for `items.gtp_corelation_statistics.gtp_user_message_stats`

Read-Only:

* `collector` (Number)
* `drop` (Number)
* `tool_pass` (Number)
<a id="nestedatt--items--gtp_interface_statistics"></a>
### Nested Schema for `items.gtp_interface_statistics`

Read-Only:

* `dropped_bytes` (Number)
* `dropped_pkts` (Number)
* `gtp_c_packets` (Number)
* `interface_type` (String) - interface type
* `rx_bytes` (Number)
* `rx_pkts` (Number)
* `tx_bytes` (Number)
* `tx_pkts` (Number)
<a id="nestedatt--items--gtp_pfcp_statistics"></a>
### Nested Schema for `items.gtp_pfcp_statistics`

Read-Only:

* `asso_rel_req` (Number)
* `asso_rel_rsp` (Number)
* `asso_setup_req` (Number)
* `asso_setup_rsp` (Number)
* `asso_upd_req` (Number)
* `asso_upd_rsp` (Number)
* `bundle_msg` (Number)
* `heart_beat_req` (Number)
* `heart_beat_rsp` (Number)
* `node_rpt_req` (Number)
* `node_rpt_rsp` (Number)
* `pfd_mgmt_req` (Number)
* `pfd_mgmt_rsp` (Number)
* `res_msg_type16_to49` (Number)
* `res_msg_type58_to255` (Number)
* `sess_set_del_req` (Number)
* `sess_set_del_rsp` (Number)
* `ver_not_supp_rsp` (Number)
<a id="nestedatt--items--gtp_session_statistics"></a>
### Nested Schema for `items.gtp_session_statistics`

Read-Only:

* `interface_type` (String) - interface type
* `sessions` (Number)
* `tunnels` (Number)

