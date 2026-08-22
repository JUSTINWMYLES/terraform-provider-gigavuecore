---
page_title: "gigavuecore_gpfcp_profile Resource - gigavuecore"
subcategory: ""
description: |-
  new in H 6.8
---

# gigavuecore_gpfcp_profile Resource

new in H 6.8

## Example Usage

```terraform
resource "gigavuecore_gpfcp_profile" "example" {
  alias = null
  comment = null
  g_profiles = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the Gpfcp Profile
* `comment` (String, optional) - Description of the Gpfcp Profile
* `g_profiles` (List(Object({comment, g_interface, ip_interface, node_type, port_list})), required)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `comment` (String, computed) - Description of the Gpfcp Profile

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_gpfcp_profile.example {alias}
```
