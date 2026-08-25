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
  alias                           = null
  ap_rules                        = {}
  cluster_id                      = null
  comment                         = null
  dst_ports                       = []
  egress_gigastream               = []
  enable                          = null
  encap_tunnel                    = null
  flex_inline                     = {}
  flex_inline_failover            = null
  flex_inline_vlan_id             = null
  flow_rules                      = {}
  flow_sample5_g_overlap_rules    = {}
  flow_sample5_g_rules            = {}
  flow_sample_diameter_rules      = {}
  flow_sample_overlap_rules       = {}
  flow_sample_rules               = {}
  flow_sample_sip_rules           = {}
  flow_whitelist5_g_overlap_rules = {}
  flow_whitelist5_g_rules         = {}
  flow_whitelist_overlap_rules    = {}
  flow_whitelist_rules            = {}
  fstype                          = {}
  gs_rules                        = {}
  gsop                            = null
  inline_traffic_path             = null
  inline_traffic_type             = null
  ip_rewrite                      = {}
  mod_time                        = null
  null_dst_port                   = null
  order                           = null
  rewrite                         = {}
  roles                           = {}
  rule_matching                   = null
  rules                           = {}
  rx_cluster_ports                = []
  src_ports                       = []
  sub_type                        = null
  traffic_type                    = null
  tx_cluster_ports                = []
  type                            = null
  vlan_tag                        = {}
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - unique map alias
* `ap_rules` (Attributes, optional) - pass and drop application profile rules (see [below for nested schema](#nestedatt--ap_rules))
* `cluster_id` (String, optional) - id of the defining cluster
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
* `inline_traffic_path` (String, optional) - Only applicable for 'inline' map types, in which case defaults to 'normal'. When set to 'bypass', the 'dstPorts' must be empty
* `inline_traffic_type` (String, optional) - Only applicable for 'inline/passAll' map types, in which case defaults to 'symmetric'. For 'asymmetric' maps, 'srcPort' must be of type 'inline-net'
* `ip_rewrite` (Attributes, optional) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--ip_rewrite))
* `mod_time` (Number, optional) - Last modification time of the map in milliseconds since the epoch
* `null_dst_port` (Boolean, optional) - enabled when the dstPort is null
* `order` (Number, optional) - relative order within per-source port map chain
* `rewrite` (Attributes, optional) - Rewrite options on the packets (see [below for nested schema](#nestedatt--rewrite))
* `roles` (Attributes, optional) - Map Roles Container. Private class (see [below for nested schema](#nestedatt--roles))
* `rule_matching` (String, optional) - If set to 'blacklist', packet are passed when no 'drop' rules are matched. This field is only valid for 'regular/byRule' map types. (maps into 'no-rule-match' CLI command)
* `rules` (Attributes, optional) - Map Rules Container. Private class (see [below for nested schema](#nestedatt--rules))
* `rx_cluster_ports` (List of String, optional)
* `src_ports` (List of String, required) - list of the 'from' ports. Only port number is supported
* `sub_type` (String, optional) - 'byRule' is applicable to all map types; 'collector' is applicable to 'regular', 'inline', 'flexInline' and 'secondLevel' maps; 'passAll' is applicable to 'regular' and 'inline' maps; 'flowFilter', 'flowSample', 'flowWhitelist', 'flowSampleSip', 'flowWhitelistSip', 'flowSampleDiameter', 'flowWhitelistDiameter', 'flowSampleOverlap', 'flowWhitelistOverlap', 'flowSample5g','flowWhitelist5g','flowSample5gOverlap' and 'flowWhitelist5gOverlap' are applicable to 'secondLevel' maps
* `traffic_type` (String, optional) - Only applicable for 'firstLevel/byRule' map types, in which case defaults to 'user'
* `tx_cluster_ports` (List of String, optional)
* `type` (String, required) - 'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'flexInline' is from inline network to tools; 'transitLevel' is from vport to vport
* `vlan_tag` (Attributes, optional) (see [below for nested schema](#nestedatt--vlan_tag))

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `ap_rules` (Attributes, computed) - pass and drop application profile rules (see [below for nested schema](#nestedatt--ap_rules))
* `cluster_id` (String, computed) - id of the defining cluster
* `comment` (String, computed)
* `egress_gigastream` (List of String, computed) - Applicable to fabric map only. It is to specify gigastream(s) to use as egress when going across cluster. The format of array member is clusterID:gigastreamAlias
* `enable` (Boolean, computed) - enable/disable map, applicable only to first level maps
* `encap_tunnel` (String, computed) - tunnel alias
* `flex_inline` (Attributes, computed) - When 'oobCopy' is configured at least one of 'aToB' or 'bToA' must be configured (see [below for nested schema](#nestedatt--flex_inline))
* `flex_inline_failover` (String, computed) - only valid for flexInline maps
* `flex_inline_vlan_id` (Number, computed) - VLAN ID carried in the VLAN tag of packets coming from the inline-network port(s). valid and applicable only with 'flexInline' map type, 'srcPorts' should be either inline-network or ib-pathway or vport
* `flow_rules` (Attributes, computed) - Map Flow Rules Container. Private class (see [below for nested schema](#nestedatt--flow_rules))
* `flow_sample5_g_overlap_rules` (Attributes, computed) - Map Flow Sample 5g Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--flow_sample5_g_overlap_rules))
* `flow_sample5_g_rules` (Attributes, computed) - Map Flow Sample 5g Rules Container. Private class (see [below for nested schema](#nestedatt--flow_sample5_g_rules))
* `flow_sample_diameter_rules` (Attributes, computed) - Map Flow Sample Diameter Rules Container. Private class (see [below for nested schema](#nestedatt--flow_sample_diameter_rules))
* `flow_sample_overlap_rules` (Attributes, computed) - Map Flow Sample Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--flow_sample_overlap_rules))
* `flow_sample_rules` (Attributes, computed) - Map Flow Sample Rules Container. Private class (see [below for nested schema](#nestedatt--flow_sample_rules))
* `flow_sample_sip_rules` (Attributes, computed) - Map Flow Sample Sip Rules Container. Private class (see [below for nested schema](#nestedatt--flow_sample_sip_rules))
* `flow_whitelist5_g_overlap_rules` (Attributes, computed) - Map Flow Whitelist 5g Overlap Rule match Definition. Private class (see [below for nested schema](#nestedatt--flow_whitelist5_g_overlap_rules))
* `flow_whitelist5_g_rules` (Attributes, computed) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--flow_whitelist5_g_rules))
* `flow_whitelist_overlap_rules` (Attributes, computed) - Map Flow Whitelist Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--flow_whitelist_overlap_rules))
* `flow_whitelist_rules` (Attributes, computed) - Map Flow Whitelist Rules Container. Private class (see [below for nested schema](#nestedatt--flow_whitelist_rules))
* `fstype` (Attributes, computed) - Type of Mobility Flowsampling & properties (see [below for nested schema](#nestedatt--fstype))
* `gs_rules` (Attributes, computed) - Map GigaSMART Rules Container. Private class (see [below for nested schema](#nestedatt--gs_rules))
* `gsop` (String, computed) - Alias of referenced GSOP. Applicable to 'regular' and 'secondLevel' map types
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `inline_traffic_path` (String, computed) - Only applicable for 'inline' map types, in which case defaults to 'normal'. When set to 'bypass', the 'dstPorts' must be empty
* `inline_traffic_type` (String, computed) - Only applicable for 'inline/passAll' map types, in which case defaults to 'symmetric'. For 'asymmetric' maps, 'srcPort' must be of type 'inline-net'
* `ip_rewrite` (Attributes, computed) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--ip_rewrite))
* `mod_time` (Number, computed) - Last modification time of the map in milliseconds since the epoch
* `null_dst_port` (Boolean, computed) - enabled when the dstPort is null
* `order` (Number, computed) - relative order within per-source port map chain
* `rewrite` (Attributes, computed) - Rewrite options on the packets (see [below for nested schema](#nestedatt--rewrite))
* `roles` (Attributes, computed) - Map Roles Container. Private class (see [below for nested schema](#nestedatt--roles))
* `rule_matching` (String, computed) - If set to 'blacklist', packet are passed when no 'drop' rules are matched. This field is only valid for 'regular/byRule' map types. (maps into 'no-rule-match' CLI command)
* `rules` (Attributes, computed) - Map Rules Container. Private class (see [below for nested schema](#nestedatt--rules))
* `rx_cluster_ports` (List of String, computed)
* `sub_type` (String, computed) - 'byRule' is applicable to all map types; 'collector' is applicable to 'regular', 'inline', 'flexInline' and 'secondLevel' maps; 'passAll' is applicable to 'regular' and 'inline' maps; 'flowFilter', 'flowSample', 'flowWhitelist', 'flowSampleSip', 'flowWhitelistSip', 'flowSampleDiameter', 'flowWhitelistDiameter', 'flowSampleOverlap', 'flowWhitelistOverlap', 'flowSample5g','flowWhitelist5g','flowSample5gOverlap' and 'flowWhitelist5gOverlap' are applicable to 'secondLevel' maps
* `traffic_type` (String, computed) - Only applicable for 'firstLevel/byRule' map types, in which case defaults to 'user'
* `tx_cluster_ports` (List of String, computed)
* `updated_time` (Number, computed) - Last Updated time of the fabric map
* `vlan_tag` (Attributes, computed) (see [below for nested schema](#nestedatt--vlan_tag))

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
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--rules--drop_rules--vlan_tag))
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
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--rules--pass_rules--vlan_tag))
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
<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_map.example {alias}
```
