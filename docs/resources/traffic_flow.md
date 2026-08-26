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
  alias             = null
  comment           = null
  deployment_type   = null
  enable            = null
  flows             = []
  has_draft         = null
  priority_type     = null
  sources_and_rules = []
  tags              = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - unique map alias
* `comment` (String, optional)
* `deployment_type` (String, required) - 'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'transitLevel' is from vport to vport
* `enable` (Boolean, optional) - enable/disable map, applicable only to first level maps
* `flows` (List of Dynamic, required)
* `has_draft` (Boolean, optional)
* `priority_type` (String, optional) - Define the map priority to be LOWEST or HIGHEST. Default priority is LOWEST.
* `sources_and_rules` (Attributes List, required) (see [below for nested schema](#nestedatt--sources_and_rules))
* `tags` (Attributes List, optional) (see [below for nested schema](#nestedatt--tags))

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `comment` (String, computed)
* `config_status` (String, computed) - Configuration status of this Traffic flow.
* `created_by` (String, computed) - created by of the traffic flow
* `created_time` (Number, computed) - created name of the traffic flow
* `deployed_by` (String, computed) - deployed by of the traffic flow
* `deployed_time` (Number, computed) - deployed name of the traffic flow
* `enable` (Boolean, computed) - enable/disable map, applicable only to first level maps
* `error_message` (String, computed) - In case of configuration failure, this message provides details about the possible cause of the failure.
* `has_draft` (Boolean, computed)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `policy_id` (String, computed) - unique id
* `tags` (Attributes List, computed) (see [below for nested schema](#nestedatt--tags))
* `updated_by` (String, computed) - updated by of the traffic flow
* `updated_time` (Number, computed) - updated time of the traffic flow

<a id="nestedatt--sources_and_rules"></a>
### Nested Schema for `sources_and_rules`

Required:

* `alias` (String) - unique sourcesAndRules alias
* `rule_type` (String) - 'byRule' is applicable to all map types; 'collector' is applicable to 'regular', 'inline' maps; 'passAll' is applicable to 'regular' and 'inline' maps
* `rules` (Attributes) - Map Rules Container. Private class (see [below for nested schema](#nestedatt--sources_and_rules--rules))
Optional:

* `components` (Attributes List) (see [below for nested schema](#nestedatt--sources_and_rules--components))
* `inline_traffic_path` (String) - Only applicable for 'inline' map types, in which case defaults to 'normal'. When set to 'bypass', the 'dstPorts' must be empty
* `inline_traffic_type` (String) - Only applicable for 'inline/passAll' map types, in which case defaults to 'symmetric'. For 'asymmetric' maps, 'srcPort' must be of type 'inline-net'
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--sources_and_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--sources_and_rules--rewrite))
* `rule_matching` (String) - If set to 'blacklist', packet are passed when no 'drop' rules are matched. This field is only valid for 'regular/byRule' map types. (maps into 'no-rule-match' CLI command)
* `tags` (Attributes List) (see [below for nested schema](#nestedatt--sources_and_rules--tags))
* `traffic_type` (String) - Only applicable for 'firstLevel/byRule' map types, in which case defaults to 'user'
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--sources_and_rules--vlan_tag))
Read-Only:

* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--sources_and_rules--health_state_reasons))
<a id="nestedatt--sources_and_rules--rules"></a>
### Nested Schema for `sources_and_rules.rules`

Optional:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--sources_and_rules--rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--sources_and_rules--rules--pass_rules))
<a id="nestedatt--sources_and_rules--rules--drop_rules"></a>
### Nested Schema for `sources_and_rules.rules.drop_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)
Optional:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--sources_and_rules--rules--drop_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--sources_and_rules--rules--drop_rules--rewrite))
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--sources_and_rules--rules--drop_rules--vlan_tag))
<a id="nestedatt--sources_and_rules--rules--drop_rules--ip_rewrite"></a>
### Nested Schema for `sources_and_rules.rules.drop_rules.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--sources_and_rules--rules--drop_rules--rewrite"></a>
### Nested Schema for `sources_and_rules.rules.drop_rules.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--sources_and_rules--rules--drop_rules--vlan_tag"></a>
### Nested Schema for `sources_and_rules.rules.drop_rules.vlan_tag`

Required:

* `vlan_action` (String)
Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)
<a id="nestedatt--sources_and_rules--rules--pass_rules"></a>
### Nested Schema for `sources_and_rules.rules.pass_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)
Optional:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--sources_and_rules--rules--pass_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--sources_and_rules--rules--pass_rules--rewrite))
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--sources_and_rules--rules--pass_rules--vlan_tag))
<a id="nestedatt--sources_and_rules--rules--pass_rules--ip_rewrite"></a>
### Nested Schema for `sources_and_rules.rules.pass_rules.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--sources_and_rules--rules--pass_rules--rewrite"></a>
### Nested Schema for `sources_and_rules.rules.pass_rules.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--sources_and_rules--rules--pass_rules--vlan_tag"></a>
### Nested Schema for `sources_and_rules.rules.pass_rules.vlan_tag`

Required:

* `vlan_action` (String)
Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)
<a id="nestedatt--sources_and_rules--components"></a>
### Nested Schema for `sources_and_rules.components`

Required:

* `cluster_id` (String) - id of the defining cluster
Optional:

* `components` (Attributes List) (see [below for nested schema](#nestedatt--sources_and_rules--components--components))
<a id="nestedatt--sources_and_rules--components--components"></a>
### Nested Schema for `sources_and_rules.components.components`

Required:

* `ids` (List of Dynamic)
* `type` (String)
Read-Only:

* `metadata_details` (Attributes List) (see [below for nested schema](#nestedatt--sources_and_rules--components--components--metadata_details))
<a id="nestedatt--sources_and_rules--components--components--metadata_details"></a>
### Nested Schema for `sources_and_rules.components.components.metadata_details`

Read-Only:

* `alias` (String) - port alias
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--sources_and_rules--components--components--metadata_details--health_state_reasons))
* `port_type` (String) - port alias
* `traffic_health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--sources_and_rules--components--components--metadata_details--traffic_health_state_reasons))
<a id="nestedatt--sources_and_rules--components--components--metadata_details--health_state_reasons"></a>
### Nested Schema for `sources_and_rules.components.components.metadata_details.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--sources_and_rules--components--components--metadata_details--traffic_health_state_reasons"></a>
### Nested Schema for `sources_and_rules.components.components.metadata_details.traffic_health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--sources_and_rules--ip_rewrite"></a>
### Nested Schema for `sources_and_rules.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--sources_and_rules--rewrite"></a>
### Nested Schema for `sources_and_rules.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--sources_and_rules--tags"></a>
### Nested Schema for `sources_and_rules.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--sources_and_rules--vlan_tag"></a>
### Nested Schema for `sources_and_rules.vlan_tag`

Required:

* `vlan_action` (String)
Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)
<a id="nestedatt--sources_and_rules--health_state_reasons"></a>
### Nested Schema for `sources_and_rules.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_traffic_flow.example {alias}
```
