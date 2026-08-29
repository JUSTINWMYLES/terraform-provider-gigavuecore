---
page_title: "gigavuecore_load_all_gsops Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all GSOPs
---

# gigavuecore_load_all_gsops Data Source

Load all GSOPs

## Example Usage

```terraform
data "gigavuecore_load_all_gsops" "example" {
  cluster_id = "example"
  page       = "example"
  sort       = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String)
* `cluster_id` (String) - id of the defining cluster
* `gs_apps` (Attributes) - GigaSMART Applications for a GSOP. At least one GsApp must be defined for the config to be valid (see [below for nested schema](#nestedatt--items--gs_apps))
* `gs_group` (String) - Alias of referenced managing GsGroup
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--health_state_reasons))
<a id="nestedatt--items--gs_apps"></a>
### Nested Schema for `items.gs_apps`

Read-Only:

* `apf` (Attributes) (see [below for nested schema](#nestedatt--items--gs_apps--apf))
* `dedup` (Attributes) (see [below for nested schema](#nestedatt--items--gs_apps--dedup))
* `diameter_whitelist` (Attributes) (see [below for nested schema](#nestedatt--items--gs_apps--diameter_whitelist))
* `flow_filter` (Attributes) - GigaSMART 'Flow Filter' Application Configuration (see [below for nested schema](#nestedatt--items--gs_apps--flow_filter))
* `flow_sampling` (Attributes) - GigaSMART 'Flow Sampling' Application Configuration (see [below for nested schema](#nestedatt--items--gs_apps--flow_sampling))
* `gseries_header_add` (Attributes) - Only applicable for G-series (see [below for nested schema](#nestedatt--items--gs_apps--gseries_header_add))
* `gseries_header_remove` (Attributes) (see [below for nested schema](#nestedatt--items--gs_apps--gseries_header_remove))
* `gseries_load_balance` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--items--gs_apps--gseries_load_balance))
* `gseries_pattern_match` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--items--gs_apps--gseries_pattern_match))
* `gtp_whitelist` (Attributes) (see [below for nested schema](#nestedatt--items--gs_apps--gtp_whitelist))
* `header_add` (Attributes) - GigaSMART 'Add Header' Application Configuration (see [below for nested schema](#nestedatt--items--gs_apps--header_add))
* `header_remove` (Attributes) - GigaSMART 'Remove Header' Application Configuration (see [below for nested schema](#nestedatt--items--gs_apps--header_remove))
* `icap` (Attributes) - GigaSMART ICAP Configuration (see [below for nested schema](#nestedatt--items--gs_apps--icap))
* `inline_ssl` (Attributes) - GigaSMART Inline SSL Profile Configuration (see [below for nested schema](#nestedatt--items--gs_apps--inline_ssl))
* `load_balance` (Attributes) - GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--items--gs_apps--load_balance))
* `masking` (Attributes) - GigaSMART 'Masking' Application Configuration (see [below for nested schema](#nestedatt--items--gs_apps--masking))
* `metadata_export` (Attributes) (see [below for nested schema](#nestedatt--items--gs_apps--metadata_export))
* `netflow` (Attributes) (see [below for nested schema](#nestedatt--items--gs_apps--netflow))
* `sa_apf` (Attributes) (see [below for nested schema](#nestedatt--items--gs_apps--sa_apf))
* `sip_whitelist` (Attributes) (see [below for nested schema](#nestedatt--items--gs_apps--sip_whitelist))
* `slicing` (Attributes) - GigaSMART 'Slicing' Application Configuration (see [below for nested schema](#nestedatt--items--gs_apps--slicing))
* `ssl_decrypt` (Attributes) - GigaSMART 'SSL Decrypt' Application Configuration (see [below for nested schema](#nestedatt--items--gs_apps--ssl_decrypt))
* `trailer_add` (Attributes) - GigaSMART 'Add Trailer' Application Configuration (see [below for nested schema](#nestedatt--items--gs_apps--trailer_add))
* `trailer_remove` (Attributes) (see [below for nested schema](#nestedatt--items--gs_apps--trailer_remove))
* `tunnel_decap` (Attributes) - GigaSMART 'Decapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--items--gs_apps--tunnel_decap))
* `tunnel_encap` (Attributes) - GigaSMART 'Encapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--items--gs_apps--tunnel_encap))
<a id="nestedatt--items--gs_apps--apf"></a>
### Nested Schema for `items.gs_apps.apf`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--gs_apps--dedup"></a>
### Nested Schema for `items.gs_apps.dedup`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--gs_apps--diameter_whitelist"></a>
### Nested Schema for `items.gs_apps.diameter_whitelist`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--gs_apps--flow_filter"></a>
### Nested Schema for `items.gs_apps.flow_filter`

Read-Only:

* `type` (String)
<a id="nestedatt--items--gs_apps--flow_sampling"></a>
### Nested Schema for `items.gs_apps.flow_sampling`

Read-Only:

* `type` (String)
<a id="nestedatt--items--gs_apps--gseries_header_add"></a>
### Nested Schema for `items.gs_apps.gseries_header_add`

Read-Only:

* `types` (Set of String)
<a id="nestedatt--items--gs_apps--gseries_header_remove"></a>
### Nested Schema for `items.gs_apps.gseries_header_remove`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--gs_apps--gseries_load_balance"></a>
### Nested Schema for `items.gs_apps.gseries_load_balance`

Read-Only:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--items--gs_apps--gseries_load_balance--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--items--gs_apps--gseries_load_balance--variable_offset))
<a id="nestedatt--items--gs_apps--gseries_load_balance--fixed_offset"></a>
### Nested Schema for `items.gs_apps.gseries_load_balance.fixed_offset`

Read-Only:

* `hash` (String)
* `length` (Number)
* `offset` (Number)
<a id="nestedatt--items--gs_apps--gseries_load_balance--variable_offset"></a>
### Nested Schema for `items.gs_apps.gseries_load_balance.variable_offset`

Read-Only:

* `end_delim` (String)
* `hash` (String)
* `start_delim` (String)
* `start_field` (String)
<a id="nestedatt--items--gs_apps--gseries_pattern_match"></a>
### Nested Schema for `items.gs_apps.gseries_pattern_match`

Read-Only:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--items--gs_apps--gseries_pattern_match--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--items--gs_apps--gseries_pattern_match--variable_offset))
<a id="nestedatt--items--gs_apps--gseries_pattern_match--fixed_offset"></a>
### Nested Schema for `items.gs_apps.gseries_pattern_match.fixed_offset`

Read-Only:

* `length` (Number)
* `offset` (Number)
<a id="nestedatt--items--gs_apps--gseries_pattern_match--variable_offset"></a>
### Nested Schema for `items.gs_apps.gseries_pattern_match.variable_offset`

Read-Only:

* `end_delim` (String)
* `start_delim` (String)
<a id="nestedatt--items--gs_apps--gtp_whitelist"></a>
### Nested Schema for `items.gs_apps.gtp_whitelist`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--gs_apps--header_add"></a>
### Nested Schema for `items.gs_apps.header_add`

Read-Only:

* `vlan` (Number)
<a id="nestedatt--items--gs_apps--header_remove"></a>
### Nested Schema for `items.gs_apps.header_remove`

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
<a id="nestedatt--items--gs_apps--icap"></a>
### Nested Schema for `items.gs_apps.icap`

Read-Only:

* `icap_profile` (String) - Alias of referenced ICAP Profile
<a id="nestedatt--items--gs_apps--inline_ssl"></a>
### Nested Schema for `items.gs_apps.inline_ssl`

Read-Only:

* `inline_ssl_profile` (String) - Alias of referenced Inline SSL Profile
<a id="nestedatt--items--gs_apps--load_balance"></a>
### Nested Schema for `items.gs_apps.load_balance`

Read-Only:

* `enhanced` (Attributes) - Private class. Enhanced part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--items--gs_apps--load_balance--enhanced))
* `stateful` (Attributes) - Private class. Stateful part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--items--gs_apps--load_balance--stateful))
* `stateless` (Attributes) - Private class. Stateless part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--items--gs_apps--load_balance--stateless))
<a id="nestedatt--items--gs_apps--load_balance--enhanced"></a>
### Nested Schema for `items.gs_apps.load_balance.enhanced`

Read-Only:

* `elb_alias` (String) - elb app alias
<a id="nestedatt--items--gs_apps--load_balance--stateful"></a>
### Nested Schema for `items.gs_apps.load_balance.stateful`

Read-Only:

* `app_type` (String) - 'sapf' option is added in release H 4.3.01. 'sip' and 'tunnel' added in release H 5.1. 'diameter' is added in release H 5.5
* `diameter_key_hash_type` (String) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyHash'
* `diameter_key_multi_hash_type` (Attributes List) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyMultiHash' (see [below for nested schema](#nestedatt--items--gs_apps--load_balance--stateful--diameter_key_multi_hash_type))
* `gtp_key_hash_type` (String) - required when 'appType' == 'gtp' and 'lbType' == 'gtpKeyHash'. ignored otherwise
* `lb_type` (String) - 'gtpKeyHash' is only applicable when 'appType' == 'gtp'. 'sipKeyHash' is only applicable when 'appType' == 'sip'
* `sip_key_hash_type` (String) - required when 'appType' == 'sip' and 'lbType' == 'sipKeyHash'. ignored otherwise
<a id="nestedatt--items--gs_apps--load_balance--stateful--diameter_key_multi_hash_type"></a>
### Nested Schema for `items.gs_apps.load_balance.stateful.diameter_key_multi_hash_type`

Read-Only:

* `avp_codevalue` (Number) - required when 'key' == 'avpCode'
* `key` (String)
<a id="nestedatt--items--gs_apps--load_balance--stateless"></a>
### Nested Schema for `items.gs_apps.load_balance.stateless`

Read-Only:

* `field_location` (String) - Ignored when 'hashFields' == 'gtpuTeid'. Required otherwise
* `hash_fields` (String)
<a id="nestedatt--items--gs_apps--masking"></a>
### Nested Schema for `items.gs_apps.masking`

Read-Only:

* `content_type` (String) - content type that will trigger masking, only valid and required for protocol 'sip'
* `length` (Number) - max length is 9600 on H-series; 9000 on G-series, not applicable for protocol 'sip' valid and required otherwise
* `offset` (Number) - min offset is 0 for 'none'; 1 for protocol-referenced offsets, not applicable for protocol 'sip' valid and required otherwise
* `pattern` (String) - 1-byte hex mask, not applicable for protocol 'sip' valid and required otherwise
* `protocol` (String)
<a id="nestedatt--items--gs_apps--metadata_export"></a>
### Nested Schema for `items.gs_apps.metadata_export`

Read-Only:

* `cache` (String) - metadata cache alias. cache should have exporters defined, and the exporters should have applicationProfiles defined
<a id="nestedatt--items--gs_apps--netflow"></a>
### Nested Schema for `items.gs_apps.netflow`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--gs_apps--sa_apf"></a>
### Nested Schema for `items.gs_apps.sa_apf`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--gs_apps--sip_whitelist"></a>
### Nested Schema for `items.gs_apps.sip_whitelist`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--gs_apps--slicing"></a>
### Nested Schema for `items.gs_apps.slicing`

Read-Only:

* `enhanced` (String) - enhanced-slicing apps alias
* `offset` (Number) - min offset is 64 for 'none'; 4 for protocol-referenced offsets. required property till H 5.6
* `protocol` (String) - required property till H 5.6
<a id="nestedatt--items--gs_apps--ssl_decrypt"></a>
### Nested Schema for `items.gs_apps.ssl_decrypt`

Read-Only:

* `in_port` (Number) - Port number of 0 represents 'any' port
* `out_port` (Number) - Port number of 0 represents 'auto' port
<a id="nestedatt--items--gs_apps--trailer_add"></a>
### Nested Schema for `items.gs_apps.trailer_add`

Read-Only:

* `types` (Set of String) - 'crc' is not applicable for G-series
<a id="nestedatt--items--gs_apps--trailer_remove"></a>
### Nested Schema for `items.gs_apps.trailer_remove`

Read-Only:

* `enabled` (String)
<a id="nestedatt--items--gs_apps--tunnel_decap"></a>
### Nested Schema for `items.gs_apps.tunnel_decap`

Read-Only:

* `custom` (Attributes) (see [below for nested schema](#nestedatt--items--gs_apps--tunnel_decap--custom))
* `erspan_flow_id` (Number) - only applicable for 'erspan', A Flow ID of 0 decapsulates all ERSPAN tunnel traffic regardless of Flow ID
* `gmip_port` (Number) - only applicable for 'gmip', in which case it is required. Specifies the UDP port on which the Tunnel Network port on the receiving GigaVUE H Series is listening. Must match the configuration of the portdst configured on the sending end of the tunnel
* `l2_gre_key` (Number) - only applicable for 'l2gre', in which case it is required.
* `tls_pcapng` (Attributes) (see [below for nested schema](#nestedatt--items--gs_apps--tunnel_decap--tls_pcapng))
* `type` (String)
* `vxlan` (Attributes) (see [below for nested schema](#nestedatt--items--gs_apps--tunnel_decap--vxlan))
<a id="nestedatt--items--gs_apps--tunnel_decap--custom"></a>
### Nested Schema for `items.gs_apps.tunnel_decap.custom`

Read-Only:

* `port_dst` (Number) - when specified as 0, no validation will be done in the packet.
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
<a id="nestedatt--items--gs_apps--tunnel_decap--tls_pcapng"></a>
### Nested Schema for `items.gs_apps.tunnel_decap.tls_pcapng`

Read-Only:

* `decap_key` (String)
* `listener` (String)
<a id="nestedatt--items--gs_apps--tunnel_decap--vxlan"></a>
### Nested Schema for `items.gs_apps.tunnel_decap.vxlan`

Read-Only:

* `port_dst` (Number)
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
* `vni` (Number)
<a id="nestedatt--items--gs_apps--tunnel_encap"></a>
### Nested Schema for `items.gs_apps.tunnel_encap`

Read-Only:

* `gmip_config` (Attributes) - Configuration for GMIP Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--items--gs_apps--tunnel_encap--gmip_config))
* `l2_gre_config` (Attributes) - Configuration for l2GRE Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--items--gs_apps--tunnel_encap--l2_gre_config))
* `tls_pcapng` (Attributes) (see [below for nested schema](#nestedatt--items--gs_apps--tunnel_encap--tls_pcapng))
* `type` (String)
* `vxlan_config` (Attributes) - Configuration for Vxlan Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--items--gs_apps--tunnel_encap--vxlan_config))
<a id="nestedatt--items--gs_apps--tunnel_encap--gmip_config"></a>
### Nested Schema for `items.gs_apps.tunnel_encap.gmip_config`

Read-Only:

* `dscp` (Number)
* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `flow_label` (Number)
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `src_port` (Number)
* `ttl` (Number)
<a id="nestedatt--items--gs_apps--tunnel_encap--l2_gre_config"></a>
### Nested Schema for `items.gs_apps.tunnel_encap.l2_gre_config`

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
<a id="nestedatt--items--gs_apps--tunnel_encap--tls_pcapng"></a>
### Nested Schema for `items.gs_apps.tunnel_encap.tls_pcapng`

Read-Only:

* `exporter` (String)
* `exporter_group` (String)
<a id="nestedatt--items--gs_apps--tunnel_encap--vxlan_config"></a>
### Nested Schema for `items.gs_apps.tunnel_encap.vxlan_config`

Read-Only:

* `dscp` (Number)
* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `src_port` (Number)
* `ttl` (Number)
* `vni` (Number)
<a id="nestedatt--items--health_state_reasons"></a>
### Nested Schema for `items.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

