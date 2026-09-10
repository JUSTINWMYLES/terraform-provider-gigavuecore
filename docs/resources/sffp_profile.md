---
page_title: "gigavuecore_sffp_profile Resource - gigavuecore"
subcategory: ""
description: |-
  new in H 5.8
---

# gigavuecore_sffp_profile Resource

new in H 5.8

## Example Usage

```terraform
resource "gigavuecore_sffp_profile" "example" {
  alias = "example"
  profiles = [{
    ip_interface = "example"
    node_type    = "control"
    port_list    = [0]
    sx_interface = {
      ip_addresses = ["example"]
    }
  }]
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the SFFP Profile
* `profiles` (Attributes List, required) (see [below for nested schema](#nestedatt--profiles))

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--profiles"></a>
### Nested Schema for `profiles`

Required:

* `ip_interface` (String)
* `node_type` (String)
* `port_list` (List of Number)
* `sx_interface` (Attributes) (see [below for nested schema](#nestedatt--profiles--sx_interface))

<a id="nestedatt--profiles--sx_interface"></a>
### Nested Schema for `profiles.sx_interface`

Optional:

* `ip_addresses` (List of String)
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_sffp_profile.example {alias}
```
