---
page_title: "gigavuecore_get_flow_diameter_reports_s6_a_summary Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Flow Diameter S6a Report Summary
---

# gigavuecore_get_flow_diameter_reports_s6_a_summary Data Source

Load Flow Diameter S6a Report Summary

## Example Usage

```terraform
data "gigavuecore_get_flow_diameter_reports_s6_a_summary" "example" {
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
* `s6_a_messages_stats` (Attributes List) (see [below for nested schema](#nestedatt--items--s6_a_messages_stats))
* `s6_a_resource_summary` (Attributes) (see [below for nested schema](#nestedatt--items--s6_a_resource_summary))
* `s6_a_sessions` (Number)

<a id="nestedatt--items--s6_a_messages_stats"></a>
### Nested Schema for `items.s6_a_messages_stats`

Read-Only:

* `drop` (Number)
* `no_match` (Number)
* `no_rule` (Number)
* `no_session` (Number)
* `other` (Number)
* `s6_a_message` (String)
* `tool_pass` (Number)

<a id="nestedatt--items--s6_a_resource_summary"></a>
### Nested Schema for `items.s6_a_resource_summary`

Read-Only:

* `num_sessions` (Number)
* `session_avail` (Number)

