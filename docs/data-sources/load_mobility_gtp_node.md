---
page_title: "gigavuecore_load_mobility_gtp_node Data Source - gigavuecore"
subcategory: ""
description: |-
  Load mobility intent gtp node by alias
---

# gigavuecore_load_mobility_gtp_node Data Source

Load mobility intent gtp node by alias

## Example Usage

```terraform
data "gigavuecore_load_mobility_gtp_node" "example" {
  gtpnode_alias = null
  solution_alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `gtpnode_alias` (String, required) - Alias of the gtp node
* `solution_alias` (String, required) - Alias of the mobility solution

### Attributes

In addition to all arguments above, the following attributes are exported:

* `additional_tool_ports` (List(String), computed) - Additional Tool ports for the GTP node which takes a copy of traffic along with vport that is added into the first level maps. Its a list of ports of format cluster:port
* `alias` (String, computed) - Alias of the GTP node
* `collector_tools` (List(String), computed) - Collector Tool ports. Its a list of ports of format cluster:port
* `config_status` (String, computed)
* `config_status_reasons` (String, computed)
* `deployed` (Bool, computed) - True when the GTP node is attempted for deployment
* `gtp_random_sampling` (Object({enabled, interval}), computed) - GsGroup Gtp Random Sampling Parameters
  * `enabled` (Bool, computed) - When enabled, sampling of subscriber's sessions happens in random fashion
  * `interval` (Number, computed) - Rotation Interval in multiples of 12 (hrs)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `location` (Object({cluster_id, engine_ports}), computed) - Location of the engine port
  * `cluster_id` (String, computed)
  * `engine_ports` (List(String), computed)
* `node_override_network_ports` (List(String), computed) - Network ports for the GTP node. Its a list of ports of format cluster:port
* `node_type` (String, computed) - Type of the GTP Node
* `tags` (List(Object({tag_key, tag_values})), computed) - User defined tags (Aggregation tags)
* `traffic_sources` (List(Object({comment, expand_port_identifier, group_interfaces, ip4_frag_rule_type, network_function_interfaces, network_function_name, network_function_type, source_group_id, tags})), computed) - List of all traffic sources for the GTP node

