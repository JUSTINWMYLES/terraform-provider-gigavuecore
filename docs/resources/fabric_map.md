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
  afm_map = {}
  child_map_aliases = []
  config_status = null
  decap_aliases = []
  encap_aliases = []
  error_message = null
  parent_map_aliases = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `afm_map` (Object({alias, ap_rules, cluster_id, comment, dst_ports, egress_gigastream, enable, encap_tunnel, flex_inline, flex_inline_failover, flex_inline_vlan_id, flow_rules, flow_sample5_g_overlap_rules, flow_sample5_g_rules, flow_sample_diameter_rules, flow_sample_overlap_rules, flow_sample_rules, flow_sample_sip_rules, flow_whitelist5_g_overlap_rules, flow_whitelist5_g_rules, flow_whitelist_overlap_rules, flow_whitelist_rules, fstype, gs_rules, gsop, health_state, health_state_reasons, inline_traffic_path, inline_traffic_type, ip_rewrite, mod_time, null_dst_port, order, rewrite, roles, rule_matching, rules, rx_cluster_ports, src_ports, sub_type, traffic_type, tx_cluster_ports, type, updated_time, vlan_tag}), required)
  * `alias` (String, required) - unique map alias
  * `ap_rules` (Object({drop_rules, pass_rules}), optional) - pass and drop application profile rules
    * `drop_rules` (List(Object({application_profile, rule_id})), optional)
    * `pass_rules` (List(Object({application_profile, rule_id})), optional)
  * `cluster_id` (String, optional) - id of the defining cluster
  * `comment` (String, optional)
  * `dst_ports` (List(String), required) - List of the 'to' ports. Invalid if 'nullDstPort' is true'. Only port number is supported
  * `egress_gigastream` (List(String), optional) - Applicable to fabric map only. It is to specify gigastream(s) to use as egress when going across cluster. The format of array member is clusterID:gigastreamAlias
  * `enable` (Bool, optional) - enable/disable map, applicable only to first level maps
  * `encap_tunnel` (String, optional) - tunnel alias
  * `flex_inline` (Object({a_to_b, b_to_a, oob_copy, svt_mode, svt_tag, tag}), optional) - When 'oobCopy' is configured at least one of 'aToB' or 'bToA' must be configured
    * `a_to_b` (Object({ib_pathway, tools, type}), optional)
      * `ib_pathway` (String, optional) - ibPathway alias. Only applicable when type is 'ibPathway'
      * `tools` (List(String), optional) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'
      * `type` (String, required) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided
    * `b_to_a` (Object({ib_pathway, tools, type}), optional)
      * `ib_pathway` (String, optional) - ibPathway alias. Only applicable when type is 'ibPathway'
      * `tools` (List(String), optional) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'
      * `type` (String, required) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided
    * `oob_copy` (List(Object({direction, dst_ports, src_ports, tag})), optional)
    * `svt_mode` (Bool, optional)
    * `svt_tag` (Number, optional) - only applicable when svtMode is enabled
    * `tag` (Object({tag_protocol_id, type, vlan_id}), optional)
      * `tag_protocol_id` (String, optional) - When tool VLAN tag is added , this protocol Id will be added which egress out the traffic
      * `type` (String, required)
      * `vlan_id` (Number, optional) - only applicable when type is 'vlan'
  * `flex_inline_failover` (String, optional) - only valid for flexInline maps
  * `flex_inline_vlan_id` (Number, optional) - VLAN ID carried in the VLAN tag of packets coming from the inline-network port(s). valid and applicable only with 'flexInline' map type, 'srcPorts' should be either inline-network or ib-pathway or vport
  * `flow_rules` (Object({drop_rules, pass_rules}), optional) - Map Flow Rules Container. Private class
    * `drop_rules` (Set(Object({gtp, rule_id})), optional)
    * `pass_rules` (Set(Object({gtp, rule_id})), optional)
  * `flow_sample5_g_overlap_rules` (Object({pass_rules}), optional) - Map Flow Sample 5g Overlap Rules Container. Private class
    * `pass_rules` (Set(Object({comment, flow5_g, percentage, rule_id})), optional)
  * `flow_sample5_g_rules` (Object({pass_rules}), optional) - Map Flow Sample 5g Rules Container. Private class
    * `pass_rules` (Set(Object({comment, flow5_g, percentage, priority, rule_id})), optional)
  * `flow_sample_diameter_rules` (Object({pass_rules}), optional) - Map Flow Sample Diameter Rules Container. Private class
    * `pass_rules` (Set(Object({diameter, interface, percentage, rule_id})), optional)
  * `flow_sample_overlap_rules` (Object({pass_rules}), optional) - Map Flow Sample Overlap Rules Container. Private class
    * `pass_rules` (Set(Object({comment, gtp, percentage, periodic_recalc, priority, rule_id})), optional)
  * `flow_sample_rules` (Object({pass_rules}), optional) - Map Flow Sample Rules Container. Private class
    * `pass_rules` (Set(Object({comment, gtp, percentage, periodic_recalc, priority, rule_id})), optional)
  * `flow_sample_sip_rules` (Object({pass_rules}), optional) - Map Flow Sample Sip Rules Container. Private class
    * `pass_rules` (Set(Object({percentage, rule_id, sip})), optional)
  * `flow_whitelist5_g_overlap_rules` (Object({dnn, type}), optional) - Map Flow Whitelist 5g Overlap Rule match Definition. Private class
    * `dnn` (String, optional) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
    * `type` (String, optional) - Set 5G WL-DB lookup type
  * `flow_whitelist5_g_rules` (Object({dnn, type, whitelist_databases}), optional) - Map Flow Whitelist 5g Rule GTP match Definition. Private class
    * `dnn` (String, optional) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
    * `type` (String, optional) - Set 5G WL-DB lookup type
    * `whitelist_databases` (List(String), optional) - Attach whitelist databases to the map
  * `flow_whitelist_overlap_rules` (Object({pass_rules}), optional) - Map Flow Whitelist Overlap Rules Container. Private class
    * `pass_rules` (Set(Object({flow5_g, gtp, rule_id, sip})), optional)
  * `flow_whitelist_rules` (Object({pass_rules}), optional) - Map Flow Whitelist Rules Container. Private class
    * `pass_rules` (Set(Object({flow5_g, gtp, rule_id, sip})), optional)
  * `fstype` (Object({offset, timer, type}), optional) - Type of Mobility Flowsampling & properties
    * `offset` (Number, optional) - Offset for Mobility Rotational Flowsampling
    * `timer` (Number, optional) - Timer for Mobility Rotational Flowsampling in minutes
    * `type` (String, optional)
  * `gs_rules` (Object({drop_rules, pass_rules}), optional) - Map GigaSMART Rules Container. Private class
    * `drop_rules` (Set(Object({comment, matches, rule_id})), optional)
    * `pass_rules` (Set(Object({comment, matches, rule_id})), optional)
  * `gsop` (String, optional) - Alias of referenced GSOP. Applicable to 'regular' and 'secondLevel' map types
  * `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
  * `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
  * `inline_traffic_path` (String, optional) - Only applicable for 'inline' map types, in which case defaults to 'normal'. When set to 'bypass', the 'dstPorts' must be empty
  * `inline_traffic_type` (String, optional) - Only applicable for 'inline/passAll' map types, in which case defaults to 'symmetric'. For 'asymmetric' maps, 'srcPort' must be of type 'inline-net'
  * `ip_rewrite` (Object({dst_ip, src_ip}), optional) - IpRewrite options on the packets
    * `dst_ip` (String, optional)
    * `src_ip` (String, optional)
  * `mod_time` (Number, optional) - Last modification time of the map in milliseconds since the epoch
  * `null_dst_port` (Bool, optional) - enabled when the dstPort is null
  * `order` (Number, optional) - relative order within per-source port map chain
  * `rewrite` (Object({dst_mac, src_mac}), optional) - Rewrite options on the packets
    * `dst_mac` (String, optional)
    * `src_mac` (String, optional)
  * `roles` (Object({editors, listeners, owners, viewers}), optional) - Map Roles Container. Private class
    * `editors` (List(String), optional)
    * `listeners` (List(String), optional)
    * `owners` (List(String), optional)
    * `viewers` (List(String), optional)
  * `rule_matching` (String, optional) - If set to 'blacklist', packet are passed when no 'drop' rules are matched. This field is only valid for 'regular/byRule' map types. (maps into 'no-rule-match' CLI command)
  * `rules` (Object({drop_rules, pass_rules}), optional) - Map Rules Container. Private class
    * `drop_rules` (Set(Object({bidi, comment, ip_rewrite, matches, rewrite, rule_id, vlan_tag})), optional)
    * `pass_rules` (Set(Object({bidi, comment, ip_rewrite, matches, rewrite, rule_id, vlan_tag})), optional)
  * `rx_cluster_ports` (List(String), optional)
  * `src_ports` (List(String), required) - list of the 'from' ports. Only port number is supported
  * `sub_type` (String, optional) - 'byRule' is applicable to all map types; 'collector' is applicable to 'regular', 'inline', 'flexInline' and 'secondLevel' maps; 'passAll' is applicable to 'regular' and 'inline' maps; 'flowFilter', 'flowSample', 'flowWhitelist', 'flowSampleSip', 'flowWhitelistSip', 'flowSampleDiameter', 'flowWhitelistDiameter', 'flowSampleOverlap', 'flowWhitelistOverlap', 'flowSample5g','flowWhitelist5g','flowSample5gOverlap' and 'flowWhitelist5gOverlap' are applicable to 'secondLevel' maps
  * `traffic_type` (String, optional) - Only applicable for 'firstLevel/byRule' map types, in which case defaults to 'user'
  * `tx_cluster_ports` (List(String), optional)
  * `type` (String, required) - 'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'flexInline' is from inline network to tools; 'transitLevel' is from vport to vport
  * `updated_time` (Number, computed) - Last Updated time of the fabric map
  * `vlan_tag` (Object({tag_protocol_id, vlan_action, vlan_id}), optional)
    * `tag_protocol_id` (String, optional)
    * `vlan_action` (String, required)
    * `vlan_id` (Number, optional)
* `child_map_aliases` (List(String), optional) - Aliases of the fabric maps that support this fabric map. Applicable only to user-defined fabric maps. For internally generated fabric maps, there is no child maps. This list is empty.
* `config_status` (String, optional) - Configuration status of this fabric map.
* `decap_aliases` (List(String), optional) - Aliases of the decap tunnel endpoints supporting this fabric map. Applicable only to internally generated fabric maps.
* `encap_aliases` (List(String), optional) - Aliases of the encap tunnel endpoints supporting this fabric map. Applicable only to internally generated fabric maps.
* `error_message` (String, optional) - In case of configuration failure, this message provides details about the possible cause of the failure.
* `parent_map_aliases` (List(String), optional) - Aliases of the fabric maps that this fabric map supports. Applicable only to internally generated fabric maps that support user-defined fabric maps. For user-defined fabric maps, there is no parent maps. This list is empty.

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `alias` (String, computed)
* `child_map_aliases` (List(String), computed) - Aliases of the fabric maps that support this fabric map. Applicable only to user-defined fabric maps. For internally generated fabric maps, there is no child maps. This list is empty.
* `config_status` (String, computed) - Configuration status of this fabric map.
* `creation_time` (Number, computed) - Fabric map creation timestamp in UTC
* `decap_aliases` (List(String), computed) - Aliases of the decap tunnel endpoints supporting this fabric map. Applicable only to internally generated fabric maps.
* `encap_aliases` (List(String), computed) - Aliases of the encap tunnel endpoints supporting this fabric map. Applicable only to internally generated fabric maps.
* `error_message` (String, computed) - In case of configuration failure, this message provides details about the possible cause of the failure.
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `parent_map_aliases` (List(String), computed) - Aliases of the fabric maps that this fabric map supports. Applicable only to internally generated fabric maps that support user-defined fabric maps. For user-defined fabric maps, there is no parent maps. This list is empty.
* `pending_reason` (String, computed)
* `updated_time` (Number, computed) - Fabric map updated timestamp in UTC

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_fabric_map.example {alias}
```
