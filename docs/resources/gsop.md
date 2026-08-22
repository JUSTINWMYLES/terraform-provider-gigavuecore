---
page_title: "gigavuecore_gsop Resource - gigavuecore"
subcategory: ""
description: |-
  Find GSOP by alias
---

# gigavuecore_gsop Resource

Find GSOP by alias

## Example Usage

```terraform
resource "gigavuecore_gsop" "example" {
  alias = null
  cluster_id = null
  gs_apps = {}
  gs_group = null
  health_state = null
  health_state_reasons = []
}
```

## Schema

### Arguments

The following arguments are supported:

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

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `cluster_id` (String, computed) - id of the defining cluster
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_gsop.example {alias}
```
