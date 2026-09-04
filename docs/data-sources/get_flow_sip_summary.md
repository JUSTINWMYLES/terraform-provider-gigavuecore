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
  alias             = "example"
  caller_id_pattern = "example"
  cluster_id        = "example"
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
* `rtp_resource_summary` (Attributes, computed) (see [below for nested schema](#nestedatt--rtp_resource_summary))
* `rtp_sessions` (Number, computed)
* `sip_messages_stats` (Attributes List, computed) (see [below for nested schema](#nestedatt--sip_messages_stats))
* `sip_resource_summary` (Attributes, computed) (see [below for nested schema](#nestedatt--sip_resource_summary))
* `sip_sessions` (Number, computed)

<a id="nestedatt--rtp_resource_summary"></a>
### Nested Schema for `rtp_resource_summary`

Read-Only:

* `data_pools_avail` (Number)
* `data_pools_in_use` (Number)

<a id="nestedatt--sip_messages_stats"></a>
### Nested Schema for `sip_messages_stats`

Read-Only:

* `drop` (Number)
* `no_match` (Number)
* `no_rule` (Number)
* `no_session` (Number)
* `other` (Number)
* `sip_message` (String)
* `tool_pass` (Number)

<a id="nestedatt--sip_resource_summary"></a>
### Nested Schema for `sip_resource_summary`

Read-Only:

* `num_sessions` (Number)
* `session_avail` (Number)
* `sip_parse_errors` (Number)

