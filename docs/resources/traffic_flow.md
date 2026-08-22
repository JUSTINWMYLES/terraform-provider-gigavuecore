---
page_title: "gigavuecore_traffic_flow Resource - gigavuecore"
subcategory: ""
description: |-
  Get Traffic Flows by alias
---

# gigavuecore_traffic_flow Resource

Get Traffic Flows by alias

## Example Usage

```terraform
resource "gigavuecore_traffic_flow" "example" {
  alias = null
  comment = null
  deployment_type = null
  enable = null
  flows = []
  has_draft = null
  priority_type = null
  sources_and_rules = []
  tags = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - unique map alias
* `comment` (String, optional)
* `deployment_type` (String, required) - 'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'transitLevel' is from vport to vport
* `enable` (Bool, optional) - enable/disable map, applicable only to first level maps
* `flows` (List(Dynamic), required)
* `has_draft` (Bool, optional)
* `priority_type` (String, optional) - Define the map priority to be LOWEST or HIGHEST. Default priority is LOWEST.
* `sources_and_rules` (List(Object({alias, components, health_state, health_state_reasons, inline_traffic_path, inline_traffic_type, ip_rewrite, rewrite, rule_matching, rule_type, rules, tags, traffic_type, vlan_tag})), required)
* `tags` (List(Object({tag_key, tag_values})), optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `comment` (String, computed)
* `config_status` (String, computed) - Configuration status of this Traffic flow.
* `created_by` (String, computed) - created by of the traffic flow
* `created_time` (Number, computed) - created name of the traffic flow
* `deployed_by` (String, computed) - deployed by of the traffic flow
* `deployed_time` (Number, computed) - deployed name of the traffic flow
* `enable` (Bool, computed) - enable/disable map, applicable only to first level maps
* `error_message` (String, computed) - In case of configuration failure, this message provides details about the possible cause of the failure.
* `has_draft` (Bool, computed)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `policy_id` (String, computed) - unique id
* `tags` (List(Object({tag_key, tag_values})), computed)
* `updated_by` (String, computed) - updated by of the traffic flow
* `updated_time` (Number, computed) - updated time of the traffic flow

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_traffic_flow.example {alias}
```
