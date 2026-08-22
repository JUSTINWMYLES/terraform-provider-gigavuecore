---
page_title: "gigavuecore_intent_policy Resource - gigavuecore"
subcategory: ""
description: |-
  Get All policies
---

# gigavuecore_intent_policy Resource

Get All policies

## Example Usage

```terraform
resource "gigavuecore_intent_policy" "example" {
  comment = null
  deployed = null
  deployment_error = null
  deployment_percent = null
  dest_port_timestamp = null
  health_state = null
  health_state_reasons = []
  name = null
  policy_id = null
  policy_timestamp = null
  priority = null
  rules = []
  src_port_timestamp = null
  src_ports_info = {}
  tags = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `comment` (String, optional)
* `deployed` (Bool, optional) - policy is deployed or not
* `deployment_error` (String, optional) - policy deployment error message
* `deployment_percent` (String, optional) - policy deployment percentage
* `dest_port_timestamp` (String, optional)
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), optional)
* `name` (String, required) - policy name
* `policy_id` (String, optional) - generated unique policy ID
* `policy_timestamp` (String, optional)
* `priority` (Bool, optional)
* `rules` (List(Object({comment, gs_operations, health_state, health_state_reasons, high_priority_drop, matches, no_expansion_tags, rule_id, rule_name, tools, type})), optional)
* `src_port_timestamp` (String, optional)
* `src_ports_info` (Object({comment, ports, template_ids}), optional)
  * `comment` (String, optional)
  * `ports` (List(String), optional)
  * `template_ids` (List(String), optional)
* `tags` (List(Object({tag_key, tag_values})), optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `id` (String, computed)
* `policies` (List(Object({comment, deployed, deployment_error, deployment_percent, dest_port_timestamp, health_state, health_state_reasons, name, policy_id, policy_timestamp, priority, rules, src_port_timestamp, src_ports_info, tags})), computed)

