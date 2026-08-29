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
  alias                 = "example"
  cluster_id            = "example"
  config_status         = "example"
  config_status_reasons = [ "example" ]
  decap = [{
    additonalgsops = {
      apf = {
        enabled = "example"
      }
      dedup = {
        enabled = "example"
      }
      diameter_whitelist = {
        enabled = "example"
      }
      flow_filter = {
        type = "example"
      }
      flow_sampling = {
        type = "example"
      }
      gseries_header_add = {
        types = [ "example" ]
      }
      gseries_header_remove = {
        enabled = "example"
      }
      gseries_load_balance = {
        fixed_offset = {
          hash   = "example"
          length = 0
          offset = 0
        }
        variable_offset = {
          end_delim   = "example"
          hash        = "example"
          start_delim = "example"
          start_field = "example"
        }
      }
      gseries_pattern_match = {
        fixed_offset = {
          length = 0
          offset = 0
        }
        variable_offset = {
          end_delim   = "example"
          start_delim = "example"
        }
      }
      gtp_whitelist = {
        enabled = "example"
      }
      header_add = {
        vlan = 0
      }
      header_remove = {
        ah1                = "example"
        ah2                = "example"
        custom_len         = 0
        erspan_flow_id     = 0
        fp_dst_switch_id   = 0
        fp_src_switch_id   = 0
        header_count       = 0
        offset             = "example"
        offset_range_value = 0
        protocol           = "example"
        timestamp_format   = "example"
        vlan_header        = "example"
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
          app_type               = "example"
          diameter_key_hash_type = "example"
          diameter_key_multi_hash_type = [{
            avp_codevalue = 0
            key           = "example"
          }]
          gtp_key_hash_type = "example"
          lb_type           = "example"
          sip_key_hash_type = "example"
        }
        stateless = {
          field_location = "example"
          hash_fields    = "example"
        }
      }
      masking = {
        content_type = "example"
        length       = 0
        offset       = 0
        pattern      = "example"
        protocol     = "example"
      }
      metadata_export = {
        cache = "example"
      }
      netflow = {
        enabled = "example"
      }
      sa_apf = {
        enabled = "example"
      }
      sip_whitelist = {
        enabled = "example"
      }
      slicing = {
        enhanced = "example"
        offset   = 0
        protocol = "example"
      }
      ssl_decrypt = {
        in_port  = 0
        out_port = 0
      }
      trailer_add = {
        types = [ "example" ]
      }
      trailer_remove = {
        enabled = "example"
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
        type = "example"
        vxlan = {
          port_dst = 0
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
          ttl        = 0
        }
        l2_gre_config = {
          dscp          = 0
          dst_ip        = "example"
          flow_label    = 0
          key           = 0
          pg_dst        = "example"
          prec          = 0
          session_field = "example"
          session_pos   = "example"
          ttl           = 0
        }
        tls_pcapng = {
          exporter       = "example"
          exporter_group = "example"
        }
        type = "example"
        vxlan_config = {
          dscp     = 0
          dst_ip   = "example"
          dst_port = 0
          src_port = 0
          ttl      = 0
          vni      = 0
        }
      }
    }
    application_ports = [ 0 ]
    destination_port  = [ "example" ]
    gsop_alias        = "example"
    listener_alias    = "example"
    map_alias         = "example"
    rules = {
      drop_rules = [{
        bidi    = true
        comment = "example"
        ip_rewrite = {
          dst_ip = "example"
          src_ip = "example"
        }
        matches = [ "example" ]
        rewrite = {
          dst_mac = "example"
          src_mac = "example"
        }
        rule_id = 0
        vlan_tag = {
          tag_protocol_id = "example"
          vlan_action     = "example"
          vlan_id         = 0
        }
      }]
      pass_rules = [{
        bidi    = true
        comment = "example"
        ip_rewrite = {
          dst_ip = "example"
          src_ip = "example"
        }
        matches = [ "example" ]
        rewrite = {
          dst_mac = "example"
          src_mac = "example"
        }
        rule_id = 0
        vlan_tag = {
          tag_protocol_id = "example"
          vlan_action     = "example"
          vlan_id         = 0
        }
      }]
    }
  }]
  description = "example"
  dscp        = 0
  encap = {
    additonalgsops = {
      apf = {
        enabled = "example"
      }
      dedup = {
        enabled = "example"
      }
      diameter_whitelist = {
        enabled = "example"
      }
      flow_filter = {
        type = "example"
      }
      flow_sampling = {
        type = "example"
      }
      gseries_header_add = {
        types = [ "example" ]
      }
      gseries_header_remove = {
        enabled = "example"
      }
      gseries_load_balance = {
        fixed_offset = {
          hash   = "example"
          length = 0
          offset = 0
        }
        variable_offset = {
          end_delim   = "example"
          hash        = "example"
          start_delim = "example"
          start_field = "example"
        }
      }
      gseries_pattern_match = {
        fixed_offset = {
          length = 0
          offset = 0
        }
        variable_offset = {
          end_delim   = "example"
          start_delim = "example"
        }
      }
      gtp_whitelist = {
        enabled = "example"
      }
      header_add = {
        vlan = 0
      }
      header_remove = {
        ah1                = "example"
        ah2                = "example"
        custom_len         = 0
        erspan_flow_id     = 0
        fp_dst_switch_id   = 0
        fp_src_switch_id   = 0
        header_count       = 0
        offset             = "example"
        offset_range_value = 0
        protocol           = "example"
        timestamp_format   = "example"
        vlan_header        = "example"
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
          app_type               = "example"
          diameter_key_hash_type = "example"
          diameter_key_multi_hash_type = [{
            avp_codevalue = 0
            key           = "example"
          }]
          gtp_key_hash_type = "example"
          lb_type           = "example"
          sip_key_hash_type = "example"
        }
        stateless = {
          field_location = "example"
          hash_fields    = "example"
        }
      }
      masking = {
        content_type = "example"
        length       = 0
        offset       = 0
        pattern      = "example"
        protocol     = "example"
      }
      metadata_export = {
        cache = "example"
      }
      netflow = {
        enabled = "example"
      }
      sa_apf = {
        enabled = "example"
      }
      sip_whitelist = {
        enabled = "example"
      }
      slicing = {
        enhanced = "example"
        offset   = 0
        protocol = "example"
      }
      ssl_decrypt = {
        in_port  = 0
        out_port = 0
      }
      trailer_add = {
        types = [ "example" ]
      }
      trailer_remove = {
        enabled = "example"
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
        type = "example"
        vxlan = {
          port_dst = 0
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
          ttl        = 0
        }
        l2_gre_config = {
          dscp          = 0
          dst_ip        = "example"
          flow_label    = 0
          key           = 0
          pg_dst        = "example"
          prec          = 0
          session_field = "example"
          session_pos   = "example"
          ttl           = 0
        }
        tls_pcapng = {
          exporter       = "example"
          exporter_group = "example"
        }
        type = "example"
        vxlan_config = {
          dscp     = 0
          dst_ip   = "example"
          dst_port = 0
          src_port = 0
          ttl      = 0
          vni      = 0
        }
      }
    }
    export_configs = [{
      exporter_alias          = "example"
      remote_application_port = 0
      remote_ip               = "example"
      source_application_port = 0
    }]
    exporter_group_alias = "example"
    gsop_alias           = "example"
    map_alias            = "example"
    rules = {
      drop_rules = [{
        bidi    = true
        comment = "example"
        ip_rewrite = {
          dst_ip = "example"
          src_ip = "example"
        }
        matches = [ "example" ]
        rewrite = {
          dst_mac = "example"
          src_mac = "example"
        }
        rule_id = 0
        vlan_tag = {
          tag_protocol_id = "example"
          vlan_action     = "example"
          vlan_id         = 0
        }
      }]
      pass_rules = [{
        bidi    = true
        comment = "example"
        ip_rewrite = {
          dst_ip = "example"
          src_ip = "example"
        }
        matches = [ "example" ]
        rewrite = {
          dst_mac = "example"
          src_mac = "example"
        }
        rule_id = 0
        vlan_tag = {
          tag_protocol_id = "example"
          vlan_action     = "example"
          vlan_id         = 0
        }
      }]
    }
    source_port = [ "example" ]
  }
  gsgroup          = "example"
  ip_interface_in  = "example"
  ip_interface_out = "example"
  local_key_alias  = "example"
  remote_key_alias = "example"
  ssl_profile = {
    cipher  = "example"
    mtls    = "example"
    version = "example"
  }
  ssl_profile_alias = "example"
  tcp_profile = {
    keep_alive_timer = 0
    selective_ack    = "example"
    syn_retries      = 0
  }
  tcp_profile_alias = "example"
  traffic_dir       = "example"
  ttl               = 0
  tunnel_type       = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional)
* `cluster_id` (String, required) - Target Cluster ID
* `config_status` (String, optional)
* `config_status_reasons` (List of String, optional)
* `decap` (Attributes List, optional) (see [below for nested schema](#nestedatt--decap))
* `description` (String, optional)
* `dscp` (Number, optional)
* `encap` (Attributes, optional) (see [below for nested schema](#nestedatt--encap))
* `gsgroup` (String, optional)
* `ip_interface_in` (String, optional) - ipInterface for Decap source
* `ip_interface_out` (String, optional) - ipInterface for Encap destination
* `local_key_alias` (String, optional)
* `remote_key_alias` (String, optional)
* `ssl_profile` (Attributes, optional) (see [below for nested schema](#nestedatt--ssl_profile))
* `ssl_profile_alias` (String, optional)
* `tcp_profile` (Attributes, optional) (see [below for nested schema](#nestedatt--tcp_profile))
* `tcp_profile_alias` (String, optional)
* `traffic_dir` (String, optional) - IN for Decap , OUT for Encap
* `ttl` (Number, optional)
* `tunnel_type` (String, optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `alias` (String, computed)
* `config_status` (String, computed)
* `config_status_reasons` (List of String, computed)
* `decap` (Attributes List, computed) (see [below for nested schema](#nestedatt--decap))
* `description` (String, computed)
* `dscp` (Number, computed)
* `encap` (Attributes, computed) (see [below for nested schema](#nestedatt--encap))
* `gsgroup` (String, computed)
* `ip_interface_in` (String, computed) - ipInterface for Decap source
* `ip_interface_out` (String, computed) - ipInterface for Encap destination
* `local_key_alias` (String, computed)
* `remote_key_alias` (String, computed)
* `ssl_profile` (Attributes, computed) (see [below for nested schema](#nestedatt--ssl_profile))
* `ssl_profile_alias` (String, computed)
* `tcp_profile` (Attributes, computed) (see [below for nested schema](#nestedatt--tcp_profile))
* `tcp_profile_alias` (String, computed)
* `traffic_dir` (String, computed) - IN for Decap , OUT for Encap
* `ttl` (Number, computed)
* `tunnel_type` (String, computed)

<a id="nestedatt--decap"></a>
### Nested Schema for `decap`

Optional:

* `additonalgsops` (Attributes) - GigaSMART Applications for a GSOP. At least one GsApp must be defined for the config to be valid (see [below for nested schema](#nestedatt--decap--additonalgsops))
* `application_ports` (List of Number)
* `destination_port` (List of String) - Tools connected to the cluster
* `gsop_alias` (String)
* `listener_alias` (String)
* `map_alias` (String)
* `rules` (Attributes) - Map Rules Container. Private class (see [below for nested schema](#nestedatt--decap--rules))
<a id="nestedatt--decap--additonalgsops"></a>
### Nested Schema for `decap.additonalgsops`

Optional:

* `apf` (Attributes) (see [below for nested schema](#nestedatt--decap--additonalgsops--apf))
* `dedup` (Attributes) (see [below for nested schema](#nestedatt--decap--additonalgsops--dedup))
* `diameter_whitelist` (Attributes) (see [below for nested schema](#nestedatt--decap--additonalgsops--diameter_whitelist))
* `flow_filter` (Attributes) - GigaSMART 'Flow Filter' Application Configuration (see [below for nested schema](#nestedatt--decap--additonalgsops--flow_filter))
* `flow_sampling` (Attributes) - GigaSMART 'Flow Sampling' Application Configuration (see [below for nested schema](#nestedatt--decap--additonalgsops--flow_sampling))
* `gseries_header_add` (Attributes) - Only applicable for G-series (see [below for nested schema](#nestedatt--decap--additonalgsops--gseries_header_add))
* `gseries_header_remove` (Attributes) (see [below for nested schema](#nestedatt--decap--additonalgsops--gseries_header_remove))
* `gseries_load_balance` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--decap--additonalgsops--gseries_load_balance))
* `gseries_pattern_match` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--decap--additonalgsops--gseries_pattern_match))
* `gtp_whitelist` (Attributes) (see [below for nested schema](#nestedatt--decap--additonalgsops--gtp_whitelist))
* `header_add` (Attributes) - GigaSMART 'Add Header' Application Configuration (see [below for nested schema](#nestedatt--decap--additonalgsops--header_add))
* `header_remove` (Attributes) - GigaSMART 'Remove Header' Application Configuration (see [below for nested schema](#nestedatt--decap--additonalgsops--header_remove))
* `icap` (Attributes) - GigaSMART ICAP Configuration (see [below for nested schema](#nestedatt--decap--additonalgsops--icap))
* `inline_ssl` (Attributes) - GigaSMART Inline SSL Profile Configuration (see [below for nested schema](#nestedatt--decap--additonalgsops--inline_ssl))
* `load_balance` (Attributes) - GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--decap--additonalgsops--load_balance))
* `masking` (Attributes) - GigaSMART 'Masking' Application Configuration (see [below for nested schema](#nestedatt--decap--additonalgsops--masking))
* `metadata_export` (Attributes) (see [below for nested schema](#nestedatt--decap--additonalgsops--metadata_export))
* `netflow` (Attributes) (see [below for nested schema](#nestedatt--decap--additonalgsops--netflow))
* `sa_apf` (Attributes) (see [below for nested schema](#nestedatt--decap--additonalgsops--sa_apf))
* `sip_whitelist` (Attributes) (see [below for nested schema](#nestedatt--decap--additonalgsops--sip_whitelist))
* `slicing` (Attributes) - GigaSMART 'Slicing' Application Configuration (see [below for nested schema](#nestedatt--decap--additonalgsops--slicing))
* `ssl_decrypt` (Attributes) - GigaSMART 'SSL Decrypt' Application Configuration (see [below for nested schema](#nestedatt--decap--additonalgsops--ssl_decrypt))
* `trailer_add` (Attributes) - GigaSMART 'Add Trailer' Application Configuration (see [below for nested schema](#nestedatt--decap--additonalgsops--trailer_add))
* `trailer_remove` (Attributes) (see [below for nested schema](#nestedatt--decap--additonalgsops--trailer_remove))
* `tunnel_decap` (Attributes) - GigaSMART 'Decapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--decap--additonalgsops--tunnel_decap))
* `tunnel_encap` (Attributes) - GigaSMART 'Encapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--decap--additonalgsops--tunnel_encap))
<a id="nestedatt--decap--additonalgsops--apf"></a>
### Nested Schema for `decap.additonalgsops.apf`

Optional:

* `enabled` (String)
<a id="nestedatt--decap--additonalgsops--dedup"></a>
### Nested Schema for `decap.additonalgsops.dedup`

Optional:

* `enabled` (String)
<a id="nestedatt--decap--additonalgsops--diameter_whitelist"></a>
### Nested Schema for `decap.additonalgsops.diameter_whitelist`

Optional:

* `enabled` (String)
<a id="nestedatt--decap--additonalgsops--flow_filter"></a>
### Nested Schema for `decap.additonalgsops.flow_filter`

Required:

* `type` (String)
<a id="nestedatt--decap--additonalgsops--flow_sampling"></a>
### Nested Schema for `decap.additonalgsops.flow_sampling`

Required:

* `type` (String)
<a id="nestedatt--decap--additonalgsops--gseries_header_add"></a>
### Nested Schema for `decap.additonalgsops.gseries_header_add`

Required:

* `types` (Set of String)
<a id="nestedatt--decap--additonalgsops--gseries_header_remove"></a>
### Nested Schema for `decap.additonalgsops.gseries_header_remove`

Optional:

* `enabled` (String)
<a id="nestedatt--decap--additonalgsops--gseries_load_balance"></a>
### Nested Schema for `decap.additonalgsops.gseries_load_balance`

Optional:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--decap--additonalgsops--gseries_load_balance--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--decap--additonalgsops--gseries_load_balance--variable_offset))
<a id="nestedatt--decap--additonalgsops--gseries_load_balance--fixed_offset"></a>
### Nested Schema for `decap.additonalgsops.gseries_load_balance.fixed_offset`

Required:

* `hash` (String)
* `length` (Number)
* `offset` (Number)
<a id="nestedatt--decap--additonalgsops--gseries_load_balance--variable_offset"></a>
### Nested Schema for `decap.additonalgsops.gseries_load_balance.variable_offset`

Required:

* `end_delim` (String)
* `hash` (String)
* `start_delim` (String)
* `start_field` (String)
<a id="nestedatt--decap--additonalgsops--gseries_pattern_match"></a>
### Nested Schema for `decap.additonalgsops.gseries_pattern_match`

Optional:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--decap--additonalgsops--gseries_pattern_match--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--decap--additonalgsops--gseries_pattern_match--variable_offset))
<a id="nestedatt--decap--additonalgsops--gseries_pattern_match--fixed_offset"></a>
### Nested Schema for `decap.additonalgsops.gseries_pattern_match.fixed_offset`

Required:

* `length` (Number)
* `offset` (Number)
<a id="nestedatt--decap--additonalgsops--gseries_pattern_match--variable_offset"></a>
### Nested Schema for `decap.additonalgsops.gseries_pattern_match.variable_offset`

Required:

* `end_delim` (String)
* `start_delim` (String)
<a id="nestedatt--decap--additonalgsops--gtp_whitelist"></a>
### Nested Schema for `decap.additonalgsops.gtp_whitelist`

Optional:

* `enabled` (String)
<a id="nestedatt--decap--additonalgsops--header_add"></a>
### Nested Schema for `decap.additonalgsops.header_add`

Required:

* `vlan` (Number)
<a id="nestedatt--decap--additonalgsops--header_remove"></a>
### Nested Schema for `decap.additonalgsops.header_remove`

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
<a id="nestedatt--decap--additonalgsops--icap"></a>
### Nested Schema for `decap.additonalgsops.icap`

Required:

* `icap_profile` (String) - Alias of referenced ICAP Profile
<a id="nestedatt--decap--additonalgsops--inline_ssl"></a>
### Nested Schema for `decap.additonalgsops.inline_ssl`

Required:

* `inline_ssl_profile` (String) - Alias of referenced Inline SSL Profile
<a id="nestedatt--decap--additonalgsops--load_balance"></a>
### Nested Schema for `decap.additonalgsops.load_balance`

Optional:

* `enhanced` (Attributes) - Private class. Enhanced part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--decap--additonalgsops--load_balance--enhanced))
* `stateful` (Attributes) - Private class. Stateful part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--decap--additonalgsops--load_balance--stateful))
* `stateless` (Attributes) - Private class. Stateless part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--decap--additonalgsops--load_balance--stateless))
<a id="nestedatt--decap--additonalgsops--load_balance--enhanced"></a>
### Nested Schema for `decap.additonalgsops.load_balance.enhanced`

Required:

* `elb_alias` (String) - elb app alias
<a id="nestedatt--decap--additonalgsops--load_balance--stateful"></a>
### Nested Schema for `decap.additonalgsops.load_balance.stateful`

Required:

* `app_type` (String) - 'sapf' option is added in release H 4.3.01. 'sip' and 'tunnel' added in release H 5.1. 'diameter' is added in release H 5.5
* `lb_type` (String) - 'gtpKeyHash' is only applicable when 'appType' == 'gtp'. 'sipKeyHash' is only applicable when 'appType' == 'sip'
Optional:

* `diameter_key_hash_type` (String) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyHash'
* `diameter_key_multi_hash_type` (Attributes List) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyMultiHash' (see [below for nested schema](#nestedatt--decap--additonalgsops--load_balance--stateful--diameter_key_multi_hash_type))
* `gtp_key_hash_type` (String) - required when 'appType' == 'gtp' and 'lbType' == 'gtpKeyHash'. ignored otherwise
* `sip_key_hash_type` (String) - required when 'appType' == 'sip' and 'lbType' == 'sipKeyHash'. ignored otherwise
<a id="nestedatt--decap--additonalgsops--load_balance--stateful--diameter_key_multi_hash_type"></a>
### Nested Schema for `decap.additonalgsops.load_balance.stateful.diameter_key_multi_hash_type`

Required:

* `key` (String)
Optional:

* `avp_codevalue` (Number) - required when 'key' == 'avpCode'
<a id="nestedatt--decap--additonalgsops--load_balance--stateless"></a>
### Nested Schema for `decap.additonalgsops.load_balance.stateless`

Required:

* `hash_fields` (String)
Optional:

* `field_location` (String) - Ignored when 'hashFields' == 'gtpuTeid'. Required otherwise
<a id="nestedatt--decap--additonalgsops--masking"></a>
### Nested Schema for `decap.additonalgsops.masking`

Required:

* `protocol` (String)
Optional:

* `content_type` (String) - content type that will trigger masking, only valid and required for protocol 'sip'
* `length` (Number) - max length is 9600 on H-series; 9000 on G-series, not applicable for protocol 'sip' valid and required otherwise
* `offset` (Number) - min offset is 0 for 'none'; 1 for protocol-referenced offsets, not applicable for protocol 'sip' valid and required otherwise
* `pattern` (String) - 1-byte hex mask, not applicable for protocol 'sip' valid and required otherwise
<a id="nestedatt--decap--additonalgsops--metadata_export"></a>
### Nested Schema for `decap.additonalgsops.metadata_export`

Optional:

* `cache` (String) - metadata cache alias. cache should have exporters defined, and the exporters should have applicationProfiles defined
<a id="nestedatt--decap--additonalgsops--netflow"></a>
### Nested Schema for `decap.additonalgsops.netflow`

Optional:

* `enabled` (String)
<a id="nestedatt--decap--additonalgsops--sa_apf"></a>
### Nested Schema for `decap.additonalgsops.sa_apf`

Optional:

* `enabled` (String)
<a id="nestedatt--decap--additonalgsops--sip_whitelist"></a>
### Nested Schema for `decap.additonalgsops.sip_whitelist`

Optional:

* `enabled` (String)
<a id="nestedatt--decap--additonalgsops--slicing"></a>
### Nested Schema for `decap.additonalgsops.slicing`

Required:

* `offset` (Number) - min offset is 64 for 'none'; 4 for protocol-referenced offsets. required property till H 5.6
* `protocol` (String) - required property till H 5.6
Optional:

* `enhanced` (String) - enhanced-slicing apps alias
<a id="nestedatt--decap--additonalgsops--ssl_decrypt"></a>
### Nested Schema for `decap.additonalgsops.ssl_decrypt`

Optional:

* `in_port` (Number) - Port number of 0 represents 'any' port
* `out_port` (Number) - Port number of 0 represents 'auto' port
<a id="nestedatt--decap--additonalgsops--trailer_add"></a>
### Nested Schema for `decap.additonalgsops.trailer_add`

Required:

* `types` (Set of String) - 'crc' is not applicable for G-series
<a id="nestedatt--decap--additonalgsops--trailer_remove"></a>
### Nested Schema for `decap.additonalgsops.trailer_remove`

Optional:

* `enabled` (String)
<a id="nestedatt--decap--additonalgsops--tunnel_decap"></a>
### Nested Schema for `decap.additonalgsops.tunnel_decap`

Required:

* `type` (String)
Optional:

* `custom` (Attributes) (see [below for nested schema](#nestedatt--decap--additonalgsops--tunnel_decap--custom))
* `erspan_flow_id` (Number) - only applicable for 'erspan', A Flow ID of 0 decapsulates all ERSPAN tunnel traffic regardless of Flow ID
* `gmip_port` (Number) - only applicable for 'gmip', in which case it is required. Specifies the UDP port on which the Tunnel Network port on the receiving GigaVUE H Series is listening. Must match the configuration of the portdst configured on the sending end of the tunnel
* `l2_gre_key` (Number) - only applicable for 'l2gre', in which case it is required.
* `tls_pcapng` (Attributes) (see [below for nested schema](#nestedatt--decap--additonalgsops--tunnel_decap--tls_pcapng))
* `vxlan` (Attributes) (see [below for nested schema](#nestedatt--decap--additonalgsops--tunnel_decap--vxlan))
<a id="nestedatt--decap--additonalgsops--tunnel_decap--custom"></a>
### Nested Schema for `decap.additonalgsops.tunnel_decap.custom`

Required:

* `port_dst` (Number) - when specified as 0, no validation will be done in the packet.
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
<a id="nestedatt--decap--additonalgsops--tunnel_decap--tls_pcapng"></a>
### Nested Schema for `decap.additonalgsops.tunnel_decap.tls_pcapng`

Optional:

* `decap_key` (String)
* `listener` (String)
<a id="nestedatt--decap--additonalgsops--tunnel_decap--vxlan"></a>
### Nested Schema for `decap.additonalgsops.tunnel_decap.vxlan`

Required:

* `port_dst` (Number)
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
* `vni` (Number)
<a id="nestedatt--decap--additonalgsops--tunnel_encap"></a>
### Nested Schema for `decap.additonalgsops.tunnel_encap`

Required:

* `type` (String)
Optional:

* `gmip_config` (Attributes) - Configuration for GMIP Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--decap--additonalgsops--tunnel_encap--gmip_config))
* `l2_gre_config` (Attributes) - Configuration for l2GRE Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--decap--additonalgsops--tunnel_encap--l2_gre_config))
* `tls_pcapng` (Attributes) (see [below for nested schema](#nestedatt--decap--additonalgsops--tunnel_encap--tls_pcapng))
* `vxlan_config` (Attributes) - Configuration for Vxlan Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--decap--additonalgsops--tunnel_encap--vxlan_config))
<a id="nestedatt--decap--additonalgsops--tunnel_encap--gmip_config"></a>
### Nested Schema for `decap.additonalgsops.tunnel_encap.gmip_config`

Required:

* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `src_port` (Number)
Optional:

* `dscp` (Number)
* `flow_label` (Number)
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `ttl` (Number)
<a id="nestedatt--decap--additonalgsops--tunnel_encap--l2_gre_config"></a>
### Nested Schema for `decap.additonalgsops.tunnel_encap.l2_gre_config`

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
<a id="nestedatt--decap--additonalgsops--tunnel_encap--tls_pcapng"></a>
### Nested Schema for `decap.additonalgsops.tunnel_encap.tls_pcapng`

Optional:

* `exporter` (String)
* `exporter_group` (String)
<a id="nestedatt--decap--additonalgsops--tunnel_encap--vxlan_config"></a>
### Nested Schema for `decap.additonalgsops.tunnel_encap.vxlan_config`

Required:

* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `src_port` (Number)
* `vni` (Number)
Optional:

* `dscp` (Number)
* `ttl` (Number)
<a id="nestedatt--decap--rules"></a>
### Nested Schema for `decap.rules`

Optional:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--decap--rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--decap--rules--pass_rules))
<a id="nestedatt--decap--rules--drop_rules"></a>
### Nested Schema for `decap.rules.drop_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)
Optional:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--decap--rules--drop_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--decap--rules--drop_rules--rewrite))
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--decap--rules--drop_rules--vlan_tag))
<a id="nestedatt--decap--rules--drop_rules--ip_rewrite"></a>
### Nested Schema for `decap.rules.drop_rules.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--decap--rules--drop_rules--rewrite"></a>
### Nested Schema for `decap.rules.drop_rules.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--decap--rules--drop_rules--vlan_tag"></a>
### Nested Schema for `decap.rules.drop_rules.vlan_tag`

Required:

* `vlan_action` (String)
Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)
<a id="nestedatt--decap--rules--pass_rules"></a>
### Nested Schema for `decap.rules.pass_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)
Optional:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--decap--rules--pass_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--decap--rules--pass_rules--rewrite))
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--decap--rules--pass_rules--vlan_tag))
<a id="nestedatt--decap--rules--pass_rules--ip_rewrite"></a>
### Nested Schema for `decap.rules.pass_rules.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--decap--rules--pass_rules--rewrite"></a>
### Nested Schema for `decap.rules.pass_rules.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--decap--rules--pass_rules--vlan_tag"></a>
### Nested Schema for `decap.rules.pass_rules.vlan_tag`

Required:

* `vlan_action` (String)
Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)
<a id="nestedatt--encap"></a>
### Nested Schema for `encap`

Optional:

* `additonalgsops` (Attributes) - GigaSMART Applications for a GSOP. At least one GsApp must be defined for the config to be valid (see [below for nested schema](#nestedatt--encap--additonalgsops))
* `export_configs` (Attributes List) (see [below for nested schema](#nestedatt--encap--export_configs))
* `exporter_group_alias` (String)
* `gsop_alias` (String)
* `map_alias` (String)
* `rules` (Attributes) - Map Rules Container. Private class (see [below for nested schema](#nestedatt--encap--rules))
* `source_port` (List of String) - source port for encap traffic , it would be on the same cluster as GsEngine
<a id="nestedatt--encap--additonalgsops"></a>
### Nested Schema for `encap.additonalgsops`

Optional:

* `apf` (Attributes) (see [below for nested schema](#nestedatt--encap--additonalgsops--apf))
* `dedup` (Attributes) (see [below for nested schema](#nestedatt--encap--additonalgsops--dedup))
* `diameter_whitelist` (Attributes) (see [below for nested schema](#nestedatt--encap--additonalgsops--diameter_whitelist))
* `flow_filter` (Attributes) - GigaSMART 'Flow Filter' Application Configuration (see [below for nested schema](#nestedatt--encap--additonalgsops--flow_filter))
* `flow_sampling` (Attributes) - GigaSMART 'Flow Sampling' Application Configuration (see [below for nested schema](#nestedatt--encap--additonalgsops--flow_sampling))
* `gseries_header_add` (Attributes) - Only applicable for G-series (see [below for nested schema](#nestedatt--encap--additonalgsops--gseries_header_add))
* `gseries_header_remove` (Attributes) (see [below for nested schema](#nestedatt--encap--additonalgsops--gseries_header_remove))
* `gseries_load_balance` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--encap--additonalgsops--gseries_load_balance))
* `gseries_pattern_match` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--encap--additonalgsops--gseries_pattern_match))
* `gtp_whitelist` (Attributes) (see [below for nested schema](#nestedatt--encap--additonalgsops--gtp_whitelist))
* `header_add` (Attributes) - GigaSMART 'Add Header' Application Configuration (see [below for nested schema](#nestedatt--encap--additonalgsops--header_add))
* `header_remove` (Attributes) - GigaSMART 'Remove Header' Application Configuration (see [below for nested schema](#nestedatt--encap--additonalgsops--header_remove))
* `icap` (Attributes) - GigaSMART ICAP Configuration (see [below for nested schema](#nestedatt--encap--additonalgsops--icap))
* `inline_ssl` (Attributes) - GigaSMART Inline SSL Profile Configuration (see [below for nested schema](#nestedatt--encap--additonalgsops--inline_ssl))
* `load_balance` (Attributes) - GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--encap--additonalgsops--load_balance))
* `masking` (Attributes) - GigaSMART 'Masking' Application Configuration (see [below for nested schema](#nestedatt--encap--additonalgsops--masking))
* `metadata_export` (Attributes) (see [below for nested schema](#nestedatt--encap--additonalgsops--metadata_export))
* `netflow` (Attributes) (see [below for nested schema](#nestedatt--encap--additonalgsops--netflow))
* `sa_apf` (Attributes) (see [below for nested schema](#nestedatt--encap--additonalgsops--sa_apf))
* `sip_whitelist` (Attributes) (see [below for nested schema](#nestedatt--encap--additonalgsops--sip_whitelist))
* `slicing` (Attributes) - GigaSMART 'Slicing' Application Configuration (see [below for nested schema](#nestedatt--encap--additonalgsops--slicing))
* `ssl_decrypt` (Attributes) - GigaSMART 'SSL Decrypt' Application Configuration (see [below for nested schema](#nestedatt--encap--additonalgsops--ssl_decrypt))
* `trailer_add` (Attributes) - GigaSMART 'Add Trailer' Application Configuration (see [below for nested schema](#nestedatt--encap--additonalgsops--trailer_add))
* `trailer_remove` (Attributes) (see [below for nested schema](#nestedatt--encap--additonalgsops--trailer_remove))
* `tunnel_decap` (Attributes) - GigaSMART 'Decapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--encap--additonalgsops--tunnel_decap))
* `tunnel_encap` (Attributes) - GigaSMART 'Encapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--encap--additonalgsops--tunnel_encap))
<a id="nestedatt--encap--additonalgsops--apf"></a>
### Nested Schema for `encap.additonalgsops.apf`

Optional:

* `enabled` (String)
<a id="nestedatt--encap--additonalgsops--dedup"></a>
### Nested Schema for `encap.additonalgsops.dedup`

Optional:

* `enabled` (String)
<a id="nestedatt--encap--additonalgsops--diameter_whitelist"></a>
### Nested Schema for `encap.additonalgsops.diameter_whitelist`

Optional:

* `enabled` (String)
<a id="nestedatt--encap--additonalgsops--flow_filter"></a>
### Nested Schema for `encap.additonalgsops.flow_filter`

Required:

* `type` (String)
<a id="nestedatt--encap--additonalgsops--flow_sampling"></a>
### Nested Schema for `encap.additonalgsops.flow_sampling`

Required:

* `type` (String)
<a id="nestedatt--encap--additonalgsops--gseries_header_add"></a>
### Nested Schema for `encap.additonalgsops.gseries_header_add`

Required:

* `types` (Set of String)
<a id="nestedatt--encap--additonalgsops--gseries_header_remove"></a>
### Nested Schema for `encap.additonalgsops.gseries_header_remove`

Optional:

* `enabled` (String)
<a id="nestedatt--encap--additonalgsops--gseries_load_balance"></a>
### Nested Schema for `encap.additonalgsops.gseries_load_balance`

Optional:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--encap--additonalgsops--gseries_load_balance--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--encap--additonalgsops--gseries_load_balance--variable_offset))
<a id="nestedatt--encap--additonalgsops--gseries_load_balance--fixed_offset"></a>
### Nested Schema for `encap.additonalgsops.gseries_load_balance.fixed_offset`

Required:

* `hash` (String)
* `length` (Number)
* `offset` (Number)
<a id="nestedatt--encap--additonalgsops--gseries_load_balance--variable_offset"></a>
### Nested Schema for `encap.additonalgsops.gseries_load_balance.variable_offset`

Required:

* `end_delim` (String)
* `hash` (String)
* `start_delim` (String)
* `start_field` (String)
<a id="nestedatt--encap--additonalgsops--gseries_pattern_match"></a>
### Nested Schema for `encap.additonalgsops.gseries_pattern_match`

Optional:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--encap--additonalgsops--gseries_pattern_match--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--encap--additonalgsops--gseries_pattern_match--variable_offset))
<a id="nestedatt--encap--additonalgsops--gseries_pattern_match--fixed_offset"></a>
### Nested Schema for `encap.additonalgsops.gseries_pattern_match.fixed_offset`

Required:

* `length` (Number)
* `offset` (Number)
<a id="nestedatt--encap--additonalgsops--gseries_pattern_match--variable_offset"></a>
### Nested Schema for `encap.additonalgsops.gseries_pattern_match.variable_offset`

Required:

* `end_delim` (String)
* `start_delim` (String)
<a id="nestedatt--encap--additonalgsops--gtp_whitelist"></a>
### Nested Schema for `encap.additonalgsops.gtp_whitelist`

Optional:

* `enabled` (String)
<a id="nestedatt--encap--additonalgsops--header_add"></a>
### Nested Schema for `encap.additonalgsops.header_add`

Required:

* `vlan` (Number)
<a id="nestedatt--encap--additonalgsops--header_remove"></a>
### Nested Schema for `encap.additonalgsops.header_remove`

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
<a id="nestedatt--encap--additonalgsops--icap"></a>
### Nested Schema for `encap.additonalgsops.icap`

Required:

* `icap_profile` (String) - Alias of referenced ICAP Profile
<a id="nestedatt--encap--additonalgsops--inline_ssl"></a>
### Nested Schema for `encap.additonalgsops.inline_ssl`

Required:

* `inline_ssl_profile` (String) - Alias of referenced Inline SSL Profile
<a id="nestedatt--encap--additonalgsops--load_balance"></a>
### Nested Schema for `encap.additonalgsops.load_balance`

Optional:

* `enhanced` (Attributes) - Private class. Enhanced part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--encap--additonalgsops--load_balance--enhanced))
* `stateful` (Attributes) - Private class. Stateful part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--encap--additonalgsops--load_balance--stateful))
* `stateless` (Attributes) - Private class. Stateless part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--encap--additonalgsops--load_balance--stateless))
<a id="nestedatt--encap--additonalgsops--load_balance--enhanced"></a>
### Nested Schema for `encap.additonalgsops.load_balance.enhanced`

Required:

* `elb_alias` (String) - elb app alias
<a id="nestedatt--encap--additonalgsops--load_balance--stateful"></a>
### Nested Schema for `encap.additonalgsops.load_balance.stateful`

Required:

* `app_type` (String) - 'sapf' option is added in release H 4.3.01. 'sip' and 'tunnel' added in release H 5.1. 'diameter' is added in release H 5.5
* `lb_type` (String) - 'gtpKeyHash' is only applicable when 'appType' == 'gtp'. 'sipKeyHash' is only applicable when 'appType' == 'sip'
Optional:

* `diameter_key_hash_type` (String) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyHash'
* `diameter_key_multi_hash_type` (Attributes List) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyMultiHash' (see [below for nested schema](#nestedatt--encap--additonalgsops--load_balance--stateful--diameter_key_multi_hash_type))
* `gtp_key_hash_type` (String) - required when 'appType' == 'gtp' and 'lbType' == 'gtpKeyHash'. ignored otherwise
* `sip_key_hash_type` (String) - required when 'appType' == 'sip' and 'lbType' == 'sipKeyHash'. ignored otherwise
<a id="nestedatt--encap--additonalgsops--load_balance--stateful--diameter_key_multi_hash_type"></a>
### Nested Schema for `encap.additonalgsops.load_balance.stateful.diameter_key_multi_hash_type`

Required:

* `key` (String)
Optional:

* `avp_codevalue` (Number) - required when 'key' == 'avpCode'
<a id="nestedatt--encap--additonalgsops--load_balance--stateless"></a>
### Nested Schema for `encap.additonalgsops.load_balance.stateless`

Required:

* `hash_fields` (String)
Optional:

* `field_location` (String) - Ignored when 'hashFields' == 'gtpuTeid'. Required otherwise
<a id="nestedatt--encap--additonalgsops--masking"></a>
### Nested Schema for `encap.additonalgsops.masking`

Required:

* `protocol` (String)
Optional:

* `content_type` (String) - content type that will trigger masking, only valid and required for protocol 'sip'
* `length` (Number) - max length is 9600 on H-series; 9000 on G-series, not applicable for protocol 'sip' valid and required otherwise
* `offset` (Number) - min offset is 0 for 'none'; 1 for protocol-referenced offsets, not applicable for protocol 'sip' valid and required otherwise
* `pattern` (String) - 1-byte hex mask, not applicable for protocol 'sip' valid and required otherwise
<a id="nestedatt--encap--additonalgsops--metadata_export"></a>
### Nested Schema for `encap.additonalgsops.metadata_export`

Optional:

* `cache` (String) - metadata cache alias. cache should have exporters defined, and the exporters should have applicationProfiles defined
<a id="nestedatt--encap--additonalgsops--netflow"></a>
### Nested Schema for `encap.additonalgsops.netflow`

Optional:

* `enabled` (String)
<a id="nestedatt--encap--additonalgsops--sa_apf"></a>
### Nested Schema for `encap.additonalgsops.sa_apf`

Optional:

* `enabled` (String)
<a id="nestedatt--encap--additonalgsops--sip_whitelist"></a>
### Nested Schema for `encap.additonalgsops.sip_whitelist`

Optional:

* `enabled` (String)
<a id="nestedatt--encap--additonalgsops--slicing"></a>
### Nested Schema for `encap.additonalgsops.slicing`

Required:

* `offset` (Number) - min offset is 64 for 'none'; 4 for protocol-referenced offsets. required property till H 5.6
* `protocol` (String) - required property till H 5.6
Optional:

* `enhanced` (String) - enhanced-slicing apps alias
<a id="nestedatt--encap--additonalgsops--ssl_decrypt"></a>
### Nested Schema for `encap.additonalgsops.ssl_decrypt`

Optional:

* `in_port` (Number) - Port number of 0 represents 'any' port
* `out_port` (Number) - Port number of 0 represents 'auto' port
<a id="nestedatt--encap--additonalgsops--trailer_add"></a>
### Nested Schema for `encap.additonalgsops.trailer_add`

Required:

* `types` (Set of String) - 'crc' is not applicable for G-series
<a id="nestedatt--encap--additonalgsops--trailer_remove"></a>
### Nested Schema for `encap.additonalgsops.trailer_remove`

Optional:

* `enabled` (String)
<a id="nestedatt--encap--additonalgsops--tunnel_decap"></a>
### Nested Schema for `encap.additonalgsops.tunnel_decap`

Required:

* `type` (String)
Optional:

* `custom` (Attributes) (see [below for nested schema](#nestedatt--encap--additonalgsops--tunnel_decap--custom))
* `erspan_flow_id` (Number) - only applicable for 'erspan', A Flow ID of 0 decapsulates all ERSPAN tunnel traffic regardless of Flow ID
* `gmip_port` (Number) - only applicable for 'gmip', in which case it is required. Specifies the UDP port on which the Tunnel Network port on the receiving GigaVUE H Series is listening. Must match the configuration of the portdst configured on the sending end of the tunnel
* `l2_gre_key` (Number) - only applicable for 'l2gre', in which case it is required.
* `tls_pcapng` (Attributes) (see [below for nested schema](#nestedatt--encap--additonalgsops--tunnel_decap--tls_pcapng))
* `vxlan` (Attributes) (see [below for nested schema](#nestedatt--encap--additonalgsops--tunnel_decap--vxlan))
<a id="nestedatt--encap--additonalgsops--tunnel_decap--custom"></a>
### Nested Schema for `encap.additonalgsops.tunnel_decap.custom`

Required:

* `port_dst` (Number) - when specified as 0, no validation will be done in the packet.
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
<a id="nestedatt--encap--additonalgsops--tunnel_decap--tls_pcapng"></a>
### Nested Schema for `encap.additonalgsops.tunnel_decap.tls_pcapng`

Optional:

* `decap_key` (String)
* `listener` (String)
<a id="nestedatt--encap--additonalgsops--tunnel_decap--vxlan"></a>
### Nested Schema for `encap.additonalgsops.tunnel_decap.vxlan`

Required:

* `port_dst` (Number)
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
* `vni` (Number)
<a id="nestedatt--encap--additonalgsops--tunnel_encap"></a>
### Nested Schema for `encap.additonalgsops.tunnel_encap`

Required:

* `type` (String)
Optional:

* `gmip_config` (Attributes) - Configuration for GMIP Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--encap--additonalgsops--tunnel_encap--gmip_config))
* `l2_gre_config` (Attributes) - Configuration for l2GRE Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--encap--additonalgsops--tunnel_encap--l2_gre_config))
* `tls_pcapng` (Attributes) (see [below for nested schema](#nestedatt--encap--additonalgsops--tunnel_encap--tls_pcapng))
* `vxlan_config` (Attributes) - Configuration for Vxlan Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--encap--additonalgsops--tunnel_encap--vxlan_config))
<a id="nestedatt--encap--additonalgsops--tunnel_encap--gmip_config"></a>
### Nested Schema for `encap.additonalgsops.tunnel_encap.gmip_config`

Required:

* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `src_port` (Number)
Optional:

* `dscp` (Number)
* `flow_label` (Number)
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `ttl` (Number)
<a id="nestedatt--encap--additonalgsops--tunnel_encap--l2_gre_config"></a>
### Nested Schema for `encap.additonalgsops.tunnel_encap.l2_gre_config`

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
<a id="nestedatt--encap--additonalgsops--tunnel_encap--tls_pcapng"></a>
### Nested Schema for `encap.additonalgsops.tunnel_encap.tls_pcapng`

Optional:

* `exporter` (String)
* `exporter_group` (String)
<a id="nestedatt--encap--additonalgsops--tunnel_encap--vxlan_config"></a>
### Nested Schema for `encap.additonalgsops.tunnel_encap.vxlan_config`

Required:

* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `src_port` (Number)
* `vni` (Number)
Optional:

* `dscp` (Number)
* `ttl` (Number)
<a id="nestedatt--encap--export_configs"></a>
### Nested Schema for `encap.export_configs`

Optional:

* `exporter_alias` (String)
* `remote_application_port` (Number)
* `remote_ip` (String)
* `source_application_port` (Number)
<a id="nestedatt--encap--rules"></a>
### Nested Schema for `encap.rules`

Optional:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--encap--rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--encap--rules--pass_rules))
<a id="nestedatt--encap--rules--drop_rules"></a>
### Nested Schema for `encap.rules.drop_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)
Optional:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--encap--rules--drop_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--encap--rules--drop_rules--rewrite))
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--encap--rules--drop_rules--vlan_tag))
<a id="nestedatt--encap--rules--drop_rules--ip_rewrite"></a>
### Nested Schema for `encap.rules.drop_rules.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--encap--rules--drop_rules--rewrite"></a>
### Nested Schema for `encap.rules.drop_rules.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--encap--rules--drop_rules--vlan_tag"></a>
### Nested Schema for `encap.rules.drop_rules.vlan_tag`

Required:

* `vlan_action` (String)
Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)
<a id="nestedatt--encap--rules--pass_rules"></a>
### Nested Schema for `encap.rules.pass_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)
Optional:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--encap--rules--pass_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--encap--rules--pass_rules--rewrite))
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--encap--rules--pass_rules--vlan_tag))
<a id="nestedatt--encap--rules--pass_rules--ip_rewrite"></a>
### Nested Schema for `encap.rules.pass_rules.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--encap--rules--pass_rules--rewrite"></a>
### Nested Schema for `encap.rules.pass_rules.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--encap--rules--pass_rules--vlan_tag"></a>
### Nested Schema for `encap.rules.pass_rules.vlan_tag`

Required:

* `vlan_action` (String)
Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)
<a id="nestedatt--ssl_profile"></a>
### Nested Schema for `ssl_profile`

Optional:

* `cipher` (String)
* `mtls` (String)
* `version` (String)
<a id="nestedatt--tcp_profile"></a>
### Nested Schema for `tcp_profile`

Optional:

* `keep_alive_timer` (Number)
* `selective_ack` (String)
* `syn_retries` (Number)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_tunnel_application.example {alias}
```
