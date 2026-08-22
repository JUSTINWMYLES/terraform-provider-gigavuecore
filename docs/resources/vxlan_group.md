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
  alias = null
  box_id = null
  comment = null
  vxlan_ids = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `box_id` (String, optional) - device box id. valid range 1 - 64.
* `comment` (String, optional)
* `vxlan_ids` (List(Number), optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `circuit_tunnel_vxlan_groups` (List(Object({alias, box_id, comment, vxlan_ids})), computed)
* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `id` (String, computed)

