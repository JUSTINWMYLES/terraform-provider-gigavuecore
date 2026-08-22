---
page_title: "gigavuecore_load_management_interface_neighbors Data Source - gigavuecore"
subcategory: ""
description: |-
  since FM 5.8
---

# gigavuecore_load_management_interface_neighbors Data Source

since FM 5.8

## Example Usage

```terraform
data "gigavuecore_load_management_interface_neighbors" "example" {
  box_id = null
  cluster_id = null
  interface_name = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (Number, required) - boxId
* `cluster_id` (String, required) - Target Cluster ID
* `interface_name` (String, required) - Management Interface Name

### Attributes

In addition to all arguments above, the following attributes are exported:

* `cdp_neighbors` (List(Object({device_id, iface_addr, last_update, mgmt_addr, net_prefix_addr, net_prefix_mask, platform, port_id, port_vlan_id, sw_version, sys_cap_available, ttl})), computed)
* `lldp_neighbors` (List(Object({chassis_id, last_update, link_agg_port_id, link_agg_status, mgmt_addr, mgmt_vlan_id, mtu, port_descr, port_id, port_vlan_id, sys_cap_available, sys_cap_enabled, sys_descr, sys_name, ttl, vlan_name})), computed)

