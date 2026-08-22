---
page_title: "gigavuecore_solution Resource - gigavuecore"
subcategory: ""
description: |-
  Load Apps Visibility Solutions by alias
---

# gigavuecore_solution Resource

Load Apps Visibility Solutions by alias

## Example Usage

```terraform
resource "gigavuecore_solution" "example" {
  app_export_config = {}
  app_filter_config = {}
  associated_monitor_solution_alias = null
  cluster_id = null
  config_status = null
  delete_monitor_sol = null
  egress_map_aliases_to_delete = []
  exporter_aliases_to_delete = []
  ingress_map_aliases_to_delete = []
  ingress_traffic_configs = []
  monitor_solution_config = {}
  solution_alias = null
  solution_desc = null
  solution_status = null
  solution_type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `app_export_config` (Object({cache_config, cache_config_alias, destination_configs, gsop_alias, gsop_config}), optional)
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
  * `cache_config_alias` (String, optional)
  * `destination_configs` (List(Object({application_names, destination_name, export_ip_interface, export_meta_app_profile, export_meta_app_profile_alias, exporter_alias, exporter_config})), optional)
  * `gsop_alias` (String, optional)
  * `gsop_config` (Object({alias, cluster_id, gs_apps, gs_group, health_state, health_state_reasons}), optional) - GigaSMART Operation
    * `alias` (String, required)
    * `cluster_id` (String, optional) - id of the defining cluster
    * `gs_apps` (Object({apf, dedup, diameter_whitelist, flow_filter, flow_sampling, gseries_header_add, gseries_header_remove, gseries_load_balance, gseries_pattern_match, gtp_whitelist, header_add, header_remove, icap, inline_ssl, load_balance, masking, metadata_export, netflow, sa_apf, sip_whitelist, slicing, ssl_decrypt, trailer_add, trailer_remove, tunnel_decap, tunnel_encap}), required) - GigaSMART Applications for a GSOP. At least one GsApp must be defined for the config to be valid
      * `apf` (Object({enabled}), optional)
        * `enabled` (String, optional)
      * `dedup` (Object({enabled}), optional)
        * `enabled` (String, optional)
      * `diameter_whitelist` (Object({enabled}), optional)
        * `enabled` (String, optional)
      * `flow_filter` (Object({type}), optional) - GigaSMART 'Flow Filter' Application Configuration
        * `type` (String, required)
      * `flow_sampling` (Object({type}), optional) - GigaSMART 'Flow Sampling' Application Configuration
        * `type` (String, required)
      * `gseries_header_add` (Object({types}), optional) - Only applicable for G-series
        * `types` (Set(String), required)
      * `gseries_header_remove` (Object({enabled}), optional)
        * `enabled` (String, optional)
      * `gseries_load_balance` (Object({fixed_offset, variable_offset}), optional) - Only applicable for G-series per-rule GSOP
        * `fixed_offset` (Object({hash, length, offset}), optional) - Per-rule Fixed Offset Load Balancing config for G-seres devices
          * `hash` (String, required)
          * `length` (Number, required)
          * `offset` (Number, required)
        * `variable_offset` (Object({end_delim, hash, start_delim, start_field}), optional) - Per-rule Variable Offset Load Balancing config for G-seres devices
          * `end_delim` (String, required)
          * `hash` (String, required)
          * `start_delim` (String, required)
          * `start_field` (String, required)
      * `gseries_pattern_match` (Object({fixed_offset, variable_offset}), optional) - Only applicable for G-series per-rule GSOP
        * `fixed_offset` (Object({length, offset}), optional) - Per-rule Fixed Offset Pattern Match config for G-seres devices
          * `length` (Number, required)
          * `offset` (Number, required)
        * `variable_offset` (Object({end_delim, start_delim}), optional) - Per-rule Variable Offset Pattern Match config for G-seres devices
          * `end_delim` (String, required)
          * `start_delim` (String, required)
      * `gtp_whitelist` (Object({enabled}), optional)
        * `enabled` (String, optional)
      * `header_add` (Object({vlan}), optional) - GigaSMART 'Add Header' Application Configuration
        * `vlan` (Number, required)
      * `header_remove` (Object({ah1, ah2, custom_len, erspan_flow_id, fp_dst_switch_id, fp_src_switch_id, header_count, offset, offset_range_value, protocol, timestamp_format, vlan_header, vxlan_id}), optional) - GigaSMART 'Remove Header' Application Configuration
        * `ah1` (String, optional) - only valid and required for 'generic'. First anchor header, from which the header to be stripped is occurred.
        * `ah2` (String, optional) - only valid and required for 'generic'. next anchor header.
        * `custom_len` (Number, optional) - only valid for 'generic'. length of unknown header.
        * `erspan_flow_id` (Number, optional) - only applicable for 'erspan', specifies the flow id to strip. Value of 0 will strip all flow ids
        * `fp_dst_switch_id` (Number, optional) - Only valid and required for 'fabricPath', 12 bit destination switch id
        * `fp_src_switch_id` (Number, optional) - Only valid and required for 'fabricPath', 12 bit source switch id
        * `header_count` (Number, optional) - only valid for 'generic'. Number of headers to be stripped.
        * `offset` (String, optional) - only valid and required for 'generic'.'start': strip from start of ah1, 'end': strip from end of ah1 or any other offset in the range of length of ah1, 'offsetRangeValue' is required if offset is 'offsetRange'.
        * `offset_range_value` (Number, optional) - only valid and required when offset is 'offsetRange', integer within range of size of header
        * `protocol` (String, required) - 'gre' and 'fabricPath' are only applicable for H-series
        * `timestamp_format` (String, optional) - Timestamp format. Only valid and required for 'fm6000Ts'
        * `vlan_header` (String, optional) - Only valid when protocol is 'vlan'. Specifies the target vlan header to strip. Defaults to 'outer'
        * `vxlan_id` (Number, optional) - 24-bit value. Only valid when protocol is 'vxlan'. Specifies the vxlan id to strip. Value of '0' will strip all vxlan ids
      * `icap` (Object({icap_profile}), optional) - GigaSMART ICAP Configuration
        * `icap_profile` (String, required) - Alias of referenced ICAP Profile
      * `inline_ssl` (Object({inline_ssl_profile}), optional) - GigaSMART Inline SSL Profile Configuration
        * `inline_ssl_profile` (String, required) - Alias of referenced Inline SSL Profile
      * `load_balance` (Object({enhanced, stateful, stateless}), optional) - GigaSMART 'Load Balancing' Application Configuration
        * `enhanced` (Object({elb_alias}), optional) - Private class. Enhanced part of the GigaSMART 'Load Balancing' Application Configuration
          * `elb_alias` (String, required) - elb app alias
        * `stateful` (Object({app_type, diameter_key_hash_type, diameter_key_multi_hash_type, gtp_key_hash_type, lb_type, sip_key_hash_type}), optional) - Private class. Stateful part of the GigaSMART 'Load Balancing' Application Configuration
          * `app_type` (String, required) - 'sapf' option is added in release H 4.3.01. 'sip' and 'tunnel' added in release H 5.1. 'diameter' is added in release H 5.5
          * `diameter_key_hash_type` (String, optional) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyHash'
          * `diameter_key_multi_hash_type` (List(Object({avp_codevalue, key})), optional) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyMultiHash'
          * `gtp_key_hash_type` (String, optional) - required when 'appType' == 'gtp' and 'lbType' == 'gtpKeyHash'. ignored otherwise
          * `lb_type` (String, required) - 'gtpKeyHash' is only applicable when 'appType' == 'gtp'. 'sipKeyHash' is only applicable when 'appType' == 'sip'
          * `sip_key_hash_type` (String, optional) - required when 'appType' == 'sip' and 'lbType' == 'sipKeyHash'. ignored otherwise
        * `stateless` (Object({field_location, hash_fields}), optional) - Private class. Stateless part of the GigaSMART 'Load Balancing' Application Configuration
          * `field_location` (String, optional) - Ignored when 'hashFields' == 'gtpuTeid'. Required otherwise
          * `hash_fields` (String, required)
      * `masking` (Object({content_type, length, offset, pattern, protocol}), optional) - GigaSMART 'Masking' Application Configuration
        * `content_type` (String, optional) - content type that will trigger masking, only valid and required for protocol 'sip'
        * `length` (Number, optional) - max length is 9600 on H-series; 9000 on G-series, not applicable for protocol 'sip' valid and required otherwise
        * `offset` (Number, optional) - min offset is 0 for 'none'; 1 for protocol-referenced offsets, not applicable for protocol 'sip' valid and required otherwise
        * `pattern` (String, optional) - 1-byte hex mask, not applicable for protocol 'sip' valid and required otherwise
        * `protocol` (String, required)
      * `metadata_export` (Object({cache}), optional)
        * `cache` (String, optional) - metadata cache alias. cache should have exporters defined, and the exporters should have applicationProfiles defined
      * `netflow` (Object({enabled}), optional)
        * `enabled` (String, optional)
      * `sa_apf` (Object({enabled}), optional)
        * `enabled` (String, optional)
      * `sip_whitelist` (Object({enabled}), optional)
        * `enabled` (String, optional)
      * `slicing` (Object({enhanced, offset, protocol}), optional) - GigaSMART 'Slicing' Application Configuration
        * `enhanced` (String, optional) - enhanced-slicing apps alias
        * `offset` (Number, required) - min offset is 64 for 'none'; 4 for protocol-referenced offsets. required property till H 5.6
        * `protocol` (String, required) - required property till H 5.6
      * `ssl_decrypt` (Object({in_port, out_port}), optional) - GigaSMART 'SSL Decrypt' Application Configuration
        * `in_port` (Number, optional) - Port number of 0 represents 'any' port
        * `out_port` (Number, optional) - Port number of 0 represents 'auto' port
      * `trailer_add` (Object({types}), optional) - GigaSMART 'Add Trailer' Application Configuration
        * `types` (Set(String), required) - 'crc' is not applicable for G-series
      * `trailer_remove` (Object({enabled}), optional)
        * `enabled` (String, optional)
      * `tunnel_decap` (Object({custom, erspan_flow_id, gmip_port, l2_gre_key, tls_pcapng, type, vxlan}), optional) - GigaSMART 'Decapsulate Tunnel' Application Configuration
        * `custom` (Object({port_dst, port_src}), optional)
          * `port_dst` (Number, required) - when specified as 0, no validation will be done in the packet.
          * `port_src` (Number, required) - when specified as 0, no validation will be done in the packet.
        * `erspan_flow_id` (Number, optional) - only applicable for 'erspan', A Flow ID of 0 decapsulates all ERSPAN tunnel traffic regardless of Flow ID
        * `gmip_port` (Number, optional) - only applicable for 'gmip', in which case it is required. Specifies the UDP port on which the Tunnel Network port on the receiving GigaVUE H Series is listening. Must match the configuration of the portdst configured on the sending end of the tunnel
        * `l2_gre_key` (Number, optional) - only applicable for 'l2gre', in which case it is required.
        * `tls_pcapng` (Object({decap_key, listener}), optional)
          * `decap_key` (String, optional)
          * `listener` (String, optional)
        * `type` (String, required)
        * `vxlan` (Object({port_dst, port_src, vni}), optional)
          * `port_dst` (Number, required)
          * `port_src` (Number, required) - when specified as 0, no validation will be done in the packet.
          * `vni` (Number, required)
      * `tunnel_encap` (Object({gmip_config, l2_gre_config, tls_pcapng, type, vxlan_config}), optional) - GigaSMART 'Encapsulate Tunnel' Application Configuration
        * `gmip_config` (Object({dscp, dst_ip, dst_port, flow_label, prec, src_port, ttl}), optional) - Configuration for GMIP Tunnel Encapsulate GigaSmaprt App
          * `dscp` (Number, optional)
          * `dst_ip` (String, required) - IP Destination. IPv4 or IPv6.
          * `dst_port` (Number, required)
          * `flow_label` (Number, optional)
          * `prec` (Number, optional) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
          * `src_port` (Number, required)
          * `ttl` (Number, optional)
        * `l2_gre_config` (Object({dscp, dst_ip, flow_label, key, pg_dst, prec, session_field, session_pos, ttl}), optional) - Configuration for l2GRE Tunnel Encapsulate GigaSmaprt App
          * `dscp` (Number, optional)
          * `dst_ip` (String, optional) - ip destination, mutually exclusive with 'pgDst'. IPv4 or IPv6.
          * `flow_label` (Number, optional)
          * `key` (Number, required)
          * `pg_dst` (String, optional) - port group destination alias, mutually exclusive with 'dstIp'
          * `prec` (Number, optional) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
          * `session_field` (String, optional) - required with stateful loadBalance when 'appType' is 'tunnel'
          * `session_pos` (String, optional) - required if 'sessionField' is specified
          * `ttl` (Number, optional)
        * `tls_pcapng` (Object({exporter, exporter_group}), optional)
          * `exporter` (String, optional)
          * `exporter_group` (String, optional)
        * `type` (String, required)
        * `vxlan_config` (Object({dscp, dst_ip, dst_port, src_port, ttl, vni}), optional) - Configuration for Vxlan Tunnel Encapsulate GigaSmaprt App
          * `dscp` (Number, optional)
          * `dst_ip` (String, required) - IP Destination. IPv4 or IPv6.
          * `dst_port` (Number, required)
          * `src_port` (Number, required)
          * `ttl` (Number, optional)
          * `vni` (Number, required)
    * `gs_group` (String, required) - Alias of referenced managing GsGroup
    * `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
    * `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), optional)
* `app_filter_config` (Object({egress_traffic_config, sapf_profile, sapf_profile_alias}), optional)
  * `egress_traffic_config` (List(Object({drop_app_profile_alias, egress_traffic_map, gs_apps, gsop_alias, gsop_config, map_alias, pass_app_profile_alias, priority})), optional)
  * `sapf_profile` (Object({alias, bidi, buffering, cluster_id, packet_count, session_fields, timeout}), optional) - Session-Aware APF
    * `alias` (String, required)
    * `bidi` (Bool, optional) - include reverse traffic for the selected sessionFields. Not applicable for 'mplsLabel', 'gtpuTeid', 'vlanId'
    * `buffering` (Object({buffer_count_before_match, enabled, protocol}), optional) - Session-Aware APF Buffereing settings
      * `buffer_count_before_match` (Number, optional) - Maximum number of packets BSAPF will buffer per session before APF match
      * `enabled` (Bool, required)
      * `protocol` (String, optional) - When buffering is enabled, changing protocol requires reboot before new value takes effect.
    * `cluster_id` (String, optional) - id of the defining cluster
    * `packet_count` (Number, optional) - For each session, forward this number of packets. Valid range is 2-100; 0 disables packetCount.
    * `session_fields` (Set(Object({pos, type})), required) - A sessionField cannot contain overlapping session attribute, position pairs (e.g. ipv4-5tuple pos 1 and ipv4-src pos 1 is not allowed). Max number of session fields is 20
    * `timeout` (Number, optional) - in seconds
  * `sapf_profile_alias` (String, optional)
* `associated_monitor_solution_alias` (String, optional) - monitor object alias
* `cluster_id` (String, optional) - cluster id for which solution is getting created
* `config_status` (String, optional) - configuration status
* `delete_monitor_sol` (Bool, optional)
* `egress_map_aliases_to_delete` (List(String), optional)
* `exporter_aliases_to_delete` (List(String), optional)
* `ingress_map_aliases_to_delete` (List(String), optional)
* `ingress_traffic_configs` (List(Object({ingress_traffic_map, map_alias})), optional) - ingress traffic configuration
* `monitor_solution_config` (Object({action, cluster_id, config_status, error_message, exporter_alias, exporter_config, gs_group, health_state, health_state_reasons, monitor_ip_interface, ref_sols, solution_alias, user_defined_application_profile, user_defined_applications, vport_alias, vport_config}), optional)
  * `action` (String, optional)
  * `cluster_id` (String, optional) - cluster id
  * `config_status` (String, optional)
  * `error_message` (String, optional)
  * `exporter_alias` (String, optional) - exporter alias associated with monitor
  * `exporter_config` (Object({alias, application_profiles, cef, description, destination, max_pkt_size, mobility_sam, monitor, netflow, snmp, source, type}), optional)
    * `alias` (String, required)
    * `application_profiles` (List(String), optional) - application profile aliases to attach to the exporter
    * `cef` (Object({active_timeout, inactive_timeout}), optional)
      * `active_timeout` (Number, optional) - in seconds
      * `inactive_timeout` (Number, optional) - in seconds
    * `description` (String, optional)
    * `destination` (Object({dscp, ipv4_address, l4_port_dst, l4_port_src, l4_protocol, ttl}), optional)
      * `dscp` (Number, optional)
      * `ipv4_address` (String, optional) - ipv4 address
      * `l4_port_dst` (Number, optional)
      * `l4_port_src` (Number, optional)
      * `l4_protocol` (String, optional)
      * `ttl` (Number, optional)
    * `max_pkt_size` (Number, optional)
    * `mobility_sam` (Object({encoding, encoding_format, event_enable, trigger}), optional)
      * `encoding` (String, optional)
      * `encoding_format` (String, optional)
      * `event_enable` (Object({modify, update}), optional)
        * `modify` (Bool, optional)
        * `update` (Bool, optional)
      * `trigger` (String, optional)
    * `monitor` (Object({timeout}), optional)
      * `timeout` (Number, optional) - how often to export in seconds
    * `netflow` (Object({active_timeout, inactive_timeout, template_refresh, template_type, version}), optional)
      * `active_timeout` (Number, optional) - in seconds
      * `inactive_timeout` (Number, optional) - in seconds
      * `template_refresh` (Number, optional) - template refresh interval in seconds
      * `template_type` (String, optional)
      * `version` (String, optional)
    * `snmp` (Object({enabled}), optional)
      * `enabled` (Bool, optional) - snmp reverse lookup enable/disable
    * `source` (Object({ip_interface}), optional)
      * `ip_interface` (String, optional)
    * `type` (String, optional)
  * `gs_group` (String, optional) - gsgroup alias associated with monitor
  * `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
  * `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
  * `monitor_ip_interface` (String, optional) - ip interface alias associated with monitor
  * `ref_sols` (List(Object({alias, associated_map, type})), optional)
  * `solution_alias` (String, optional) - solution alias
  * `user_defined_application_profile` (Dynamic, computed)
  * `user_defined_applications` (List(List(String)), computed)
  * `vport_alias` (String, optional) - vport alias associated with monitor
  * `vport_config` (Object({alias, deferred_binding, fail_over_action, gs_group, health_state, health_state_reasons, inline_status, inner_traffic_path, metadata_monitoring, mode, outer_traffic_path, sa_apf_profile}), optional) - GigaSMART vPort
    * `alias` (String, required)
    * `deferred_binding` (Bool, optional) - enable/disable deferred-binding
    * `fail_over_action` (String, optional)
    * `gs_group` (String, required) - Alias of referenced managing GsGroup
    * `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
    * `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
    * `inline_status` (String, optional)
    * `inner_traffic_path` (String, optional) - Similar to inline-network traffic-path, applicable for inner map
    * `metadata_monitoring` (Object({action, exporters}), optional)
      * `action` (String, optional) - metadata monitoring action
      * `exporters` (List(String), optional)
    * `mode` (String, optional)
    * `outer_traffic_path` (String, optional) - Similar to inline-tool/inline-tool-group flex-traffic path, applicable for outer map
    * `sa_apf_profile` (String, optional) - ASF session profile
* `solution_alias` (String, optional) - user defined application visibility solution alias
* `solution_desc` (String, optional) - user defined application visibility solution description
* `solution_status` (String, optional) - solution status
* `solution_type` (String, optional) - user intent solution type

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `app_export_config` (Object({cache_config, cache_config_alias, destination_configs, gsop_alias, gsop_config}), computed)
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
  * `cache_config_alias` (String, optional)
  * `destination_configs` (List(Object({application_names, destination_name, export_ip_interface, export_meta_app_profile, export_meta_app_profile_alias, exporter_alias, exporter_config})), optional)
  * `gsop_alias` (String, optional)
  * `gsop_config` (Object({alias, cluster_id, gs_apps, gs_group, health_state, health_state_reasons}), optional) - GigaSMART Operation
    * `alias` (String, required)
    * `cluster_id` (String, optional) - id of the defining cluster
    * `gs_apps` (Object({apf, dedup, diameter_whitelist, flow_filter, flow_sampling, gseries_header_add, gseries_header_remove, gseries_load_balance, gseries_pattern_match, gtp_whitelist, header_add, header_remove, icap, inline_ssl, load_balance, masking, metadata_export, netflow, sa_apf, sip_whitelist, slicing, ssl_decrypt, trailer_add, trailer_remove, tunnel_decap, tunnel_encap}), required) - GigaSMART Applications for a GSOP. At least one GsApp must be defined for the config to be valid
      * `apf` (Object({enabled}), optional)
        * `enabled` (String, optional)
      * `dedup` (Object({enabled}), optional)
        * `enabled` (String, optional)
      * `diameter_whitelist` (Object({enabled}), optional)
        * `enabled` (String, optional)
      * `flow_filter` (Object({type}), optional) - GigaSMART 'Flow Filter' Application Configuration
        * `type` (String, required)
      * `flow_sampling` (Object({type}), optional) - GigaSMART 'Flow Sampling' Application Configuration
        * `type` (String, required)
      * `gseries_header_add` (Object({types}), optional) - Only applicable for G-series
        * `types` (Set(String), required)
      * `gseries_header_remove` (Object({enabled}), optional)
        * `enabled` (String, optional)
      * `gseries_load_balance` (Object({fixed_offset, variable_offset}), optional) - Only applicable for G-series per-rule GSOP
        * `fixed_offset` (Object({hash, length, offset}), optional) - Per-rule Fixed Offset Load Balancing config for G-seres devices
          * `hash` (String, required)
          * `length` (Number, required)
          * `offset` (Number, required)
        * `variable_offset` (Object({end_delim, hash, start_delim, start_field}), optional) - Per-rule Variable Offset Load Balancing config for G-seres devices
          * `end_delim` (String, required)
          * `hash` (String, required)
          * `start_delim` (String, required)
          * `start_field` (String, required)
      * `gseries_pattern_match` (Object({fixed_offset, variable_offset}), optional) - Only applicable for G-series per-rule GSOP
        * `fixed_offset` (Object({length, offset}), optional) - Per-rule Fixed Offset Pattern Match config for G-seres devices
          * `length` (Number, required)
          * `offset` (Number, required)
        * `variable_offset` (Object({end_delim, start_delim}), optional) - Per-rule Variable Offset Pattern Match config for G-seres devices
          * `end_delim` (String, required)
          * `start_delim` (String, required)
      * `gtp_whitelist` (Object({enabled}), optional)
        * `enabled` (String, optional)
      * `header_add` (Object({vlan}), optional) - GigaSMART 'Add Header' Application Configuration
        * `vlan` (Number, required)
      * `header_remove` (Object({ah1, ah2, custom_len, erspan_flow_id, fp_dst_switch_id, fp_src_switch_id, header_count, offset, offset_range_value, protocol, timestamp_format, vlan_header, vxlan_id}), optional) - GigaSMART 'Remove Header' Application Configuration
        * `ah1` (String, optional) - only valid and required for 'generic'. First anchor header, from which the header to be stripped is occurred.
        * `ah2` (String, optional) - only valid and required for 'generic'. next anchor header.
        * `custom_len` (Number, optional) - only valid for 'generic'. length of unknown header.
        * `erspan_flow_id` (Number, optional) - only applicable for 'erspan', specifies the flow id to strip. Value of 0 will strip all flow ids
        * `fp_dst_switch_id` (Number, optional) - Only valid and required for 'fabricPath', 12 bit destination switch id
        * `fp_src_switch_id` (Number, optional) - Only valid and required for 'fabricPath', 12 bit source switch id
        * `header_count` (Number, optional) - only valid for 'generic'. Number of headers to be stripped.
        * `offset` (String, optional) - only valid and required for 'generic'.'start': strip from start of ah1, 'end': strip from end of ah1 or any other offset in the range of length of ah1, 'offsetRangeValue' is required if offset is 'offsetRange'.
        * `offset_range_value` (Number, optional) - only valid and required when offset is 'offsetRange', integer within range of size of header
        * `protocol` (String, required) - 'gre' and 'fabricPath' are only applicable for H-series
        * `timestamp_format` (String, optional) - Timestamp format. Only valid and required for 'fm6000Ts'
        * `vlan_header` (String, optional) - Only valid when protocol is 'vlan'. Specifies the target vlan header to strip. Defaults to 'outer'
        * `vxlan_id` (Number, optional) - 24-bit value. Only valid when protocol is 'vxlan'. Specifies the vxlan id to strip. Value of '0' will strip all vxlan ids
      * `icap` (Object({icap_profile}), optional) - GigaSMART ICAP Configuration
        * `icap_profile` (String, required) - Alias of referenced ICAP Profile
      * `inline_ssl` (Object({inline_ssl_profile}), optional) - GigaSMART Inline SSL Profile Configuration
        * `inline_ssl_profile` (String, required) - Alias of referenced Inline SSL Profile
      * `load_balance` (Object({enhanced, stateful, stateless}), optional) - GigaSMART 'Load Balancing' Application Configuration
        * `enhanced` (Object({elb_alias}), optional) - Private class. Enhanced part of the GigaSMART 'Load Balancing' Application Configuration
          * `elb_alias` (String, required) - elb app alias
        * `stateful` (Object({app_type, diameter_key_hash_type, diameter_key_multi_hash_type, gtp_key_hash_type, lb_type, sip_key_hash_type}), optional) - Private class. Stateful part of the GigaSMART 'Load Balancing' Application Configuration
          * `app_type` (String, required) - 'sapf' option is added in release H 4.3.01. 'sip' and 'tunnel' added in release H 5.1. 'diameter' is added in release H 5.5
          * `diameter_key_hash_type` (String, optional) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyHash'
          * `diameter_key_multi_hash_type` (List(Object({avp_codevalue, key})), optional) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyMultiHash'
          * `gtp_key_hash_type` (String, optional) - required when 'appType' == 'gtp' and 'lbType' == 'gtpKeyHash'. ignored otherwise
          * `lb_type` (String, required) - 'gtpKeyHash' is only applicable when 'appType' == 'gtp'. 'sipKeyHash' is only applicable when 'appType' == 'sip'
          * `sip_key_hash_type` (String, optional) - required when 'appType' == 'sip' and 'lbType' == 'sipKeyHash'. ignored otherwise
        * `stateless` (Object({field_location, hash_fields}), optional) - Private class. Stateless part of the GigaSMART 'Load Balancing' Application Configuration
          * `field_location` (String, optional) - Ignored when 'hashFields' == 'gtpuTeid'. Required otherwise
          * `hash_fields` (String, required)
      * `masking` (Object({content_type, length, offset, pattern, protocol}), optional) - GigaSMART 'Masking' Application Configuration
        * `content_type` (String, optional) - content type that will trigger masking, only valid and required for protocol 'sip'
        * `length` (Number, optional) - max length is 9600 on H-series; 9000 on G-series, not applicable for protocol 'sip' valid and required otherwise
        * `offset` (Number, optional) - min offset is 0 for 'none'; 1 for protocol-referenced offsets, not applicable for protocol 'sip' valid and required otherwise
        * `pattern` (String, optional) - 1-byte hex mask, not applicable for protocol 'sip' valid and required otherwise
        * `protocol` (String, required)
      * `metadata_export` (Object({cache}), optional)
        * `cache` (String, optional) - metadata cache alias. cache should have exporters defined, and the exporters should have applicationProfiles defined
      * `netflow` (Object({enabled}), optional)
        * `enabled` (String, optional)
      * `sa_apf` (Object({enabled}), optional)
        * `enabled` (String, optional)
      * `sip_whitelist` (Object({enabled}), optional)
        * `enabled` (String, optional)
      * `slicing` (Object({enhanced, offset, protocol}), optional) - GigaSMART 'Slicing' Application Configuration
        * `enhanced` (String, optional) - enhanced-slicing apps alias
        * `offset` (Number, required) - min offset is 64 for 'none'; 4 for protocol-referenced offsets. required property till H 5.6
        * `protocol` (String, required) - required property till H 5.6
      * `ssl_decrypt` (Object({in_port, out_port}), optional) - GigaSMART 'SSL Decrypt' Application Configuration
        * `in_port` (Number, optional) - Port number of 0 represents 'any' port
        * `out_port` (Number, optional) - Port number of 0 represents 'auto' port
      * `trailer_add` (Object({types}), optional) - GigaSMART 'Add Trailer' Application Configuration
        * `types` (Set(String), required) - 'crc' is not applicable for G-series
      * `trailer_remove` (Object({enabled}), optional)
        * `enabled` (String, optional)
      * `tunnel_decap` (Object({custom, erspan_flow_id, gmip_port, l2_gre_key, tls_pcapng, type, vxlan}), optional) - GigaSMART 'Decapsulate Tunnel' Application Configuration
        * `custom` (Object({port_dst, port_src}), optional)
          * `port_dst` (Number, required) - when specified as 0, no validation will be done in the packet.
          * `port_src` (Number, required) - when specified as 0, no validation will be done in the packet.
        * `erspan_flow_id` (Number, optional) - only applicable for 'erspan', A Flow ID of 0 decapsulates all ERSPAN tunnel traffic regardless of Flow ID
        * `gmip_port` (Number, optional) - only applicable for 'gmip', in which case it is required. Specifies the UDP port on which the Tunnel Network port on the receiving GigaVUE H Series is listening. Must match the configuration of the portdst configured on the sending end of the tunnel
        * `l2_gre_key` (Number, optional) - only applicable for 'l2gre', in which case it is required.
        * `tls_pcapng` (Object({decap_key, listener}), optional)
          * `decap_key` (String, optional)
          * `listener` (String, optional)
        * `type` (String, required)
        * `vxlan` (Object({port_dst, port_src, vni}), optional)
          * `port_dst` (Number, required)
          * `port_src` (Number, required) - when specified as 0, no validation will be done in the packet.
          * `vni` (Number, required)
      * `tunnel_encap` (Object({gmip_config, l2_gre_config, tls_pcapng, type, vxlan_config}), optional) - GigaSMART 'Encapsulate Tunnel' Application Configuration
        * `gmip_config` (Object({dscp, dst_ip, dst_port, flow_label, prec, src_port, ttl}), optional) - Configuration for GMIP Tunnel Encapsulate GigaSmaprt App
          * `dscp` (Number, optional)
          * `dst_ip` (String, required) - IP Destination. IPv4 or IPv6.
          * `dst_port` (Number, required)
          * `flow_label` (Number, optional)
          * `prec` (Number, optional) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
          * `src_port` (Number, required)
          * `ttl` (Number, optional)
        * `l2_gre_config` (Object({dscp, dst_ip, flow_label, key, pg_dst, prec, session_field, session_pos, ttl}), optional) - Configuration for l2GRE Tunnel Encapsulate GigaSmaprt App
          * `dscp` (Number, optional)
          * `dst_ip` (String, optional) - ip destination, mutually exclusive with 'pgDst'. IPv4 or IPv6.
          * `flow_label` (Number, optional)
          * `key` (Number, required)
          * `pg_dst` (String, optional) - port group destination alias, mutually exclusive with 'dstIp'
          * `prec` (Number, optional) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
          * `session_field` (String, optional) - required with stateful loadBalance when 'appType' is 'tunnel'
          * `session_pos` (String, optional) - required if 'sessionField' is specified
          * `ttl` (Number, optional)
        * `tls_pcapng` (Object({exporter, exporter_group}), optional)
          * `exporter` (String, optional)
          * `exporter_group` (String, optional)
        * `type` (String, required)
        * `vxlan_config` (Object({dscp, dst_ip, dst_port, src_port, ttl, vni}), optional) - Configuration for Vxlan Tunnel Encapsulate GigaSmaprt App
          * `dscp` (Number, optional)
          * `dst_ip` (String, required) - IP Destination. IPv4 or IPv6.
          * `dst_port` (Number, required)
          * `src_port` (Number, required)
          * `ttl` (Number, optional)
          * `vni` (Number, required)
    * `gs_group` (String, required) - Alias of referenced managing GsGroup
    * `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
    * `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), optional)
* `app_filter_config` (Object({egress_traffic_config, sapf_profile, sapf_profile_alias}), computed)
  * `egress_traffic_config` (List(Object({drop_app_profile_alias, egress_traffic_map, gs_apps, gsop_alias, gsop_config, map_alias, pass_app_profile_alias, priority})), optional)
  * `sapf_profile` (Object({alias, bidi, buffering, cluster_id, packet_count, session_fields, timeout}), optional) - Session-Aware APF
    * `alias` (String, required)
    * `bidi` (Bool, optional) - include reverse traffic for the selected sessionFields. Not applicable for 'mplsLabel', 'gtpuTeid', 'vlanId'
    * `buffering` (Object({buffer_count_before_match, enabled, protocol}), optional) - Session-Aware APF Buffereing settings
      * `buffer_count_before_match` (Number, optional) - Maximum number of packets BSAPF will buffer per session before APF match
      * `enabled` (Bool, required)
      * `protocol` (String, optional) - When buffering is enabled, changing protocol requires reboot before new value takes effect.
    * `cluster_id` (String, optional) - id of the defining cluster
    * `packet_count` (Number, optional) - For each session, forward this number of packets. Valid range is 2-100; 0 disables packetCount.
    * `session_fields` (Set(Object({pos, type})), required) - A sessionField cannot contain overlapping session attribute, position pairs (e.g. ipv4-5tuple pos 1 and ipv4-src pos 1 is not allowed). Max number of session fields is 20
    * `timeout` (Number, optional) - in seconds
  * `sapf_profile_alias` (String, optional)
* `associated_monitor_solution_alias` (String, computed) - monitor object alias
* `cluster_id` (String, computed) - cluster id for which solution is getting created
* `config_status` (String, computed) - configuration status
* `delete_monitor_sol` (Bool, computed)
* `egress_map_aliases_to_delete` (List(String), computed)
* `exporter_aliases_to_delete` (List(String), computed)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `ingress_map_aliases_to_delete` (List(String), computed)
* `ingress_traffic_configs` (List(Object({ingress_traffic_map, map_alias})), computed) - ingress traffic configuration
* `monitor_solution_config` (Object({action, cluster_id, config_status, error_message, exporter_alias, exporter_config, gs_group, health_state, health_state_reasons, monitor_ip_interface, ref_sols, solution_alias, user_defined_application_profile, user_defined_applications, vport_alias, vport_config}), computed)
  * `action` (String, optional)
  * `cluster_id` (String, optional) - cluster id
  * `config_status` (String, optional)
  * `error_message` (String, optional)
  * `exporter_alias` (String, optional) - exporter alias associated with monitor
  * `exporter_config` (Object({alias, application_profiles, cef, description, destination, max_pkt_size, mobility_sam, monitor, netflow, snmp, source, type}), optional)
    * `alias` (String, required)
    * `application_profiles` (List(String), optional) - application profile aliases to attach to the exporter
    * `cef` (Object({active_timeout, inactive_timeout}), optional)
      * `active_timeout` (Number, optional) - in seconds
      * `inactive_timeout` (Number, optional) - in seconds
    * `description` (String, optional)
    * `destination` (Object({dscp, ipv4_address, l4_port_dst, l4_port_src, l4_protocol, ttl}), optional)
      * `dscp` (Number, optional)
      * `ipv4_address` (String, optional) - ipv4 address
      * `l4_port_dst` (Number, optional)
      * `l4_port_src` (Number, optional)
      * `l4_protocol` (String, optional)
      * `ttl` (Number, optional)
    * `max_pkt_size` (Number, optional)
    * `mobility_sam` (Object({encoding, encoding_format, event_enable, trigger}), optional)
      * `encoding` (String, optional)
      * `encoding_format` (String, optional)
      * `event_enable` (Object({modify, update}), optional)
        * `modify` (Bool, optional)
        * `update` (Bool, optional)
      * `trigger` (String, optional)
    * `monitor` (Object({timeout}), optional)
      * `timeout` (Number, optional) - how often to export in seconds
    * `netflow` (Object({active_timeout, inactive_timeout, template_refresh, template_type, version}), optional)
      * `active_timeout` (Number, optional) - in seconds
      * `inactive_timeout` (Number, optional) - in seconds
      * `template_refresh` (Number, optional) - template refresh interval in seconds
      * `template_type` (String, optional)
      * `version` (String, optional)
    * `snmp` (Object({enabled}), optional)
      * `enabled` (Bool, optional) - snmp reverse lookup enable/disable
    * `source` (Object({ip_interface}), optional)
      * `ip_interface` (String, optional)
    * `type` (String, optional)
  * `gs_group` (String, optional) - gsgroup alias associated with monitor
  * `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
  * `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
  * `monitor_ip_interface` (String, optional) - ip interface alias associated with monitor
  * `ref_sols` (List(Object({alias, associated_map, type})), optional)
  * `solution_alias` (String, optional) - solution alias
  * `user_defined_application_profile` (Dynamic, computed)
  * `user_defined_applications` (List(List(String)), computed)
  * `vport_alias` (String, optional) - vport alias associated with monitor
  * `vport_config` (Object({alias, deferred_binding, fail_over_action, gs_group, health_state, health_state_reasons, inline_status, inner_traffic_path, metadata_monitoring, mode, outer_traffic_path, sa_apf_profile}), optional) - GigaSMART vPort
    * `alias` (String, required)
    * `deferred_binding` (Bool, optional) - enable/disable deferred-binding
    * `fail_over_action` (String, optional)
    * `gs_group` (String, required) - Alias of referenced managing GsGroup
    * `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
    * `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
    * `inline_status` (String, optional)
    * `inner_traffic_path` (String, optional) - Similar to inline-network traffic-path, applicable for inner map
    * `metadata_monitoring` (Object({action, exporters}), optional)
      * `action` (String, optional) - metadata monitoring action
      * `exporters` (List(String), optional)
    * `mode` (String, optional)
    * `outer_traffic_path` (String, optional) - Similar to inline-tool/inline-tool-group flex-traffic path, applicable for outer map
    * `sa_apf_profile` (String, optional) - ASF session profile
* `solution_alias` (String, computed) - user defined application visibility solution alias
* `solution_desc` (String, computed) - user defined application visibility solution description
* `solution_status` (String, computed) - solution status
* `solution_type` (String, computed) - user intent solution type

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_solution.example {solution_alias}
```
