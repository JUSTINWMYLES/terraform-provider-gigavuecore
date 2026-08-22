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
  alias = null
  cluster_id = null
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
* `s6_a_messages_stats` (List(Object({drop, no_match, no_rule, no_session, other, s6_a_message, tool_pass})), computed)
* `s6_a_resource_summary` (Object({num_sessions, session_avail}), computed)
  * `num_sessions` (Number, computed)
  * `session_avail` (Number, computed)
* `s6_a_sessions` (Number, computed)

