---
page_title: "gigavuecore_load_fabric_path Data Source - gigavuecore"
subcategory: ""
description: |-
  Load fabricPath details for the given srcClusterID and dstClusterID
---

# gigavuecore_load_fabric_path Data Source

Load fabricPath details for the given srcClusterID and dstClusterID

## Example Usage

```terraform
data "gigavuecore_load_fabric_path" "example" {
  details = null
  dst_cluster_id = null
  src_cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `details` (Bool, optional) - load comprehensive response on Fabric Path, including detailed information on topology link endpoints
* `dst_cluster_id` (String, required) - destination Cluster ID
* `src_cluster_id` (String, required) - source Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `giga_topology_fabric_paths_details` (List(Object({alias, config_state, dst_cluster_uuid, dst_vertices, gfp_id, operation_state, path_type, referenced, src_cluster_uuid, src_vertices, topology_links, version})), computed)

