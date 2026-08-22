---
page_title: "gigavuecore_load_all_map_alias_chains Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all aliased map chains
---

# gigavuecore_load_all_map_alias_chains Data Source

Load all aliased map chains

## Example Usage

```terraform
data "gigavuecore_load_all_map_alias_chains" "example" {
  cluster_id = null
  map_alias = null
  page = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `map_alias` (String, optional) - Filters the response to include only the mapChain that includes the map of the given alias. Note that the map can be either cluster map or fabric map.
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `map_alias_chains` (List(Object({cluster_id, collector, health_state, health_state_reasons, map_chain_id, ordered_map_aliases, src_ports, src_ports_as_id})), computed)

