---
page_title: "gigavuecore_redundancy_profile Resource - gigavuecore"
subcategory: ""
description: |-
  Find Redundancy Profile by alias
---

# gigavuecore_redundancy_profile Resource

Find Redundancy Profile by alias

## Example Usage

```terraform
resource "gigavuecore_redundancy_profile" "example" {
  alias = null
  protection_role = null
  signaling_port = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Inline Network alias. Unique within a cluster
* `protection_role` (String, optional) - only applicable for 'protected' inline networks
* `signaling_port` (String, required) - only applicable for 'protected' inline networks

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `protection_role` (String, computed) - only applicable for 'protected' inline networks

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_redundancy_profile.example {alias}
```
