---
page_title: "gigavuecore_load_all_ib_pathways Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Inter-broker Pathways
---

# gigavuecore_load_all_ib_pathways Data Source

Load all Inter-broker Pathways

## Example Usage

```terraform
data "gigavuecore_load_all_ib_pathways" "example" {
  page = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `ib_pathways` (Attributes List, computed) (see [below for nested schema](#nestedatt--ib_pathways))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--ib_pathways"></a>
### Nested Schema for `ib_pathways`

Read-Only:

* `alias` (String) - Inline resilient inter-broker pathway alias
* `comment` (String)
* `min_ports_up` (Number) - minimum number of ports in the 'up' state needed to declare the ib-pathway to be in the 'up' state.
* `operational_state` (String)
* `ports` (List of String) - list of local network ports with same speed
* `traffic_path` (String)

