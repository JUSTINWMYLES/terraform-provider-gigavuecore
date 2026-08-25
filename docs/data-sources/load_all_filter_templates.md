---
page_title: "gigavuecore_load_all_filter_templates Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Filter Templates
---

# gigavuecore_load_all_filter_templates Data Source

Load all Filter Templates

## Example Usage

```terraform
data "gigavuecore_load_all_filter_templates" "example" {
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
* `filter_templates` (Attributes List, computed) (see [below for nested schema](#nestedatt--filter_templates))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--filter_templates"></a>
### Nested Schema for `filter_templates`

Read-Only:

* `alias` (String) - Filter Template alias
* `qualifier_set` (List of String)
* `qualifiers` (List of String) - 'ip6fl', 'ipfrag', 'tcpctl', 'tos', 'ttl' are available only for default filter templates. 'circuit-id' is readonly and will be available on all filter templates by default

