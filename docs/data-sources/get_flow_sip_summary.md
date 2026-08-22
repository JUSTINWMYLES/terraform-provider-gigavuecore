---
page_title: "gigavuecore_get_flow_sip_summary Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Flow Sip Report Summary
---

# gigavuecore_get_flow_sip_summary Data Source

Load Flow Sip Report Summary

## Example Usage

```terraform
data "gigavuecore_get_flow_sip_summary" "example" {
  alias = null
  caller_id_pattern = null
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `caller_id_pattern` (String, optional) - callerid pattern based active flows
* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `gsgroup` (String, computed) - alias of gsgroup
* `rtp_resource_summary` (Object({data_pools_avail, data_pools_in_use}), computed)
  * `data_pools_avail` (Number, computed)
  * `data_pools_in_use` (Number, computed)
* `rtp_sessions` (Number, computed)
* `sip_messages_stats` (List(Object({drop, no_match, no_rule, no_session, other, sip_message, tool_pass})), computed)
* `sip_resource_summary` (Object({num_sessions, session_avail, sip_parse_errors}), computed)
  * `num_sessions` (Number, computed)
  * `session_avail` (Number, computed)
  * `sip_parse_errors` (Number, computed)
* `sip_sessions` (Number, computed)

