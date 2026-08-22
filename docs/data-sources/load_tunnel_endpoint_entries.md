---
page_title: "gigavuecore_load_tunnel_endpoint_entries Data Source - gigavuecore"
subcategory: ""
description: |-
  Load known Tunnel Endpoints
---

# gigavuecore_load_tunnel_endpoint_entries Data Source

Load known Tunnel Endpoints

## Example Usage

```terraform
data "gigavuecore_load_tunnel_endpoint_entries" "example" {
  cluster_id = null
  page = null
  sort = null
  tunnel_type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Cluster ID to filter by
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:100) to prevent reading entire DB by accident
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `tunnel_type` (String, optional) - Tunnel type to filter by

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `tunnel_endpoints` (List(Object({tunnel_endpoint, tunneled_port})), computed)

