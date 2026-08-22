---
page_title: "gigavuecore_get_managed_clusters Data Source - gigavuecore"
subcategory: ""
description: |-
  get cluster-grouped managed device list
---

# gigavuecore_get_managed_clusters Data Source

get cluster-grouped managed device list

## Example Usage

```terraform
data "gigavuecore_get_managed_clusters" "example" {
  cluster_id = null
  page = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - if provided, only requested cluster is returned
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `clusters` (List(Object({cluster_id, cluster_vip, family, health_state, health_state_reasons, leader_id, master_id, members})), computed)
* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type

