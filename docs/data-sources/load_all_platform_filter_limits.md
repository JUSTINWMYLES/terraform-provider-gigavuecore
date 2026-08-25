---
page_title: "gigavuecore_load_all_platform_filter_limits Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all platform Filter Template Limits
---

# gigavuecore_load_all_platform_filter_limits Data Source

Load all platform Filter Template Limits

## Example Usage

```terraform
data "gigavuecore_load_all_platform_filter_limits" "example" {
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
* `platforms_filter_limits` (Attributes List, computed) (see [below for nested schema](#nestedatt--platforms_filter_limits))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--platforms_filter_limits"></a>
### Nested Schema for `platforms_filter_limits`

Read-Only:

* `alias` (String) - alias of filter template
* `cluster_id` (String) - id of the defining cluster
* `platform_limits` (Attributes List) (see [below for nested schema](#nestedatt--platforms_filter_limits--platform_limits))
<a id="nestedatt--platforms_filter_limits--platform_limits"></a>
### Nested Schema for `platforms_filter_limits.platform_limits`

Read-Only:

* `card_limits` (Attributes List) (see [below for nested schema](#nestedatt--platforms_filter_limits--platform_limits--card_limits))
* `platform` (String)
<a id="nestedatt--platforms_filter_limits--platform_limits--card_limits"></a>
### Nested Schema for `platforms_filter_limits.platform_limits.card_limits`

Read-Only:

* `card` (String) - Card Type
* `limit` (Number)

