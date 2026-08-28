---
page_title: "gigavuecore_map_chain Resource - gigavuecore"
subcategory: ""
description: |-
  Load map chains either by mapChainId or srcPortsAsId
---

# gigavuecore_map_chain Resource

Load map chains either by mapChainId or srcPortsAsId

## Example Usage

```terraform
resource "gigavuecore_map_chain" "example" {
  cluster_id          = null
  id                  = null
  map_chain_id        = null
  ordered_map_aliases = []
  src_ports_as_id     = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - id of the defining cluster
* `id` (String, required)
* `map_chain_id` (String, optional) - mapChain ID - replaces srcPortsAsId - should be used anywhere that requires srcPortsAsId
* `ordered_map_aliases` (Set of String, optional) - ordered list of map aliases in the chain. Can be a combination of cluster map and fabric map aliases.
* `src_ports_as_id` (String, optional) - (Deprecated - use mapChainId instead) mapChain ID

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `collector` (Attributes, computed) (see [below for nested schema](#nestedatt--collector))
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `map_chain_id` (String, computed) - mapChain ID - replaces srcPortsAsId - should be used anywhere that requires srcPortsAsId
* `maps` (Attributes Set, computed) (see [below for nested schema](#nestedatt--maps))
* `src_ports` (List of String, computed) - network ports for the maps chain
* `src_ports_as_id` (String, computed) - (Deprecated - use mapChainId instead) mapChain ID

<a id="nestedatt--collector"></a>
### Nested Schema for `collector`

Read-Only:

* `alias` (String) - unique map alias
* `ap_rules` (Attributes) - pass and drop application profile rules (see [below for nested schema](#nestedatt--collector--ap_rules))
* `cluster_id` (String) - id of the defining cluster
* `comment` (String)
* `dst_ports` (List of String) - List of the 'to' ports. Invalid if 'nullDstPort' is true'. Only port number is supported
* `egress_gigastream` (List of String) - Applicable to fabric map only. It is to specify gigastream(s) to use as egress when going across cluster. The format of array member is clusterID:gigastreamAlias
* `enable` (Boolean) - enable/disable map, applicable only to first level maps
* `encap_tunnel` (String) - tunnel alias
* `flex_inline` (Attributes) - When 'oobCopy' is configured at least one of 'aToB' or 'bToA' must be configured (see [below for nested schema](#nestedatt--collector--flex_inline))
* `flex_inline_failover` (String) - only valid for flexInline maps
* `flex_inline_vlan_id` (Number) - VLAN ID carried in the VLAN tag of packets coming from the inline-network port(s). valid and applicable only with 'flexInline' map type, 'srcPorts' should be either inline-network or ib-pathway or vport
* `flow_rules` (Attributes) - Map Flow Rules Container. Private class (see [below for nested schema](#nestedatt--collector--flow_rules))
* `flow_sample5_g_overlap_rules` (Attributes) - Map Flow Sample 5g Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--collector--flow_sample5_g_overlap_rules))
* `flow_sample5_g_rules` (Attributes) - Map Flow Sample 5g Rules Container. Private class (see [below for nested schema](#nestedatt--collector--flow_sample5_g_rules))
* `flow_sample_diameter_rules` (Attributes) - Map Flow Sample Diameter Rules Container. Private class (see [below for nested schema](#nestedatt--collector--flow_sample_diameter_rules))
* `flow_sample_overlap_rules` (Attributes) - Map Flow Sample Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--collector--flow_sample_overlap_rules))
* `flow_sample_rules` (Attributes) - Map Flow Sample Rules Container. Private class (see [below for nested schema](#nestedatt--collector--flow_sample_rules))
* `flow_sample_sip_rules` (Attributes) - Map Flow Sample Sip Rules Container. Private class (see [below for nested schema](#nestedatt--collector--flow_sample_sip_rules))
* `flow_whitelist5_g_overlap_rules` (Attributes) - Map Flow Whitelist 5g Overlap Rule match Definition. Private class (see [below for nested schema](#nestedatt--collector--flow_whitelist5_g_overlap_rules))
* `flow_whitelist5_g_rules` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--collector--flow_whitelist5_g_rules))
* `flow_whitelist_overlap_rules` (Attributes) - Map Flow Whitelist Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--collector--flow_whitelist_overlap_rules))
* `flow_whitelist_rules` (Attributes) - Map Flow Whitelist Rules Container. Private class (see [below for nested schema](#nestedatt--collector--flow_whitelist_rules))
* `fstype` (Attributes) - Type of Mobility Flowsampling & properties (see [below for nested schema](#nestedatt--collector--fstype))
* `gs_rules` (Attributes) - Map GigaSMART Rules Container. Private class (see [below for nested schema](#nestedatt--collector--gs_rules))
* `gsop` (String) - Alias of referenced GSOP. Applicable to 'regular' and 'secondLevel' map types
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--collector--health_state_reasons))
* `inline_traffic_path` (String) - Only applicable for 'inline' map types, in which case defaults to 'normal'. When set to 'bypass', the 'dstPorts' must be empty
* `inline_traffic_type` (String) - Only applicable for 'inline/passAll' map types, in which case defaults to 'symmetric'. For 'asymmetric' maps, 'srcPort' must be of type 'inline-net'
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--collector--ip_rewrite))
* `mod_time` (Number) - Last modification time of the map in milliseconds since the epoch
* `null_dst_port` (Boolean) - enabled when the dstPort is null
* `order` (Number) - relative order within per-source port map chain
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--collector--rewrite))
* `roles` (Attributes) - Map Roles Container. Private class (see [below for nested schema](#nestedatt--collector--roles))
* `rule_matching` (String) - If set to 'blacklist', packet are passed when no 'drop' rules are matched. This field is only valid for 'regular/byRule' map types. (maps into 'no-rule-match' CLI command)
* `rules` (Attributes) - Map Rules Container. Private class (see [below for nested schema](#nestedatt--collector--rules))
* `rx_cluster_ports` (List of String)
* `src_ports` (List of String) - list of the 'from' ports. Only port number is supported
* `sub_type` (String) - 'byRule' is applicable to all map types; 'collector' is applicable to 'regular', 'inline', 'flexInline' and 'secondLevel' maps; 'passAll' is applicable to 'regular' and 'inline' maps; 'flowFilter', 'flowSample', 'flowWhitelist', 'flowSampleSip', 'flowWhitelistSip', 'flowSampleDiameter', 'flowWhitelistDiameter', 'flowSampleOverlap', 'flowWhitelistOverlap', 'flowSample5g','flowWhitelist5g','flowSample5gOverlap' and 'flowWhitelist5gOverlap' are applicable to 'secondLevel' maps
* `traffic_type` (String) - Only applicable for 'firstLevel/byRule' map types, in which case defaults to 'user'
* `tx_cluster_ports` (List of String)
* `type` (String) - 'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'flexInline' is from inline network to tools; 'transitLevel' is from vport to vport
* `updated_time` (Number) - Last Updated time of the fabric map
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--collector--vlan_tag))
<a id="nestedatt--collector--ap_rules"></a>
### Nested Schema for `collector.ap_rules`

Read-Only:

* `drop_rules` (Attributes List) (see [below for nested schema](#nestedatt--collector--ap_rules--drop_rules))
* `pass_rules` (Attributes List) (see [below for nested schema](#nestedatt--collector--ap_rules--pass_rules))
<a id="nestedatt--collector--ap_rules--drop_rules"></a>
### Nested Schema for `collector.ap_rules.drop_rules`

Read-Only:

* `application_profile` (String) - application profile alias
* `rule_id` (Number) - application profile rule Id, should not have same id as gsRules
<a id="nestedatt--collector--ap_rules--pass_rules"></a>
### Nested Schema for `collector.ap_rules.pass_rules`

Read-Only:

* `application_profile` (String) - application profile alias
* `rule_id` (Number) - application profile rule Id, should not have same id as gsRules
<a id="nestedatt--collector--flex_inline"></a>
### Nested Schema for `collector.flex_inline`

Read-Only:

* `a_to_b` (Attributes) (see [below for nested schema](#nestedatt--collector--flex_inline--a_to_b))
* `b_to_a` (Attributes) (see [below for nested schema](#nestedatt--collector--flex_inline--b_to_a))
* `oob_copy` (Attributes List) (see [below for nested schema](#nestedatt--collector--flex_inline--oob_copy))
* `svt_mode` (Boolean)
* `svt_tag` (Number) - only applicable when svtMode is enabled
* `tag` (Attributes) (see [below for nested schema](#nestedatt--collector--flex_inline--tag))
<a id="nestedatt--collector--flex_inline--a_to_b"></a>
### Nested Schema for `collector.flex_inline.a_to_b`

Read-Only:

* `ib_pathway` (String) - ibPathway alias. Only applicable when type is 'ibPathway'
* `tools` (List of String) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'
* `type` (String) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided
<a id="nestedatt--collector--flex_inline--b_to_a"></a>
### Nested Schema for `collector.flex_inline.b_to_a`

Read-Only:

* `ib_pathway` (String) - ibPathway alias. Only applicable when type is 'ibPathway'
* `tools` (List of String) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'
* `type` (String) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided
<a id="nestedatt--collector--flex_inline--oob_copy"></a>
### Nested Schema for `collector.flex_inline.oob_copy`

Read-Only:

* `direction` (String)
* `dst_ports` (List of String) - list of destination tool ports
* `src_ports` (List of String) - inline network or an item from a-to-b and b-to-a lists
* `tag` (Attributes) (see [below for nested schema](#nestedatt--collector--flex_inline--oob_copy--tag))
<a id="nestedatt--collector--flex_inline--oob_copy--tag"></a>
### Nested Schema for `collector.flex_inline.oob_copy.tag`

Read-Only:

* `type` (String)
<a id="nestedatt--collector--flex_inline--tag"></a>
### Nested Schema for `collector.flex_inline.tag`

Read-Only:

* `tag_protocol_id` (String) - When tool VLAN tag is added , this protocol Id will be added which egress out the traffic
* `type` (String)
* `vlan_id` (Number) - only applicable when type is 'vlan'
<a id="nestedatt--collector--flow_rules"></a>
### Nested Schema for `collector.flow_rules`

Read-Only:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--collector--flow_rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--collector--flow_rules--pass_rules))
<a id="nestedatt--collector--flow_rules--drop_rules"></a>
### Nested Schema for `collector.flow_rules.drop_rules`

Read-Only:

* `gtp` (Attributes) - Map Flow Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--collector--flow_rules--drop_rules--gtp))
* `rule_id` (Number)
<a id="nestedatt--collector--flow_rules--drop_rules--gtp"></a>
### Nested Schema for `collector.flow_rules.drop_rules.gtp`

Read-Only:

* `imei` (String) - mutually exclusive with 'imsi' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - mutually exclusive with 'imei' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - mutually exclusive with 'imsi' and 'imei'. If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--collector--flow_rules--pass_rules"></a>
### Nested Schema for `collector.flow_rules.pass_rules`

Read-Only:

* `gtp` (Attributes) - Map Flow Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--collector--flow_rules--pass_rules--gtp))
* `rule_id` (Number)
<a id="nestedatt--collector--flow_rules--pass_rules--gtp"></a>
### Nested Schema for `collector.flow_rules.pass_rules.gtp`

Read-Only:

* `imei` (String) - mutually exclusive with 'imsi' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - mutually exclusive with 'imei' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - mutually exclusive with 'imsi' and 'imei'. If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--collector--flow_sample5_g_overlap_rules"></a>
### Nested Schema for `collector.flow_sample5_g_overlap_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--collector--flow_sample5_g_overlap_rules--pass_rules))
<a id="nestedatt--collector--flow_sample5_g_overlap_rules--pass_rules"></a>
### Nested Schema for `collector.flow_sample5_g_overlap_rules.pass_rules`

Read-Only:

* `comment` (String)
* `flow5_g` (Attributes) - Map Flow Sample Overlap Rule 5g match Definition. Private class (see [below for nested schema](#nestedatt--collector--flow_sample5_g_overlap_rules--pass_rules--flow5_g))
* `percentage` (Number)
* `rule_id` (Number)
<a id="nestedatt--collector--flow_sample5_g_overlap_rules--pass_rules--flow5_g"></a>
### Nested Schema for `collector.flow_sample5_g_overlap_rules.pass_rules.flow5_g`

Read-Only:

* `dnn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `gpsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nsiid` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `pei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `supi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--collector--flow_sample5_g_rules"></a>
### Nested Schema for `collector.flow_sample5_g_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--collector--flow_sample5_g_rules--pass_rules))
<a id="nestedatt--collector--flow_sample5_g_rules--pass_rules"></a>
### Nested Schema for `collector.flow_sample5_g_rules.pass_rules`

Read-Only:

* `comment` (String)
* `flow5_g` (Attributes) - Map Flow Sample Rule 5g match Definition. Private class (see [below for nested schema](#nestedatt--collector--flow_sample5_g_rules--pass_rules--flow5_g))
* `percentage` (Number)
* `priority` (Number)
* `rule_id` (Number)
<a id="nestedatt--collector--flow_sample5_g_rules--pass_rules--flow5_g"></a>
### Nested Schema for `collector.flow_sample5_g_rules.pass_rules.flow5_g`

Read-Only:

* `dnn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `gpsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nsiid` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `pei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `supi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--collector--flow_sample_diameter_rules"></a>
### Nested Schema for `collector.flow_sample_diameter_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--collector--flow_sample_diameter_rules--pass_rules))
<a id="nestedatt--collector--flow_sample_diameter_rules--pass_rules"></a>
### Nested Schema for `collector.flow_sample_diameter_rules.pass_rules`

Read-Only:

* `diameter` (Attributes) - Map Flow Sample Diameter Rule Definition (see [below for nested schema](#nestedatt--collector--flow_sample_diameter_rules--pass_rules--diameter))
* `interface` (String) - interface type
* `percentage` (Number)
* `rule_id` (Number)
<a id="nestedatt--collector--flow_sample_diameter_rules--pass_rules--diameter"></a>
### Nested Schema for `collector.flow_sample_diameter_rules.pass_rules.diameter`

Read-Only:

* `user_name` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--collector--flow_sample_overlap_rules"></a>
### Nested Schema for `collector.flow_sample_overlap_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--collector--flow_sample_overlap_rules--pass_rules))
<a id="nestedatt--collector--flow_sample_overlap_rules--pass_rules"></a>
### Nested Schema for `collector.flow_sample_overlap_rules.pass_rules`

Read-Only:

* `comment` (String)
* `gtp` (Attributes) - Map Flow Sample Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--collector--flow_sample_overlap_rules--pass_rules--gtp))
* `percentage` (Number)
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number) - maximum value is equal to the number of rules upon completion of the request
* `rule_id` (Number)
<a id="nestedatt--collector--flow_sample_overlap_rules--pass_rules--gtp"></a>
### Nested Schema for `collector.flow_sample_overlap_rules.pass_rules.gtp`

Read-Only:

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
<a id="nestedatt--collector--flow_sample_rules"></a>
### Nested Schema for `collector.flow_sample_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--collector--flow_sample_rules--pass_rules))
<a id="nestedatt--collector--flow_sample_rules--pass_rules"></a>
### Nested Schema for `collector.flow_sample_rules.pass_rules`

Read-Only:

* `comment` (String)
* `gtp` (Attributes) - Map Flow Sample Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--collector--flow_sample_rules--pass_rules--gtp))
* `percentage` (Number)
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number) - maximum value is equal to the number of rules upon completion of the request
* `rule_id` (Number)
<a id="nestedatt--collector--flow_sample_rules--pass_rules--gtp"></a>
### Nested Schema for `collector.flow_sample_rules.pass_rules.gtp`

Read-Only:

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
<a id="nestedatt--collector--flow_sample_sip_rules"></a>
### Nested Schema for `collector.flow_sample_sip_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--collector--flow_sample_sip_rules--pass_rules))
<a id="nestedatt--collector--flow_sample_sip_rules--pass_rules"></a>
### Nested Schema for `collector.flow_sample_sip_rules.pass_rules`

Read-Only:

* `percentage` (Number)
* `rule_id` (Number)
* `sip` (Attributes) (see [below for nested schema](#nestedatt--collector--flow_sample_sip_rules--pass_rules--sip))
<a id="nestedatt--collector--flow_sample_sip_rules--pass_rules--sip"></a>
### Nested Schema for `collector.flow_sample_sip_rules.pass_rules.sip`

Read-Only:

* `callee_id` (String) - sip callee id
* `callee_id_range` (Attributes) (see [below for nested schema](#nestedatt--collector--flow_sample_sip_rules--pass_rules--sip--callee_id_range))
* `caller_id` (String) - sip caller id
* `caller_id_range` (Attributes) (see [below for nested schema](#nestedatt--collector--flow_sample_sip_rules--pass_rules--sip--caller_id_range))
* `id_range` (Attributes) (see [below for nested schema](#nestedatt--collector--flow_sample_sip_rules--pass_rules--sip--id_range))
<a id="nestedatt--collector--flow_sample_sip_rules--pass_rules--sip--callee_id_range"></a>
### Nested Schema for `collector.flow_sample_sip_rules.pass_rules.sip.callee_id_range`

Read-Only:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--collector--flow_sample_sip_rules--pass_rules--sip--caller_id_range"></a>
### Nested Schema for `collector.flow_sample_sip_rules.pass_rules.sip.caller_id_range`

Read-Only:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--collector--flow_sample_sip_rules--pass_rules--sip--id_range"></a>
### Nested Schema for `collector.flow_sample_sip_rules.pass_rules.sip.id_range`

Read-Only:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--collector--flow_whitelist5_g_overlap_rules"></a>
### Nested Schema for `collector.flow_whitelist5_g_overlap_rules`

Read-Only:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
<a id="nestedatt--collector--flow_whitelist5_g_rules"></a>
### Nested Schema for `collector.flow_whitelist5_g_rules`

Read-Only:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--collector--flow_whitelist_overlap_rules"></a>
### Nested Schema for `collector.flow_whitelist_overlap_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--collector--flow_whitelist_overlap_rules--pass_rules))
<a id="nestedatt--collector--flow_whitelist_overlap_rules--pass_rules"></a>
### Nested Schema for `collector.flow_whitelist_overlap_rules.pass_rules`

Read-Only:

* `flow5_g` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--collector--flow_whitelist_overlap_rules--pass_rules--flow5_g))
* `gtp` (Attributes) - Map Flow Whitelist Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--collector--flow_whitelist_overlap_rules--pass_rules--gtp))
* `rule_id` (Number)
* `sip` (Attributes) - Map Flow Whitelist Rule Sip match definition (see [below for nested schema](#nestedatt--collector--flow_whitelist_overlap_rules--pass_rules--sip))
<a id="nestedatt--collector--flow_whitelist_overlap_rules--pass_rules--flow5_g"></a>
### Nested Schema for `collector.flow_whitelist_overlap_rules.pass_rules.flow5_g`

Read-Only:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--collector--flow_whitelist_overlap_rules--pass_rules--gtp"></a>
### Nested Schema for `collector.flow_whitelist_overlap_rules.pass_rules.gtp`

Read-Only:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--collector--flow_whitelist_overlap_rules--pass_rules--sip"></a>
### Nested Schema for `collector.flow_whitelist_overlap_rules.pass_rules.sip`

Read-Only:

* `type` (String) - all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address
<a id="nestedatt--collector--flow_whitelist_rules"></a>
### Nested Schema for `collector.flow_whitelist_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--collector--flow_whitelist_rules--pass_rules))
<a id="nestedatt--collector--flow_whitelist_rules--pass_rules"></a>
### Nested Schema for `collector.flow_whitelist_rules.pass_rules`

Read-Only:

* `flow5_g` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--collector--flow_whitelist_rules--pass_rules--flow5_g))
* `gtp` (Attributes) - Map Flow Whitelist Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--collector--flow_whitelist_rules--pass_rules--gtp))
* `rule_id` (Number)
* `sip` (Attributes) - Map Flow Whitelist Rule Sip match definition (see [below for nested schema](#nestedatt--collector--flow_whitelist_rules--pass_rules--sip))
<a id="nestedatt--collector--flow_whitelist_rules--pass_rules--flow5_g"></a>
### Nested Schema for `collector.flow_whitelist_rules.pass_rules.flow5_g`

Read-Only:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--collector--flow_whitelist_rules--pass_rules--gtp"></a>
### Nested Schema for `collector.flow_whitelist_rules.pass_rules.gtp`

Read-Only:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--collector--flow_whitelist_rules--pass_rules--sip"></a>
### Nested Schema for `collector.flow_whitelist_rules.pass_rules.sip`

Read-Only:

* `type` (String) - all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address
<a id="nestedatt--collector--fstype"></a>
### Nested Schema for `collector.fstype`

Read-Only:

* `offset` (Number) - Offset for Mobility Rotational Flowsampling
* `timer` (Number) - Timer for Mobility Rotational Flowsampling in minutes
* `type` (String)
<a id="nestedatt--collector--gs_rules"></a>
### Nested Schema for `collector.gs_rules`

Read-Only:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--collector--gs_rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--collector--gs_rules--pass_rules))
<a id="nestedatt--collector--gs_rules--drop_rules"></a>
### Nested Schema for `collector.gs_rules.drop_rules`

Read-Only:

* `comment` (String)
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number)
<a id="nestedatt--collector--gs_rules--pass_rules"></a>
### Nested Schema for `collector.gs_rules.pass_rules`

Read-Only:

* `comment` (String)
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number)
<a id="nestedatt--collector--health_state_reasons"></a>
### Nested Schema for `collector.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--collector--ip_rewrite"></a>
### Nested Schema for `collector.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--collector--rewrite"></a>
### Nested Schema for `collector.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--collector--roles"></a>
### Nested Schema for `collector.roles`

Read-Only:

* `editors` (List of String)
* `listeners` (List of String)
* `owners` (List of String)
* `viewers` (List of String)
<a id="nestedatt--collector--rules"></a>
### Nested Schema for `collector.rules`

Read-Only:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--collector--rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--collector--rules--pass_rules))
<a id="nestedatt--collector--rules--drop_rules"></a>
### Nested Schema for `collector.rules.drop_rules`

Read-Only:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--collector--rules--drop_rules--ip_rewrite))
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--collector--rules--drop_rules--rewrite))
* `rule_id` (Number)
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--collector--rules--drop_rules--vlan_tag))
<a id="nestedatt--collector--rules--drop_rules--ip_rewrite"></a>
### Nested Schema for `collector.rules.drop_rules.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--collector--rules--drop_rules--rewrite"></a>
### Nested Schema for `collector.rules.drop_rules.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--collector--rules--drop_rules--vlan_tag"></a>
### Nested Schema for `collector.rules.drop_rules.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)
<a id="nestedatt--collector--rules--pass_rules"></a>
### Nested Schema for `collector.rules.pass_rules`

Read-Only:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--collector--rules--pass_rules--ip_rewrite))
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--collector--rules--pass_rules--rewrite))
* `rule_id` (Number)
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--collector--rules--pass_rules--vlan_tag))
<a id="nestedatt--collector--rules--pass_rules--ip_rewrite"></a>
### Nested Schema for `collector.rules.pass_rules.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--collector--rules--pass_rules--rewrite"></a>
### Nested Schema for `collector.rules.pass_rules.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--collector--rules--pass_rules--vlan_tag"></a>
### Nested Schema for `collector.rules.pass_rules.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)
<a id="nestedatt--collector--vlan_tag"></a>
### Nested Schema for `collector.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)
<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--maps"></a>
### Nested Schema for `maps`

Read-Only:

* `alias` (String) - unique map alias
* `ap_rules` (Attributes) - pass and drop application profile rules (see [below for nested schema](#nestedatt--maps--ap_rules))
* `cluster_id` (String) - id of the defining cluster
* `comment` (String)
* `dst_ports` (List of String) - List of the 'to' ports. Invalid if 'nullDstPort' is true'. Only port number is supported
* `egress_gigastream` (List of String) - Applicable to fabric map only. It is to specify gigastream(s) to use as egress when going across cluster. The format of array member is clusterID:gigastreamAlias
* `enable` (Boolean) - enable/disable map, applicable only to first level maps
* `encap_tunnel` (String) - tunnel alias
* `flex_inline` (Attributes) - When 'oobCopy' is configured at least one of 'aToB' or 'bToA' must be configured (see [below for nested schema](#nestedatt--maps--flex_inline))
* `flex_inline_failover` (String) - only valid for flexInline maps
* `flex_inline_vlan_id` (Number) - VLAN ID carried in the VLAN tag of packets coming from the inline-network port(s). valid and applicable only with 'flexInline' map type, 'srcPorts' should be either inline-network or ib-pathway or vport
* `flow_rules` (Attributes) - Map Flow Rules Container. Private class (see [below for nested schema](#nestedatt--maps--flow_rules))
* `flow_sample5_g_overlap_rules` (Attributes) - Map Flow Sample 5g Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--maps--flow_sample5_g_overlap_rules))
* `flow_sample5_g_rules` (Attributes) - Map Flow Sample 5g Rules Container. Private class (see [below for nested schema](#nestedatt--maps--flow_sample5_g_rules))
* `flow_sample_diameter_rules` (Attributes) - Map Flow Sample Diameter Rules Container. Private class (see [below for nested schema](#nestedatt--maps--flow_sample_diameter_rules))
* `flow_sample_overlap_rules` (Attributes) - Map Flow Sample Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--maps--flow_sample_overlap_rules))
* `flow_sample_rules` (Attributes) - Map Flow Sample Rules Container. Private class (see [below for nested schema](#nestedatt--maps--flow_sample_rules))
* `flow_sample_sip_rules` (Attributes) - Map Flow Sample Sip Rules Container. Private class (see [below for nested schema](#nestedatt--maps--flow_sample_sip_rules))
* `flow_whitelist5_g_overlap_rules` (Attributes) - Map Flow Whitelist 5g Overlap Rule match Definition. Private class (see [below for nested schema](#nestedatt--maps--flow_whitelist5_g_overlap_rules))
* `flow_whitelist5_g_rules` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--maps--flow_whitelist5_g_rules))
* `flow_whitelist_overlap_rules` (Attributes) - Map Flow Whitelist Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--maps--flow_whitelist_overlap_rules))
* `flow_whitelist_rules` (Attributes) - Map Flow Whitelist Rules Container. Private class (see [below for nested schema](#nestedatt--maps--flow_whitelist_rules))
* `fstype` (Attributes) - Type of Mobility Flowsampling & properties (see [below for nested schema](#nestedatt--maps--fstype))
* `gs_rules` (Attributes) - Map GigaSMART Rules Container. Private class (see [below for nested schema](#nestedatt--maps--gs_rules))
* `gsop` (String) - Alias of referenced GSOP. Applicable to 'regular' and 'secondLevel' map types
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--maps--health_state_reasons))
* `inline_traffic_path` (String) - Only applicable for 'inline' map types, in which case defaults to 'normal'. When set to 'bypass', the 'dstPorts' must be empty
* `inline_traffic_type` (String) - Only applicable for 'inline/passAll' map types, in which case defaults to 'symmetric'. For 'asymmetric' maps, 'srcPort' must be of type 'inline-net'
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--maps--ip_rewrite))
* `mod_time` (Number) - Last modification time of the map in milliseconds since the epoch
* `null_dst_port` (Boolean) - enabled when the dstPort is null
* `order` (Number) - relative order within per-source port map chain
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--maps--rewrite))
* `roles` (Attributes) - Map Roles Container. Private class (see [below for nested schema](#nestedatt--maps--roles))
* `rule_matching` (String) - If set to 'blacklist', packet are passed when no 'drop' rules are matched. This field is only valid for 'regular/byRule' map types. (maps into 'no-rule-match' CLI command)
* `rules` (Attributes) - Map Rules Container. Private class (see [below for nested schema](#nestedatt--maps--rules))
* `rx_cluster_ports` (List of String)
* `src_ports` (List of String) - list of the 'from' ports. Only port number is supported
* `sub_type` (String) - 'byRule' is applicable to all map types; 'collector' is applicable to 'regular', 'inline', 'flexInline' and 'secondLevel' maps; 'passAll' is applicable to 'regular' and 'inline' maps; 'flowFilter', 'flowSample', 'flowWhitelist', 'flowSampleSip', 'flowWhitelistSip', 'flowSampleDiameter', 'flowWhitelistDiameter', 'flowSampleOverlap', 'flowWhitelistOverlap', 'flowSample5g','flowWhitelist5g','flowSample5gOverlap' and 'flowWhitelist5gOverlap' are applicable to 'secondLevel' maps
* `traffic_type` (String) - Only applicable for 'firstLevel/byRule' map types, in which case defaults to 'user'
* `tx_cluster_ports` (List of String)
* `type` (String) - 'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'flexInline' is from inline network to tools; 'transitLevel' is from vport to vport
* `updated_time` (Number) - Last Updated time of the fabric map
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--maps--vlan_tag))
<a id="nestedatt--maps--ap_rules"></a>
### Nested Schema for `maps.ap_rules`

Read-Only:

* `drop_rules` (Attributes List) (see [below for nested schema](#nestedatt--maps--ap_rules--drop_rules))
* `pass_rules` (Attributes List) (see [below for nested schema](#nestedatt--maps--ap_rules--pass_rules))
<a id="nestedatt--maps--ap_rules--drop_rules"></a>
### Nested Schema for `maps.ap_rules.drop_rules`

Read-Only:

* `application_profile` (String) - application profile alias
* `rule_id` (Number) - application profile rule Id, should not have same id as gsRules
<a id="nestedatt--maps--ap_rules--pass_rules"></a>
### Nested Schema for `maps.ap_rules.pass_rules`

Read-Only:

* `application_profile` (String) - application profile alias
* `rule_id` (Number) - application profile rule Id, should not have same id as gsRules
<a id="nestedatt--maps--flex_inline"></a>
### Nested Schema for `maps.flex_inline`

Read-Only:

* `a_to_b` (Attributes) (see [below for nested schema](#nestedatt--maps--flex_inline--a_to_b))
* `b_to_a` (Attributes) (see [below for nested schema](#nestedatt--maps--flex_inline--b_to_a))
* `oob_copy` (Attributes List) (see [below for nested schema](#nestedatt--maps--flex_inline--oob_copy))
* `svt_mode` (Boolean)
* `svt_tag` (Number) - only applicable when svtMode is enabled
* `tag` (Attributes) (see [below for nested schema](#nestedatt--maps--flex_inline--tag))
<a id="nestedatt--maps--flex_inline--a_to_b"></a>
### Nested Schema for `maps.flex_inline.a_to_b`

Read-Only:

* `ib_pathway` (String) - ibPathway alias. Only applicable when type is 'ibPathway'
* `tools` (List of String) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'
* `type` (String) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided
<a id="nestedatt--maps--flex_inline--b_to_a"></a>
### Nested Schema for `maps.flex_inline.b_to_a`

Read-Only:

* `ib_pathway` (String) - ibPathway alias. Only applicable when type is 'ibPathway'
* `tools` (List of String) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'
* `type` (String) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided
<a id="nestedatt--maps--flex_inline--oob_copy"></a>
### Nested Schema for `maps.flex_inline.oob_copy`

Read-Only:

* `direction` (String)
* `dst_ports` (List of String) - list of destination tool ports
* `src_ports` (List of String) - inline network or an item from a-to-b and b-to-a lists
* `tag` (Attributes) (see [below for nested schema](#nestedatt--maps--flex_inline--oob_copy--tag))
<a id="nestedatt--maps--flex_inline--oob_copy--tag"></a>
### Nested Schema for `maps.flex_inline.oob_copy.tag`

Read-Only:

* `type` (String)
<a id="nestedatt--maps--flex_inline--tag"></a>
### Nested Schema for `maps.flex_inline.tag`

Read-Only:

* `tag_protocol_id` (String) - When tool VLAN tag is added , this protocol Id will be added which egress out the traffic
* `type` (String)
* `vlan_id` (Number) - only applicable when type is 'vlan'
<a id="nestedatt--maps--flow_rules"></a>
### Nested Schema for `maps.flow_rules`

Read-Only:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--maps--flow_rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--maps--flow_rules--pass_rules))
<a id="nestedatt--maps--flow_rules--drop_rules"></a>
### Nested Schema for `maps.flow_rules.drop_rules`

Read-Only:

* `gtp` (Attributes) - Map Flow Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--maps--flow_rules--drop_rules--gtp))
* `rule_id` (Number)
<a id="nestedatt--maps--flow_rules--drop_rules--gtp"></a>
### Nested Schema for `maps.flow_rules.drop_rules.gtp`

Read-Only:

* `imei` (String) - mutually exclusive with 'imsi' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - mutually exclusive with 'imei' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - mutually exclusive with 'imsi' and 'imei'. If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--maps--flow_rules--pass_rules"></a>
### Nested Schema for `maps.flow_rules.pass_rules`

Read-Only:

* `gtp` (Attributes) - Map Flow Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--maps--flow_rules--pass_rules--gtp))
* `rule_id` (Number)
<a id="nestedatt--maps--flow_rules--pass_rules--gtp"></a>
### Nested Schema for `maps.flow_rules.pass_rules.gtp`

Read-Only:

* `imei` (String) - mutually exclusive with 'imsi' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - mutually exclusive with 'imei' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - mutually exclusive with 'imsi' and 'imei'. If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--maps--flow_sample5_g_overlap_rules"></a>
### Nested Schema for `maps.flow_sample5_g_overlap_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--maps--flow_sample5_g_overlap_rules--pass_rules))
<a id="nestedatt--maps--flow_sample5_g_overlap_rules--pass_rules"></a>
### Nested Schema for `maps.flow_sample5_g_overlap_rules.pass_rules`

Read-Only:

* `comment` (String)
* `flow5_g` (Attributes) - Map Flow Sample Overlap Rule 5g match Definition. Private class (see [below for nested schema](#nestedatt--maps--flow_sample5_g_overlap_rules--pass_rules--flow5_g))
* `percentage` (Number)
* `rule_id` (Number)
<a id="nestedatt--maps--flow_sample5_g_overlap_rules--pass_rules--flow5_g"></a>
### Nested Schema for `maps.flow_sample5_g_overlap_rules.pass_rules.flow5_g`

Read-Only:

* `dnn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `gpsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nsiid` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `pei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `supi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--maps--flow_sample5_g_rules"></a>
### Nested Schema for `maps.flow_sample5_g_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--maps--flow_sample5_g_rules--pass_rules))
<a id="nestedatt--maps--flow_sample5_g_rules--pass_rules"></a>
### Nested Schema for `maps.flow_sample5_g_rules.pass_rules`

Read-Only:

* `comment` (String)
* `flow5_g` (Attributes) - Map Flow Sample Rule 5g match Definition. Private class (see [below for nested schema](#nestedatt--maps--flow_sample5_g_rules--pass_rules--flow5_g))
* `percentage` (Number)
* `priority` (Number)
* `rule_id` (Number)
<a id="nestedatt--maps--flow_sample5_g_rules--pass_rules--flow5_g"></a>
### Nested Schema for `maps.flow_sample5_g_rules.pass_rules.flow5_g`

Read-Only:

* `dnn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `gpsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nsiid` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `pei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `supi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--maps--flow_sample_diameter_rules"></a>
### Nested Schema for `maps.flow_sample_diameter_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--maps--flow_sample_diameter_rules--pass_rules))
<a id="nestedatt--maps--flow_sample_diameter_rules--pass_rules"></a>
### Nested Schema for `maps.flow_sample_diameter_rules.pass_rules`

Read-Only:

* `diameter` (Attributes) - Map Flow Sample Diameter Rule Definition (see [below for nested schema](#nestedatt--maps--flow_sample_diameter_rules--pass_rules--diameter))
* `interface` (String) - interface type
* `percentage` (Number)
* `rule_id` (Number)
<a id="nestedatt--maps--flow_sample_diameter_rules--pass_rules--diameter"></a>
### Nested Schema for `maps.flow_sample_diameter_rules.pass_rules.diameter`

Read-Only:

* `user_name` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--maps--flow_sample_overlap_rules"></a>
### Nested Schema for `maps.flow_sample_overlap_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--maps--flow_sample_overlap_rules--pass_rules))
<a id="nestedatt--maps--flow_sample_overlap_rules--pass_rules"></a>
### Nested Schema for `maps.flow_sample_overlap_rules.pass_rules`

Read-Only:

* `comment` (String)
* `gtp` (Attributes) - Map Flow Sample Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--maps--flow_sample_overlap_rules--pass_rules--gtp))
* `percentage` (Number)
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number) - maximum value is equal to the number of rules upon completion of the request
* `rule_id` (Number)
<a id="nestedatt--maps--flow_sample_overlap_rules--pass_rules--gtp"></a>
### Nested Schema for `maps.flow_sample_overlap_rules.pass_rules.gtp`

Read-Only:

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
<a id="nestedatt--maps--flow_sample_rules"></a>
### Nested Schema for `maps.flow_sample_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--maps--flow_sample_rules--pass_rules))
<a id="nestedatt--maps--flow_sample_rules--pass_rules"></a>
### Nested Schema for `maps.flow_sample_rules.pass_rules`

Read-Only:

* `comment` (String)
* `gtp` (Attributes) - Map Flow Sample Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--maps--flow_sample_rules--pass_rules--gtp))
* `percentage` (Number)
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number) - maximum value is equal to the number of rules upon completion of the request
* `rule_id` (Number)
<a id="nestedatt--maps--flow_sample_rules--pass_rules--gtp"></a>
### Nested Schema for `maps.flow_sample_rules.pass_rules.gtp`

Read-Only:

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
<a id="nestedatt--maps--flow_sample_sip_rules"></a>
### Nested Schema for `maps.flow_sample_sip_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--maps--flow_sample_sip_rules--pass_rules))
<a id="nestedatt--maps--flow_sample_sip_rules--pass_rules"></a>
### Nested Schema for `maps.flow_sample_sip_rules.pass_rules`

Read-Only:

* `percentage` (Number)
* `rule_id` (Number)
* `sip` (Attributes) (see [below for nested schema](#nestedatt--maps--flow_sample_sip_rules--pass_rules--sip))
<a id="nestedatt--maps--flow_sample_sip_rules--pass_rules--sip"></a>
### Nested Schema for `maps.flow_sample_sip_rules.pass_rules.sip`

Read-Only:

* `callee_id` (String) - sip callee id
* `callee_id_range` (Attributes) (see [below for nested schema](#nestedatt--maps--flow_sample_sip_rules--pass_rules--sip--callee_id_range))
* `caller_id` (String) - sip caller id
* `caller_id_range` (Attributes) (see [below for nested schema](#nestedatt--maps--flow_sample_sip_rules--pass_rules--sip--caller_id_range))
* `id_range` (Attributes) (see [below for nested schema](#nestedatt--maps--flow_sample_sip_rules--pass_rules--sip--id_range))
<a id="nestedatt--maps--flow_sample_sip_rules--pass_rules--sip--callee_id_range"></a>
### Nested Schema for `maps.flow_sample_sip_rules.pass_rules.sip.callee_id_range`

Read-Only:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--maps--flow_sample_sip_rules--pass_rules--sip--caller_id_range"></a>
### Nested Schema for `maps.flow_sample_sip_rules.pass_rules.sip.caller_id_range`

Read-Only:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--maps--flow_sample_sip_rules--pass_rules--sip--id_range"></a>
### Nested Schema for `maps.flow_sample_sip_rules.pass_rules.sip.id_range`

Read-Only:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--maps--flow_whitelist5_g_overlap_rules"></a>
### Nested Schema for `maps.flow_whitelist5_g_overlap_rules`

Read-Only:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
<a id="nestedatt--maps--flow_whitelist5_g_rules"></a>
### Nested Schema for `maps.flow_whitelist5_g_rules`

Read-Only:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--maps--flow_whitelist_overlap_rules"></a>
### Nested Schema for `maps.flow_whitelist_overlap_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--maps--flow_whitelist_overlap_rules--pass_rules))
<a id="nestedatt--maps--flow_whitelist_overlap_rules--pass_rules"></a>
### Nested Schema for `maps.flow_whitelist_overlap_rules.pass_rules`

Read-Only:

* `flow5_g` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--maps--flow_whitelist_overlap_rules--pass_rules--flow5_g))
* `gtp` (Attributes) - Map Flow Whitelist Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--maps--flow_whitelist_overlap_rules--pass_rules--gtp))
* `rule_id` (Number)
* `sip` (Attributes) - Map Flow Whitelist Rule Sip match definition (see [below for nested schema](#nestedatt--maps--flow_whitelist_overlap_rules--pass_rules--sip))
<a id="nestedatt--maps--flow_whitelist_overlap_rules--pass_rules--flow5_g"></a>
### Nested Schema for `maps.flow_whitelist_overlap_rules.pass_rules.flow5_g`

Read-Only:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--maps--flow_whitelist_overlap_rules--pass_rules--gtp"></a>
### Nested Schema for `maps.flow_whitelist_overlap_rules.pass_rules.gtp`

Read-Only:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--maps--flow_whitelist_overlap_rules--pass_rules--sip"></a>
### Nested Schema for `maps.flow_whitelist_overlap_rules.pass_rules.sip`

Read-Only:

* `type` (String) - all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address
<a id="nestedatt--maps--flow_whitelist_rules"></a>
### Nested Schema for `maps.flow_whitelist_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--maps--flow_whitelist_rules--pass_rules))
<a id="nestedatt--maps--flow_whitelist_rules--pass_rules"></a>
### Nested Schema for `maps.flow_whitelist_rules.pass_rules`

Read-Only:

* `flow5_g` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--maps--flow_whitelist_rules--pass_rules--flow5_g))
* `gtp` (Attributes) - Map Flow Whitelist Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--maps--flow_whitelist_rules--pass_rules--gtp))
* `rule_id` (Number)
* `sip` (Attributes) - Map Flow Whitelist Rule Sip match definition (see [below for nested schema](#nestedatt--maps--flow_whitelist_rules--pass_rules--sip))
<a id="nestedatt--maps--flow_whitelist_rules--pass_rules--flow5_g"></a>
### Nested Schema for `maps.flow_whitelist_rules.pass_rules.flow5_g`

Read-Only:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--maps--flow_whitelist_rules--pass_rules--gtp"></a>
### Nested Schema for `maps.flow_whitelist_rules.pass_rules.gtp`

Read-Only:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--maps--flow_whitelist_rules--pass_rules--sip"></a>
### Nested Schema for `maps.flow_whitelist_rules.pass_rules.sip`

Read-Only:

* `type` (String) - all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address
<a id="nestedatt--maps--fstype"></a>
### Nested Schema for `maps.fstype`

Read-Only:

* `offset` (Number) - Offset for Mobility Rotational Flowsampling
* `timer` (Number) - Timer for Mobility Rotational Flowsampling in minutes
* `type` (String)
<a id="nestedatt--maps--gs_rules"></a>
### Nested Schema for `maps.gs_rules`

Read-Only:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--maps--gs_rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--maps--gs_rules--pass_rules))
<a id="nestedatt--maps--gs_rules--drop_rules"></a>
### Nested Schema for `maps.gs_rules.drop_rules`

Read-Only:

* `comment` (String)
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number)
<a id="nestedatt--maps--gs_rules--pass_rules"></a>
### Nested Schema for `maps.gs_rules.pass_rules`

Read-Only:

* `comment` (String)
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number)
<a id="nestedatt--maps--health_state_reasons"></a>
### Nested Schema for `maps.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--maps--ip_rewrite"></a>
### Nested Schema for `maps.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--maps--rewrite"></a>
### Nested Schema for `maps.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--maps--roles"></a>
### Nested Schema for `maps.roles`

Read-Only:

* `editors` (List of String)
* `listeners` (List of String)
* `owners` (List of String)
* `viewers` (List of String)
<a id="nestedatt--maps--rules"></a>
### Nested Schema for `maps.rules`

Read-Only:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--maps--rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--maps--rules--pass_rules))
<a id="nestedatt--maps--rules--drop_rules"></a>
### Nested Schema for `maps.rules.drop_rules`

Read-Only:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--maps--rules--drop_rules--ip_rewrite))
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--maps--rules--drop_rules--rewrite))
* `rule_id` (Number)
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--maps--rules--drop_rules--vlan_tag))
<a id="nestedatt--maps--rules--drop_rules--ip_rewrite"></a>
### Nested Schema for `maps.rules.drop_rules.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--maps--rules--drop_rules--rewrite"></a>
### Nested Schema for `maps.rules.drop_rules.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--maps--rules--drop_rules--vlan_tag"></a>
### Nested Schema for `maps.rules.drop_rules.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)
<a id="nestedatt--maps--rules--pass_rules"></a>
### Nested Schema for `maps.rules.pass_rules`

Read-Only:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--maps--rules--pass_rules--ip_rewrite))
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--maps--rules--pass_rules--rewrite))
* `rule_id` (Number)
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--maps--rules--pass_rules--vlan_tag))
<a id="nestedatt--maps--rules--pass_rules--ip_rewrite"></a>
### Nested Schema for `maps.rules.pass_rules.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--maps--rules--pass_rules--rewrite"></a>
### Nested Schema for `maps.rules.pass_rules.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--maps--rules--pass_rules--vlan_tag"></a>
### Nested Schema for `maps.rules.pass_rules.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)
<a id="nestedatt--maps--vlan_tag"></a>
### Nested Schema for `maps.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)

