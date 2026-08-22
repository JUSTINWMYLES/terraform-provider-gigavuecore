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
  alias = null
  profiles = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the SFFP Profile
* `profiles` (List(Object({ip_interface, node_type, port_list, sx_interface})), required)

### Attributes

In addition to all arguments above, the following computed attributes are exported:


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_sffp_profile.example {alias}
```
