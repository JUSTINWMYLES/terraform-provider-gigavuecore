---
page_title: "gigavuecore_update_gsop_apps Action - gigavuecore"
subcategory: ""
description: |-
  Update GSOP apps configuration
---

# gigavuecore_update_gsop_apps Action

Update GSOP apps configuration

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_update_gsop_apps" "example" {
  config {
    alias = "example"
    apf = {
      enabled = "enabled"
    }
    cluster_id = "example"
    dedup = {
      enabled = "enabled"
    }
    diameter_whitelist = {
      enabled = "enabled"
    }
    flow_filter = {
      type = "gtp"
    }
    flow_sampling = {
      type = "ip"
    }
    gseries_header_add = {
      types = ["srcid"]
    }
    gseries_header_remove = {
      enabled = "enabled"
    }
    gseries_load_balance = {
      fixed_offset = {
        hash   = "checksum"
        length = 1
        offset = 0
      }
      variable_offset = {
        end_delim   = "example"
        hash        = "checksum"
        start_delim = "example"
        start_field = "example"
      }
    }
    gseries_pattern_match = {
      fixed_offset = {
        length = 1
        offset = 0
      }
      variable_offset = {
        end_delim   = "example"
        start_delim = "example"
      }
    }
    gtp_whitelist = {
      enabled = "enabled"
    }
    header_add = {
      vlan = 0
    }
    header_remove = {
      ah1                = "none"
      ah2                = "none"
      custom_len         = 1
      erspan_flow_id     = 0
      fp_dst_switch_id   = 0
      fp_src_switch_id   = 0
      header_count       = 1
      offset             = "start"
      offset_range_value = 0
      protocol           = "gtp"
      timestamp_format   = "gigasmart"
      vlan_header        = "all"
      vxlan_id           = 0
    }
    icap = {
      icap_profile = "example"
    }
    inline_ssl = {
      inline_ssl_profile = "example"
    }
    load_balance = {
      enhanced = {
        elb_alias = "example"
      }
      stateful = {
        app_type               = "gtp"
        diameter_key_hash_type = "sessionId"
        diameter_key_multi_hash_type = [{
          avp_codevalue = 0
          key           = "sessionId"
        }]
        gtp_key_hash_type = "imsi"
        lb_type           = "leastBw"
        sip_key_hash_type = "callerId"
      }
      stateless = {
        field_location = "inner"
        hash_fields    = "ipOnly"
      }
    }
    masking = {
      content_type = "message_cpim"
      length       = 1
      offset       = 0
      pattern      = "a1"
      protocol     = "none"
    }
    metadata_export = {
      cache = "example"
    }
    netflow = {
      enabled = "enabled"
    }
    sa_apf = {
      enabled = "enabled"
    }
    sip_whitelist = {
      enabled = "enabled"
    }
    slicing = {
      enhanced = "example"
      offset   = 4
      protocol = "none"
    }
    ssl_decrypt = {
      in_port  = 0
      out_port = 0
    }
    trailer_add = {
      types = ["crc"]
    }
    trailer_remove = {
      enabled = "enabled"
    }
    tunnel_decap = {
      custom = {
        port_dst = 0
        port_src = 0
      }
      erspan_flow_id = 0
      gmip_port      = 0
      l2_gre_key     = 0
      tls_pcapng = {
        decap_key = "example"
        listener  = "example"
      }
      type = "gmip"
      vxlan = {
        port_dst = 1
        port_src = 0
        vni      = 0
      }
    }
    tunnel_encap = {
      gmip_config = {
        dscp       = 0
        dst_ip     = "example"
        dst_port   = 0
        flow_label = 0
        prec       = 0
        src_port   = 0
        ttl        = 1
      }
      l2_gre_config = {
        dscp          = 0
        dst_ip        = "example"
        flow_label    = 0
        key           = 0
        pg_dst        = "example"
        prec          = 0
        session_field = "fiveTupleIpv4"
        session_pos   = "inner"
        ttl           = 1
      }
      tls_pcapng = {
        exporter       = "example"
        exporter_group = "example"
      }
      type = "gmip"
      vxlan_config = {
        dscp     = 0
        dst_ip   = "example"
        dst_port = 4789
        src_port = 0
        ttl      = 1
        vni      = 1
      }
    }
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GSOP
* `apf` (Attributes, optional) - Only applicable for H-series (see [below for nested schema](#nestedatt--apf))
* `cluster_id` (String, required) - Target Cluster ID
* `dedup` (Attributes, optional) (see [below for nested schema](#nestedatt--dedup))
* `diameter_whitelist` (Attributes, optional) - Only applicable for H-series (see [below for nested schema](#nestedatt--diameter_whitelist))
* `flow_filter` (Attributes, optional) - GigaSMART 'Flow Filter' Application Configuration (see [below for nested schema](#nestedatt--flow_filter))
* `flow_sampling` (Attributes, optional) - GigaSMART 'Flow Sampling' Application Configuration (see [below for nested schema](#nestedatt--flow_sampling))
* `gseries_header_add` (Attributes, optional) - Only applicable for G-series (see [below for nested schema](#nestedatt--gseries_header_add))
* `gseries_header_remove` (Attributes, optional) - Only applicable for G-series (see [below for nested schema](#nestedatt--gseries_header_remove))
* `gseries_load_balance` (Attributes, optional) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--gseries_load_balance))
* `gseries_pattern_match` (Attributes, optional) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--gseries_pattern_match))
* `gtp_whitelist` (Attributes, optional) - Only applicable for H-series (see [below for nested schema](#nestedatt--gtp_whitelist))
* `header_add` (Attributes, optional) - GigaSMART 'Add Header' Application Configuration (see [below for nested schema](#nestedatt--header_add))
* `header_remove` (Attributes, optional) - GigaSMART 'Remove Header' Application Configuration (see [below for nested schema](#nestedatt--header_remove))
* `icap` (Attributes, optional) - GigaSMART ICAP Configuration (see [below for nested schema](#nestedatt--icap))
* `inline_ssl` (Attributes, optional) - GigaSMART Inline SSL Profile Configuration (see [below for nested schema](#nestedatt--inline_ssl))
* `load_balance` (Attributes, optional) - GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--load_balance))
* `masking` (Attributes, optional) - GigaSMART 'Masking' Application Configuration (see [below for nested schema](#nestedatt--masking))
* `metadata_export` (Attributes, optional) - Only applicable for H-series (see [below for nested schema](#nestedatt--metadata_export))
* `netflow` (Attributes, optional) - Only applicable for H-series (see [below for nested schema](#nestedatt--netflow))
* `sa_apf` (Attributes, optional) - Deprecated (functionality will not be supported from 5.4 release onwards); Only applicable for H-series (see [below for nested schema](#nestedatt--sa_apf))
* `sip_whitelist` (Attributes, optional) - Only applicable for H-series (see [below for nested schema](#nestedatt--sip_whitelist))
* `slicing` (Attributes, optional) - GigaSMART 'Slicing' Application Configuration (see [below for nested schema](#nestedatt--slicing))
* `ssl_decrypt` (Attributes, optional) - GigaSMART 'SSL Decrypt' Application Configuration (see [below for nested schema](#nestedatt--ssl_decrypt))
* `trailer_add` (Attributes, optional) - GigaSMART 'Add Trailer' Application Configuration (see [below for nested schema](#nestedatt--trailer_add))
* `trailer_remove` (Attributes, optional) (see [below for nested schema](#nestedatt--trailer_remove))
* `tunnel_decap` (Attributes, optional) - GigaSMART 'Decapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--tunnel_decap))
* `tunnel_encap` (Attributes, optional) - GigaSMART 'Encapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--tunnel_encap))

<a id="nestedatt--apf"></a>
### Nested Schema for `apf`

Optional:

* `enabled` (String)

<a id="nestedatt--dedup"></a>
### Nested Schema for `dedup`

Optional:

* `enabled` (String)

<a id="nestedatt--diameter_whitelist"></a>
### Nested Schema for `diameter_whitelist`

Optional:

* `enabled` (String)

<a id="nestedatt--flow_filter"></a>
### Nested Schema for `flow_filter`

Required:

* `type` (String)

<a id="nestedatt--flow_sampling"></a>
### Nested Schema for `flow_sampling`

Required:

* `type` (String)

<a id="nestedatt--gseries_header_add"></a>
### Nested Schema for `gseries_header_add`

Required:

* `types` (Set of String)

<a id="nestedatt--gseries_header_remove"></a>
### Nested Schema for `gseries_header_remove`

Optional:

* `enabled` (String)

<a id="nestedatt--gseries_load_balance"></a>
### Nested Schema for `gseries_load_balance`

Optional:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--gseries_load_balance--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--gseries_load_balance--variable_offset))

<a id="nestedatt--gseries_load_balance--fixed_offset"></a>
### Nested Schema for `gseries_load_balance.fixed_offset`

Required:

* `hash` (String)
* `length` (Number)
* `offset` (Number)

<a id="nestedatt--gseries_load_balance--variable_offset"></a>
### Nested Schema for `gseries_load_balance.variable_offset`

Required:

* `end_delim` (String)
* `hash` (String)
* `start_delim` (String)
* `start_field` (String)

<a id="nestedatt--gseries_pattern_match"></a>
### Nested Schema for `gseries_pattern_match`

Optional:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--gseries_pattern_match--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--gseries_pattern_match--variable_offset))

<a id="nestedatt--gseries_pattern_match--fixed_offset"></a>
### Nested Schema for `gseries_pattern_match.fixed_offset`

Required:

* `length` (Number)
* `offset` (Number)

<a id="nestedatt--gseries_pattern_match--variable_offset"></a>
### Nested Schema for `gseries_pattern_match.variable_offset`

Required:

* `end_delim` (String)
* `start_delim` (String)

<a id="nestedatt--gtp_whitelist"></a>
### Nested Schema for `gtp_whitelist`

Optional:

* `enabled` (String)

<a id="nestedatt--header_add"></a>
### Nested Schema for `header_add`

Required:

* `vlan` (Number)

<a id="nestedatt--header_remove"></a>
### Nested Schema for `header_remove`

Required:

* `protocol` (String) - 'gre' and 'fabricPath' are only applicable for H-series

Optional:

* `ah1` (String) - only valid and required for 'generic'. First anchor header, from which the header to be stripped is occurred.
* `ah2` (String) - only valid and required for 'generic'. next anchor header.
* `custom_len` (Number) - only valid for 'generic'. length of unknown header.
* `erspan_flow_id` (Number) - only applicable for 'erspan', specifies the flow id to strip. Value of 0 will strip all flow ids
* `fp_dst_switch_id` (Number) - Only valid and required for 'fabricPath', 12 bit destination switch id
* `fp_src_switch_id` (Number) - Only valid and required for 'fabricPath', 12 bit source switch id
* `header_count` (Number) - only valid for 'generic'. Number of headers to be stripped.
* `offset` (String) - only valid and required for 'generic'.'start': strip from start of ah1, 'end': strip from end of ah1 or any other offset in the range of length of ah1, 'offsetRangeValue' is required if offset is 'offsetRange'.
* `offset_range_value` (Number) - only valid and required when offset is 'offsetRange', integer within range of size of header
* `timestamp_format` (String) - Timestamp format. Only valid and required for 'fm6000Ts'
* `vlan_header` (String) - Only valid when protocol is 'vlan'. Specifies the target vlan header to strip. Defaults to 'outer'
* `vxlan_id` (Number) - 24-bit value. Only valid when protocol is 'vxlan'. Specifies the vxlan id to strip. Value of '0' will strip all vxlan ids

<a id="nestedatt--icap"></a>
### Nested Schema for `icap`

Required:

* `icap_profile` (String) - Alias of referenced ICAP Profile

<a id="nestedatt--inline_ssl"></a>
### Nested Schema for `inline_ssl`

Required:

* `inline_ssl_profile` (String) - Alias of referenced Inline SSL Profile

<a id="nestedatt--load_balance"></a>
### Nested Schema for `load_balance`

Optional:

* `enhanced` (Attributes) - Private class. Enhanced part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--load_balance--enhanced))
* `stateful` (Attributes) - Private class. Stateful part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--load_balance--stateful))
* `stateless` (Attributes) - Private class. Stateless part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--load_balance--stateless))

<a id="nestedatt--load_balance--enhanced"></a>
### Nested Schema for `load_balance.enhanced`

Required:

* `elb_alias` (String) - elb app alias

<a id="nestedatt--load_balance--stateful"></a>
### Nested Schema for `load_balance.stateful`

Required:

* `app_type` (String) - 'sapf' option is added in release H 4.3.01. 'sip' and 'tunnel' added in release H 5.1. 'diameter' is added in release H 5.5
* `lb_type` (String) - 'gtpKeyHash' is only applicable when 'appType' == 'gtp'. 'sipKeyHash' is only applicable when 'appType' == 'sip'

Optional:

* `diameter_key_hash_type` (String) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyHash'
* `diameter_key_multi_hash_type` (Attributes List) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyMultiHash' (see [below for nested schema](#nestedatt--load_balance--stateful--diameter_key_multi_hash_type))
* `gtp_key_hash_type` (String) - required when 'appType' == 'gtp' and 'lbType' == 'gtpKeyHash'. ignored otherwise
* `sip_key_hash_type` (String) - required when 'appType' == 'sip' and 'lbType' == 'sipKeyHash'. ignored otherwise

<a id="nestedatt--load_balance--stateful--diameter_key_multi_hash_type"></a>
### Nested Schema for `load_balance.stateful.diameter_key_multi_hash_type`

Required:

* `key` (String)

Optional:

* `avp_codevalue` (Number) - required when 'key' == 'avpCode'

<a id="nestedatt--load_balance--stateless"></a>
### Nested Schema for `load_balance.stateless`

Required:

* `hash_fields` (String)

Optional:

* `field_location` (String) - Ignored when 'hashFields' == 'gtpuTeid'. Required otherwise

<a id="nestedatt--masking"></a>
### Nested Schema for `masking`

Required:

* `protocol` (String)

Optional:

* `content_type` (String) - content type that will trigger masking, only valid and required for protocol 'sip'
* `length` (Number) - max length is 9600 on H-series; 9000 on G-series, not applicable for protocol 'sip' valid and required otherwise
* `offset` (Number) - min offset is 0 for 'none'; 1 for protocol-referenced offsets, not applicable for protocol 'sip' valid and required otherwise
* `pattern` (String) - 1-byte hex mask, not applicable for protocol 'sip' valid and required otherwise

<a id="nestedatt--metadata_export"></a>
### Nested Schema for `metadata_export`

Optional:

* `cache` (String) - metadata cache alias. cache should have exporters defined, and the exporters should have applicationProfiles defined

<a id="nestedatt--netflow"></a>
### Nested Schema for `netflow`

Optional:

* `enabled` (String)

<a id="nestedatt--sa_apf"></a>
### Nested Schema for `sa_apf`

Optional:

* `enabled` (String)

<a id="nestedatt--sip_whitelist"></a>
### Nested Schema for `sip_whitelist`

Optional:

* `enabled` (String)

<a id="nestedatt--slicing"></a>
### Nested Schema for `slicing`

Required:

* `offset` (Number) - min offset is 64 for 'none'; 4 for protocol-referenced offsets. required property till H 5.6
* `protocol` (String) - required property till H 5.6

Optional:

* `enhanced` (String) - enhanced-slicing apps alias

<a id="nestedatt--ssl_decrypt"></a>
### Nested Schema for `ssl_decrypt`

Optional:

* `in_port` (Number) - Port number of 0 represents 'any' port
* `out_port` (Number) - Port number of 0 represents 'auto' port

<a id="nestedatt--trailer_add"></a>
### Nested Schema for `trailer_add`

Required:

* `types` (Set of String) - 'crc' is not applicable for G-series

<a id="nestedatt--trailer_remove"></a>
### Nested Schema for `trailer_remove`

Optional:

* `enabled` (String)

<a id="nestedatt--tunnel_decap"></a>
### Nested Schema for `tunnel_decap`

Required:

* `type` (String)

Optional:

* `custom` (Attributes) - only applicable for 'custom', in which case it is required. (see [below for nested schema](#nestedatt--tunnel_decap--custom))
* `erspan_flow_id` (Number) - only applicable for 'erspan', A Flow ID of 0 decapsulates all ERSPAN tunnel traffic regardless of Flow ID
* `gmip_port` (Number) - only applicable for 'gmip', in which case it is required. Specifies the UDP port on which the Tunnel Network port on the receiving GigaVUE H Series is listening. Must match the configuration of the portdst configured on the sending end of the tunnel
* `l2_gre_key` (Number) - only applicable for 'l2gre', in which case it is required.
* `tls_pcapng` (Attributes) - only applicable for 'tls-pcapng', in which case it is required (see [below for nested schema](#nestedatt--tunnel_decap--tls_pcapng))
* `vxlan` (Attributes) - only applicable for 'vxlan', in which case it is required. (see [below for nested schema](#nestedatt--tunnel_decap--vxlan))

<a id="nestedatt--tunnel_decap--custom"></a>
### Nested Schema for `tunnel_decap.custom`

Required:

* `port_dst` (Number) - when specified as 0, no validation will be done in the packet.
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.

<a id="nestedatt--tunnel_decap--tls_pcapng"></a>
### Nested Schema for `tunnel_decap.tls_pcapng`

Optional:

* `decap_key` (String)
* `listener` (String)

<a id="nestedatt--tunnel_decap--vxlan"></a>
### Nested Schema for `tunnel_decap.vxlan`

Required:

* `port_dst` (Number)
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
* `vni` (Number)

<a id="nestedatt--tunnel_encap"></a>
### Nested Schema for `tunnel_encap`

Required:

* `type` (String)

Optional:

* `gmip_config` (Attributes) - Configuration for GMIP Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--tunnel_encap--gmip_config))
* `l2_gre_config` (Attributes) - Configuration for l2GRE Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--tunnel_encap--l2_gre_config))
* `tls_pcapng` (Attributes) - only applicable for 'tls-pcapng', in which case it is required (see [below for nested schema](#nestedatt--tunnel_encap--tls_pcapng))
* `vxlan_config` (Attributes) - Configuration for Vxlan Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--tunnel_encap--vxlan_config))

<a id="nestedatt--tunnel_encap--gmip_config"></a>
### Nested Schema for `tunnel_encap.gmip_config`

Required:

* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `src_port` (Number)

Optional:

* `dscp` (Number)
* `flow_label` (Number)
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `ttl` (Number)

<a id="nestedatt--tunnel_encap--l2_gre_config"></a>
### Nested Schema for `tunnel_encap.l2_gre_config`

Required:

* `key` (Number)

Optional:

* `dscp` (Number)
* `dst_ip` (String) - ip destination, mutually exclusive with 'pgDst'. IPv4 or IPv6.
* `flow_label` (Number)
* `pg_dst` (String) - port group destination alias, mutually exclusive with 'dstIp'
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `session_field` (String) - required with stateful loadBalance when 'appType' is 'tunnel'
* `session_pos` (String) - required if 'sessionField' is specified
* `ttl` (Number)

<a id="nestedatt--tunnel_encap--tls_pcapng"></a>
### Nested Schema for `tunnel_encap.tls_pcapng`

Optional:

* `exporter` (String)
* `exporter_group` (String)

<a id="nestedatt--tunnel_encap--vxlan_config"></a>
### Nested Schema for `tunnel_encap.vxlan_config`

Required:

* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `src_port` (Number)
* `vni` (Number)

Optional:

* `dscp` (Number)
* `ttl` (Number)

