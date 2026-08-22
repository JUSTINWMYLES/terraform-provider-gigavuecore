---
page_title: "gigavuecore_app_visibility Resource - gigavuecore"
subcategory: ""
description: |-
  Load an Application Intelligence Solution by alias
---

# gigavuecore_app_visibility Resource

Load an Application Intelligence Solution by alias

## Example Usage

```terraform
resource "gigavuecore_app_visibility" "example" {
  app_export_config = {}
  app_filter_config = {}
  conn_id = null
  distribute_traffic = null
  dynamic_scale_unit = null
  env_id = null
  ingress_traffic_configs = []
  monitor_solution_config = {}
  scale_unit = null
  solution_alias = null
  solution_desc = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `app_export_config` (Object({cache_config, destination_configs}), optional)
  * `cache_config` (Object({advance_hash, alias, description, dpi_inject_limit, event, exporters, flow_behavior, match, multi_collect, network_profiles, observation_domain_id, sampling, size, timeout}), optional)
    * `advance_hash` (Bool, optional) - When scheduling flows for processing, use external encapsulate packet header if disable, use inner packet header if enable.
    * `alias` (String, required)
    * `description` (String, optional)
    * `dpi_inject_limit` (Number, optional)
    * `event` (String, optional)
    * `exporters` (List(String), optional) - alias of metadata exporters to attach this cache
    * `flow_behavior` (String, optional) - direction for flow identification
    * `match` (Object({datalink, interface, ip, ipv4, ipv6, transport}), optional)
      * `datalink` (Object({mac_dst, mac_src, vlan}), optional)
        * `mac_dst` (Bool, optional)
        * `mac_src` (Bool, optional)
        * `vlan` (Bool, optional)
      * `interface` (Object({in_name_width, in_physical_width}), optional)
        * `in_name_width` (Number, optional)
        * `in_physical_width` (Number, optional)
      * `ip` (Object({version}), optional)
        * `version` (Bool, optional)
      * `ipv4` (Object({destination, dscp, fragmentation, header_len, option_map, precedence, protocol, section, source, tos, total_length, ttl}), optional)
        * `destination` (Object({prefix_min_mask}), optional)
          * `prefix_min_mask` (String, optional) - ipv4 destination prefix minimum-mask - netmask or mask length
        * `dscp` (Bool, optional)
        * `fragmentation` (Object({flags, offset}), optional)
          * `flags` (Bool, optional)
          * `offset` (Bool, optional)
        * `header_len` (Bool, optional)
        * `option_map` (Bool, optional)
        * `precedence` (Bool, optional)
        * `protocol` (Bool, optional)
        * `section` (Object({header_size, payload_size}), optional)
          * `header_size` (Number, optional)
          * `payload_size` (Number, optional)
        * `source` (Object({prefix_min_mask}), optional)
          * `prefix_min_mask` (String, optional) - ipv4 source prefix minimum-mask - netmask or mask length
        * `tos` (Bool, optional)
        * `total_length` (Bool, optional)
        * `ttl` (Bool, optional)
      * `ipv6` (Object({destination, dscp, extension_map, flow_label, fragmentation, hop_limit, length, next_header, precedence, section, source, traffic_class}), optional)
        * `destination` (Object({prefix_min_mask}), optional)
          * `prefix_min_mask` (String, optional)
        * `dscp` (Bool, optional)
        * `extension_map` (Bool, optional)
        * `flow_label` (Bool, optional)
        * `fragmentation` (Object({flags, offset}), optional)
          * `flags` (Bool, optional)
          * `offset` (Bool, optional)
        * `hop_limit` (Bool, optional)
        * `length` (Object({header, payload, total}), optional)
          * `header` (Bool, optional)
          * `payload` (Bool, optional)
          * `total` (Bool, optional)
        * `next_header` (Bool, optional)
        * `precedence` (Bool, optional)
        * `section` (Object({header_size, payload_size}), optional)
          * `header_size` (Number, optional)
          * `payload_size` (Number, optional)
        * `source` (Object({prefix_min_mask}), optional)
          * `prefix_min_mask` (String, optional) - ipv6 source prefix minimum-mask - netmask or mask length
        * `traffic_class` (Bool, optional)
      * `transport` (Object({dst_port, icmp, src_port, tcp, udp}), optional)
        * `dst_port` (Bool, optional)
        * `icmp` (Object({ipv4_code, ipv4_type, ipv6_code, ipv6_type}), optional)
          * `ipv4_code` (Bool, optional)
          * `ipv4_type` (Bool, optional)
          * `ipv6_code` (Bool, optional)
          * `ipv6_type` (Bool, optional)
        * `src_port` (Bool, optional)
        * `tcp` (Object({ack_number, dst_port, flags, header_len, seq_number, src_port, urgent_ptr, window_size}), optional)
          * `ack_number` (Bool, optional)
          * `dst_port` (Bool, optional)
          * `flags` (Bool, optional)
          * `header_len` (Bool, optional)
          * `seq_number` (Bool, optional)
          * `src_port` (Bool, optional)
          * `urgent_ptr` (Bool, optional)
          * `window_size` (Bool, optional)
        * `udp` (Object({dst_port, msg_len, src_port}), optional)
          * `dst_port` (Bool, optional)
          * `msg_len` (Bool, optional)
          * `src_port` (Bool, optional)
    * `multi_collect` (Bool, optional) - Collect all attributes as it is discovered when enable. It will export the same record once when disable.
    * `network_profiles` (List(String), optional) - alias of metadata network profiles to attach this cache
    * `observation_domain_id` (Number, optional)
    * `sampling` (Object({mode, single_sampling_rate}), optional)
      * `mode` (String, optional)
      * `single_sampling_rate` (Number, optional) - Packet interval window size. Valid values: 10-16000 (in packets)
    * `size` (Object({flows}), optional) - size of the flows
      * `flows` (Number, optional) - size of flows in millions
    * `timeout` (Object({idle}), optional)
      * `idle` (Number, optional) - idle timeout in seconds. max value 7days. default 30 min
  * `destination_configs` (List(Object({application_names, destination_name, export_meta_app_profile, export_meta_app_profile_alias, exporter_alias, exporter_config, iface})), optional)
* `app_filter_config` (Object({app_filter_tiered_batch_id, egress_traffic_configs, map_alias, sapf_profile_alias}), optional)
  * `app_filter_tiered_batch_id` (String, computed) - ID of the appFilterConfig tiered batch
  * `egress_traffic_configs` (List(Object({drop_application_names, pass_application_names, priority, tunnel_aliases})), optional)
  * `map_alias` (String, computed) - Alias of the appFiltering map
  * `sapf_profile_alias` (String, computed) - Alias of the SAPF Profile
* `conn_id` (String, required) - ID of the Connection Domain
* `distribute_traffic` (Bool, optional) - Indicates Traffic Distribution enabled or not
* `dynamic_scale_unit` (Bool, optional) - When set as true, FM automatically sets the optimal scaleUnit
* `env_id` (String, required) - ID of the Environment
* `ingress_traffic_configs` (List(Object({source_selector_aliases, src_raw_end_point_interfaces, tunnel_aliases, tunnel_interface_mappings})), optional)
* `monitor_solution_config` (Object({app_instance_config_id, app_instance_name, app_viz_tiered_batch_id, destination_config, gs_param_config_id, mgmt_interface, user_defined_application_profile, user_defined_applications}), required)
  * `app_instance_config_id` (String, computed)
  * `app_instance_name` (String, computed)
  * `app_viz_tiered_batch_id` (String, computed) - ID of the monitoringSolutionConfig tiered batch
  * `destination_config` (Object({destination_name, exporter_alias, iface}), required)
    * `destination_name` (String, computed)
    * `exporter_alias` (String, computed)
    * `iface` (String, optional) - Interface Mapping for Egress Tunnel
  * `gs_param_config_id` (String, computed) - Gs param name on which the solution is applied
  * `mgmt_interface` (String, optional) - Pass this value if you would like to use the management interface of the Node to export App Monitoring metadata to FM
  * `user_defined_application_profile` (Dynamic, computed)
  * `user_defined_applications` (List(List(String)), computed)
* `scale_unit` (Number, optional) - Number of units of memory needed for each Application
* `solution_alias` (String, required) - Alias of the solution
* `solution_desc` (String, optional) - Description of the solution

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `app_env_id` (String, computed) - ID of the Application Environment
* `app_exporter_config` (Object({app_instance_name, app_metadata_tiered_batch_id, cache_config_alias, destination_configs}), computed)
  * `app_instance_name` (String, computed) - The name of the AppMetadata app instance
  * `app_metadata_tiered_batch_id` (String, computed) - ID of the appExporterConfig tiered batch
  * `cache_config_alias` (String, computed) - Alias of the cacheConfig
  * `destination_configs` (List(Object({destination_name, exporter_alias, iface, template_name})), computed)
* `app_filter_config` (Object({app_filter_tiered_batch_id, egress_traffic_configs, map_alias, sapf_profile_alias}), computed)
  * `app_filter_tiered_batch_id` (String, computed) - ID of the appFilterConfig tiered batch
  * `egress_traffic_configs` (List(Object({drop_application_names, pass_application_names, priority, tunnel_aliases})), optional)
  * `map_alias` (String, computed) - Alias of the appFiltering map
  * `sapf_profile_alias` (String, computed) - Alias of the SAPF Profile
* `config_status` (String, computed) - configuration status
* `config_status_reasons` (String, computed) - configuration status reasons
* `dedup` (Object({app_config_id, app_instance_name, dedup_tiered_batch_id, enabled, gs_params_name}), computed)
  * `app_config_id` (String, computed)
  * `app_instance_name` (String, computed)
  * `dedup_tiered_batch_id` (String, computed)
  * `enabled` (Bool, computed)
  * `gs_params_name` (String, computed)
* `dynamic_scale_unit` (Bool, computed) - When set as true, FM automatically sets the optimal scaleUnit
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `ingress_traffic_configs` (List(Object({source_selector_aliases, src_raw_end_point_interfaces, tunnel_aliases, tunnel_interface_mappings})), computed)
* `scale_unit` (Number, computed) - Number of units of memory needed for each Application
* `solution_created_timestamp` (String, computed)
* `solution_desc` (String, computed) - Description of the solution
* `solution_modified_timestamp` (String, computed)
* `traffic_policy_graph_alias` (String, computed) - Alias of the Traffic Policy Graph

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_app_visibility.example {solution_alias}
```
