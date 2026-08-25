---
page_title: "gigavuecore_get_all_gta_profiles Data Source - gigavuecore"
subcategory: ""
description: |-
  get all gta profiles
---

# gigavuecore_get_all_gta_profiles Data Source

get all gta profiles

## Example Usage

```terraform
data "gigavuecore_get_all_gta_profiles" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - alias of the GTA profile
* `control_node` (String)
* `core_network_nodes` (Attributes) - core network nodes specified in one of 3 ways: list of addresses, range of addresses, or an address subnet (see [below for nested schema](#nestedatt--items--core_network_nodes))
* `dst_port` (Number)
* `src_port` (Number)
* `user_node` (String)
<a id="nestedatt--items--core_network_nodes"></a>
### Nested Schema for `items.core_network_nodes`

Read-Only:

* `address_list` (Attributes) (see [below for nested schema](#nestedatt--items--core_network_nodes--address_list))
* `address_range` (Attributes) (see [below for nested schema](#nestedatt--items--core_network_nodes--address_range))
* `address_subnet` (Attributes) (see [below for nested schema](#nestedatt--items--core_network_nodes--address_subnet))
<a id="nestedatt--items--core_network_nodes--address_list"></a>
### Nested Schema for `items.core_network_nodes.address_list`

Read-Only:

* `value` (List of String) - ipv4 or ipv6 addresses
<a id="nestedatt--items--core_network_nodes--address_range"></a>
### Nested Schema for `items.core_network_nodes.address_range`

Read-Only:

* `ip_ranges` (Attributes List) (see [below for nested schema](#nestedatt--items--core_network_nodes--address_range--ip_ranges))
<a id="nestedatt--items--core_network_nodes--address_range--ip_ranges"></a>
### Nested Schema for `items.core_network_nodes.address_range.ip_ranges`

Read-Only:

* `max_value` (String) - ipv4 or ipv6 maximum value
* `value` (String) - ipv4 or ipv6 minimum value
<a id="nestedatt--items--core_network_nodes--address_subnet"></a>
### Nested Schema for `items.core_network_nodes.address_subnet`

Read-Only:

* `values` (List of String) - array of ipv4 or ipv6 subnet address space. Mutually exclusive with value

