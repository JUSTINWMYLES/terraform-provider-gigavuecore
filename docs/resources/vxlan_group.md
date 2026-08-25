---
page_title: "gigavuecore_vxlan_group Resource - gigavuecore"
subcategory: ""
description: |-
  Load all Circuit Tunnel Vxlan Groups
---

# gigavuecore_vxlan_group Resource

Load all Circuit Tunnel Vxlan Groups

## Example Usage

```terraform
resource "gigavuecore_vxlan_group" "example" {
  alias     = null
  box_id    = null
  comment   = null
  vxlan_ids = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `box_id` (String, optional) - device box id. valid range 1 - 64.
* `comment` (String, optional)
* `vxlan_ids` (List of Number, optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `circuit_tunnel_vxlan_groups` (Attributes List, computed) (see [below for nested schema](#nestedatt--circuit_tunnel_vxlan_groups))
* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `id` (String, computed)

<a id="nestedatt--circuit_tunnel_vxlan_groups"></a>
### Nested Schema for `circuit_tunnel_vxlan_groups`

Read-Only:

* `alias` (String)
* `box_id` (String) - device box id. valid range 1 - 64.
* `comment` (String)
* `vxlan_ids` (List of Number)
<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type

