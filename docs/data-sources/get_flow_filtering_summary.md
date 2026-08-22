---
page_title: "gigavuecore_get_flow_filtering_summary Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Flow Filtering Report Summary
---

# gigavuecore_get_flow_filtering_summary Data Source

Load Flow Filtering Report Summary

## Example Usage

```terraform
data "gigavuecore_get_flow_filtering_summary" "example" {
  alias = null
  cluster_id = null
  gtp_imsi_pattern = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `gtp_imsi_pattern` (String, optional) - gtp imsi pattern based active flows

### Attributes

In addition to all arguments above, the following attributes are exported:

* `control_only_session` (Number, computed) - number of sessions without user bearers
* `control_tunnels` (Number, computed) - total number of control tunnels
* `control_user_tunnels` (Number, computed) - total number of control and user tunnels
* `gsgroup` (String, computed) - alias of gsgroup
* `gtp_corelation_statistics` (List(Object({gtp_control_message_stats, gtp_user_message_stats})), computed) - control messages and user data message counters
* `gtp_interface_statistics` (List(Object({dropped_bytes, dropped_pkts, gtp_c_packets, interface_type, rx_bytes, rx_pkts, tx_bytes, tx_pkts})), computed) - number of sessions and tunnels by interface
* `gtp_pfcp_statistics` (Object({asso_rel_req, asso_rel_rsp, asso_setup_req, asso_setup_rsp, asso_upd_req, asso_upd_rsp, bundle_msg, heart_beat_req, heart_beat_rsp, node_rpt_req, node_rpt_rsp, pfd_mgmt_req, pfd_mgmt_rsp, res_msg_type16_to49, res_msg_type58_to255, sess_set_del_req, sess_set_del_rsp, ver_not_supp_rsp}), computed) - PFCP message stats
  * `asso_rel_req` (Number, computed)
  * `asso_rel_rsp` (Number, computed)
  * `asso_setup_req` (Number, computed)
  * `asso_setup_rsp` (Number, computed)
  * `asso_upd_req` (Number, computed)
  * `asso_upd_rsp` (Number, computed)
  * `bundle_msg` (Number, computed)
  * `heart_beat_req` (Number, computed)
  * `heart_beat_rsp` (Number, computed)
  * `node_rpt_req` (Number, computed)
  * `node_rpt_rsp` (Number, computed)
  * `pfd_mgmt_req` (Number, computed)
  * `pfd_mgmt_rsp` (Number, computed)
  * `res_msg_type16_to49` (Number, computed)
  * `res_msg_type58_to255` (Number, computed)
  * `sess_set_del_req` (Number, computed)
  * `sess_set_del_rsp` (Number, computed)
  * `ver_not_supp_rsp` (Number, computed)
* `gtp_session_statistics` (List(Object({interface_type, sessions, tunnels})), computed)
* `pending_session` (Number, computed) - number of sessions waiting for control message response

