---
page_title: "gigavuecore_load_all_stack_links Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Stack Links
---

# gigavuecore_load_all_stack_links Data Source

Load all Stack Links

## Example Usage

```terraform
data "gigavuecore_load_all_stack_links" "example" {
  cluster_id = null
  page = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `stack_links` (List(Object({alias, cluster_id, comment, endpoint1, endpoint2, health_state, health_state_reasons, port_list1, port_list2, type})), computed)

