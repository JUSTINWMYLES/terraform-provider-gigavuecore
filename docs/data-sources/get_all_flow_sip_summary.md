---
page_title: "gigavuecore_get_all_flow_sip_summary Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Flow Sip Report Summary
---

# gigavuecore_get_all_flow_sip_summary Data Source

Load all Flow Sip Report Summary

## Example Usage

```terraform
data "gigavuecore_get_all_flow_sip_summary" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `gsgroup` (String) - alias of gsgroup
* `rtp_resource_summary` (Attributes) (see [below for nested schema](#nestedatt--items--rtp_resource_summary))
* `rtp_sessions` (Number)
* `sip_messages_stats` (Attributes List) (see [below for nested schema](#nestedatt--items--sip_messages_stats))
* `sip_resource_summary` (Attributes) (see [below for nested schema](#nestedatt--items--sip_resource_summary))
* `sip_sessions` (Number)
<a id="nestedatt--items--rtp_resource_summary"></a>
### Nested Schema for `items.rtp_resource_summary`

Read-Only:

* `data_pools_avail` (Number)
* `data_pools_in_use` (Number)
<a id="nestedatt--items--sip_messages_stats"></a>
### Nested Schema for `items.sip_messages_stats`

Read-Only:

* `drop` (Number)
* `no_match` (Number)
* `no_rule` (Number)
* `no_session` (Number)
* `other` (Number)
* `sip_message` (String)
* `tool_pass` (Number)
<a id="nestedatt--items--sip_resource_summary"></a>
### Nested Schema for `items.sip_resource_summary`

Read-Only:

* `num_sessions` (Number)
* `session_avail` (Number)
* `sip_parse_errors` (Number)

