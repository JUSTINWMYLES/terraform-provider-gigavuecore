---
page_title: "gigavuecore_get_map_groups Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Map Groups
---

# gigavuecore_get_map_groups Data Source

Load all Map Groups

## Example Usage

```terraform
data "gigavuecore_get_map_groups" "example" {
  cluster_id = null
  page       = null
  sort       = null
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

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `map_groups` (Attributes List, computed) (see [below for nested schema](#nestedatt--map_groups))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--map_groups"></a>
### Nested Schema for `map_groups`

Read-Only:

* `alias` (String) - Map Group alias
* `cluster_id` (String) - id of the defining cluster
* `comment` (String)
* `maps` (List of String) - List of maps.SecondLevel maps with flowSample-overlap or flowWhitelist-overlap subtype

