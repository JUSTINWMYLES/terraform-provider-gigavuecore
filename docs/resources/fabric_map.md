---
page_title: "gigavuecore_fabric_map Resource - gigavuecore"
subcategory: ""
description: |-
  Get a user-defined fabric map by alias
---

# gigavuecore_fabric_map Resource

Get a user-defined fabric map by alias

## Example Usage

```terraform
resource "gigavuecore_fabric_map" "example" {
  afm_map = {
    alias = "example"
    ap_rules = {
      drop_rules = [{
        application_profile = "example"
        rule_id             = 0
      }]
      pass_rules = [{
        application_profile = "example"
        rule_id             = 0
      }]
    }
    cluster_id        = "example"
    comment           = "example"
    dst_ports         = [ "example" ]
    egress_gigastream = [ "example" ]
    enable            = true
    encap_tunnel      = "example"
    flex_inline = {
      a_to_b = {
        ib_pathway = "example"
        tools      = [ "example" ]
        type       = "example"
      }
      b_to_a = {
        ib_pathway = "example"
        tools      = [ "example" ]
        type       = "example"
      }
      oob_copy = [{
        direction = "example"
        dst_ports = [ "example" ]
        src_ports = [ "example" ]
        tag = {
          type = "example"
        }
      }]
      svt_mode = true
      svt_tag  = 0
      tag = {
        tag_protocol_id = "example"
        type            = "example"
        vlan_id         = 0
      }
    }
    flex_inline_failover = "example"
    flex_inline_vlan_id  = 0
    flow_rules = {
      drop_rules = [{
        gtp = {
          imei      = "example"
          imsi      = "example"
          interface = "example"
          msisdn    = "example"
          version   = "example"
        }
        rule_id = 0
      }]
      pass_rules = [{
        gtp = {
          imei      = "example"
          imsi      = "example"
          interface = "example"
          msisdn    = "example"
          version   = "example"
        }
        rule_id = 0
      }]
    }
    flow_sample5_g_overlap_rules = {
      pass_rules = [{
        comment = "example"
        flow5_g = {
          dnn      = "example"
          gpsi     = "example"
          nas_5_qi = "example"
          nci      = "example"
          nsiid    = "example"
          pei      = "example"
          plmn_id  = "example"
          supi     = "example"
          tac      = "example"
        }
        percentage = 0
        rule_id    = 0
      }]
    }
    flow_sample5_g_rules = {
      pass_rules = [{
        comment = "example"
        flow5_g = {
          dnn     = "example"
          gpsi    = "example"
          nci     = "example"
          nsiid   = "example"
          pei     = "example"
          plmn_id = "example"
          supi    = "example"
          tac     = "example"
        }
        percentage = 0
        priority   = 0
        rule_id    = 0
      }]
    }
    flow_sample_diameter_rules = {
      pass_rules = [{
        diameter = {
          user_name = "example"
        }
        interface  = "example"
        percentage = 0
        rule_id    = 0
      }]
    }
    flow_sample_overlap_rules = {
      pass_rules = [{
        comment = "example"
        gtp = {
          apn       = "example"
          eci       = "example"
          imei      = "example"
          imsi      = "example"
          interface = "example"
          msisdn    = "example"
          nas_5_qi  = "example"
          nci       = "example"
          plmn_id   = "example"
          qci       = 0
          snssai    = "example"
          tac       = "example"
          tac_5_g   = "example"
          version   = "example"
        }
        percentage      = 0
        periodic_recalc = true
        priority        = 0
        rule_id         = 0
      }]
    }
    flow_sample_rules = {
      pass_rules = [{
        comment = "example"
        gtp = {
          apn       = "example"
          eci       = "example"
          imei      = "example"
          imsi      = "example"
          interface = "example"
          msisdn    = "example"
          nas_5_qi  = "example"
          nci       = "example"
          plmn_id   = "example"
          qci       = 0
          snssai    = "example"
          tac       = "example"
          tac_5_g   = "example"
          version   = "example"
        }
        percentage      = 0
        periodic_recalc = true
        priority        = 0
        rule_id         = 0
      }]
    }
    flow_sample_sip_rules = {
      pass_rules = [{
        percentage = 0
        rule_id    = 0
        sip = {
          callee_id = "example"
          callee_id_range = {
            max_value = "example"
            value     = "example"
          }
          caller_id = "example"
          caller_id_range = {
            max_value = "example"
            value     = "example"
          }
          id_range = {
            max_value = "example"
            value     = "example"
          }
        }
      }]
    }
    flow_whitelist5_g_overlap_rules = {
      dnn  = "example"
      type = "example"
    }
    flow_whitelist5_g_rules = {
      dnn                 = "example"
      type                = "example"
      whitelist_databases = [ "example" ]
    }
    flow_whitelist_overlap_rules = {
      pass_rules = [{
        flow5_g = {
          dnn                 = "example"
          type                = "example"
          whitelist_databases = [ "example" ]
        }
        gtp = {
          apn                 = "example"
          interface           = "example"
          type                = "example"
          version             = "example"
          whitelist_databases = [ "example" ]
        }
        rule_id = 0
        sip = {
          type = "example"
        }
      }]
    }
    flow_whitelist_rules = {
      pass_rules = [{
        flow5_g = {
          dnn                 = "example"
          type                = "example"
          whitelist_databases = [ "example" ]
        }
        gtp = {
          apn                 = "example"
          interface           = "example"
          type                = "example"
          version             = "example"
          whitelist_databases = [ "example" ]
        }
        rule_id = 0
        sip = {
          type = "example"
        }
      }]
    }
    fstype = {
      offset = 0
      timer  = 0
      type   = "example"
    }
    gs_rules = {
      drop_rules = [{
        comment = "example"
        matches = [ "example" ]
        rule_id = 0
      }]
      pass_rules = [{
        comment = "example"
        matches = [ "example" ]
        rule_id = 0
      }]
    }
    gsop                = "example"
    inline_traffic_path = "example"
    inline_traffic_type = "example"
    ip_rewrite = {
      dst_ip = "example"
      src_ip = "example"
    }
    mod_time      = 0
    null_dst_port = true
    order         = 0
    rewrite = {
      dst_mac = "example"
      src_mac = "example"
    }
    roles = {
      editors   = [ "example" ]
      listeners = [ "example" ]
      owners    = [ "example" ]
      viewers   = [ "example" ]
    }
    rule_matching = "example"
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
    rx_cluster_ports = [ "example" ]
    src_ports        = [ "example" ]
    sub_type         = "example"
    traffic_type     = "example"
    tx_cluster_ports = [ "example" ]
    type             = "example"
    vlan_tag = {
      tag_protocol_id = "example"
      vlan_action     = "example"
      vlan_id         = 0
    }
  }
  child_map_aliases  = [ "example" ]
  config_status      = "example"
  decap_aliases      = [ "example" ]
  encap_aliases      = [ "example" ]
  error_message      = "example"
  parent_map_aliases = [ "example" ]
}
```

## Schema

### Arguments

The following arguments are supported:

* `afm_map` (Attributes, required) (see [below for nested schema](#nestedatt--afm_map))
* `child_map_aliases` (List of String, optional) - Aliases of the fabric maps that support this fabric map. Applicable only to user-defined fabric maps. For internally generated fabric maps, there is no child maps. This list is empty.
* `config_status` (String, optional) - Configuration status of this fabric map.
* `decap_aliases` (List of String, optional) - Aliases of the decap tunnel endpoints supporting this fabric map. Applicable only to internally generated fabric maps.
* `encap_aliases` (List of String, optional) - Aliases of the encap tunnel endpoints supporting this fabric map. Applicable only to internally generated fabric maps.
* `error_message` (String, optional) - In case of configuration failure, this message provides details about the possible cause of the failure.
* `parent_map_aliases` (List of String, optional) - Aliases of the fabric maps that this fabric map supports. Applicable only to internally generated fabric maps that support user-defined fabric maps. For user-defined fabric maps, there is no parent maps. This list is empty.

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `alias` (String, computed) - alias of the fabric map
* `child_map_aliases` (List of String, computed) - Aliases of the fabric maps that support this fabric map. Applicable only to user-defined fabric maps. For internally generated fabric maps, there is no child maps. This list is empty.
* `config_status` (String, computed) - Configuration status of this fabric map.
* `creation_time` (Number, computed) - Fabric map creation timestamp in UTC
* `decap_aliases` (List of String, computed) - Aliases of the decap tunnel endpoints supporting this fabric map. Applicable only to internally generated fabric maps.
* `encap_aliases` (List of String, computed) - Aliases of the encap tunnel endpoints supporting this fabric map. Applicable only to internally generated fabric maps.
* `error_message` (String, computed) - In case of configuration failure, this message provides details about the possible cause of the failure.
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `parent_map_aliases` (List of String, computed) - Aliases of the fabric maps that this fabric map supports. Applicable only to internally generated fabric maps that support user-defined fabric maps. For user-defined fabric maps, there is no parent maps. This list is empty.
* `pending_reason` (String, computed)
* `updated_time` (Number, computed) - Fabric map updated timestamp in UTC

<a id="nestedatt--afm_map"></a>
### Nested Schema for `afm_map`

Required:

* `alias` (String) - unique map alias
* `dst_ports` (List of String) - List of the 'to' ports. Invalid if 'nullDstPort' is true'. Only port number is supported
* `src_ports` (List of String) - list of the 'from' ports. Only port number is supported
* `type` (String) - 'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'flexInline' is from inline network to tools; 'transitLevel' is from vport to vport
Optional:

* `ap_rules` (Attributes) - pass and drop application profile rules (see [below for nested schema](#nestedatt--afm_map--ap_rules))
* `cluster_id` (String) - id of the defining cluster
* `comment` (String)
* `egress_gigastream` (List of String) - Applicable to fabric map only. It is to specify gigastream(s) to use as egress when going across cluster. The format of array member is clusterID:gigastreamAlias
* `enable` (Boolean) - enable/disable map, applicable only to first level maps
* `encap_tunnel` (String) - tunnel alias
* `flex_inline` (Attributes) - When 'oobCopy' is configured at least one of 'aToB' or 'bToA' must be configured (see [below for nested schema](#nestedatt--afm_map--flex_inline))
* `flex_inline_failover` (String) - only valid for flexInline maps
* `flex_inline_vlan_id` (Number) - VLAN ID carried in the VLAN tag of packets coming from the inline-network port(s). valid and applicable only with 'flexInline' map type, 'srcPorts' should be either inline-network or ib-pathway or vport
* `flow_rules` (Attributes) - Map Flow Rules Container. Private class (see [below for nested schema](#nestedatt--afm_map--flow_rules))
* `flow_sample5_g_overlap_rules` (Attributes) - Map Flow Sample 5g Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--afm_map--flow_sample5_g_overlap_rules))
* `flow_sample5_g_rules` (Attributes) - Map Flow Sample 5g Rules Container. Private class (see [below for nested schema](#nestedatt--afm_map--flow_sample5_g_rules))
* `flow_sample_diameter_rules` (Attributes) - Map Flow Sample Diameter Rules Container. Private class (see [below for nested schema](#nestedatt--afm_map--flow_sample_diameter_rules))
* `flow_sample_overlap_rules` (Attributes) - Map Flow Sample Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--afm_map--flow_sample_overlap_rules))
* `flow_sample_rules` (Attributes) - Map Flow Sample Rules Container. Private class (see [below for nested schema](#nestedatt--afm_map--flow_sample_rules))
* `flow_sample_sip_rules` (Attributes) - Map Flow Sample Sip Rules Container. Private class (see [below for nested schema](#nestedatt--afm_map--flow_sample_sip_rules))
* `flow_whitelist5_g_overlap_rules` (Attributes) - Map Flow Whitelist 5g Overlap Rule match Definition. Private class (see [below for nested schema](#nestedatt--afm_map--flow_whitelist5_g_overlap_rules))
* `flow_whitelist5_g_rules` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--afm_map--flow_whitelist5_g_rules))
* `flow_whitelist_overlap_rules` (Attributes) - Map Flow Whitelist Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--afm_map--flow_whitelist_overlap_rules))
* `flow_whitelist_rules` (Attributes) - Map Flow Whitelist Rules Container. Private class (see [below for nested schema](#nestedatt--afm_map--flow_whitelist_rules))
* `fstype` (Attributes) - Type of Mobility Flowsampling & properties (see [below for nested schema](#nestedatt--afm_map--fstype))
* `gs_rules` (Attributes) - Map GigaSMART Rules Container. Private class (see [below for nested schema](#nestedatt--afm_map--gs_rules))
* `gsop` (String) - Alias of referenced GSOP. Applicable to 'regular' and 'secondLevel' map types
* `inline_traffic_path` (String) - Only applicable for 'inline' map types, in which case defaults to 'normal'. When set to 'bypass', the 'dstPorts' must be empty
* `inline_traffic_type` (String) - Only applicable for 'inline/passAll' map types, in which case defaults to 'symmetric'. For 'asymmetric' maps, 'srcPort' must be of type 'inline-net'
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--afm_map--ip_rewrite))
* `mod_time` (Number) - Last modification time of the map in milliseconds since the epoch
* `null_dst_port` (Boolean) - enabled when the dstPort is null
* `order` (Number) - relative order within per-source port map chain
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--afm_map--rewrite))
* `roles` (Attributes) - Map Roles Container. Private class (see [below for nested schema](#nestedatt--afm_map--roles))
* `rule_matching` (String) - If set to 'blacklist', packet are passed when no 'drop' rules are matched. This field is only valid for 'regular/byRule' map types. (maps into 'no-rule-match' CLI command)
* `rules` (Attributes) - Map Rules Container. Private class (see [below for nested schema](#nestedatt--afm_map--rules))
* `rx_cluster_ports` (List of String)
* `sub_type` (String) - 'byRule' is applicable to all map types; 'collector' is applicable to 'regular', 'inline', 'flexInline' and 'secondLevel' maps; 'passAll' is applicable to 'regular' and 'inline' maps; 'flowFilter', 'flowSample', 'flowWhitelist', 'flowSampleSip', 'flowWhitelistSip', 'flowSampleDiameter', 'flowWhitelistDiameter', 'flowSampleOverlap', 'flowWhitelistOverlap', 'flowSample5g','flowWhitelist5g','flowSample5gOverlap' and 'flowWhitelist5gOverlap' are applicable to 'secondLevel' maps
* `traffic_type` (String) - Only applicable for 'firstLevel/byRule' map types, in which case defaults to 'user'
* `tx_cluster_ports` (List of String)
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--afm_map--vlan_tag))
Read-Only:

* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--afm_map--health_state_reasons))
* `updated_time` (Number) - Last Updated time of the fabric map
<a id="nestedatt--afm_map--ap_rules"></a>
### Nested Schema for `afm_map.ap_rules`

Optional:

* `drop_rules` (Attributes List) (see [below for nested schema](#nestedatt--afm_map--ap_rules--drop_rules))
* `pass_rules` (Attributes List) (see [below for nested schema](#nestedatt--afm_map--ap_rules--pass_rules))
<a id="nestedatt--afm_map--ap_rules--drop_rules"></a>
### Nested Schema for `afm_map.ap_rules.drop_rules`

Required:

* `application_profile` (String) - application profile alias
* `rule_id` (Number) - application profile rule Id, should not have same id as gsRules
<a id="nestedatt--afm_map--ap_rules--pass_rules"></a>
### Nested Schema for `afm_map.ap_rules.pass_rules`

Required:

* `application_profile` (String) - application profile alias
* `rule_id` (Number) - application profile rule Id, should not have same id as gsRules
<a id="nestedatt--afm_map--flex_inline"></a>
### Nested Schema for `afm_map.flex_inline`

Optional:

* `a_to_b` (Attributes) (see [below for nested schema](#nestedatt--afm_map--flex_inline--a_to_b))
* `b_to_a` (Attributes) (see [below for nested schema](#nestedatt--afm_map--flex_inline--b_to_a))
* `oob_copy` (Attributes List) (see [below for nested schema](#nestedatt--afm_map--flex_inline--oob_copy))
* `svt_mode` (Boolean)
* `svt_tag` (Number) - only applicable when svtMode is enabled
* `tag` (Attributes) (see [below for nested schema](#nestedatt--afm_map--flex_inline--tag))
<a id="nestedatt--afm_map--flex_inline--a_to_b"></a>
### Nested Schema for `afm_map.flex_inline.a_to_b`

Required:

* `type` (String) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided
Optional:

* `ib_pathway` (String) - ibPathway alias. Only applicable when type is 'ibPathway'
* `tools` (List of String) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'
<a id="nestedatt--afm_map--flex_inline--b_to_a"></a>
### Nested Schema for `afm_map.flex_inline.b_to_a`

Required:

* `type` (String) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided
Optional:

* `ib_pathway` (String) - ibPathway alias. Only applicable when type is 'ibPathway'
* `tools` (List of String) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'
<a id="nestedatt--afm_map--flex_inline--oob_copy"></a>
### Nested Schema for `afm_map.flex_inline.oob_copy`

Required:

* `dst_ports` (List of String) - list of destination tool ports
* `src_ports` (List of String) - inline network or an item from a-to-b and b-to-a lists
Optional:

* `direction` (String)
* `tag` (Attributes) (see [below for nested schema](#nestedatt--afm_map--flex_inline--oob_copy--tag))
<a id="nestedatt--afm_map--flex_inline--oob_copy--tag"></a>
### Nested Schema for `afm_map.flex_inline.oob_copy.tag`

Required:

* `type` (String)
<a id="nestedatt--afm_map--flex_inline--tag"></a>
### Nested Schema for `afm_map.flex_inline.tag`

Required:

* `type` (String)
Optional:

* `tag_protocol_id` (String) - When tool VLAN tag is added , this protocol Id will be added which egress out the traffic
* `vlan_id` (Number) - only applicable when type is 'vlan'
<a id="nestedatt--afm_map--flow_rules"></a>
### Nested Schema for `afm_map.flow_rules`

Optional:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--afm_map--flow_rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--afm_map--flow_rules--pass_rules))
<a id="nestedatt--afm_map--flow_rules--drop_rules"></a>
### Nested Schema for `afm_map.flow_rules.drop_rules`

Required:

* `gtp` (Attributes) - Map Flow Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--afm_map--flow_rules--drop_rules--gtp))
* `rule_id` (Number)
<a id="nestedatt--afm_map--flow_rules--drop_rules--gtp"></a>
### Nested Schema for `afm_map.flow_rules.drop_rules.gtp`

Optional:

* `imei` (String) - mutually exclusive with 'imsi' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - mutually exclusive with 'imei' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - mutually exclusive with 'imsi' and 'imei'. If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--afm_map--flow_rules--pass_rules"></a>
### Nested Schema for `afm_map.flow_rules.pass_rules`

Required:

* `gtp` (Attributes) - Map Flow Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--afm_map--flow_rules--pass_rules--gtp))
* `rule_id` (Number)
<a id="nestedatt--afm_map--flow_rules--pass_rules--gtp"></a>
### Nested Schema for `afm_map.flow_rules.pass_rules.gtp`

Optional:

* `imei` (String) - mutually exclusive with 'imsi' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - mutually exclusive with 'imei' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - mutually exclusive with 'imsi' and 'imei'. If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--afm_map--flow_sample5_g_overlap_rules"></a>
### Nested Schema for `afm_map.flow_sample5_g_overlap_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--afm_map--flow_sample5_g_overlap_rules--pass_rules))
<a id="nestedatt--afm_map--flow_sample5_g_overlap_rules--pass_rules"></a>
### Nested Schema for `afm_map.flow_sample5_g_overlap_rules.pass_rules`

Required:

* `flow5_g` (Attributes) - Map Flow Sample Overlap Rule 5g match Definition. Private class (see [below for nested schema](#nestedatt--afm_map--flow_sample5_g_overlap_rules--pass_rules--flow5_g))
* `percentage` (Number)
* `rule_id` (Number)
Optional:

* `comment` (String)
<a id="nestedatt--afm_map--flow_sample5_g_overlap_rules--pass_rules--flow5_g"></a>
### Nested Schema for `afm_map.flow_sample5_g_overlap_rules.pass_rules.flow5_g`

Optional:

* `dnn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `gpsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nsiid` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `pei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `supi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--afm_map--flow_sample5_g_rules"></a>
### Nested Schema for `afm_map.flow_sample5_g_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--afm_map--flow_sample5_g_rules--pass_rules))
<a id="nestedatt--afm_map--flow_sample5_g_rules--pass_rules"></a>
### Nested Schema for `afm_map.flow_sample5_g_rules.pass_rules`

Required:

* `flow5_g` (Attributes) - Map Flow Sample Rule 5g match Definition. Private class (see [below for nested schema](#nestedatt--afm_map--flow_sample5_g_rules--pass_rules--flow5_g))
* `percentage` (Number)
* `rule_id` (Number)
Optional:

* `comment` (String)
* `priority` (Number)
<a id="nestedatt--afm_map--flow_sample5_g_rules--pass_rules--flow5_g"></a>
### Nested Schema for `afm_map.flow_sample5_g_rules.pass_rules.flow5_g`

Optional:

* `dnn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `gpsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nsiid` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `pei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `supi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--afm_map--flow_sample_diameter_rules"></a>
### Nested Schema for `afm_map.flow_sample_diameter_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--afm_map--flow_sample_diameter_rules--pass_rules))
<a id="nestedatt--afm_map--flow_sample_diameter_rules--pass_rules"></a>
### Nested Schema for `afm_map.flow_sample_diameter_rules.pass_rules`

Required:

* `diameter` (Attributes) - Map Flow Sample Diameter Rule Definition (see [below for nested schema](#nestedatt--afm_map--flow_sample_diameter_rules--pass_rules--diameter))
* `interface` (String) - interface type
* `percentage` (Number)
* `rule_id` (Number)
<a id="nestedatt--afm_map--flow_sample_diameter_rules--pass_rules--diameter"></a>
### Nested Schema for `afm_map.flow_sample_diameter_rules.pass_rules.diameter`

Optional:

* `user_name` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--afm_map--flow_sample_overlap_rules"></a>
### Nested Schema for `afm_map.flow_sample_overlap_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--afm_map--flow_sample_overlap_rules--pass_rules))
<a id="nestedatt--afm_map--flow_sample_overlap_rules--pass_rules"></a>
### Nested Schema for `afm_map.flow_sample_overlap_rules.pass_rules`

Required:

* `gtp` (Attributes) - Map Flow Sample Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--afm_map--flow_sample_overlap_rules--pass_rules--gtp))
* `percentage` (Number)
* `rule_id` (Number)
Optional:

* `comment` (String)
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number) - maximum value is equal to the number of rules upon completion of the request
<a id="nestedatt--afm_map--flow_sample_overlap_rules--pass_rules--gtp"></a>
### Nested Schema for `afm_map.flow_sample_overlap_rules.pass_rules.gtp`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `eci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `qci` (Number) - QoS Class Indicator
* `snssai` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac_5_g` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--afm_map--flow_sample_rules"></a>
### Nested Schema for `afm_map.flow_sample_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--afm_map--flow_sample_rules--pass_rules))
<a id="nestedatt--afm_map--flow_sample_rules--pass_rules"></a>
### Nested Schema for `afm_map.flow_sample_rules.pass_rules`

Required:

* `gtp` (Attributes) - Map Flow Sample Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--afm_map--flow_sample_rules--pass_rules--gtp))
* `percentage` (Number)
* `rule_id` (Number)
Optional:

* `comment` (String)
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number) - maximum value is equal to the number of rules upon completion of the request
<a id="nestedatt--afm_map--flow_sample_rules--pass_rules--gtp"></a>
### Nested Schema for `afm_map.flow_sample_rules.pass_rules.gtp`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `eci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `qci` (Number) - QoS Class Indicator
* `snssai` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac_5_g` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--afm_map--flow_sample_sip_rules"></a>
### Nested Schema for `afm_map.flow_sample_sip_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--afm_map--flow_sample_sip_rules--pass_rules))
<a id="nestedatt--afm_map--flow_sample_sip_rules--pass_rules"></a>
### Nested Schema for `afm_map.flow_sample_sip_rules.pass_rules`

Required:

* `percentage` (Number)
* `rule_id` (Number)
* `sip` (Attributes) (see [below for nested schema](#nestedatt--afm_map--flow_sample_sip_rules--pass_rules--sip))
<a id="nestedatt--afm_map--flow_sample_sip_rules--pass_rules--sip"></a>
### Nested Schema for `afm_map.flow_sample_sip_rules.pass_rules.sip`

Optional:

* `callee_id` (String) - sip callee id
* `callee_id_range` (Attributes) (see [below for nested schema](#nestedatt--afm_map--flow_sample_sip_rules--pass_rules--sip--callee_id_range))
* `caller_id` (String) - sip caller id
* `caller_id_range` (Attributes) (see [below for nested schema](#nestedatt--afm_map--flow_sample_sip_rules--pass_rules--sip--caller_id_range))
* `id_range` (Attributes) (see [below for nested schema](#nestedatt--afm_map--flow_sample_sip_rules--pass_rules--sip--id_range))
<a id="nestedatt--afm_map--flow_sample_sip_rules--pass_rules--sip--callee_id_range"></a>
### Nested Schema for `afm_map.flow_sample_sip_rules.pass_rules.sip.callee_id_range`

Required:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--afm_map--flow_sample_sip_rules--pass_rules--sip--caller_id_range"></a>
### Nested Schema for `afm_map.flow_sample_sip_rules.pass_rules.sip.caller_id_range`

Required:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--afm_map--flow_sample_sip_rules--pass_rules--sip--id_range"></a>
### Nested Schema for `afm_map.flow_sample_sip_rules.pass_rules.sip.id_range`

Required:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--afm_map--flow_whitelist5_g_overlap_rules"></a>
### Nested Schema for `afm_map.flow_whitelist5_g_overlap_rules`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
<a id="nestedatt--afm_map--flow_whitelist5_g_rules"></a>
### Nested Schema for `afm_map.flow_whitelist5_g_rules`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--afm_map--flow_whitelist_overlap_rules"></a>
### Nested Schema for `afm_map.flow_whitelist_overlap_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--afm_map--flow_whitelist_overlap_rules--pass_rules))
<a id="nestedatt--afm_map--flow_whitelist_overlap_rules--pass_rules"></a>
### Nested Schema for `afm_map.flow_whitelist_overlap_rules.pass_rules`

Required:

* `rule_id` (Number)
Optional:

* `flow5_g` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--afm_map--flow_whitelist_overlap_rules--pass_rules--flow5_g))
* `gtp` (Attributes) - Map Flow Whitelist Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--afm_map--flow_whitelist_overlap_rules--pass_rules--gtp))
* `sip` (Attributes) - Map Flow Whitelist Rule Sip match definition (see [below for nested schema](#nestedatt--afm_map--flow_whitelist_overlap_rules--pass_rules--sip))
<a id="nestedatt--afm_map--flow_whitelist_overlap_rules--pass_rules--flow5_g"></a>
### Nested Schema for `afm_map.flow_whitelist_overlap_rules.pass_rules.flow5_g`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--afm_map--flow_whitelist_overlap_rules--pass_rules--gtp"></a>
### Nested Schema for `afm_map.flow_whitelist_overlap_rules.pass_rules.gtp`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--afm_map--flow_whitelist_overlap_rules--pass_rules--sip"></a>
### Nested Schema for `afm_map.flow_whitelist_overlap_rules.pass_rules.sip`

Optional:

* `type` (String) - all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address
<a id="nestedatt--afm_map--flow_whitelist_rules"></a>
### Nested Schema for `afm_map.flow_whitelist_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--afm_map--flow_whitelist_rules--pass_rules))
<a id="nestedatt--afm_map--flow_whitelist_rules--pass_rules"></a>
### Nested Schema for `afm_map.flow_whitelist_rules.pass_rules`

Required:

* `rule_id` (Number)
Optional:

* `flow5_g` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--afm_map--flow_whitelist_rules--pass_rules--flow5_g))
* `gtp` (Attributes) - Map Flow Whitelist Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--afm_map--flow_whitelist_rules--pass_rules--gtp))
* `sip` (Attributes) - Map Flow Whitelist Rule Sip match definition (see [below for nested schema](#nestedatt--afm_map--flow_whitelist_rules--pass_rules--sip))
<a id="nestedatt--afm_map--flow_whitelist_rules--pass_rules--flow5_g"></a>
### Nested Schema for `afm_map.flow_whitelist_rules.pass_rules.flow5_g`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--afm_map--flow_whitelist_rules--pass_rules--gtp"></a>
### Nested Schema for `afm_map.flow_whitelist_rules.pass_rules.gtp`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--afm_map--flow_whitelist_rules--pass_rules--sip"></a>
### Nested Schema for `afm_map.flow_whitelist_rules.pass_rules.sip`

Optional:

* `type` (String) - all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address
<a id="nestedatt--afm_map--fstype"></a>
### Nested Schema for `afm_map.fstype`

Optional:

* `offset` (Number) - Offset for Mobility Rotational Flowsampling
* `timer` (Number) - Timer for Mobility Rotational Flowsampling in minutes
* `type` (String)
<a id="nestedatt--afm_map--gs_rules"></a>
### Nested Schema for `afm_map.gs_rules`

Optional:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--afm_map--gs_rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--afm_map--gs_rules--pass_rules))
<a id="nestedatt--afm_map--gs_rules--drop_rules"></a>
### Nested Schema for `afm_map.gs_rules.drop_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number)
Optional:

* `comment` (String)
<a id="nestedatt--afm_map--gs_rules--pass_rules"></a>
### Nested Schema for `afm_map.gs_rules.pass_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number)
Optional:

* `comment` (String)
<a id="nestedatt--afm_map--ip_rewrite"></a>
### Nested Schema for `afm_map.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--afm_map--rewrite"></a>
### Nested Schema for `afm_map.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--afm_map--roles"></a>
### Nested Schema for `afm_map.roles`

Optional:

* `editors` (List of String)
* `listeners` (List of String)
* `owners` (List of String)
* `viewers` (List of String)
<a id="nestedatt--afm_map--rules"></a>
### Nested Schema for `afm_map.rules`

Optional:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--afm_map--rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--afm_map--rules--pass_rules))
<a id="nestedatt--afm_map--rules--drop_rules"></a>
### Nested Schema for `afm_map.rules.drop_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)
Optional:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--afm_map--rules--drop_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--afm_map--rules--drop_rules--rewrite))
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--afm_map--rules--drop_rules--vlan_tag))
<a id="nestedatt--afm_map--rules--drop_rules--ip_rewrite"></a>
### Nested Schema for `afm_map.rules.drop_rules.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--afm_map--rules--drop_rules--rewrite"></a>
### Nested Schema for `afm_map.rules.drop_rules.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--afm_map--rules--drop_rules--vlan_tag"></a>
### Nested Schema for `afm_map.rules.drop_rules.vlan_tag`

Required:

* `vlan_action` (String)
Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)
<a id="nestedatt--afm_map--rules--pass_rules"></a>
### Nested Schema for `afm_map.rules.pass_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)
Optional:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--afm_map--rules--pass_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--afm_map--rules--pass_rules--rewrite))
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--afm_map--rules--pass_rules--vlan_tag))
<a id="nestedatt--afm_map--rules--pass_rules--ip_rewrite"></a>
### Nested Schema for `afm_map.rules.pass_rules.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--afm_map--rules--pass_rules--rewrite"></a>
### Nested Schema for `afm_map.rules.pass_rules.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--afm_map--rules--pass_rules--vlan_tag"></a>
### Nested Schema for `afm_map.rules.pass_rules.vlan_tag`

Required:

* `vlan_action` (String)
Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)
<a id="nestedatt--afm_map--vlan_tag"></a>
### Nested Schema for `afm_map.vlan_tag`

Required:

* `vlan_action` (String)
Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)
<a id="nestedatt--afm_map--health_state_reasons"></a>
### Nested Schema for `afm_map.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_fabric_map.example {alias}
```
