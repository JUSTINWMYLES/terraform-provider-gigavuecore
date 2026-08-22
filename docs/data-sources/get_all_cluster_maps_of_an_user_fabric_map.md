---
page_title: "gigavuecore_get_all_cluster_maps_of_an_user_fabric_map Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all cluster maps of a user-defined fabric map
---

# gigavuecore_get_all_cluster_maps_of_an_user_fabric_map Data Source

Get all cluster maps of a user-defined fabric map

## Example Usage

```terraform
data "gigavuecore_get_all_cluster_maps_of_an_user_fabric_map" "example" {
  alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the fabric map

### Attributes

In addition to all arguments above, the following attributes are exported:

* `cluster_maps` (List(Object({afm_map, cluster_id, config_status, error_message, fabric_map_aliases})), computed)
* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type

