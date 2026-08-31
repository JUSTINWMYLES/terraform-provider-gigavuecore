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
  alias        = "example"
  control_node = "example"
  core_network_nodes = {
    address_list = {
      value = [ "example" ]
    }
    address_range = {
      ip_ranges = [{
        max_value = "example"
        value     = "example"
      }]
    }
    address_subnet = {
      values = [ "example" ]
    }
  }
  dst_port  = 0
  src_port  = 0
  user_node = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional) - alias of the GTA profile
* `control_node` (String, optional)
* `core_network_nodes` (Attributes, optional) - core network nodes specified in one of 3 ways: list of addresses, range of addresses, or an address subnet (see [below for nested schema](#nestedatt--core_network_nodes))
* `dst_port` (Number, optional)
* `src_port` (Number, optional)
* `user_node` (String, optional)

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--core_network_nodes"></a>
### Nested Schema for `core_network_nodes`

Optional:

* `address_list` (Attributes) (see [below for nested schema](#nestedatt--core_network_nodes--address_list))
* `address_range` (Attributes) (see [below for nested schema](#nestedatt--core_network_nodes--address_range))
* `address_subnet` (Attributes) (see [below for nested schema](#nestedatt--core_network_nodes--address_subnet))

<a id="nestedatt--core_network_nodes--address_list"></a>
### Nested Schema for `core_network_nodes.address_list`

Required:

* `value` (List of String) - ipv4 or ipv6 addresses

<a id="nestedatt--core_network_nodes--address_range"></a>
### Nested Schema for `core_network_nodes.address_range`

Required:

* `ip_ranges` (Attributes List) (see [below for nested schema](#nestedatt--core_network_nodes--address_range--ip_ranges))

<a id="nestedatt--core_network_nodes--address_range--ip_ranges"></a>
### Nested Schema for `core_network_nodes.address_range.ip_ranges`

Required:

* `max_value` (String) - ipv4 or ipv6 maximum value
* `value` (String) - ipv4 or ipv6 minimum value

<a id="nestedatt--core_network_nodes--address_subnet"></a>
### Nested Schema for `core_network_nodes.address_subnet`

Required:

* `values` (List of String) - array of ipv4 or ipv6 subnet address space. Mutually exclusive with value
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_gta_profile.example {alias}
```
