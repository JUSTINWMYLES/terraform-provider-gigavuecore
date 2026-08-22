---
page_title: "gigavuecore_load_all_system_interfaces Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all System Interfaces
---

# gigavuecore_load_all_system_interfaces Data Source

Load all System Interfaces

## Example Usage

```terraform
data "gigavuecore_load_all_system_interfaces" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({admin_status, autoconf_enabled, autoconf_privacy, autoconf_route, cluster_id, comment, dhcp_enabled, dhcpv6_enabled, duplex, hw_address, if_index, if_source, ip_address, ipv6_addresses, ipv6_enabled, link_status, mtu, name, netmask, speed, type})), computed)

