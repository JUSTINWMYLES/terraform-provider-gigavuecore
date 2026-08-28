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
  alias    = null
  profiles = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the SFFP Profile
* `profiles` (Attributes List, required) (see [below for nested schema](#nestedatt--profiles))

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

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_sffp_profile.example {alias}
```
