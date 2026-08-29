---
page_title: "gigavuecore_load_mobility_control_node Data Source - gigavuecore"
subcategory: ""
description: |-
  Load mobility intent control node by alias
---

# gigavuecore_load_mobility_control_node Data Source

Load mobility intent control node by alias

## Example Usage

```terraform
data "gigavuecore_load_mobility_control_node" "example" {
  cpnode_alias   = "example"
  solution_alias = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cpnode_alias` (String, required) - Alias of the control node
* `solution_alias` (String, required) - Alias of the mobility solution

### Attributes

In addition to all arguments above, the following attributes are exported:

* `additional_tool_ports` (List of String, computed) - Additional Tool ports for the Control node which takes a copy of traffic along with vport that is added into the first level maps. Its a list of ports of format cluster:port
* `alias` (String, computed) - Alias of the Control Node
* `app5_g_http2_ports` (List of Number, computed) - List of TCP ports
* `app_tcp` (Attributes, computed) - TCP Loadbalancing properties for Control 5G node (PCPN\_5G) (see [below for nested schema](#nestedatt--app_tcp))
* `collector_tools` (List of String, computed) - Collector Tool ports. Its a list of ports of format cluster:port
* `config_status` (String, computed)
* `config_status_reasons` (String, computed)
* `control_metadata_ip_interface_alias` (String, computed) - Alias of the ip interface used for exporting control json records
* `deployed` (Boolean, computed) - True when the Control node attempted for deployment
* `gtp_control_sample` (Boolean, computed) - When enabled, GTP Control plane traffic for the subscribers contained within the subscriber sample defined in GTP flowsample maps will be sent to tools. When disabled, all GTP Control plane traffic for all subscribers in GTP flowsample maps will be sent to tools.
* `gtp_random_sampling` (Attributes, computed) - GsGroup Gtp Random Sampling Parameters (see [below for nested schema](#nestedatt--gtp_random_sampling))
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `ip_interface_alias` (String, computed) - Alias of the ip interface used by Control Node
* `location` (Attributes, computed) - Location of the engine port (see [below for nested schema](#nestedatt--location))
* `node_override_network_ports` (List of String, computed) - Network ports for the Control node. Its a list of ports of format cluster:port
* `node_type` (String, computed) - Type of the Control Node
* `number_of5_g_sessions` (Number, computed) - Number of 5G sessions to allocate for Control 5G Node (PCPN\_5G)
* `number_of_lte_sessions` (Number, computed) - Number of LTE sessions to allocate for Control LTE Node (PCPN\_LTE)
* `tags` (Attributes List, computed) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--tags))
* `traffic_sources` (Attributes List, computed) - List of all traffic sources for the Control node (see [below for nested schema](#nestedatt--traffic_sources))

<a id="nestedatt--app_tcp"></a>
### Nested Schema for `app_tcp`

Read-Only:

* `application` (String) - To choose the action on Unknown Application Data
* `load_balance` (Boolean) - When true it enables TCP loadbalancing on the Tool Ports
* `tcp_control` (String) - To choose the action on TCP Control messages
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

