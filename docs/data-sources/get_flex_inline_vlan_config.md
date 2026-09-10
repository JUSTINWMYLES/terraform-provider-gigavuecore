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
  cluster_id           = "example"
  inline_network_alias = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Gets vlan configs for the provided cluster
* `inline_network_alias` (String, optional) - if provided, returns vlan configs only for that inline network construct

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `inline_network_alias` (String) - Alias of the inline network
* `map_alias` (String) - Alias of the map
* `vlan_configs` (Attributes List) - Array holding vlan ids (see [below for nested schema](#nestedatt--items--vlan_configs))

<a id="nestedatt--items--vlan_configs"></a>
### Nested Schema for `items.vlan_configs`

Read-Only:

* `flex_inline_tag_type` (String) - type of vlan id, auto (or) vlan
* `flex_inline_tag_vlan_id` (Number) - Tool side vlan id
* `flex_inline_vlan_id` (Number) - Network side vlan id
* `network_alias` (String) - Alias of the network source
* `tag_protocol_id` (String) - When tool VLAN tag is added , this protocol Id will be added which egress out the traffic

