---
page_title: "gigavuecore_get_flex_inline_vlan_config Data Source - gigavuecore"
subcategory: ""
description: |-
  Gets all vlan configs for the given cluster
---

# gigavuecore_get_flex_inline_vlan_config Data Source

Gets all vlan configs for the given cluster

## Example Usage

```terraform
data "gigavuecore_get_flex_inline_vlan_config" "example" {
  cluster_id = null
  inline_network_alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Gets vlan configs for the provided cluster
* `inline_network_alias` (String, optional) - if provided, returns vlan configs only for that inline network construct

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({inline_network_alias, map_alias, vlan_configs})), computed)

