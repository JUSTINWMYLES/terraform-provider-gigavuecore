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
* `enable` (Boolean, computed) - enable/disable map, applicable only to first level maps
* `error_message` (String, computed) - In case of configuration failure, this message provides details about the possible cause of the failure.
* `flows` (List of Dynamic, computed)
* `has_draft` (Boolean, computed)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `policy_id` (String, computed) - unique id
* `sources_and_rules` (Attributes List, computed) (see [below for nested schema](#nestedatt--sources_and_rules))
* `tags` (Attributes List, computed) (see [below for nested schema](#nestedatt--tags))
* `updated_by` (String, computed) - updated by of the traffic flow
* `updated_time` (Number, computed) - updated time of the traffic flow

<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--sources_and_rules"></a>
### Nested Schema for `sources_and_rules`

Read-Only:

* `alias` (String) - unique sourcesAndRules alias
* `components` (Attributes List) (see [below for nested schema](#nestedatt--sources_and_rules--components))
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--sources_and_rules--health_state_reasons))
* `inline_traffic_path` (String) - Only applicable for 'inline' map types, in which case defaults to 'normal'. When set to 'bypass', the 'dstPorts' must be empty
* `inline_traffic_type` (String) - Only applicable for 'inline/passAll' map types, in which case defaults to 'symmetric'. For 'asymmetric' maps, 'srcPort' must be of type 'inline-net'
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--sources_and_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--sources_and_rules--rewrite))
* `rule_matching` (String) - If set to 'blacklist', packet are passed when no 'drop' rules are matched. This field is only valid for 'regular/byRule' map types. (maps into 'no-rule-match' CLI command)
* `rule_type` (String) - 'byRule' is applicable to all map types; 'collector' is applicable to 'regular', 'inline' maps; 'passAll' is applicable to 'regular' and 'inline' maps
* `rules` (Attributes) - Map Rules Container. Private class (see [below for nested schema](#nestedatt--sources_and_rules--rules))
* `tags` (Attributes List) (see [below for nested schema](#nestedatt--sources_and_rules--tags))
* `traffic_type` (String) - Only applicable for 'firstLevel/byRule' map types, in which case defaults to 'user'
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--sources_and_rules--vlan_tag))
<a id="nestedatt--sources_and_rules--components"></a>
### Nested Schema for `sources_and_rules.components`

Read-Only:

* `cluster_id` (String) - id of the defining cluster
* `components` (Attributes List) (see [below for nested schema](#nestedatt--sources_and_rules--components--components))
<a id="nestedatt--sources_and_rules--components--components"></a>
### Nested Schema for `sources_and_rules.components.components`

Read-Only:

* `ids` (List of Dynamic)
* `metadata_details` (Attributes List) (see [below for nested schema](#nestedatt--sources_and_rules--components--components--metadata_details))
* `type` (String)
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
<a id="nestedatt--sources_and_rules--health_state_reasons"></a>
### Nested Schema for `sources_and_rules.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--sources_and_rules--ip_rewrite"></a>
### Nested Schema for `sources_and_rules.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--sources_and_rules--rewrite"></a>
### Nested Schema for `sources_and_rules.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--sources_and_rules--rules"></a>
### Nested Schema for `sources_and_rules.rules`

Read-Only:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--sources_and_rules--rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--sources_and_rules--rules--pass_rules))
<a id="nestedatt--sources_and_rules--rules--drop_rules"></a>
### Nested Schema for `sources_and_rules.rules.drop_rules`

Read-Only:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--sources_and_rules--rules--drop_rules--ip_rewrite))
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--sources_and_rules--rules--drop_rules--rewrite))
* `rule_id` (Number)
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--sources_and_rules--rules--drop_rules--vlan_tag))
<a id="nestedatt--sources_and_rules--rules--drop_rules--ip_rewrite"></a>
### Nested Schema for `sources_and_rules.rules.drop_rules.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--sources_and_rules--rules--drop_rules--rewrite"></a>
### Nested Schema for `sources_and_rules.rules.drop_rules.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--sources_and_rules--rules--drop_rules--vlan_tag"></a>
### Nested Schema for `sources_and_rules.rules.drop_rules.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)
<a id="nestedatt--sources_and_rules--rules--pass_rules"></a>
### Nested Schema for `sources_and_rules.rules.pass_rules`

Read-Only:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--sources_and_rules--rules--pass_rules--ip_rewrite))
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--sources_and_rules--rules--pass_rules--rewrite))
* `rule_id` (Number)
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--sources_and_rules--rules--pass_rules--vlan_tag))
<a id="nestedatt--sources_and_rules--rules--pass_rules--ip_rewrite"></a>
### Nested Schema for `sources_and_rules.rules.pass_rules.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--sources_and_rules--rules--pass_rules--rewrite"></a>
### Nested Schema for `sources_and_rules.rules.pass_rules.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--sources_and_rules--rules--pass_rules--vlan_tag"></a>
### Nested Schema for `sources_and_rules.rules.pass_rules.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)
<a id="nestedatt--sources_and_rules--tags"></a>
### Nested Schema for `sources_and_rules.tags`

Read-Only:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--sources_and_rules--vlan_tag"></a>
### Nested Schema for `sources_and_rules.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)
<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag

