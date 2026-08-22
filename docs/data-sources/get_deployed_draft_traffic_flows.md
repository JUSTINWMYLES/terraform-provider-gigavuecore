---
page_title: "gigavuecore_get_deployed_draft_traffic_flows Data Source - gigavuecore"
subcategory: ""
description: |-
  Get deployed draft Traffic Flows by alias
---

# gigavuecore_get_deployed_draft_traffic_flows Data Source

Get deployed draft Traffic Flows by alias

## Example Usage

```terraform
data "gigavuecore_get_deployed_draft_traffic_flows" "example" {
  alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - unique map alias

### Attributes

In addition to all arguments above, the following attributes are exported:

* `comment` (String, computed)
* `config_status` (String, computed) - Configuration status of this Traffic flow.
* `created_by` (String, computed) - created by of the traffic flow
* `created_time` (Number, computed) - created name of the traffic flow
* `deployed_by` (String, computed) - deployed by of the traffic flow
* `deployed_time` (Number, computed) - deployed name of the traffic flow
* `deployment_type` (String, computed) - 'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'transitLevel' is from vport to vport
* `enable` (Bool, computed) - enable/disable map, applicable only to first level maps
* `error_message` (String, computed) - In case of configuration failure, this message provides details about the possible cause of the failure.
* `flows` (List(Dynamic), computed)
* `has_draft` (Bool, computed)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `policy_id` (String, computed) - unique id
* `sources_and_rules` (List(Object({alias, components, health_state, health_state_reasons, inline_traffic_path, inline_traffic_type, ip_rewrite, rewrite, rule_matching, rule_type, rules, tags, traffic_type, vlan_tag})), computed)
* `tags` (List(Object({tag_key, tag_values})), computed)
* `updated_by` (String, computed) - updated by of the traffic flow
* `updated_time` (Number, computed) - updated time of the traffic flow

