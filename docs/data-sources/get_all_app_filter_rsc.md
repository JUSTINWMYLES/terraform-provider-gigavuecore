---
page_title: "gigavuecore_get_all_app_filter_rsc Data Source - gigavuecore"
subcategory: ""
description: |-
  Load All Application Filter Resources
---

# gigavuecore_get_all_app_filter_rsc Data Source

Load All Application Filter Resources

## Example Usage

```terraform
data "gigavuecore_get_all_app_filter_rsc" "example" {
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

* `app_filter_rscs` (Attributes List, computed) (see [below for nested schema](#nestedatt--app_filter_rscs))
* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))

<a id="nestedatt--app_filter_rscs"></a>
### Nested Schema for `app_filter_rscs`

Read-Only:

* `app_filter_resources` (Attributes List) (see [below for nested schema](#nestedatt--app_filter_rscs--app_filter_resources))
* `slot_id` (String)
<a id="nestedatt--app_filter_rscs--app_filter_resources"></a>
### Nested Schema for `app_filter_rscs.app_filter_resources`

Read-Only:

* `app` (String) - vfp client application
* `free_pool_used` (Number) - vfp resource used from free pool
* `rsvd` (Number) - vfp resource reserved for app
* `rsvd_used` (Number) - resource used from 'rsvd'
<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type

