---
page_title: "gigavuecore_get_cluster_map_of_an_user_fabric_map Data Source - gigavuecore"
subcategory: ""
description: |-
  Get a cluster map of a user-defined fabric map
---

# gigavuecore_get_cluster_map_of_an_user_fabric_map Data Source

Get a cluster map of a user-defined fabric map

## Example Usage

```terraform
data "gigavuecore_get_cluster_map_of_an_user_fabric_map" "example" {
  alias = null
  cm_alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the fabric map
* `cm_alias` (String, required) - alias of the cluster map

### Attributes

In addition to all arguments above, the following attributes are exported:

* `afm_map` (Object({alias, ap_rules, cluster_id, comment, dst_ports, egress_gigastream, enable, encap_tunnel, flex_inline, flex_inline_failover, flex_inline_vlan_id, flow_rules, flow_sample5_g_overlap_rules, flow_sample5_g_rules, flow_sample_diameter_rules, flow_sample_overlap_rules, flow_sample_rules, flow_sample_sip_rules, flow_whitelist5_g_overlap_rules, flow_whitelist5_g_rules, flow_whitelist_overlap_rules, flow_whitelist_rules, fstype, gs_rules, gsop, health_state, health_state_reasons, inline_traffic_path, inline_traffic_type, ip_rewrite, mod_time, null_dst_port, order, rewrite, roles, rule_matching, rules, rx_cluster_ports, src_ports, sub_type, traffic_type, tx_cluster_ports, type, updated_time, vlan_tag}), computed)
  * `alias` (String, computed) - unique map alias
  * `ap_rules` (Object({drop_rules, pass_rules}), computed) - pass and drop application profile rules
    * `drop_rules` (List(Object({application_profile, rule_id})), computed)
    * `pass_rules` (List(Object({application_profile, rule_id})), computed)
  * `cluster_id` (String, computed) - id of the defining cluster
  * `comment` (String, computed)
  * `dst_ports` (List(String), computed) - List of the 'to' ports. Invalid if 'nullDstPort' is true'. Only port number is supported
  * `egress_gigastream` (List(String), computed) - Applicable to fabric map only. It is to specify gigastream(s) to use as egress when going across cluster. The format of array member is clusterID:gigastreamAlias
  * `enable` (Bool, computed) - enable/disable map, applicable only to first level maps
  * `encap_tunnel` (String, computed) - tunnel alias
  * `flex_inline` (Object({a_to_b, b_to_a, oob_copy, svt_mode, svt_tag, tag}), computed) - When 'oobCopy' is configured at least one of 'aToB' or 'bToA' must be configured
    * `a_to_b` (Object({ib_pathway, tools, type}), computed)
      * `ib_pathway` (String, computed) - ibPathway alias. Only applicable when type is 'ibPathway'
      * `tools` (List(String), computed) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'
      * `type` (String, computed) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided
    * `b_to_a` (Object({ib_pathway, tools, type}), computed)
      * `ib_pathway` (String, computed) - ibPathway alias. Only applicable when type is 'ibPathway'
      * `tools` (List(String), computed) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'
      * `type` (String, computed) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided
    * `oob_copy` (List(Object({direction, dst_ports, src_ports, tag})), computed)
    * `svt_mode` (Bool, computed)
    * `svt_tag` (Number, computed) - only applicable when svtMode is enabled
    * `tag` (Object({tag_protocol_id, type, vlan_id}), computed)
      * `tag_protocol_id` (String, computed) - When tool VLAN tag is added , this protocol Id will be added which egress out the traffic
      * `type` (String, computed)
      * `vlan_id` (Number, computed) - only applicable when type is 'vlan'
  * `flex_inline_failover` (String, computed) - only valid for flexInline maps
  * `flex_inline_vlan_id` (Number, computed) - VLAN ID carried in the VLAN tag of packets coming from the inline-network port(s). valid and applicable only with 'flexInline' map type, 'srcPorts' should be either inline-network or ib-pathway or vport
  * `flow_rules` (Object({drop_rules, pass_rules}), computed) - Map Flow Rules Container. Private class
    * `drop_rules` (Set(Object({gtp, rule_id})), computed)
    * `pass_rules` (Set(Object({gtp, rule_id})), computed)
  * `flow_sample5_g_overlap_rules` (Object({pass_rules}), computed) - Map Flow Sample 5g Overlap Rules Container. Private class
    * `pass_rules` (Set(Object({comment, flow5_g, percentage, rule_id})), computed)
  * `flow_sample5_g_rules` (Object({pass_rules}), computed) - Map Flow Sample 5g Rules Container. Private class
    * `pass_rules` (Set(Object({comment, flow5_g, percentage, priority, rule_id})), computed)
  * `flow_sample_diameter_rules` (Object({pass_rules}), computed) - Map Flow Sample Diameter Rules Container. Private class
    * `pass_rules` (Set(Object({diameter, interface, percentage, rule_id})), computed)
  * `flow_sample_overlap_rules` (Object({pass_rules}), computed) - Map Flow Sample Overlap Rules Container. Private class
    * `pass_rules` (Set(Object({comment, gtp, percentage, periodic_recalc, priority, rule_id})), computed)
  * `flow_sample_rules` (Object({pass_rules}), computed) - Map Flow Sample Rules Container. Private class
    * `pass_rules` (Set(Object({comment, gtp, percentage, periodic_recalc, priority, rule_id})), computed)
  * `flow_sample_sip_rules` (Object({pass_rules}), computed) - Map Flow Sample Sip Rules Container. Private class
    * `pass_rules` (Set(Object({percentage, rule_id, sip})), computed)
  * `flow_whitelist5_g_overlap_rules` (Object({dnn, type}), computed) - Map Flow Whitelist 5g Overlap Rule match Definition. Private class
    * `dnn` (String, computed) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
    * `type` (String, computed) - Set 5G WL-DB lookup type
  * `flow_whitelist5_g_rules` (Object({dnn, type, whitelist_databases}), computed) - Map Flow Whitelist 5g Rule GTP match Definition. Private class
    * `dnn` (String, computed) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
    * `type` (String, computed) - Set 5G WL-DB lookup type
    * `whitelist_databases` (List(String), computed) - Attach whitelist databases to the map
  * `flow_whitelist_overlap_rules` (Object({pass_rules}), computed) - Map Flow Whitelist Overlap Rules Container. Private class
    * `pass_rules` (Set(Object({flow5_g, gtp, rule_id, sip})), computed)
  * `flow_whitelist_rules` (Object({pass_rules}), computed) - Map Flow Whitelist Rules Container. Private class
    * `pass_rules` (Set(Object({flow5_g, gtp, rule_id, sip})), computed)
  * `fstype` (Object({offset, timer, type}), computed) - Type of Mobility Flowsampling & properties
    * `offset` (Number, computed) - Offset for Mobility Rotational Flowsampling
    * `timer` (Number, computed) - Timer for Mobility Rotational Flowsampling in minutes
    * `type` (String, computed)
  * `gs_rules` (Object({drop_rules, pass_rules}), computed) - Map GigaSMART Rules Container. Private class
    * `drop_rules` (Set(Object({comment, matches, rule_id})), computed)
    * `pass_rules` (Set(Object({comment, matches, rule_id})), computed)
  * `gsop` (String, computed) - Alias of referenced GSOP. Applicable to 'regular' and 'secondLevel' map types
  * `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
  * `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
  * `inline_traffic_path` (String, computed) - Only applicable for 'inline' map types, in which case defaults to 'normal'. When set to 'bypass', the 'dstPorts' must be empty
  * `inline_traffic_type` (String, computed) - Only applicable for 'inline/passAll' map types, in which case defaults to 'symmetric'. For 'asymmetric' maps, 'srcPort' must be of type 'inline-net'
  * `ip_rewrite` (Object({dst_ip, src_ip}), computed) - IpRewrite options on the packets
    * `dst_ip` (String, computed)
    * `src_ip` (String, computed)
  * `mod_time` (Number, computed) - Last modification time of the map in milliseconds since the epoch
  * `null_dst_port` (Bool, computed) - enabled when the dstPort is null
  * `order` (Number, computed) - relative order within per-source port map chain
  * `rewrite` (Object({dst_mac, src_mac}), computed) - Rewrite options on the packets
    * `dst_mac` (String, computed)
    * `src_mac` (String, computed)
  * `roles` (Object({editors, listeners, owners, viewers}), computed) - Map Roles Container. Private class
    * `editors` (List(String), computed)
    * `listeners` (List(String), computed)
    * `owners` (List(String), computed)
    * `viewers` (List(String), computed)
  * `rule_matching` (String, computed) - If set to 'blacklist', packet are passed when no 'drop' rules are matched. This field is only valid for 'regular/byRule' map types. (maps into 'no-rule-match' CLI command)
  * `rules` (Object({drop_rules, pass_rules}), computed) - Map Rules Container. Private class
    * `drop_rules` (Set(Object({bidi, comment, ip_rewrite, matches, rewrite, rule_id, vlan_tag})), computed)
    * `pass_rules` (Set(Object({bidi, comment, ip_rewrite, matches, rewrite, rule_id, vlan_tag})), computed)
  * `rx_cluster_ports` (List(String), computed)
  * `src_ports` (List(String), computed) - list of the 'from' ports. Only port number is supported
  * `sub_type` (String, computed) - 'byRule' is applicable to all map types; 'collector' is applicable to 'regular', 'inline', 'flexInline' and 'secondLevel' maps; 'passAll' is applicable to 'regular' and 'inline' maps; 'flowFilter', 'flowSample', 'flowWhitelist', 'flowSampleSip', 'flowWhitelistSip', 'flowSampleDiameter', 'flowWhitelistDiameter', 'flowSampleOverlap', 'flowWhitelistOverlap', 'flowSample5g','flowWhitelist5g','flowSample5gOverlap' and 'flowWhitelist5gOverlap' are applicable to 'secondLevel' maps
  * `traffic_type` (String, computed) - Only applicable for 'firstLevel/byRule' map types, in which case defaults to 'user'
  * `tx_cluster_ports` (List(String), computed)
  * `type` (String, computed) - 'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'flexInline' is from inline network to tools; 'transitLevel' is from vport to vport
  * `updated_time` (Number, computed) - Last Updated time of the fabric map
  * `vlan_tag` (Object({tag_protocol_id, vlan_action, vlan_id}), computed)
    * `tag_protocol_id` (String, computed)
    * `vlan_action` (String, computed)
    * `vlan_id` (Number, computed)
* `cluster_id` (String, computed) - id of the cluster in which the map is created.
* `config_status` (String, computed) - Configuration status of this map.
* `error_message` (String, computed) - In case of configuration failure, this message provides details about the possible cause of the failure.
* `fabric_map_aliases` (List(String), computed) - Aliases of the fabric maps that this map supports.

