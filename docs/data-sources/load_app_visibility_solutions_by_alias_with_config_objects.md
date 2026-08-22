---
page_title: "gigavuecore_load_app_visibility_solutions_by_alias_with_config_objects Data Source - gigavuecore"
subcategory: ""
description: |-
  Load an Application Intelligence Solution by alias with config objects
---

# gigavuecore_load_app_visibility_solutions_by_alias_with_config_objects Data Source

Load an Application Intelligence Solution by alias with config objects

## Example Usage

```terraform
data "gigavuecore_load_app_visibility_solutions_by_alias_with_config_objects" "example" {
  solution_alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `solution_alias` (String, required) - Alias of the solution

### Attributes

In addition to all arguments above, the following attributes are exported:

* `app_env_config` (Object({app_info, argv, name, nr_cores, nr_s_cores}), computed)
  * `app_info` (String, computed)
  * `argv` (String, computed)
  * `name` (String, computed)
  * `nr_cores` (Number, computed)
  * `nr_s_cores` (Number, computed)
* `app_env_id` (String, computed) - ID of the Application Environment
* `app_exporter_config` (Object({app_instance_name, app_metadata_tiered_batch_id, cache_config, cache_config_alias, destination_configs}), computed)
  * `app_instance_name` (String, computed)
  * `app_metadata_tiered_batch_id` (String, computed) - ID of the appExporterConfig tiered batch
  * `cache_config` (Object({advance_hash, alias, description, dpi_inject_limit, event, exporters, flow_behavior, match, multi_collect, network_profiles, observation_domain_id, sampling, size, timeout}), computed)
    * `advance_hash` (Bool, computed) - When scheduling flows for processing, use external encapsulate packet header if disable, use inner packet header if enable.
    * `alias` (String, computed)
    * `description` (String, computed)
    * `dpi_inject_limit` (Number, computed)
    * `event` (String, computed)
    * `exporters` (List(String), computed) - alias of metadata exporters to attach this cache
    * `flow_behavior` (String, computed) - direction for flow identification
    * `match` (Object({datalink, interface, ip, ipv4, ipv6, transport}), computed)
      * `datalink` (Object({mac_dst, mac_src, vlan}), computed)
        * `mac_dst` (Bool, computed)
        * `mac_src` (Bool, computed)
        * `vlan` (Bool, computed)
      * `interface` (Object({in_name_width, in_physical_width}), computed)
        * `in_name_width` (Number, computed)
        * `in_physical_width` (Number, computed)
      * `ip` (Object({version}), computed)
        * `version` (Bool, computed)
      * `ipv4` (Object({destination, dscp, fragmentation, header_len, option_map, precedence, protocol, section, source, tos, total_length, ttl}), computed)
        * `destination` (Object({prefix_min_mask}), computed)
          * `prefix_min_mask` (String, computed) - ipv4 destination prefix minimum-mask - netmask or mask length
        * `dscp` (Bool, computed)
        * `fragmentation` (Object({flags, offset}), computed)
          * `flags` (Bool, computed)
          * `offset` (Bool, computed)
        * `header_len` (Bool, computed)
        * `option_map` (Bool, computed)
        * `precedence` (Bool, computed)
        * `protocol` (Bool, computed)
        * `section` (Object({header_size, payload_size}), computed)
          * `header_size` (Number, computed)
          * `payload_size` (Number, computed)
        * `source` (Object({prefix_min_mask}), computed)
          * `prefix_min_mask` (String, computed) - ipv4 source prefix minimum-mask - netmask or mask length
        * `tos` (Bool, computed)
        * `total_length` (Bool, computed)
        * `ttl` (Bool, computed)
      * `ipv6` (Object({destination, dscp, extension_map, flow_label, fragmentation, hop_limit, length, next_header, precedence, section, source, traffic_class}), computed)
        * `destination` (Object({prefix_min_mask}), computed)
          * `prefix_min_mask` (String, computed)
        * `dscp` (Bool, computed)
        * `extension_map` (Bool, computed)
        * `flow_label` (Bool, computed)
        * `fragmentation` (Object({flags, offset}), computed)
          * `flags` (Bool, computed)
          * `offset` (Bool, computed)
        * `hop_limit` (Bool, computed)
        * `length` (Object({header, payload, total}), computed)
          * `header` (Bool, computed)
          * `payload` (Bool, computed)
          * `total` (Bool, computed)
        * `next_header` (Bool, computed)
        * `precedence` (Bool, computed)
        * `section` (Object({header_size, payload_size}), computed)
          * `header_size` (Number, computed)
          * `payload_size` (Number, computed)
        * `source` (Object({prefix_min_mask}), computed)
          * `prefix_min_mask` (String, computed) - ipv6 source prefix minimum-mask - netmask or mask length
        * `traffic_class` (Bool, computed)
      * `transport` (Object({dst_port, icmp, src_port, tcp, udp}), computed)
        * `dst_port` (Bool, computed)
        * `icmp` (Object({ipv4_code, ipv4_type, ipv6_code, ipv6_type}), computed)
          * `ipv4_code` (Bool, computed)
          * `ipv4_type` (Bool, computed)
          * `ipv6_code` (Bool, computed)
          * `ipv6_type` (Bool, computed)
        * `src_port` (Bool, computed)
        * `tcp` (Object({ack_number, dst_port, flags, header_len, seq_number, src_port, urgent_ptr, window_size}), computed)
          * `ack_number` (Bool, computed)
          * `dst_port` (Bool, computed)
          * `flags` (Bool, computed)
          * `header_len` (Bool, computed)
          * `seq_number` (Bool, computed)
          * `src_port` (Bool, computed)
          * `urgent_ptr` (Bool, computed)
          * `window_size` (Bool, computed)
        * `udp` (Object({dst_port, msg_len, src_port}), computed)
          * `dst_port` (Bool, computed)
          * `msg_len` (Bool, computed)
          * `src_port` (Bool, computed)
    * `multi_collect` (Bool, computed) - Collect all attributes as it is discovered when enable. It will export the same record once when disable.
    * `network_profiles` (List(String), computed) - alias of metadata network profiles to attach this cache
    * `observation_domain_id` (Number, computed)
    * `sampling` (Object({mode, single_sampling_rate}), computed)
      * `mode` (String, computed)
      * `single_sampling_rate` (Number, computed) - Packet interval window size. Valid values: 10-16000 (in packets)
    * `size` (Object({flows}), computed) - size of the flows
      * `flows` (Number, computed) - size of flows in millions
    * `timeout` (Object({idle}), computed)
      * `idle` (Number, computed) - idle timeout in seconds. max value 7days. default 30 min
  * `cache_config_alias` (String, computed) - Alias of the cacheConfig
  * `destination_configs` (List(Object({application_names, destination_name, export_meta_app_profile, export_meta_app_profile_alias, exporter_alias, exporter_config, iface})), computed)
* `app_filter_config` (Object({app_filter_tiered_batch_id, egress_traffic_configs, map_alias, sapf_profile, sapf_profile_alias}), computed)
  * `app_filter_tiered_batch_id` (String, computed) - ID of the appFilterConfig tiered batch
  * `egress_traffic_configs` (List(Object({drop_app_profile_alias, drop_application_names, pass_app_profile_alias, pass_application_names, priority, rule_sets, tunnel_aliases, tunnel_interface_mappings})), computed)
  * `map_alias` (String, computed) - Alias of the appFiltering map
  * `sapf_profile` (Object({alias, bidi, buffering, session_fields, timeout}), computed) - Session-Aware APF
    * `alias` (String, computed)
    * `bidi` (Bool, computed) - include reverse traffic for the selected sessionFields. Not applicable for 'mplsLabel', 'gtpuTeid', 'vlanId'
    * `buffering` (Object({buffer_count_before_match, enabled, protocol}), computed) - Session-Aware APF Buffereing settings
      * `buffer_count_before_match` (Number, computed) - Maximum number of packets BSAPF will buffer per session before APF match
      * `enabled` (Bool, computed)
      * `protocol` (String, computed) - When buffering is enabled, changing protocol requires reboot before new value takes effect.
    * `session_fields` (Set(Object({pos, type})), computed) - A sessionField cannot contain overlapping session attribute, position pairs (e.g. ipv4-5tuple pos 1 and ipv4-src pos 1 is not allowed). Max number of session fields is 20
    * `timeout` (Number, computed) - in seconds
  * `sapf_profile_alias` (String, computed) - Alias of the SAPF Profile
* `config_status` (String, computed) - configuration status
* `config_status_reason` (String, computed) - configuration status reasons
* `conn_id` (String, computed) - ID of the Connection Domain
* `dedup` (Object({app_config_id, app_instance_name, dedup_tiered_batch_id, enabled, gs_params_name}), computed)
  * `app_config_id` (String, computed)
  * `app_instance_name` (String, computed)
  * `dedup_tiered_batch_id` (String, computed)
  * `enabled` (Bool, computed)
  * `gs_params_name` (String, computed)
* `dynamic_scale_unit` (Bool, computed) - When set as true, FM automatically sets the optimal scaleUnit
* `env_id` (String, computed) - ID of the Environment
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `ingress_traffic_configs` (List(Object({source_selector_aliases, src_raw_end_point_interfaces, tunnel_aliases, tunnel_interface_mappings})), computed)
* `monitor_solution_config` (Object({app_instance_config, app_instance_config_id, app_instance_name, app_viz_tiered_batch_id, destination_config, mgmt_interface}), computed)
  * `app_instance_config` (Object({action, exporter}), computed)
    * `action` (Bool, computed)
    * `exporter` (String, computed)
  * `app_instance_config_id` (String, computed)
  * `app_instance_name` (String, computed)
  * `app_viz_tiered_batch_id` (String, computed) - ID of the monitoringSolutionConfig tiered batch
  * `destination_config` (Object({destination_name, exporter_alias, exporter_config}), computed)
    * `destination_name` (String, computed)
    * `exporter_alias` (String, computed)
    * `exporter_config` (Object({alias, destination, monitor, type}), computed)
      * `alias` (String, computed)
      * `destination` (Object({ipv4_address, ipv6_address}), computed)
        * `ipv4_address` (String, computed)
        * `ipv6_address` (String, computed)
      * `monitor` (Object({timeout}), computed)
        * `timeout` (Number, computed)
      * `type` (String, computed)
  * `mgmt_interface` (String, computed) - Pass this value if you would like to use the management interface of the Node to export App Monitoring metadata to FM
* `scale_unit` (Number, computed) - Number of units of memory needed for each Application
* `solution_created_timestamp` (String, computed)
* `solution_desc` (String, computed) - Description of the solution
* `solution_modified_timestamp` (String, computed)
* `traffic_policy_graph_alias` (String, computed) - Alias of the Traffic Policy Graph

