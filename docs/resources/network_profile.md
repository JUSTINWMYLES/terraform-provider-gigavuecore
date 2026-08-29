---
page_title: "gigavuecore_network_profile Resource - gigavuecore"
subcategory: ""
description: |-
  Load Metadata Network Profile by alias
---

# gigavuecore_network_profile Resource

Load Metadata Network Profile by alias

## Example Usage

```terraform
resource "gigavuecore_network_profile" "example" {
  alias        = "example"
  description  = "example"
  ipv4_subnets = [ "example" ]
  ipv6_masks   = [ "example" ]
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - network profile alias
* `description` (String, optional)
* `ipv4_subnets` (List of String, optional) - ipv4 subnets, to identify client
* `ipv6_masks` (List of String, optional) - ipv6 subnets, to identify client

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `description` (String, computed)
* `ipv4_subnets` (List of String, computed) - ipv4 subnets, to identify client
* `ipv6_masks` (List of String, computed) - ipv6 subnets, to identify client


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_network_profile.example {alias}
```
