---
page_title: "gigavuecore_port_packet_threshold Data Source - gigavuecore"
subcategory: ""
description: |-
  new in FM 5.7
---

# gigavuecore_port_packet_threshold Data Source

new in FM 5.7

## Example Usage

```terraform
data "gigavuecore_port_packet_threshold" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({alarm_count, behind_nat, box_id, cluster_id, cluster_mode, cluster_mode_alias, config_refresh_statistics, config_save_counters, device_id, device_ip, device_ips, disc_outcome, family, global_node_id, health_state, hostname, leader_pref, licensed, master_pref, model, neighbour_links, neighbour_nodes, oper_status, serial_number, stats_collection_counters, suppressed, sw_version, tags, threshold_config, topo_node_id, uboot_version})), computed)

