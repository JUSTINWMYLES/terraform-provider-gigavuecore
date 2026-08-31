---
page_title: "gigavuecore_map Resource - gigavuecore"
subcategory: ""
description: |-
  Find map by alias
---

# gigavuecore_map Resource

Find map by alias

## Example Usage

```terraform
resource "gigavuecore_map" "example" {
  alias = "example"
  ap_rules = {
    drop_rules = [{
      application_profile = "example"
      rule_id             = 1
    }]
    pass_rules = [{
      application_profile = "example"
      rule_id             = 1
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
      type       = "bypass"
    }
    b_to_a = {
      ib_pathway = "example"
      tools      = [ "example" ]
      type       = "bypass"
    }
    oob_copy = [{
      direction = "aToB"
      dst_ports = [ "example" ]
      src_ports = [ "example" ]
      tag = {
        type = "none"
      }
    }]
    svt_mode = true
    svt_tag  = 0
    tag = {
      tag_protocol_id = "0x8100"
      type            = "auto"
      vlan_id         = 0
    }
  }
  flex_inline_failover = "bypass"
  flex_inline_vlan_id  = 1
  flow_rules = {
    drop_rules = [{
      gtp = {
        imei      = "example"
        imsi      = "example"
        interface = "Gn"
        msisdn    = "example"
        version   = "any"
      }
      rule_id = 1
    }]
    pass_rules = [{
      gtp = {
        imei      = "example"
        imsi      = "example"
        interface = "Gn"
        msisdn    = "example"
        version   = "any"
      }
      rule_id = 1
    }]
  }
  flow_sample5_g_overlap_rules = {
    pass_rules = [{
      comment = "example"
      flow5_g = {
        dnn      = "example"
        gpsi     = "example"
        nas_5_qi = "0"
        nci      = "example"
        nsiid    = "0"
        pei      = "example"
        plmn_id  = "example"
        supi     = "example"
        tac      = "example"
      }
      percentage = 0
      rule_id    = 1
    }]
  }
  flow_sample5_g_rules = {
    pass_rules = [{
      comment = "example"
      flow5_g = {
        dnn     = "example"
        gpsi    = "example"
        nci     = "example"
        nsiid   = "0"
        pei     = "example"
        plmn_id = "example"
        supi    = "example"
        tac     = "example"
      }
      percentage = 0
      priority   = 1
      rule_id    = 1
    }]
  }
  flow_sample_diameter_rules = {
    pass_rules = [{
      diameter = {
        user_name = "example"
      }
      interface  = "s6a"
      percentage = 0
      rule_id    = 1
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
        interface = "Gn"
        msisdn    = "example"
        nas_5_qi  = "0"
        nci       = "example"
        plmn_id   = "example"
        qci       = 0
        snssai    = "0"
        tac       = "abc1"
        tac_5_g   = "example"
        version   = "any"
      }
      percentage      = 0
      periodic_recalc = true
      priority        = 1
      rule_id         = 1
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
        interface = "Gn"
        msisdn    = "example"
        nas_5_qi  = "0"
        nci       = "example"
        plmn_id   = "example"
        qci       = 0
        snssai    = "0"
        tac       = "abc1"
        tac_5_g   = "example"
        version   = "any"
      }
      percentage      = 0
      periodic_recalc = true
      priority        = 1
      rule_id         = 1
    }]
  }
  flow_sample_sip_rules = {
    pass_rules = [{
      percentage = 0
      rule_id    = 1
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
        interface           = "Gn"
        type                = "example"
        version             = "v1"
        whitelist_databases = [ "example" ]
      }
      rule_id = 1
      sip = {
        type = "all"
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
        interface           = "Gn"
        type                = "example"
        version             = "v1"
        whitelist_databases = [ "example" ]
      }
      rule_id = 1
      sip = {
        type = "all"
      }
    }]
  }
  fstype = {
    offset = 1
    timer  = 15
    type   = "_default"
  }
  gs_rules = {
    drop_rules = [{
      comment = "example"
      matches = [ "example" ]
      rule_id = 1
    }]
    pass_rules = [{
      comment = "example"
      matches = [ "example" ]
      rule_id = 1
    }]
  }
  gsop         = "example"
  health_state = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  inline_traffic_path = "normal"
  inline_traffic_type = "symmetric"
  ip_rewrite = {
    dst_ip = "example"
    src_ip = "example"
  }
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
  rule_matching = "normal"
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
      rule_id = 1
      vlan_tag = {
        tag_protocol_id = "0x8100"
        vlan_action     = "add"
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
      rule_id = 1
      vlan_tag = {
        tag_protocol_id = "0x8100"
        vlan_action     = "add"
        vlan_id         = 0
      }
    }]
  }
  src_ports    = [ "example" ]
  sub_type     = "byRule"
  traffic_type = "control"
  type         = "regular"
  updated_time = 1.0
  vlan_tag = {
    tag_protocol_id = "0x8100"
    vlan_action     = "add"
    vlan_id         = 0
  }
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - unique map alias
* `ap_rules` (Attributes, optional) - pass and drop application profile rules (see [below for nested schema](#nestedatt--ap_rules))
* `cluster_id` (String, required) - id of the defining cluster
* `comment` (String, optional)
* `dst_ports` (List of String, required) - List of the 'to' ports. Invalid if 'nullDstPort' is true'. Only port number is supported
* `egress_gigastream` (List of String, optional) - Applicable to fabric map only. It is to specify gigastream(s) to use as egress when going across cluster. The format of array member is clusterID:gigastreamAlias
* `enable` (Boolean, optional) - enable/disable map, applicable only to first level maps
* `encap_tunnel` (String, optional) - tunnel alias
* `flex_inline` (Attributes, optional) - When 'oobCopy' is configured at least one of 'aToB' or 'bToA' must be configured (see [below for nested schema](#nestedatt--flex_inline))
* `flex_inline_failover` (String, optional) - only valid for flexInline maps
* `flex_inline_vlan_id` (Number, optional) - VLAN ID carried in the VLAN tag of packets coming from the inline-network port(s). valid and applicable only with 'flexInline' map type, 'srcPorts' should be either inline-network or ib-pathway or vport
* `flow_rules` (Attributes, optional) - Map Flow Rules Container. Private class (see [below for nested schema](#nestedatt--flow_rules))
* `flow_sample5_g_overlap_rules` (Attributes, optional) - Map Flow Sample 5g Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--flow_sample5_g_overlap_rules))
* `flow_sample5_g_rules` (Attributes, optional) - Map Flow Sample 5g Rules Container. Private class (see [below for nested schema](#nestedatt--flow_sample5_g_rules))
* `flow_sample_diameter_rules` (Attributes, optional) - Map Flow Sample Diameter Rules Container. Private class (see [below for nested schema](#nestedatt--flow_sample_diameter_rules))
* `flow_sample_overlap_rules` (Attributes, optional) - Map Flow Sample Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--flow_sample_overlap_rules))
* `flow_sample_rules` (Attributes, optional) - Map Flow Sample Rules Container. Private class (see [below for nested schema](#nestedatt--flow_sample_rules))
* `flow_sample_sip_rules` (Attributes, optional) - Map Flow Sample Sip Rules Container. Private class (see [below for nested schema](#nestedatt--flow_sample_sip_rules))
* `flow_whitelist5_g_overlap_rules` (Attributes, optional) - Map Flow Whitelist 5g Overlap Rule match Definition. Private class (see [below for nested schema](#nestedatt--flow_whitelist5_g_overlap_rules))
* `flow_whitelist5_g_rules` (Attributes, optional) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--flow_whitelist5_g_rules))
* `flow_whitelist_overlap_rules` (Attributes, optional) - Map Flow Whitelist Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--flow_whitelist_overlap_rules))
* `flow_whitelist_rules` (Attributes, optional) - Map Flow Whitelist Rules Container. Private class (see [below for nested schema](#nestedatt--flow_whitelist_rules))
* `fstype` (Attributes, optional) - Type of Mobility Flowsampling & properties (see [below for nested schema](#nestedatt--fstype))
* `gs_rules` (Attributes, optional) - Map GigaSMART Rules Container. Private class (see [below for nested schema](#nestedatt--gs_rules))
* `gsop` (String, optional) - Alias of referenced GSOP. Applicable to 'regular' and 'secondLevel' map types
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `inline_traffic_path` (String, optional) - Only applicable for 'inline' map types, in which case defaults to 'normal'. When set to 'bypass', the 'dstPorts' must be empty
* `inline_traffic_type` (String, optional) - Only applicable for 'inline/passAll' map types, in which case defaults to 'symmetric'. For 'asymmetric' maps, 'srcPort' must be of type 'inline-net'
* `ip_rewrite` (Attributes, optional) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--ip_rewrite))
* `null_dst_port` (Boolean, optional) - enabled when the dstPort is null
* `order` (Number, optional) - relative order within per-source port map chain
* `rewrite` (Attributes, optional) - Rewrite options on the packets (see [below for nested schema](#nestedatt--rewrite))
* `roles` (Attributes, optional) - Map Roles Container. Private class (see [below for nested schema](#nestedatt--roles))
* `rule_matching` (String, optional) - If set to 'blacklist', packet are passed when no 'drop' rules are matched. This field is only valid for 'regular/byRule' map types. (maps into 'no-rule-match' CLI command)
* `rules` (Attributes, optional) - Map Rules Container. Private class (see [below for nested schema](#nestedatt--rules))
* `src_ports` (List of String, required) - list of the 'from' ports. Only port number is supported
* `sub_type` (String, optional) - 'byRule' is applicable to all map types; 'collector' is applicable to 'regular', 'inline', 'flexInline' and 'secondLevel' maps; 'passAll' is applicable to 'regular' and 'inline' maps; 'flowFilter', 'flowSample', 'flowWhitelist', 'flowSampleSip', 'flowWhitelistSip', 'flowSampleDiameter', 'flowWhitelistDiameter', 'flowSampleOverlap', 'flowWhitelistOverlap', 'flowSample5g','flowWhitelist5g','flowSample5gOverlap' and 'flowWhitelist5gOverlap' are applicable to 'secondLevel' maps
* `traffic_type` (String, optional) - Only applicable for 'firstLevel/byRule' map types, in which case defaults to 'user'
* `type` (String, required) - 'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'flexInline' is from inline network to tools; 'transitLevel' is from vport to vport
* `updated_time` (Number, optional) - Last Updated time of the fabric map
* `vlan_tag` (Attributes, optional) - This field is only valid for 'regular/byRule', 'collector' map types (see [below for nested schema](#nestedatt--vlan_tag))

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `mod_time` (Number, computed) - Last modification time of the map in milliseconds since the epoch
* `rx_cluster_ports` (List of String, computed)
* `tx_cluster_ports` (List of String, computed)

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--ap_rules"></a>
### Nested Schema for `ap_rules`

Optional:

* `drop_rules` (Attributes List) (see [below for nested schema](#nestedatt--ap_rules--drop_rules))
* `pass_rules` (Attributes List) (see [below for nested schema](#nestedatt--ap_rules--pass_rules))

<a id="nestedatt--ap_rules--drop_rules"></a>
### Nested Schema for `ap_rules.drop_rules`

Required:

* `application_profile` (String) - application profile alias
* `rule_id` (Number) - application profile rule Id, should not have same id as gsRules

<a id="nestedatt--ap_rules--pass_rules"></a>
### Nested Schema for `ap_rules.pass_rules`

Required:

* `application_profile` (String) - application profile alias
* `rule_id` (Number) - application profile rule Id, should not have same id as gsRules

<a id="nestedatt--flex_inline"></a>
### Nested Schema for `flex_inline`

Optional:

* `a_to_b` (Attributes) (see [below for nested schema](#nestedatt--flex_inline--a_to_b))
* `b_to_a` (Attributes) (see [below for nested schema](#nestedatt--flex_inline--b_to_a))
* `oob_copy` (Attributes List) (see [below for nested schema](#nestedatt--flex_inline--oob_copy))
* `svt_mode` (Boolean)
* `svt_tag` (Number) - only applicable when svtMode is enabled
* `tag` (Attributes) (see [below for nested schema](#nestedatt--flex_inline--tag))

<a id="nestedatt--flex_inline--a_to_b"></a>
### Nested Schema for `flex_inline.a_to_b`

Required:

* `type` (String) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided

Optional:

* `ib_pathway` (String) - ibPathway alias. Only applicable when type is 'ibPathway'
* `tools` (List of String) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'

<a id="nestedatt--flex_inline--b_to_a"></a>
### Nested Schema for `flex_inline.b_to_a`

Required:

* `type` (String) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided

Optional:

* `ib_pathway` (String) - ibPathway alias. Only applicable when type is 'ibPathway'
* `tools` (List of String) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'

<a id="nestedatt--flex_inline--oob_copy"></a>
### Nested Schema for `flex_inline.oob_copy`

Required:

* `dst_ports` (List of String) - list of destination tool ports
* `src_ports` (List of String) - inline network or an item from a-to-b and b-to-a lists

Optional:

* `direction` (String)
* `tag` (Attributes) (see [below for nested schema](#nestedatt--flex_inline--oob_copy--tag))

<a id="nestedatt--flex_inline--oob_copy--tag"></a>
### Nested Schema for `flex_inline.oob_copy.tag`

Required:

* `type` (String)

<a id="nestedatt--flex_inline--tag"></a>
### Nested Schema for `flex_inline.tag`

Required:

* `type` (String)

Optional:

* `tag_protocol_id` (String) - When tool VLAN tag is added , this protocol Id will be added which egress out the traffic
* `vlan_id` (Number) - only applicable when type is 'vlan'

<a id="nestedatt--flow_rules"></a>
### Nested Schema for `flow_rules`

Optional:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--flow_rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--flow_rules--pass_rules))

<a id="nestedatt--flow_rules--drop_rules"></a>
### Nested Schema for `flow_rules.drop_rules`

Required:

* `gtp` (Attributes) - Map Flow Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--flow_rules--drop_rules--gtp))
* `rule_id` (Number)

<a id="nestedatt--flow_rules--drop_rules--gtp"></a>
### Nested Schema for `flow_rules.drop_rules.gtp`

Optional:

* `imei` (String) - mutually exclusive with 'imsi' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - mutually exclusive with 'imei' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - mutually exclusive with 'imsi' and 'imei'. If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface

<a id="nestedatt--flow_rules--pass_rules"></a>
### Nested Schema for `flow_rules.pass_rules`

Required:

* `gtp` (Attributes) - Map Flow Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--flow_rules--pass_rules--gtp))
* `rule_id` (Number)

<a id="nestedatt--flow_rules--pass_rules--gtp"></a>
### Nested Schema for `flow_rules.pass_rules.gtp`

Optional:

* `imei` (String) - mutually exclusive with 'imsi' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - mutually exclusive with 'imei' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - mutually exclusive with 'imsi' and 'imei'. If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface

<a id="nestedatt--flow_sample5_g_overlap_rules"></a>
### Nested Schema for `flow_sample5_g_overlap_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--flow_sample5_g_overlap_rules--pass_rules))

<a id="nestedatt--flow_sample5_g_overlap_rules--pass_rules"></a>
### Nested Schema for `flow_sample5_g_overlap_rules.pass_rules`

Required:

* `flow5_g` (Attributes) - Map Flow Sample Overlap Rule 5g match Definition. Private class (see [below for nested schema](#nestedatt--flow_sample5_g_overlap_rules--pass_rules--flow5_g))
* `percentage` (Number)
* `rule_id` (Number)

Optional:

* `comment` (String)

<a id="nestedatt--flow_sample5_g_overlap_rules--pass_rules--flow5_g"></a>
### Nested Schema for `flow_sample5_g_overlap_rules.pass_rules.flow5_g`

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

<a id="nestedatt--flow_sample5_g_rules"></a>
### Nested Schema for `flow_sample5_g_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--flow_sample5_g_rules--pass_rules))

<a id="nestedatt--flow_sample5_g_rules--pass_rules"></a>
### Nested Schema for `flow_sample5_g_rules.pass_rules`

Required:

* `flow5_g` (Attributes) - Map Flow Sample Rule 5g match Definition. Private class (see [below for nested schema](#nestedatt--flow_sample5_g_rules--pass_rules--flow5_g))
* `percentage` (Number)
* `rule_id` (Number)

Optional:

* `comment` (String)
* `priority` (Number)

<a id="nestedatt--flow_sample5_g_rules--pass_rules--flow5_g"></a>
### Nested Schema for `flow_sample5_g_rules.pass_rules.flow5_g`

Optional:

* `dnn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `gpsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nsiid` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `pei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `supi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix

<a id="nestedatt--flow_sample_diameter_rules"></a>
### Nested Schema for `flow_sample_diameter_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--flow_sample_diameter_rules--pass_rules))

<a id="nestedatt--flow_sample_diameter_rules--pass_rules"></a>
### Nested Schema for `flow_sample_diameter_rules.pass_rules`

Required:

* `diameter` (Attributes) - Map Flow Sample Diameter Rule Definition (see [below for nested schema](#nestedatt--flow_sample_diameter_rules--pass_rules--diameter))
* `interface` (String) - interface type
* `percentage` (Number)
* `rule_id` (Number)

<a id="nestedatt--flow_sample_diameter_rules--pass_rules--diameter"></a>
### Nested Schema for `flow_sample_diameter_rules.pass_rules.diameter`

Optional:

* `user_name` (String) - If '\*' is added at the end of the value, it is treated as prefix

<a id="nestedatt--flow_sample_overlap_rules"></a>
### Nested Schema for `flow_sample_overlap_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--flow_sample_overlap_rules--pass_rules))

<a id="nestedatt--flow_sample_overlap_rules--pass_rules"></a>
### Nested Schema for `flow_sample_overlap_rules.pass_rules`

Required:

* `gtp` (Attributes) - Map Flow Sample Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--flow_sample_overlap_rules--pass_rules--gtp))
* `percentage` (Number)
* `rule_id` (Number)

Optional:

* `comment` (String)
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number) - maximum value is equal to the number of rules upon completion of the request

<a id="nestedatt--flow_sample_overlap_rules--pass_rules--gtp"></a>
### Nested Schema for `flow_sample_overlap_rules.pass_rules.gtp`

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

<a id="nestedatt--flow_sample_rules"></a>
### Nested Schema for `flow_sample_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--flow_sample_rules--pass_rules))

<a id="nestedatt--flow_sample_rules--pass_rules"></a>
### Nested Schema for `flow_sample_rules.pass_rules`

Required:

* `gtp` (Attributes) - Map Flow Sample Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--flow_sample_rules--pass_rules--gtp))
* `percentage` (Number)
* `rule_id` (Number)

Optional:

* `comment` (String)
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number) - maximum value is equal to the number of rules upon completion of the request

<a id="nestedatt--flow_sample_rules--pass_rules--gtp"></a>
### Nested Schema for `flow_sample_rules.pass_rules.gtp`

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

<a id="nestedatt--flow_sample_sip_rules"></a>
### Nested Schema for `flow_sample_sip_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--flow_sample_sip_rules--pass_rules))

<a id="nestedatt--flow_sample_sip_rules--pass_rules"></a>
### Nested Schema for `flow_sample_sip_rules.pass_rules`

Required:

* `percentage` (Number)
* `rule_id` (Number)
* `sip` (Attributes) (see [below for nested schema](#nestedatt--flow_sample_sip_rules--pass_rules--sip))

<a id="nestedatt--flow_sample_sip_rules--pass_rules--sip"></a>
### Nested Schema for `flow_sample_sip_rules.pass_rules.sip`

Optional:

* `callee_id` (String) - sip callee id
* `callee_id_range` (Attributes) (see [below for nested schema](#nestedatt--flow_sample_sip_rules--pass_rules--sip--callee_id_range))
* `caller_id` (String) - sip caller id
* `caller_id_range` (Attributes) (see [below for nested schema](#nestedatt--flow_sample_sip_rules--pass_rules--sip--caller_id_range))
* `id_range` (Attributes) (see [below for nested schema](#nestedatt--flow_sample_sip_rules--pass_rules--sip--id_range))

<a id="nestedatt--flow_sample_sip_rules--pass_rules--sip--callee_id_range"></a>
### Nested Schema for `flow_sample_sip_rules.pass_rules.sip.callee_id_range`

Required:

* `max_value` (String)
* `value` (String)

<a id="nestedatt--flow_sample_sip_rules--pass_rules--sip--caller_id_range"></a>
### Nested Schema for `flow_sample_sip_rules.pass_rules.sip.caller_id_range`

Required:

* `max_value` (String)
* `value` (String)

<a id="nestedatt--flow_sample_sip_rules--pass_rules--sip--id_range"></a>
### Nested Schema for `flow_sample_sip_rules.pass_rules.sip.id_range`

Required:

* `max_value` (String)
* `value` (String)

<a id="nestedatt--flow_whitelist5_g_overlap_rules"></a>
### Nested Schema for `flow_whitelist5_g_overlap_rules`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type

<a id="nestedatt--flow_whitelist5_g_rules"></a>
### Nested Schema for `flow_whitelist5_g_rules`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map

<a id="nestedatt--flow_whitelist_overlap_rules"></a>
### Nested Schema for `flow_whitelist_overlap_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--flow_whitelist_overlap_rules--pass_rules))

<a id="nestedatt--flow_whitelist_overlap_rules--pass_rules"></a>
### Nested Schema for `flow_whitelist_overlap_rules.pass_rules`

Required:

* `rule_id` (Number)

Optional:

* `flow5_g` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--flow_whitelist_overlap_rules--pass_rules--flow5_g))
* `gtp` (Attributes) - Map Flow Whitelist Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--flow_whitelist_overlap_rules--pass_rules--gtp))
* `sip` (Attributes) - Map Flow Whitelist Rule Sip match definition (see [below for nested schema](#nestedatt--flow_whitelist_overlap_rules--pass_rules--sip))

<a id="nestedatt--flow_whitelist_overlap_rules--pass_rules--flow5_g"></a>
### Nested Schema for `flow_whitelist_overlap_rules.pass_rules.flow5_g`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map

<a id="nestedatt--flow_whitelist_overlap_rules--pass_rules--gtp"></a>
### Nested Schema for `flow_whitelist_overlap_rules.pass_rules.gtp`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map

<a id="nestedatt--flow_whitelist_overlap_rules--pass_rules--sip"></a>
### Nested Schema for `flow_whitelist_overlap_rules.pass_rules.sip`

Optional:

* `type` (String) - all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address

<a id="nestedatt--flow_whitelist_rules"></a>
### Nested Schema for `flow_whitelist_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--flow_whitelist_rules--pass_rules))

<a id="nestedatt--flow_whitelist_rules--pass_rules"></a>
### Nested Schema for `flow_whitelist_rules.pass_rules`

Required:

* `rule_id` (Number)

Optional:

* `flow5_g` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--flow_whitelist_rules--pass_rules--flow5_g))
* `gtp` (Attributes) - Map Flow Whitelist Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--flow_whitelist_rules--pass_rules--gtp))
* `sip` (Attributes) - Map Flow Whitelist Rule Sip match definition (see [below for nested schema](#nestedatt--flow_whitelist_rules--pass_rules--sip))

<a id="nestedatt--flow_whitelist_rules--pass_rules--flow5_g"></a>
### Nested Schema for `flow_whitelist_rules.pass_rules.flow5_g`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map

<a id="nestedatt--flow_whitelist_rules--pass_rules--gtp"></a>
### Nested Schema for `flow_whitelist_rules.pass_rules.gtp`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map

<a id="nestedatt--flow_whitelist_rules--pass_rules--sip"></a>
### Nested Schema for `flow_whitelist_rules.pass_rules.sip`

Optional:

* `type` (String) - all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address

<a id="nestedatt--fstype"></a>
### Nested Schema for `fstype`

Optional:

* `offset` (Number) - Offset for Mobility Rotational Flowsampling
* `timer` (Number) - Timer for Mobility Rotational Flowsampling in minutes
* `type` (String)

<a id="nestedatt--gs_rules"></a>
### Nested Schema for `gs_rules`

Optional:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--gs_rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--gs_rules--pass_rules))

<a id="nestedatt--gs_rules--drop_rules"></a>
### Nested Schema for `gs_rules.drop_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number)

Optional:

* `comment` (String)

<a id="nestedatt--gs_rules--pass_rules"></a>
### Nested Schema for `gs_rules.pass_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number)

Optional:

* `comment` (String)

<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

<a id="nestedatt--ip_rewrite"></a>
### Nested Schema for `ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)

<a id="nestedatt--rewrite"></a>
### Nested Schema for `rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)

<a id="nestedatt--roles"></a>
### Nested Schema for `roles`

Optional:

* `editors` (List of String)
* `listeners` (List of String)
* `owners` (List of String)
* `viewers` (List of String)

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Optional:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--rules--pass_rules))

<a id="nestedatt--rules--drop_rules"></a>
### Nested Schema for `rules.drop_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)

Optional:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--rules--drop_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--rules--drop_rules--rewrite))
* `vlan_tag` (Attributes) - This field is only  applicable for pass rules (see [below for nested schema](#nestedatt--rules--drop_rules--vlan_tag))

<a id="nestedatt--rules--drop_rules--ip_rewrite"></a>
### Nested Schema for `rules.drop_rules.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)

<a id="nestedatt--rules--drop_rules--rewrite"></a>
### Nested Schema for `rules.drop_rules.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)

<a id="nestedatt--rules--drop_rules--vlan_tag"></a>
### Nested Schema for `rules.drop_rules.vlan_tag`

Required:

* `vlan_action` (String)

Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)

<a id="nestedatt--rules--pass_rules"></a>
### Nested Schema for `rules.pass_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)

Optional:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--rules--pass_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--rules--pass_rules--rewrite))
* `vlan_tag` (Attributes) - This field is only  applicable for pass rules (see [below for nested schema](#nestedatt--rules--pass_rules--vlan_tag))

<a id="nestedatt--rules--pass_rules--ip_rewrite"></a>
### Nested Schema for `rules.pass_rules.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)

<a id="nestedatt--rules--pass_rules--rewrite"></a>
### Nested Schema for `rules.pass_rules.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)

<a id="nestedatt--rules--pass_rules--vlan_tag"></a>
### Nested Schema for `rules.pass_rules.vlan_tag`

Required:

* `vlan_action` (String)

Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)

<a id="nestedatt--vlan_tag"></a>
### Nested Schema for `vlan_tag`

Required:

* `vlan_action` (String)

Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_map.example {alias}/{cluster_id}
```
