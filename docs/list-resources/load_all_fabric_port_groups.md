---
page_title: "gigavuecore_load_all_fabric_port_groups List Resource - gigavuecore"
subcategory: ""
description: |-
  Get all the Fabric Port Groups
---

# gigavuecore_load_all_fabric_port_groups List Resource

Get all the Fabric Port Groups

## Example Usage

```terraform
list "gigavuecore_load_all_fabric_port_groups" "example" {
  provider = gigavuecore
  limit    = 100
}

```
## Schema

### Identity Attributes

The following identity attributes are exported for each matching result:

* `alias` (String, computed) - Alias of the Fabric Port Group


