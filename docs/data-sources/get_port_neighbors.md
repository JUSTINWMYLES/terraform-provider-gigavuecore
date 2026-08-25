---
page_title: "gigavuecore_get_port_neighbors Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Port Neighbors
---

# gigavuecore_get_port_neighbors Data Source

Get Port Neighbors

## Example Usage

```terraform
data "gigavuecore_get_port_neighbors" "example" {
  cluster_id = null
  page       = null
  port_id    = null
  slot_id    = null
  sort       = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `port_id` (String, optional) - target device Port Id. Ignored if 'slotId' parameter is provided.
* `slot_id` (String, optional) - target device Slot Id. Takes precedence over 'portId' parameter
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `port_neighbors` (Attributes List, computed) (see [below for nested schema](#nestedatt--port_neighbors))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--port_neighbors"></a>
### Nested Schema for `port_neighbors`

Read-Only:

* `neighbors` (List of Dynamic)
* `port_id` (String) - id of the port this neighbor details are for

