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
  box_id         = null
  cluster_id     = null
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

* `cdp_neighbors` (Attributes List, computed) (see [below for nested schema](#nestedatt--cdp_neighbors))
* `lldp_neighbors` (Attributes List, computed) (see [below for nested schema](#nestedatt--lldp_neighbors))

<a id="nestedatt--cdp_neighbors"></a>
### Nested Schema for `cdp_neighbors`

Read-Only:

* `device_id` (String)
* `iface_addr` (String) - ipv4
* `last_update` (Number) - Timestamp of the last update. In UTC milliseconds
* `mgmt_addr` (String) - ipv4
* `net_prefix_addr` (String) - ipv4
* `net_prefix_mask` (String) - ipv4
* `platform` (String)
* `port_id` (String)
* `port_vlan_id` (Number)
* `sw_version` (String)
* `sys_cap_available` (Number)
* `ttl` (Number)
<a id="nestedatt--lldp_neighbors"></a>
### Nested Schema for `lldp_neighbors`

Read-Only:

* `chassis_id` (String)
* `last_update` (Number) - Timestamp of the last update. In UTC milliseconds
* `link_agg_port_id` (Number)
* `link_agg_status` (Number)
* `mgmt_addr` (String) - ipv4 or ipv6
* `mgmt_vlan_id` (Number)
* `mtu` (Number)
* `port_descr` (String)
* `port_id` (String)
* `port_vlan_id` (Number)
* `sys_cap_available` (Number)
* `sys_cap_enabled` (Number)
* `sys_descr` (String)
* `sys_name` (String)
* `ttl` (Number)
* `vlan_name` (String)

