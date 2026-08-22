---
page_title: "gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all cluster circuit tunnel endpoints of an internal fabric map
---

# gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map Data Source

Get all cluster circuit tunnel endpoints of an internal fabric map

## Example Usage

```terraform
data "gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map" "example" {
  alias = null
  ifm_alias = null
  mode = null
  type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the fabric map
* `ifm_alias` (String, required) - alias of the internal fabric map
* `mode` (String, optional) - tunnel mode: 'encap' or 'decap'.
* `type` (String, optional) - tunnel type: 'circuit' or 'vxlan'.

### Attributes

In addition to all arguments above, the following attributes are exported:

* `cluster_maps` (List(Object({circuit_tunnel, cluster_id, config_status, error_message, fabric_map_aliases})), computed)
* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type

