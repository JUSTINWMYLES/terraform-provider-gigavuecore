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
  upnode_alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `solution_alias` (String, required) - Alias of the mobility solution
* `upnode_alias` (String, required) - Alias of the user node

### Attributes

In addition to all arguments above, the following attributes are exported:

* `additional_tool_ports` (List(String), computed) - Additional Tool ports for the User node which takes a copy of traffic along with vport that is added into the first level maps. Its a list of ports of format cluster:port
* `alias` (String, computed) - Alias of the User node
* `collector_tools` (List(String), computed) - Collector Tool ports. Its a list of ports of format cluster:port
* `config_status` (String, computed)
* `config_status_reasons` (String, computed)
* `deployed` (Bool, computed) - True when the User node is attempted for deployment
* `gtp_control_sample` (Bool, computed) - When enabled, GTP Control plane traffic for the subscribers contained within the subscriber sample defined in GTP flowsample maps will be sent to tools. When disabled, all GTP Control plane traffic for all subscribers in GTP flowsample maps will be sent to tools.
* `gtp_random_sampling` (Object({enabled, interval}), computed) - GsGroup Gtp Random Sampling Parameters
  * `enabled` (Bool, computed) - When enabled, sampling of subscriber's sessions happens in random fashion
  * `interval` (Number, computed) - Rotation Interval in multiples of 12 (hrs)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `ip_interface_alias` (String, computed) - Alias of the ip interface used by User Node
* `location` (Object({cluster_id, engine_ports}), computed) - Location of the engine port
  * `cluster_id` (String, computed)
  * `engine_ports` (List(String), computed)
* `node_override_network_ports` (List(String), computed) - Network ports for the User node. Its a list of ports of format cluster:port
* `node_type` (String, computed) - Type of the User Node
* `number_of_lte_sessions` (Number, computed) - Number of LTE sessions to allocate for UPN Node
* `stand_alone_mode` (Bool, computed) - Enable or Disable standAloneMode in User Nodes
* `tags` (List(Object({tag_key, tag_values})), computed) - User defined tags (Aggregation tags)
* `traffic_sources` (List(Object({comment, expand_port_identifier, group_interfaces, ip4_frag_rule_type, network_function_interfaces, network_function_name, network_function_type, source_group_id, tags})), computed) - List of all traffic sources for the User node

