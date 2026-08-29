---
page_title: "gigavuecore_get_all_tunnel_application Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all Tunnel Apps
---

# gigavuecore_get_all_tunnel_application Data Source

Get all Tunnel Apps

## Example Usage

```terraform
data "gigavuecore_get_all_tunnel_application" "example" {
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String)
* `cluster_id` (String)
* `config_status` (String)
* `config_status_reasons` (List of String)
* `decap` (Attributes List) (see [below for nested schema](#nestedatt--items--decap))
* `description` (String)
* `dscp` (Number)
* `encap` (Attributes) (see [below for nested schema](#nestedatt--items--encap))
* `gsgroup` (String)
* `ip_interface_in` (String) - ipInterface for Decap source
* `ip_interface_out` (String) - ipInterface for Encap destination
* `local_key_alias` (String)
* `remote_key_alias` (String)
* `ssl_profile` (Attributes) (see [below for nested schema](#nestedatt--items--ssl_profile))
* `ssl_profile_alias` (String)
* `tcp_profile` (Attributes) (see [below for nested schema](#nestedatt--items--tcp_profile))
* `tcp_profile_alias` (String)
* `traffic_dir` (String) - IN for Decap , OUT for Encap
* `ttl` (Number)
* `tunnel_type` (String)
<a id="nestedatt--items--decap"></a>
### Nested Schema for `items.decap`

Read-Only:

* `additonalgsops` (Attributes) - GigaSMART Applications for a GSOP. At least one GsApp must be defined for the config to be valid (see [below for nested schema](#nestedatt--items--decap--additonalgsops))
* `application_ports` (List of Number)
* `destination_port` (List of String) - Tools connected to the cluster
* `gsop_alias` (String)
* `listener_alias` (String)
* `map_alias` (String)
* `rules` (Attributes) - Map Rules Container. Private class (see [below for nested schema](#nestedatt--items--decap--rules))
<a id="nestedatt--items--decap--additonalgsops"></a>
### Nested Schema for `items.decap.additonalgsops`

Read-Only:

* `apf` (Attributes) (see [below for nested schema](#nestedatt--items--decap--additonalgsops--apf))
* `dedup` (Attributes) (see [below for nested schema](#nestedatt--items--decap--additonalgsops--dedup))
* `diameter_whitelist` (Attributes) (see [below for nested schema](#nestedatt--items--decap--additonalgsops--diameter_whitelist))
* `flow_filter` (Attributes) - GigaSMART 'Flow Filter' Application Configuration (see [below for nested schema](#nestedatt--items--decap--additonalgsops--flow_filter))
* `flow_sampling` (Attributes) - GigaSMART 'Flow Sampling' Application Configuration (see [below for nested schema](#nestedatt--items--decap--additonalgsops--flow_sampling))
* `gseries_header_add` (Attributes) - Only applicable for G-series (see [below for nested schema](#nestedatt--items--decap--additonalgsops--gseries_header_add))
* `gseries_header_remove` (Attributes) (see [below for nested schema](#nestedatt--items--decap--additonalgsops--gseries_header_remove))
* `gseries_load_balance` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--items--decap--additonalgsops--gseries_load_balance))
* `gseries_pattern_match` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--items--decap--additonalgsops--gseries_pattern_match))
* `gtp_whitelist` (Attributes) (see [below for nested schema](#nestedatt--items--decap--additonalgsops--gtp_whitelist))
* `header_add` (Attributes) - GigaSMART 'Add Header' Application Configuration (see [below for nested schema](#nestedatt--items--decap--additonalgsops--header_add))
* `header_remove` (Attributes) - GigaSMART 'Remove Header' Application Configuration (see [below for nested schema](#nestedatt--items--decap--additonalgsops--header_remove))
* `icap` (Attributes) - GigaSMART ICAP Configuration (see [below for nested schema](#nestedatt--items--decap--additonalgsops--icap))
* `inline_ssl` (Attributes) - GigaSMART Inline SSL Profile Configuration (see [below for nested schema](#nestedatt--items--decap--additonalgsops--inline_ssl))
* `load_balance` (Attributes) - GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--items--decap--additonalgsops--load_balance))
* `masking` (Attributes) - GigaSMART 'Masking' Application Configuration (see [below for nested schema](#nestedatt--items--decap--additonalgsops--masking))
* `metadata_export` (Attributes) (see [below for nested schema](#nestedatt--items--decap--additonalgsops--metadata_export))
* `netflow` (Attributes) (see [below for nested schema](#nestedatt--items--decap--additonalgsops--netflow))
* `sa_apf` (Attributes) (see [below for nested schema](#nestedatt--items--decap--additonalgsops--sa_apf))
* `sip_whitelist` (Attributes) (see [below for nested schema](#nestedatt--items--decap--additonalgsops--sip_whitelist))
* `slicing` (Attributes) - GigaSMART 'Slicing' Application Configuration (see [below for nested schema](#nestedatt--items--decap--additonalgsops--slicing))
* `ssl_decrypt` (Attributes) - GigaSMART 'SSL Decrypt' Application Configuration (see [below for nested schema](#nestedatt--items--decap--additonalgsops--ssl_decrypt))
* `trailer_add` (Attributes) - GigaSMART 'Add Trailer' Application Configuration (see [below for nested schema](#nestedatt--items--decap--additonalgsops--trailer_add))
* `trailer_remove` (Attributes) (see [below for nested schema](#nestedatt--items--decap--additonalgsops--trailer_remove))
* `tunnel_decap` (Attributes) - GigaSMART 'Decapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--items--decap--additonalgsops--tunnel_decap))
* `tunnel_encap` (Attributes) - GigaSMART 'Encapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--items--decap--additonalgsops--tunnel_encap))
<a id="nestedatt--items--decap--additonalgsops--apf"></a>
### Nested Schema for `items.decap.additonalgsops.apf`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--decap--additonalgsops--dedup"></a>
### Nested Schema for `items.decap.additonalgsops.dedup`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--decap--additonalgsops--diameter_whitelist"></a>
### Nested Schema for `items.decap.additonalgsops.diameter_whitelist`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--decap--additonalgsops--flow_filter"></a>
### Nested Schema for `items.decap.additonalgsops.flow_filter`

Read-Only:

* `type` (String)
<a id="nestedatt--items--decap--additonalgsops--flow_sampling"></a>
### Nested Schema for `items.decap.additonalgsops.flow_sampling`

Read-Only:

* `type` (String)
<a id="nestedatt--items--decap--additonalgsops--gseries_header_add"></a>
### Nested Schema for `items.decap.additonalgsops.gseries_header_add`

Read-Only:

* `types` (Set of String)
<a id="nestedatt--items--decap--additonalgsops--gseries_header_remove"></a>
### Nested Schema for `items.decap.additonalgsops.gseries_header_remove`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--decap--additonalgsops--gseries_load_balance"></a>
### Nested Schema for `items.decap.additonalgsops.gseries_load_balance`

Read-Only:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--items--decap--additonalgsops--gseries_load_balance--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--items--decap--additonalgsops--gseries_load_balance--variable_offset))
<a id="nestedatt--items--decap--additonalgsops--gseries_load_balance--fixed_offset"></a>
### Nested Schema for `items.decap.additonalgsops.gseries_load_balance.fixed_offset`

Read-Only:

* `hash` (String)
* `length` (Number)
* `offset` (Number)
<a id="nestedatt--items--decap--additonalgsops--gseries_load_balance--variable_offset"></a>
### Nested Schema for `items.decap.additonalgsops.gseries_load_balance.variable_offset`

Read-Only:

* `end_delim` (String)
* `hash` (String)
* `start_delim` (String)
* `start_field` (String)
<a id="nestedatt--items--decap--additonalgsops--gseries_pattern_match"></a>
### Nested Schema for `items.decap.additonalgsops.gseries_pattern_match`

Read-Only:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--items--decap--additonalgsops--gseries_pattern_match--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--items--decap--additonalgsops--gseries_pattern_match--variable_offset))
<a id="nestedatt--items--decap--additonalgsops--gseries_pattern_match--fixed_offset"></a>
### Nested Schema for `items.decap.additonalgsops.gseries_pattern_match.fixed_offset`

Read-Only:

* `length` (Number)
* `offset` (Number)
<a id="nestedatt--items--decap--additonalgsops--gseries_pattern_match--variable_offset"></a>
### Nested Schema for `items.decap.additonalgsops.gseries_pattern_match.variable_offset`

Read-Only:

* `end_delim` (String)
* `start_delim` (String)
<a id="nestedatt--items--decap--additonalgsops--gtp_whitelist"></a>
### Nested Schema for `items.decap.additonalgsops.gtp_whitelist`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--decap--additonalgsops--header_add"></a>
### Nested Schema for `items.decap.additonalgsops.header_add`

Read-Only:

* `vlan` (Number)
<a id="nestedatt--items--decap--additonalgsops--header_remove"></a>
### Nested Schema for `items.decap.additonalgsops.header_remove`

Read-Only:

* `ah1` (String) - only valid and required for 'generic'. First anchor header, from which the header to be stripped is occurred.
* `ah2` (String) - only valid and required for 'generic'. next anchor header.
* `custom_len` (Number) - only valid for 'generic'. length of unknown header.
* `erspan_flow_id` (Number) - only applicable for 'erspan', specifies the flow id to strip. Value of 0 will strip all flow ids
* `fp_dst_switch_id` (Number) - Only valid and required for 'fabricPath', 12 bit destination switch id
* `fp_src_switch_id` (Number) - Only valid and required for 'fabricPath', 12 bit source switch id
* `header_count` (Number) - only valid for 'generic'. Number of headers to be stripped.
* `offset` (String) - only valid and required for 'generic'.'start': strip from start of ah1, 'end': strip from end of ah1 or any other offset in the range of length of ah1, 'offsetRangeValue' is required if offset is 'offsetRange'.
* `offset_range_value` (Number) - only valid and required when offset is 'offsetRange', integer within range of size of header
* `protocol` (String) - 'gre' and 'fabricPath' are only applicable for H-series
* `timestamp_format` (String) - Timestamp format. Only valid and required for 'fm6000Ts'
* `vlan_header` (String) - Only valid when protocol is 'vlan'. Specifies the target vlan header to strip. Defaults to 'outer'
* `vxlan_id` (Number) - 24-bit value. Only valid when protocol is 'vxlan'. Specifies the vxlan id to strip. Value of '0' will strip all vxlan ids
<a id="nestedatt--items--decap--additonalgsops--icap"></a>
### Nested Schema for `items.decap.additonalgsops.icap`

Read-Only:

* `icap_profile` (String) - Alias of referenced ICAP Profile
<a id="nestedatt--items--decap--additonalgsops--inline_ssl"></a>
### Nested Schema for `items.decap.additonalgsops.inline_ssl`

Read-Only:

* `inline_ssl_profile` (String) - Alias of referenced Inline SSL Profile
<a id="nestedatt--items--decap--additonalgsops--load_balance"></a>
### Nested Schema for `items.decap.additonalgsops.load_balance`

Read-Only:

* `enhanced` (Attributes) - Private class. Enhanced part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--items--decap--additonalgsops--load_balance--enhanced))
* `stateful` (Attributes) - Private class. Stateful part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--items--decap--additonalgsops--load_balance--stateful))
* `stateless` (Attributes) - Private class. Stateless part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--items--decap--additonalgsops--load_balance--stateless))
<a id="nestedatt--items--decap--additonalgsops--load_balance--enhanced"></a>
### Nested Schema for `items.decap.additonalgsops.load_balance.enhanced`

Read-Only:

* `elb_alias` (String) - elb app alias
<a id="nestedatt--items--decap--additonalgsops--load_balance--stateful"></a>
### Nested Schema for `items.decap.additonalgsops.load_balance.stateful`

Read-Only:

* `app_type` (String) - 'sapf' option is added in release H 4.3.01. 'sip' and 'tunnel' added in release H 5.1. 'diameter' is added in release H 5.5
* `diameter_key_hash_type` (String) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyHash'
* `diameter_key_multi_hash_type` (Attributes List) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyMultiHash' (see [below for nested schema](#nestedatt--items--decap--additonalgsops--load_balance--stateful--diameter_key_multi_hash_type))
* `gtp_key_hash_type` (String) - required when 'appType' == 'gtp' and 'lbType' == 'gtpKeyHash'. ignored otherwise
* `lb_type` (String) - 'gtpKeyHash' is only applicable when 'appType' == 'gtp'. 'sipKeyHash' is only applicable when 'appType' == 'sip'
* `sip_key_hash_type` (String) - required when 'appType' == 'sip' and 'lbType' == 'sipKeyHash'. ignored otherwise
<a id="nestedatt--items--decap--additonalgsops--load_balance--stateful--diameter_key_multi_hash_type"></a>
### Nested Schema for `items.decap.additonalgsops.load_balance.stateful.diameter_key_multi_hash_type`

Read-Only:

* `avp_codevalue` (Number) - required when 'key' == 'avpCode'
* `key` (String)
<a id="nestedatt--items--decap--additonalgsops--load_balance--stateless"></a>
### Nested Schema for `items.decap.additonalgsops.load_balance.stateless`

Read-Only:

* `field_location` (String) - Ignored when 'hashFields' == 'gtpuTeid'. Required otherwise
* `hash_fields` (String)
<a id="nestedatt--items--decap--additonalgsops--masking"></a>
### Nested Schema for `items.decap.additonalgsops.masking`

Read-Only:

* `content_type` (String) - content type that will trigger masking, only valid and required for protocol 'sip'
* `length` (Number) - max length is 9600 on H-series; 9000 on G-series, not applicable for protocol 'sip' valid and required otherwise
* `offset` (Number) - min offset is 0 for 'none'; 1 for protocol-referenced offsets, not applicable for protocol 'sip' valid and required otherwise
* `pattern` (String) - 1-byte hex mask, not applicable for protocol 'sip' valid and required otherwise
* `protocol` (String)
<a id="nestedatt--items--decap--additonalgsops--metadata_export"></a>
### Nested Schema for `items.decap.additonalgsops.metadata_export`

Read-Only:

* `cache` (String) - metadata cache alias. cache should have exporters defined, and the exporters should have applicationProfiles defined
<a id="nestedatt--items--decap--additonalgsops--netflow"></a>
### Nested Schema for `items.decap.additonalgsops.netflow`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--decap--additonalgsops--sa_apf"></a>
### Nested Schema for `items.decap.additonalgsops.sa_apf`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--decap--additonalgsops--sip_whitelist"></a>
### Nested Schema for `items.decap.additonalgsops.sip_whitelist`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--decap--additonalgsops--slicing"></a>
### Nested Schema for `items.decap.additonalgsops.slicing`

Read-Only:

* `enhanced` (String) - enhanced-slicing apps alias
* `offset` (Number) - min offset is 64 for 'none'; 4 for protocol-referenced offsets. required property till H 5.6
* `protocol` (String) - required property till H 5.6
<a id="nestedatt--items--decap--additonalgsops--ssl_decrypt"></a>
### Nested Schema for `items.decap.additonalgsops.ssl_decrypt`

Read-Only:

* `in_port` (Number) - Port number of 0 represents 'any' port
* `out_port` (Number) - Port number of 0 represents 'auto' port
<a id="nestedatt--items--decap--additonalgsops--trailer_add"></a>
### Nested Schema for `items.decap.additonalgsops.trailer_add`

Read-Only:

* `types` (Set of String) - 'crc' is not applicable for G-series
<a id="nestedatt--items--decap--additonalgsops--trailer_remove"></a>
### Nested Schema for `items.decap.additonalgsops.trailer_remove`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--decap--additonalgsops--tunnel_decap"></a>
### Nested Schema for `items.decap.additonalgsops.tunnel_decap`

Read-Only:

* `custom` (Attributes) (see [below for nested schema](#nestedatt--items--decap--additonalgsops--tunnel_decap--custom))
* `erspan_flow_id` (Number) - only applicable for 'erspan', A Flow ID of 0 decapsulates all ERSPAN tunnel traffic regardless of Flow ID
* `gmip_port` (Number) - only applicable for 'gmip', in which case it is required. Specifies the UDP port on which the Tunnel Network port on the receiving GigaVUE H Series is listening. Must match the configuration of the portdst configured on the sending end of the tunnel
* `l2_gre_key` (Number) - only applicable for 'l2gre', in which case it is required.
* `tls_pcapng` (Attributes) (see [below for nested schema](#nestedatt--items--decap--additonalgsops--tunnel_decap--tls_pcapng))
* `type` (String)
* `vxlan` (Attributes) (see [below for nested schema](#nestedatt--items--decap--additonalgsops--tunnel_decap--vxlan))
<a id="nestedatt--items--decap--additonalgsops--tunnel_decap--custom"></a>
### Nested Schema for `items.decap.additonalgsops.tunnel_decap.custom`

Read-Only:

* `port_dst` (Number) - when specified as 0, no validation will be done in the packet.
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
<a id="nestedatt--items--decap--additonalgsops--tunnel_decap--tls_pcapng"></a>
### Nested Schema for `items.decap.additonalgsops.tunnel_decap.tls_pcapng`

Read-Only:

* `decap_key` (String)
* `listener` (String)
<a id="nestedatt--items--decap--additonalgsops--tunnel_decap--vxlan"></a>
### Nested Schema for `items.decap.additonalgsops.tunnel_decap.vxlan`

Read-Only:

* `port_dst` (Number)
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
* `vni` (Number)
<a id="nestedatt--items--decap--additonalgsops--tunnel_encap"></a>
### Nested Schema for `items.decap.additonalgsops.tunnel_encap`

Read-Only:

* `gmip_config` (Attributes) - Configuration for GMIP Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--items--decap--additonalgsops--tunnel_encap--gmip_config))
* `l2_gre_config` (Attributes) - Configuration for l2GRE Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--items--decap--additonalgsops--tunnel_encap--l2_gre_config))
* `tls_pcapng` (Attributes) (see [below for nested schema](#nestedatt--items--decap--additonalgsops--tunnel_encap--tls_pcapng))
* `type` (String)
* `vxlan_config` (Attributes) - Configuration for Vxlan Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--items--decap--additonalgsops--tunnel_encap--vxlan_config))
<a id="nestedatt--items--decap--additonalgsops--tunnel_encap--gmip_config"></a>
### Nested Schema for `items.decap.additonalgsops.tunnel_encap.gmip_config`

Read-Only:

* `dscp` (Number)
* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `flow_label` (Number)
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `src_port` (Number)
* `ttl` (Number)
<a id="nestedatt--items--decap--additonalgsops--tunnel_encap--l2_gre_config"></a>
### Nested Schema for `items.decap.additonalgsops.tunnel_encap.l2_gre_config`

Read-Only:

* `dscp` (Number)
* `dst_ip` (String) - ip destination, mutually exclusive with 'pgDst'. IPv4 or IPv6.
* `flow_label` (Number)
* `key` (Number)
* `pg_dst` (String) - port group destination alias, mutually exclusive with 'dstIp'
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `session_field` (String) - required with stateful loadBalance when 'appType' is 'tunnel'
* `session_pos` (String) - required if 'sessionField' is specified
* `ttl` (Number)
<a id="nestedatt--items--decap--additonalgsops--tunnel_encap--tls_pcapng"></a>
### Nested Schema for `items.decap.additonalgsops.tunnel_encap.tls_pcapng`

Read-Only:

* `exporter` (String)
* `exporter_group` (String)
<a id="nestedatt--items--decap--additonalgsops--tunnel_encap--vxlan_config"></a>
### Nested Schema for `items.decap.additonalgsops.tunnel_encap.vxlan_config`

Read-Only:

* `dscp` (Number)
* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `src_port` (Number)
* `ttl` (Number)
* `vni` (Number)
<a id="nestedatt--items--decap--rules"></a>
### Nested Schema for `items.decap.rules`

Read-Only:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--items--decap--rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--items--decap--rules--pass_rules))
<a id="nestedatt--items--decap--rules--drop_rules"></a>
### Nested Schema for `items.decap.rules.drop_rules`

Read-Only:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--items--decap--rules--drop_rules--ip_rewrite))
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--items--decap--rules--drop_rules--rewrite))
* `rule_id` (Number)
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--items--decap--rules--drop_rules--vlan_tag))
<a id="nestedatt--items--decap--rules--drop_rules--ip_rewrite"></a>
### Nested Schema for `items.decap.rules.drop_rules.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--items--decap--rules--drop_rules--rewrite"></a>
### Nested Schema for `items.decap.rules.drop_rules.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--items--decap--rules--drop_rules--vlan_tag"></a>
### Nested Schema for `items.decap.rules.drop_rules.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)
<a id="nestedatt--items--decap--rules--pass_rules"></a>
### Nested Schema for `items.decap.rules.pass_rules`

Read-Only:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--items--decap--rules--pass_rules--ip_rewrite))
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--items--decap--rules--pass_rules--rewrite))
* `rule_id` (Number)
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--items--decap--rules--pass_rules--vlan_tag))
<a id="nestedatt--items--decap--rules--pass_rules--ip_rewrite"></a>
### Nested Schema for `items.decap.rules.pass_rules.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--items--decap--rules--pass_rules--rewrite"></a>
### Nested Schema for `items.decap.rules.pass_rules.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--items--decap--rules--pass_rules--vlan_tag"></a>
### Nested Schema for `items.decap.rules.pass_rules.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)
<a id="nestedatt--items--encap"></a>
### Nested Schema for `items.encap`

Read-Only:

* `additonalgsops` (Attributes) - GigaSMART Applications for a GSOP. At least one GsApp must be defined for the config to be valid (see [below for nested schema](#nestedatt--items--encap--additonalgsops))
* `export_configs` (Attributes List) (see [below for nested schema](#nestedatt--items--encap--export_configs))
* `exporter_group_alias` (String)
* `gsop_alias` (String)
* `map_alias` (String)
* `rules` (Attributes) - Map Rules Container. Private class (see [below for nested schema](#nestedatt--items--encap--rules))
* `source_port` (List of String) - source port for encap traffic , it would be on the same cluster as GsEngine
<a id="nestedatt--items--encap--additonalgsops"></a>
### Nested Schema for `items.encap.additonalgsops`

Read-Only:

* `apf` (Attributes) (see [below for nested schema](#nestedatt--items--encap--additonalgsops--apf))
* `dedup` (Attributes) (see [below for nested schema](#nestedatt--items--encap--additonalgsops--dedup))
* `diameter_whitelist` (Attributes) (see [below for nested schema](#nestedatt--items--encap--additonalgsops--diameter_whitelist))
* `flow_filter` (Attributes) - GigaSMART 'Flow Filter' Application Configuration (see [below for nested schema](#nestedatt--items--encap--additonalgsops--flow_filter))
* `flow_sampling` (Attributes) - GigaSMART 'Flow Sampling' Application Configuration (see [below for nested schema](#nestedatt--items--encap--additonalgsops--flow_sampling))
* `gseries_header_add` (Attributes) - Only applicable for G-series (see [below for nested schema](#nestedatt--items--encap--additonalgsops--gseries_header_add))
* `gseries_header_remove` (Attributes) (see [below for nested schema](#nestedatt--items--encap--additonalgsops--gseries_header_remove))
* `gseries_load_balance` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--items--encap--additonalgsops--gseries_load_balance))
* `gseries_pattern_match` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--items--encap--additonalgsops--gseries_pattern_match))
* `gtp_whitelist` (Attributes) (see [below for nested schema](#nestedatt--items--encap--additonalgsops--gtp_whitelist))
* `header_add` (Attributes) - GigaSMART 'Add Header' Application Configuration (see [below for nested schema](#nestedatt--items--encap--additonalgsops--header_add))
* `header_remove` (Attributes) - GigaSMART 'Remove Header' Application Configuration (see [below for nested schema](#nestedatt--items--encap--additonalgsops--header_remove))
* `icap` (Attributes) - GigaSMART ICAP Configuration (see [below for nested schema](#nestedatt--items--encap--additonalgsops--icap))
* `inline_ssl` (Attributes) - GigaSMART Inline SSL Profile Configuration (see [below for nested schema](#nestedatt--items--encap--additonalgsops--inline_ssl))
* `load_balance` (Attributes) - GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--items--encap--additonalgsops--load_balance))
* `masking` (Attributes) - GigaSMART 'Masking' Application Configuration (see [below for nested schema](#nestedatt--items--encap--additonalgsops--masking))
* `metadata_export` (Attributes) (see [below for nested schema](#nestedatt--items--encap--additonalgsops--metadata_export))
* `netflow` (Attributes) (see [below for nested schema](#nestedatt--items--encap--additonalgsops--netflow))
* `sa_apf` (Attributes) (see [below for nested schema](#nestedatt--items--encap--additonalgsops--sa_apf))
* `sip_whitelist` (Attributes) (see [below for nested schema](#nestedatt--items--encap--additonalgsops--sip_whitelist))
* `slicing` (Attributes) - GigaSMART 'Slicing' Application Configuration (see [below for nested schema](#nestedatt--items--encap--additonalgsops--slicing))
* `ssl_decrypt` (Attributes) - GigaSMART 'SSL Decrypt' Application Configuration (see [below for nested schema](#nestedatt--items--encap--additonalgsops--ssl_decrypt))
* `trailer_add` (Attributes) - GigaSMART 'Add Trailer' Application Configuration (see [below for nested schema](#nestedatt--items--encap--additonalgsops--trailer_add))
* `trailer_remove` (Attributes) (see [below for nested schema](#nestedatt--items--encap--additonalgsops--trailer_remove))
* `tunnel_decap` (Attributes) - GigaSMART 'Decapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--items--encap--additonalgsops--tunnel_decap))
* `tunnel_encap` (Attributes) - GigaSMART 'Encapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--items--encap--additonalgsops--tunnel_encap))
<a id="nestedatt--items--encap--additonalgsops--apf"></a>
### Nested Schema for `items.encap.additonalgsops.apf`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--encap--additonalgsops--dedup"></a>
### Nested Schema for `items.encap.additonalgsops.dedup`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--encap--additonalgsops--diameter_whitelist"></a>
### Nested Schema for `items.encap.additonalgsops.diameter_whitelist`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--encap--additonalgsops--flow_filter"></a>
### Nested Schema for `items.encap.additonalgsops.flow_filter`

Read-Only:

* `type` (String)
<a id="nestedatt--items--encap--additonalgsops--flow_sampling"></a>
### Nested Schema for `items.encap.additonalgsops.flow_sampling`

Read-Only:

* `type` (String)
<a id="nestedatt--items--encap--additonalgsops--gseries_header_add"></a>
### Nested Schema for `items.encap.additonalgsops.gseries_header_add`

Read-Only:

* `types` (Set of String)
<a id="nestedatt--items--encap--additonalgsops--gseries_header_remove"></a>
### Nested Schema for `items.encap.additonalgsops.gseries_header_remove`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--encap--additonalgsops--gseries_load_balance"></a>
### Nested Schema for `items.encap.additonalgsops.gseries_load_balance`

Read-Only:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--items--encap--additonalgsops--gseries_load_balance--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--items--encap--additonalgsops--gseries_load_balance--variable_offset))
<a id="nestedatt--items--encap--additonalgsops--gseries_load_balance--fixed_offset"></a>
### Nested Schema for `items.encap.additonalgsops.gseries_load_balance.fixed_offset`

Read-Only:

* `hash` (String)
* `length` (Number)
* `offset` (Number)
<a id="nestedatt--items--encap--additonalgsops--gseries_load_balance--variable_offset"></a>
### Nested Schema for `items.encap.additonalgsops.gseries_load_balance.variable_offset`

Read-Only:

* `end_delim` (String)
* `hash` (String)
* `start_delim` (String)
* `start_field` (String)
<a id="nestedatt--items--encap--additonalgsops--gseries_pattern_match"></a>
### Nested Schema for `items.encap.additonalgsops.gseries_pattern_match`

Read-Only:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--items--encap--additonalgsops--gseries_pattern_match--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--items--encap--additonalgsops--gseries_pattern_match--variable_offset))
<a id="nestedatt--items--encap--additonalgsops--gseries_pattern_match--fixed_offset"></a>
### Nested Schema for `items.encap.additonalgsops.gseries_pattern_match.fixed_offset`

Read-Only:

* `length` (Number)
* `offset` (Number)
<a id="nestedatt--items--encap--additonalgsops--gseries_pattern_match--variable_offset"></a>
### Nested Schema for `items.encap.additonalgsops.gseries_pattern_match.variable_offset`

Read-Only:

* `end_delim` (String)
* `start_delim` (String)
<a id="nestedatt--items--encap--additonalgsops--gtp_whitelist"></a>
### Nested Schema for `items.encap.additonalgsops.gtp_whitelist`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--encap--additonalgsops--header_add"></a>
### Nested Schema for `items.encap.additonalgsops.header_add`

Read-Only:

* `vlan` (Number)
<a id="nestedatt--items--encap--additonalgsops--header_remove"></a>
### Nested Schema for `items.encap.additonalgsops.header_remove`

Read-Only:

* `ah1` (String) - only valid and required for 'generic'. First anchor header, from which the header to be stripped is occurred.
* `ah2` (String) - only valid and required for 'generic'. next anchor header.
* `custom_len` (Number) - only valid for 'generic'. length of unknown header.
* `erspan_flow_id` (Number) - only applicable for 'erspan', specifies the flow id to strip. Value of 0 will strip all flow ids
* `fp_dst_switch_id` (Number) - Only valid and required for 'fabricPath', 12 bit destination switch id
* `fp_src_switch_id` (Number) - Only valid and required for 'fabricPath', 12 bit source switch id
* `header_count` (Number) - only valid for 'generic'. Number of headers to be stripped.
* `offset` (String) - only valid and required for 'generic'.'start': strip from start of ah1, 'end': strip from end of ah1 or any other offset in the range of length of ah1, 'offsetRangeValue' is required if offset is 'offsetRange'.
* `offset_range_value` (Number) - only valid and required when offset is 'offsetRange', integer within range of size of header
* `protocol` (String) - 'gre' and 'fabricPath' are only applicable for H-series
* `timestamp_format` (String) - Timestamp format. Only valid and required for 'fm6000Ts'
* `vlan_header` (String) - Only valid when protocol is 'vlan'. Specifies the target vlan header to strip. Defaults to 'outer'
* `vxlan_id` (Number) - 24-bit value. Only valid when protocol is 'vxlan'. Specifies the vxlan id to strip. Value of '0' will strip all vxlan ids
<a id="nestedatt--items--encap--additonalgsops--icap"></a>
### Nested Schema for `items.encap.additonalgsops.icap`

Read-Only:

* `icap_profile` (String) - Alias of referenced ICAP Profile
<a id="nestedatt--items--encap--additonalgsops--inline_ssl"></a>
### Nested Schema for `items.encap.additonalgsops.inline_ssl`

Read-Only:

* `inline_ssl_profile` (String) - Alias of referenced Inline SSL Profile
<a id="nestedatt--items--encap--additonalgsops--load_balance"></a>
### Nested Schema for `items.encap.additonalgsops.load_balance`

Read-Only:

* `enhanced` (Attributes) - Private class. Enhanced part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--items--encap--additonalgsops--load_balance--enhanced))
* `stateful` (Attributes) - Private class. Stateful part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--items--encap--additonalgsops--load_balance--stateful))
* `stateless` (Attributes) - Private class. Stateless part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--items--encap--additonalgsops--load_balance--stateless))
<a id="nestedatt--items--encap--additonalgsops--load_balance--enhanced"></a>
### Nested Schema for `items.encap.additonalgsops.load_balance.enhanced`

Read-Only:

* `elb_alias` (String) - elb app alias
<a id="nestedatt--items--encap--additonalgsops--load_balance--stateful"></a>
### Nested Schema for `items.encap.additonalgsops.load_balance.stateful`

Read-Only:

* `app_type` (String) - 'sapf' option is added in release H 4.3.01. 'sip' and 'tunnel' added in release H 5.1. 'diameter' is added in release H 5.5
* `diameter_key_hash_type` (String) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyHash'
* `diameter_key_multi_hash_type` (Attributes List) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyMultiHash' (see [below for nested schema](#nestedatt--items--encap--additonalgsops--load_balance--stateful--diameter_key_multi_hash_type))
* `gtp_key_hash_type` (String) - required when 'appType' == 'gtp' and 'lbType' == 'gtpKeyHash'. ignored otherwise
* `lb_type` (String) - 'gtpKeyHash' is only applicable when 'appType' == 'gtp'. 'sipKeyHash' is only applicable when 'appType' == 'sip'
* `sip_key_hash_type` (String) - required when 'appType' == 'sip' and 'lbType' == 'sipKeyHash'. ignored otherwise
<a id="nestedatt--items--encap--additonalgsops--load_balance--stateful--diameter_key_multi_hash_type"></a>
### Nested Schema for `items.encap.additonalgsops.load_balance.stateful.diameter_key_multi_hash_type`

Read-Only:

* `avp_codevalue` (Number) - required when 'key' == 'avpCode'
* `key` (String)
<a id="nestedatt--items--encap--additonalgsops--load_balance--stateless"></a>
### Nested Schema for `items.encap.additonalgsops.load_balance.stateless`

Read-Only:

* `field_location` (String) - Ignored when 'hashFields' == 'gtpuTeid'. Required otherwise
* `hash_fields` (String)
<a id="nestedatt--items--encap--additonalgsops--masking"></a>
### Nested Schema for `items.encap.additonalgsops.masking`

Read-Only:

* `content_type` (String) - content type that will trigger masking, only valid and required for protocol 'sip'
* `length` (Number) - max length is 9600 on H-series; 9000 on G-series, not applicable for protocol 'sip' valid and required otherwise
* `offset` (Number) - min offset is 0 for 'none'; 1 for protocol-referenced offsets, not applicable for protocol 'sip' valid and required otherwise
* `pattern` (String) - 1-byte hex mask, not applicable for protocol 'sip' valid and required otherwise
* `protocol` (String)
<a id="nestedatt--items--encap--additonalgsops--metadata_export"></a>
### Nested Schema for `items.encap.additonalgsops.metadata_export`

Read-Only:

* `cache` (String) - metadata cache alias. cache should have exporters defined, and the exporters should have applicationProfiles defined
<a id="nestedatt--items--encap--additonalgsops--netflow"></a>
### Nested Schema for `items.encap.additonalgsops.netflow`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--encap--additonalgsops--sa_apf"></a>
### Nested Schema for `items.encap.additonalgsops.sa_apf`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--encap--additonalgsops--sip_whitelist"></a>
### Nested Schema for `items.encap.additonalgsops.sip_whitelist`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--encap--additonalgsops--slicing"></a>
### Nested Schema for `items.encap.additonalgsops.slicing`

Read-Only:

* `enhanced` (String) - enhanced-slicing apps alias
* `offset` (Number) - min offset is 64 for 'none'; 4 for protocol-referenced offsets. required property till H 5.6
* `protocol` (String) - required property till H 5.6
<a id="nestedatt--items--encap--additonalgsops--ssl_decrypt"></a>
### Nested Schema for `items.encap.additonalgsops.ssl_decrypt`

Read-Only:

* `in_port` (Number) - Port number of 0 represents 'any' port
* `out_port` (Number) - Port number of 0 represents 'auto' port
<a id="nestedatt--items--encap--additonalgsops--trailer_add"></a>
### Nested Schema for `items.encap.additonalgsops.trailer_add`

Read-Only:

* `types` (Set of String) - 'crc' is not applicable for G-series
<a id="nestedatt--items--encap--additonalgsops--trailer_remove"></a>
### Nested Schema for `items.encap.additonalgsops.trailer_remove`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--encap--additonalgsops--tunnel_decap"></a>
### Nested Schema for `items.encap.additonalgsops.tunnel_decap`

Read-Only:

* `custom` (Attributes) (see [below for nested schema](#nestedatt--items--encap--additonalgsops--tunnel_decap--custom))
* `erspan_flow_id` (Number) - only applicable for 'erspan', A Flow ID of 0 decapsulates all ERSPAN tunnel traffic regardless of Flow ID
* `gmip_port` (Number) - only applicable for 'gmip', in which case it is required. Specifies the UDP port on which the Tunnel Network port on the receiving GigaVUE H Series is listening. Must match the configuration of the portdst configured on the sending end of the tunnel
* `l2_gre_key` (Number) - only applicable for 'l2gre', in which case it is required.
* `tls_pcapng` (Attributes) (see [below for nested schema](#nestedatt--items--encap--additonalgsops--tunnel_decap--tls_pcapng))
* `type` (String)
* `vxlan` (Attributes) (see [below for nested schema](#nestedatt--items--encap--additonalgsops--tunnel_decap--vxlan))
<a id="nestedatt--items--encap--additonalgsops--tunnel_decap--custom"></a>
### Nested Schema for `items.encap.additonalgsops.tunnel_decap.custom`

Read-Only:

* `port_dst` (Number) - when specified as 0, no validation will be done in the packet.
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
<a id="nestedatt--items--encap--additonalgsops--tunnel_decap--tls_pcapng"></a>
### Nested Schema for `items.encap.additonalgsops.tunnel_decap.tls_pcapng`

Read-Only:

* `decap_key` (String)
* `listener` (String)
<a id="nestedatt--items--encap--additonalgsops--tunnel_decap--vxlan"></a>
### Nested Schema for `items.encap.additonalgsops.tunnel_decap.vxlan`

Read-Only:

* `port_dst` (Number)
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
* `vni` (Number)
<a id="nestedatt--items--encap--additonalgsops--tunnel_encap"></a>
### Nested Schema for `items.encap.additonalgsops.tunnel_encap`

Read-Only:

* `gmip_config` (Attributes) - Configuration for GMIP Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--items--encap--additonalgsops--tunnel_encap--gmip_config))
* `l2_gre_config` (Attributes) - Configuration for l2GRE Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--items--encap--additonalgsops--tunnel_encap--l2_gre_config))
* `tls_pcapng` (Attributes) (see [below for nested schema](#nestedatt--items--encap--additonalgsops--tunnel_encap--tls_pcapng))
* `type` (String)
* `vxlan_config` (Attributes) - Configuration for Vxlan Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--items--encap--additonalgsops--tunnel_encap--vxlan_config))
<a id="nestedatt--items--encap--additonalgsops--tunnel_encap--gmip_config"></a>
### Nested Schema for `items.encap.additonalgsops.tunnel_encap.gmip_config`

Read-Only:

* `dscp` (Number)
* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `flow_label` (Number)
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `src_port` (Number)
* `ttl` (Number)
<a id="nestedatt--items--encap--additonalgsops--tunnel_encap--l2_gre_config"></a>
### Nested Schema for `items.encap.additonalgsops.tunnel_encap.l2_gre_config`

Read-Only:

* `dscp` (Number)
* `dst_ip` (String) - ip destination, mutually exclusive with 'pgDst'. IPv4 or IPv6.
* `flow_label` (Number)
* `key` (Number)
* `pg_dst` (String) - port group destination alias, mutually exclusive with 'dstIp'
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `session_field` (String) - required with stateful loadBalance when 'appType' is 'tunnel'
* `session_pos` (String) - required if 'sessionField' is specified
* `ttl` (Number)
<a id="nestedatt--items--encap--additonalgsops--tunnel_encap--tls_pcapng"></a>
### Nested Schema for `items.encap.additonalgsops.tunnel_encap.tls_pcapng`

Read-Only:

* `exporter` (String)
* `exporter_group` (String)
<a id="nestedatt--items--encap--additonalgsops--tunnel_encap--vxlan_config"></a>
### Nested Schema for `items.encap.additonalgsops.tunnel_encap.vxlan_config`

Read-Only:

* `dscp` (Number)
* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `src_port` (Number)
* `ttl` (Number)
* `vni` (Number)
<a id="nestedatt--items--encap--export_configs"></a>
### Nested Schema for `items.encap.export_configs`

Read-Only:

* `exporter_alias` (String)
* `remote_application_port` (Number)
* `remote_ip` (String)
* `source_application_port` (Number)
<a id="nestedatt--items--encap--rules"></a>
### Nested Schema for `items.encap.rules`

Read-Only:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--items--encap--rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--items--encap--rules--pass_rules))
<a id="nestedatt--items--encap--rules--drop_rules"></a>
### Nested Schema for `items.encap.rules.drop_rules`

Read-Only:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--items--encap--rules--drop_rules--ip_rewrite))
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--items--encap--rules--drop_rules--rewrite))
* `rule_id` (Number)
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--items--encap--rules--drop_rules--vlan_tag))
<a id="nestedatt--items--encap--rules--drop_rules--ip_rewrite"></a>
### Nested Schema for `items.encap.rules.drop_rules.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--items--encap--rules--drop_rules--rewrite"></a>
### Nested Schema for `items.encap.rules.drop_rules.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--items--encap--rules--drop_rules--vlan_tag"></a>
### Nested Schema for `items.encap.rules.drop_rules.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)
<a id="nestedatt--items--encap--rules--pass_rules"></a>
### Nested Schema for `items.encap.rules.pass_rules`

Read-Only:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--items--encap--rules--pass_rules--ip_rewrite))
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--items--encap--rules--pass_rules--rewrite))
* `rule_id` (Number)
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--items--encap--rules--pass_rules--vlan_tag))
<a id="nestedatt--items--encap--rules--pass_rules--ip_rewrite"></a>
### Nested Schema for `items.encap.rules.pass_rules.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--items--encap--rules--pass_rules--rewrite"></a>
### Nested Schema for `items.encap.rules.pass_rules.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--items--encap--rules--pass_rules--vlan_tag"></a>
### Nested Schema for `items.encap.rules.pass_rules.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)
<a id="nestedatt--items--ssl_profile"></a>
### Nested Schema for `items.ssl_profile`

Read-Only:

* `cipher` (String)
* `mtls` (String)
* `version` (String)
<a id="nestedatt--items--tcp_profile"></a>
### Nested Schema for `items.tcp_profile`

Read-Only:

* `keep_alive_timer` (Number)
* `selective_ack` (String)
* `syn_retries` (Number)

