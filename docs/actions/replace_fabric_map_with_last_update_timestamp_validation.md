---
page_title: "gigavuecore_replace_fabric_map_with_last_update_timestamp_validation Action - gigavuecore"
subcategory: ""
description: |-
  since FM 6.8.00
---

# gigavuecore_replace_fabric_map_with_last_update_timestamp_validation Action

since FM 6.8.00

## Example Usage

```terraform
action "gigavuecore_replace_fabric_map_with_last_update_timestamp_validation" "example" {
  config {
    alias = "example"
    ap_rules = null
    body_alias = "example"
    cluster_id = "example"
    comment = "example"
    dst_ports = [ "example" ]
    egress_gigastream = [ "example" ]
    enable = true
    encap_tunnel = "example"
    flex_inline = null
    flex_inline_failover = "example"
    flex_inline_vlan_id = 1
    flow_rules = null
    flow_sample5_g_overlap_rules = null
    flow_sample5_g_rules = null
    flow_sample_diameter_rules = null
    flow_sample_overlap_rules = null
    flow_sample_rules = null
    flow_sample_sip_rules = null
    flow_whitelist5_g_overlap_rules = null
    flow_whitelist5_g_rules = null
    flow_whitelist_overlap_rules = null
    flow_whitelist_rules = null
    fstype = null
    gs_rules = null
    gsop = "example"
    health_state = "example"
    health_state_reasons = null
    inline_traffic_path = "example"
    inline_traffic_type = "example"
    ip_rewrite = null
    mod_time = 1
    null_dst_port = true
    order = 1
    rewrite = null
    roles = null
    rule_matching = "example"
    rules = null
    rx_cluster_ports = [ "example" ]
    src_ports = [ "example" ]
    sub_type = "example"
    traffic_type = "example"
    tx_cluster_ports = [ "example" ]
    type = "example"
    updated_time = 1.0
    vlan_tag = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the fabric map to be replaced
* `ap_rules` (Dynamic, optional) - pass and drop application profile rules
* `body_alias` (String, required) - unique map alias
* `cluster_id` (String, optional) - id of the defining cluster
* `comment` (String, optional)
* `dst_ports` (List(String), required) - List of the 'to' ports. Invalid if 'nullDstPort' is true'. Only port number is supported
* `egress_gigastream` (List(String), optional) - Applicable to fabric map only. It is to specify gigastream(s) to use as egress when going across cluster. The format of array member is clusterID:gigastreamAlias
* `enable` (Bool, optional) - enable/disable map, applicable only to first level maps
* `encap_tunnel` (String, optional) - tunnel alias
* `flex_inline` (Dynamic, optional) - When 'oobCopy' is configured at least one of 'aToB' or 'bToA' must be configured
* `flex_inline_failover` (String, optional) - only valid for flexInline maps
* `flex_inline_vlan_id` (Number, optional) - VLAN ID carried in the VLAN tag of packets coming from the inline-network port(s). valid and applicable only with 'flexInline' map type, 'srcPorts' should be either inline-network or ib-pathway or vport
* `flow_rules` (Dynamic, optional) - Map Flow Rules Container. Private class
* `flow_sample5_g_overlap_rules` (Dynamic, optional) - Map Flow Sample 5g Overlap Rules Container. Private class
* `flow_sample5_g_rules` (Dynamic, optional) - Map Flow Sample 5g Rules Container. Private class
* `flow_sample_diameter_rules` (Dynamic, optional) - Map Flow Sample Diameter Rules Container. Private class
* `flow_sample_overlap_rules` (Dynamic, optional) - Map Flow Sample Overlap Rules Container. Private class
* `flow_sample_rules` (Dynamic, optional) - Map Flow Sample Rules Container. Private class
* `flow_sample_sip_rules` (Dynamic, optional) - Map Flow Sample Sip Rules Container. Private class
* `flow_whitelist5_g_overlap_rules` (Dynamic, optional) - Map Flow Whitelist 5g Overlap Rule match Definition. Private class
* `flow_whitelist5_g_rules` (Dynamic, optional) - Map Flow Whitelist 5g Rule GTP match Definition. Private class
* `flow_whitelist_overlap_rules` (Dynamic, optional) - Map Flow Whitelist Overlap Rules Container. Private class
* `flow_whitelist_rules` (Dynamic, optional) - Map Flow Whitelist Rules Container. Private class
* `fstype` (Dynamic, optional) - Type of Mobility Flowsampling & properties
* `gs_rules` (Dynamic, optional) - Map GigaSMART Rules Container. Private class
* `gsop` (String, optional) - Alias of referenced GSOP. Applicable to 'regular' and 'secondLevel' map types
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Dynamic), optional)
* `inline_traffic_path` (String, optional) - Only applicable for 'inline' map types, in which case defaults to 'normal'. When set to 'bypass', the 'dstPorts' must be empty
* `inline_traffic_type` (String, optional) - Only applicable for 'inline/passAll' map types, in which case defaults to 'symmetric'. For 'asymmetric' maps, 'srcPort' must be of type 'inline-net'
* `ip_rewrite` (Dynamic, optional) - IpRewrite options on the packets
* `mod_time` (Number, optional) - Last modification time of the map in milliseconds since the epoch
* `null_dst_port` (Bool, optional) - enabled when the dstPort is null
* `order` (Number, optional) - relative order within per-source port map chain
* `rewrite` (Dynamic, optional) - Rewrite options on the packets
* `roles` (Dynamic, optional) - Map Roles Container. Private class
* `rule_matching` (String, optional) - If set to 'blacklist', packet are passed when no 'drop' rules are matched. This field is only valid for 'regular/byRule' map types. (maps into 'no-rule-match' CLI command)
* `rules` (Dynamic, optional) - Map Rules Container. Private class
* `rx_cluster_ports` (List(String), optional)
* `src_ports` (List(String), required) - list of the 'from' ports. Only port number is supported
* `sub_type` (String, optional) - 'byRule' is applicable to all map types; 'collector' is applicable to 'regular', 'inline', 'flexInline' and 'secondLevel' maps; 'passAll' is applicable to 'regular' and 'inline' maps; 'flowFilter', 'flowSample', 'flowWhitelist', 'flowSampleSip', 'flowWhitelistSip', 'flowSampleDiameter', 'flowWhitelistDiameter', 'flowSampleOverlap', 'flowWhitelistOverlap', 'flowSample5g','flowWhitelist5g','flowSample5gOverlap' and 'flowWhitelist5gOverlap' are applicable to 'secondLevel' maps
* `traffic_type` (String, optional) - Only applicable for 'firstLevel/byRule' map types, in which case defaults to 'user'
* `tx_cluster_ports` (List(String), optional)
* `type` (String, required) - 'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'flexInline' is from inline network to tools; 'transitLevel' is from vport to vport
* `updated_time` (Number, optional) - Last Updated time of the fabric map
* `vlan_tag` (Dynamic, optional)
