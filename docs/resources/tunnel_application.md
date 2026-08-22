---
page_title: "gigavuecore_tunnel_application Resource - gigavuecore"
subcategory: ""
description: |-
  Get Tunnel App for given alias
---

# gigavuecore_tunnel_application Resource

Get Tunnel App for given alias

## Example Usage

```terraform
resource "gigavuecore_tunnel_application" "example" {
  alias = null
  cluster_id = null
  config_status = null
  config_status_reasons = []
  decap = []
  description = null
  dscp = null
  encap = {}
  gsgroup = null
  ip_interface_in = null
  ip_interface_out = null
  local_key_alias = null
  remote_key_alias = null
  ssl_profile = {}
  ssl_profile_alias = null
  tcp_profile = {}
  tcp_profile_alias = null
  traffic_dir = null
  ttl = null
  tunnel_type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional)
* `cluster_id` (String, optional)
* `config_status` (String, optional)
* `config_status_reasons` (List(String), optional)
* `decap` (List(Object({additonalgsops, application_ports, destination_port, gsop_alias, listener_alias, map_alias, rules})), optional)
* `description` (String, optional)
* `dscp` (Number, optional)
* `encap` (Object({additonalgsops, export_configs, exporter_group_alias, gsop_alias, map_alias, rules, source_port}), optional)
  * `additonalgsops` (Object({apf, dedup, diameter_whitelist, flow_filter, flow_sampling, gseries_header_add, gseries_header_remove, gseries_load_balance, gseries_pattern_match, gtp_whitelist, header_add, header_remove, icap, inline_ssl, load_balance, masking, metadata_export, netflow, sa_apf, sip_whitelist, slicing, ssl_decrypt, trailer_add, trailer_remove, tunnel_decap, tunnel_encap}), optional) - GigaSMART Applications for a GSOP. At least one GsApp must be defined for the config to be valid
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
  * `export_configs` (List(Object({exporter_alias, remote_application_port, remote_ip, source_application_port})), optional)
  * `exporter_group_alias` (String, optional)
  * `gsop_alias` (String, optional)
  * `map_alias` (String, optional)
  * `rules` (Object({drop_rules, pass_rules}), optional) - Map Rules Container. Private class
    * `drop_rules` (Set(Object({bidi, comment, ip_rewrite, matches, rewrite, rule_id, vlan_tag})), optional)
    * `pass_rules` (Set(Object({bidi, comment, ip_rewrite, matches, rewrite, rule_id, vlan_tag})), optional)
  * `source_port` (List(String), optional) - source port for encap traffic , it would be on the same cluster as GsEngine
* `gsgroup` (String, optional)
* `ip_interface_in` (String, optional) - ipInterface for Decap source
* `ip_interface_out` (String, optional) - ipInterface for Encap destination
* `local_key_alias` (String, optional)
* `remote_key_alias` (String, optional)
* `ssl_profile` (Object({cipher, mtls, version}), optional)
  * `cipher` (String, optional)
  * `mtls` (String, optional)
  * `version` (String, optional)
* `ssl_profile_alias` (String, optional)
* `tcp_profile` (Object({keep_alive_timer, selective_ack, syn_retries}), optional)
  * `keep_alive_timer` (Number, optional)
  * `selective_ack` (String, optional)
  * `syn_retries` (Number, optional)
* `tcp_profile_alias` (String, optional)
* `traffic_dir` (String, optional) - IN for Decap , OUT for Encap
* `ttl` (Number, optional)
* `tunnel_type` (String, optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `alias` (String, computed)
* `cluster_id` (String, computed)
* `config_status` (String, computed)
* `config_status_reasons` (List(String), computed)
* `decap` (List(Object({additonalgsops, application_ports, destination_port, gsop_alias, listener_alias, map_alias, rules})), computed)
* `description` (String, computed)
* `dscp` (Number, computed)
* `encap` (Object({additonalgsops, export_configs, exporter_group_alias, gsop_alias, map_alias, rules, source_port}), computed)
  * `additonalgsops` (Object({apf, dedup, diameter_whitelist, flow_filter, flow_sampling, gseries_header_add, gseries_header_remove, gseries_load_balance, gseries_pattern_match, gtp_whitelist, header_add, header_remove, icap, inline_ssl, load_balance, masking, metadata_export, netflow, sa_apf, sip_whitelist, slicing, ssl_decrypt, trailer_add, trailer_remove, tunnel_decap, tunnel_encap}), optional) - GigaSMART Applications for a GSOP. At least one GsApp must be defined for the config to be valid
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
  * `export_configs` (List(Object({exporter_alias, remote_application_port, remote_ip, source_application_port})), optional)
  * `exporter_group_alias` (String, optional)
  * `gsop_alias` (String, optional)
  * `map_alias` (String, optional)
  * `rules` (Object({drop_rules, pass_rules}), optional) - Map Rules Container. Private class
    * `drop_rules` (Set(Object({bidi, comment, ip_rewrite, matches, rewrite, rule_id, vlan_tag})), optional)
    * `pass_rules` (Set(Object({bidi, comment, ip_rewrite, matches, rewrite, rule_id, vlan_tag})), optional)
  * `source_port` (List(String), optional) - source port for encap traffic , it would be on the same cluster as GsEngine
* `gsgroup` (String, computed)
* `ip_interface_in` (String, computed) - ipInterface for Decap source
* `ip_interface_out` (String, computed) - ipInterface for Encap destination
* `local_key_alias` (String, computed)
* `remote_key_alias` (String, computed)
* `ssl_profile` (Object({cipher, mtls, version}), computed)
  * `cipher` (String, optional)
  * `mtls` (String, optional)
  * `version` (String, optional)
* `ssl_profile_alias` (String, computed)
* `tcp_profile` (Object({keep_alive_timer, selective_ack, syn_retries}), computed)
  * `keep_alive_timer` (Number, optional)
  * `selective_ack` (String, optional)
  * `syn_retries` (Number, optional)
* `tcp_profile_alias` (String, computed)
* `traffic_dir` (String, computed) - IN for Decap , OUT for Encap
* `ttl` (Number, computed)
* `tunnel_type` (String, computed)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_tunnel_application.example {alias}
```
