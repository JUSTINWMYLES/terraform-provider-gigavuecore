---
page_title: "gigavuecore_load_mobility_sam_node Data Source - gigavuecore"
subcategory: ""
description: |-
  Load mobility intent sam node by alias
---

# gigavuecore_load_mobility_sam_node Data Source

Load mobility intent sam node by alias

## Example Usage

```terraform
data "gigavuecore_load_mobility_sam_node" "example" {
  samnode_alias = null
  solution_alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `samnode_alias` (String, required) - Alias of the sam node
* `solution_alias` (String, required) - Alias of the mobility solution

### Attributes

In addition to all arguments above, the following attributes are exported:

* `alias` (String, computed) - Alias of the SAM Exporter node
* `app_profile_config` (Object({application_id, applications, counter, flow, gtpu, inner_ipv4, inner_ipv6, outer_ipv4, outer_ipv6, timestamp, transport}), computed)
  * `application_id` (Bool, computed)
  * `applications` (Object({attributes, is_user_defined, name}), computed)
    * `attributes` (List(Object({name, value})), computed)
    * `is_user_defined` (Bool, computed) - Default value is false. Must be set to true only for App Intel solution while configuring user defined apps
    * `name` (String, computed) - application name
  * `counter` (Object({bytes, bytes_long, packets, packets_long}), computed)
    * `bytes` (Bool, computed)
    * `bytes_long` (Bool, computed)
    * `packets` (Bool, computed)
    * `packets_long` (Bool, computed)
  * `flow` (Object({end_reason}), computed)
    * `end_reason` (Bool, computed)
  * `gtpu` (Object({qfi, teid}), computed)
    * `qfi` (Bool, computed)
    * `teid` (Bool, computed)
  * `inner_ipv4` (Object({destination, protocol, source}), computed)
    * `destination` (Bool, computed)
    * `protocol` (Bool, computed)
    * `source` (Bool, computed)
  * `inner_ipv6` (Object({destination, next_header, source}), computed)
    * `destination` (Bool, computed)
    * `next_header` (Bool, computed)
    * `source` (Bool, computed)
  * `outer_ipv4` (Object({destination, source}), computed)
    * `destination` (Bool, computed)
    * `source` (Bool, computed)
  * `outer_ipv6` (Object({destination, source}), computed)
    * `destination` (Bool, computed)
    * `source` (Bool, computed)
  * `timestamp` (Object({flow_end_msec, flow_endsec, flow_start_msec, flow_startsec}), computed)
    * `flow_end_msec` (Bool, computed)
    * `flow_endsec` (Bool, computed)
    * `flow_start_msec` (Bool, computed)
    * `flow_startsec` (Bool, computed)
  * `transport` (Object({dst_port, src_port}), computed)
    * `dst_port` (Bool, computed)
    * `src_port` (Bool, computed)
* `config_status` (String, computed)
* `config_status_reasons` (String, computed)
* `control_plane_setting` (Object({encoding, encoding_format, event_enable, trigger}), computed)
  * `encoding` (String, computed)
  * `encoding_format` (String, computed)
  * `event_enable` (Object({modify, update}), computed)
    * `modify` (Bool, computed)
    * `update` (Bool, computed)
  * `trigger` (String, computed)
* `deployed` (Bool, computed) - True when the SAM node is attempted for deployment
* `deployment_details` (List(Object({engine_port, sam_node_alias})), computed)
* `engine_meta_data_cache_configs` (List(Object({engine_port, event, flow_behavior, flows_size, idle_timeout, match, observation_domain_id})), computed)
* `engine_source_mappings` (List(Object({engine_port, network_source})), computed)
* `exporter_config` (Object({active_timeout, inactive_timeout, record_type}), computed)
  * `active_timeout` (String, computed)
  * `inactive_timeout` (Number, computed)
  * `record_type` (String, computed)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `ip_interface_alias` (String, computed) - Alias of the ip interface used by SAM Exporter Node
* `location` (Object({cluster_id, engine_ports}), computed) - Location of the engine port
  * `cluster_id` (String, computed)
  * `engine_ports` (List(String), computed)
* `node_override_network_ports` (List(String), computed) - Network ports for the SAM Exporter node. Its a list of ports of format cluster:port
* `node_type` (String, computed) - Type of the SAM Exporter Node
* `param_configs` (List(Object({engine_port, resource_metadata})), computed)
* `smaf_details` (List(Object({control_application, management_address, user_application})), computed)
* `tags` (List(Object({tag_key, tag_values})), computed) - User defined tags (Aggregation tags)
* `traffic_sources` (List(Object({comment, expand_port_identifier, group_interfaces, ip4_frag_rule_type, network_function_interfaces, network_function_name, network_function_type, source_group_id, tags})), computed) - List of all traffic sources for the SAM Exporter node

