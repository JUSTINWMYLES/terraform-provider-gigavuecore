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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `admin_status` (Boolean)
* `autoconf_enabled` (Boolean)
* `autoconf_privacy` (Boolean)
* `autoconf_route` (Boolean)
* `cluster_id` (String) - id of the defining cluster
* `comment` (String)
* `dhcp_enabled` (Boolean)
* `dhcpv6_enabled` (Boolean)
* `duplex` (String)
* `hw_address` (String)
* `if_index` (Number)
* `if_source` (String)
* `ip_address` (String) - ipv4
* `ipv6_addresses` (List of String)
* `ipv6_enabled` (Boolean)
* `link_status` (Boolean)
* `mtu` (Number)
* `name` (String) - Interface name
* `netmask` (String)
* `speed` (String)
* `type` (String)

