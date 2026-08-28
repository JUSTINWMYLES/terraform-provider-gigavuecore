---
page_title: "gigavuecore_get_flow_diameter_report_s6_a_summary Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Flow Diameter S6a Report Summary
---

# gigavuecore_get_flow_diameter_report_s6_a_summary Data Source

Load Flow Diameter S6a Report Summary

## Example Usage

```terraform
data "gigavuecore_get_flow_diameter_report_s6_a_summary" "example" {
  alias             = null
  cluster_id        = null
  user_name_pattern = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `user_name_pattern` (String, optional) - username pattern based active flows

### Attributes

In addition to all arguments above, the following attributes are exported:

* `gsgroup` (String, computed) - alias of gsgroup
* `s6_a_messages_stats` (Attributes List, computed) (see [below for nested schema](#nestedatt--s6_a_messages_stats))
* `s6_a_resource_summary` (Attributes, computed) (see [below for nested schema](#nestedatt--s6_a_resource_summary))
* `s6_a_sessions` (Number, computed)

<a id="nestedatt--s6_a_messages_stats"></a>
### Nested Schema for `s6_a_messages_stats`

Read-Only:

* `drop` (Number)
* `no_match` (Number)
* `no_rule` (Number)
* `no_session` (Number)
* `other` (Number)
* `s6_a_message` (String)
* `tool_pass` (Number)
<a id="nestedatt--s6_a_resource_summary"></a>
### Nested Schema for `s6_a_resource_summary`

Read-Only:

* `num_sessions` (Number)
* `session_avail` (Number)

