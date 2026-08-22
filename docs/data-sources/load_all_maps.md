---
page_title: "gigavuecore_load_all_maps Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all maps
---

# gigavuecore_load_all_maps Data Source

Load all maps

## Example Usage

```terraform
data "gigavuecore_load_all_maps" "example" {
  cluster_id = null
  map_types = []
  page = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `map_types` (List(String), optional) - Comma-separated list of map types that are to be considered for filtering.
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `maps` (List(Object({alias, ap_rules, cluster_id, comment, dst_ports, egress_gigastream, enable, encap_tunnel, flex_inline, flex_inline_failover, flex_inline_vlan_id, flow_rules, flow_sample5_g_overlap_rules, flow_sample5_g_rules, flow_sample_diameter_rules, flow_sample_overlap_rules, flow_sample_rules, flow_sample_sip_rules, flow_whitelist5_g_overlap_rules, flow_whitelist5_g_rules, flow_whitelist_overlap_rules, flow_whitelist_rules, fstype, gs_rules, gsop, health_state, health_state_reasons, inline_traffic_path, inline_traffic_type, ip_rewrite, mod_time, null_dst_port, order, rewrite, roles, rule_matching, rules, rx_cluster_ports, src_ports, sub_type, traffic_type, tx_cluster_ports, type, updated_time, vlan_tag})), computed)

