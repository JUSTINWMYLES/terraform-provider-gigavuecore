---
page_title: "gigavuecore_load_mobility_user_node Data Source - gigavuecore"
subcategory: ""
description: |-
  Load mobility intent user node by alias
---

# gigavuecore_load_mobility_user_node Data Source

Load mobility intent user node by alias

## Example Usage

```terraform
data "gigavuecore_load_mobility_user_node" "example" {
  solution_alias = null
  upnode_alias   = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `solution_alias` (String, required) - Alias of the mobility solution
* `upnode_alias` (String, required) - Alias of the user node

### Attributes

In addition to all arguments above, the following attributes are exported:

* `additional_tool_ports` (List of String, computed) - Additional Tool ports for the User node which takes a copy of traffic along with vport that is added into the first level maps. Its a list of ports of format cluster:port
* `alias` (String, computed) - Alias of the User node
* `collector_tools` (List of String, computed) - Collector Tool ports. Its a list of ports of format cluster:port
* `config_status` (String, computed)
* `config_status_reasons` (String, computed)
* `deployed` (Boolean, computed) - True when the User node is attempted for deployment
* `gtp_control_sample` (Boolean, computed) - When enabled, GTP Control plane traffic for the subscribers contained within the subscriber sample defined in GTP flowsample maps will be sent to tools. When disabled, all GTP Control plane traffic for all subscribers in GTP flowsample maps will be sent to tools.
* `gtp_random_sampling` (Attributes, computed) - GsGroup Gtp Random Sampling Parameters (see [below for nested schema](#nestedatt--gtp_random_sampling))
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `ip_interface_alias` (String, computed) - Alias of the ip interface used by User Node
* `location` (Attributes, computed) - Location of the engine port (see [below for nested schema](#nestedatt--location))
* `node_override_network_ports` (List of String, computed) - Network ports for the User node. Its a list of ports of format cluster:port
* `node_type` (String, computed) - Type of the User Node
* `number_of_lte_sessions` (Number, computed) - Number of LTE sessions to allocate for UPN Node
* `stand_alone_mode` (Boolean, computed) - Enable or Disable standAloneMode in User Nodes
* `tags` (Attributes List, computed) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--tags))
* `traffic_sources` (Attributes List, computed) - List of all traffic sources for the User node (see [below for nested schema](#nestedatt--traffic_sources))

<a id="nestedatt--gtp_random_sampling"></a>
### Nested Schema for `gtp_random_sampling`

Read-Only:

* `enabled` (Boolean) - When enabled, sampling of subscriber's sessions happens in random fashion
* `interval` (Number) - Rotation Interval in multiples of 12 (hrs)
<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--location"></a>
### Nested Schema for `location`

Read-Only:

* `cluster_id` (String)
* `engine_ports` (List of String)
<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--traffic_sources"></a>
### Nested Schema for `traffic_sources`

Read-Only:

* `comment` (String)
* `expand_port_identifier` (Boolean) - if enabled two rules added will be created with source port and destination port along with the other identifiers.
* `group_interfaces` (Boolean) - Enable if the user needs to combine rules in tunnelIdentifiers into single ingress map
* `ip4_frag_rule_type` (String) - if provided, the ip frag rule of given type would be appended to the rules.
* `network_function_interfaces` (Attributes List) (see [below for nested schema](#nestedatt--traffic_sources--network_function_interfaces))
* `network_function_name` (String)
* `network_function_type` (String)
* `source_group_id` (String) - if provided a vport with this sourceGroupId would be created else default vport would be used
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--traffic_sources--tags))
<a id="nestedatt--traffic_sources--network_function_interfaces"></a>
### Nested Schema for `traffic_sources.network_function_interfaces`

Read-Only:

* `comment` (String)
* `interface_type` (String)
* `source_override_dst_ports` (List of String) - Tool ports which takes a copy of traffic along with vport that is added into the first level maps. Its a list of ports of format cluster:port
* `source_override_network_ports` (List of String) - Network ports for the network element. Its a list of ports of format cluster:port
* `tunnel_identifiers` (List of Dynamic)
<a id="nestedatt--traffic_sources--tags"></a>
### Nested Schema for `traffic_sources.tags`

Read-Only:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag

