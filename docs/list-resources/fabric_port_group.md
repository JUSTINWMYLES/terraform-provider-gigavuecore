---
page_title: "gigavuecore_fabric_port_group List Resource - gigavuecore"
subcategory: ""
description: |-
  Get all the Fabric Port Groups
---

# gigavuecore_fabric_port_group List Resource

Get all the Fabric Port Groups

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

## Example Usage

```terraform
list "gigavuecore_fabric_port_group" "example" {
  provider = gigavuecore
  limit    = 100
}
```
## Schema

### Identity Attributes

The following identity attributes are exported for each matching result:

* `alias` (String, computed) - Alias of the Fabric Port Group


