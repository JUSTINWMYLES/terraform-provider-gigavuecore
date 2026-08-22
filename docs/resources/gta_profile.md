---
page_title: "gigavuecore_gta_profile Resource - gigavuecore"
subcategory: ""
description: |-
  get gta profile
---

# gigavuecore_gta_profile Resource

get gta profile

## Example Usage

```terraform
resource "gigavuecore_gta_profile" "example" {
  alias = null
  control_node = null
  core_network_nodes = {}
  dst_port = null
  src_port = null
  user_node = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional) - alias of the GTA profile
* `control_node` (String, optional)
* `core_network_nodes` (Object({address_list, address_range, address_subnet}), optional) - core network nodes specified in one of 3 ways: list of addresses, range of addresses, or an address subnet
  * `address_list` (Object({value}), optional)
    * `value` (List(String), required) - ipv4 or ipv6 addresses
  * `address_range` (Object({ip_ranges}), optional)
    * `ip_ranges` (List(Object({max_value, value})), required)
  * `address_subnet` (Object({values}), optional)
    * `values` (List(String), required) - array of ipv4 or ipv6 subnet address space. Mutually exclusive with value
* `dst_port` (Number, optional)
* `src_port` (Number, optional)
* `user_node` (String, optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `alias` (String, computed) - alias of the GTA profile
* `control_node` (String, computed)
* `core_network_nodes` (Object({address_list, address_range, address_subnet}), computed) - core network nodes specified in one of 3 ways: list of addresses, range of addresses, or an address subnet
  * `address_list` (Object({value}), optional)
    * `value` (List(String), required) - ipv4 or ipv6 addresses
  * `address_range` (Object({ip_ranges}), optional)
    * `ip_ranges` (List(Object({max_value, value})), required)
  * `address_subnet` (Object({values}), optional)
    * `values` (List(String), required) - array of ipv4 or ipv6 subnet address space. Mutually exclusive with value
* `dst_port` (Number, computed)
* `src_port` (Number, computed)
* `user_node` (String, computed)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_gta_profile.example {alias}
```
