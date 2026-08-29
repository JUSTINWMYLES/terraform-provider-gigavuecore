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
  cluster_id = "example"
  page       = "example"
  sort       = "example"
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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `app_filter_resources` (Attributes List) (see [below for nested schema](#nestedatt--items--app_filter_resources))
* `slot_id` (String)
<a id="nestedatt--items--app_filter_resources"></a>
### Nested Schema for `items.app_filter_resources`

Read-Only:

* `app` (String) - vfp client application
* `free_pool_used` (Number) - vfp resource used from free pool
* `rsvd` (Number) - vfp resource reserved for app
* `rsvd_used` (Number) - resource used from 'rsvd'

