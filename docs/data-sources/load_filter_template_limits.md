---
page_title: "gigavuecore_load_filter_template_limits Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Filter Template Limits
---

# gigavuecore_load_filter_template_limits Data Source

Load all Filter Template Limits

## Example Usage

```terraform
data "gigavuecore_load_filter_template_limits" "example" {
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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `cluster_id` (String) - id of the defining cluster
* `filter_limits` (Attributes List) (see [below for nested schema](#nestedatt--items--filter_limits))
* `slot_id` (String)
<a id="nestedatt--items--filter_limits"></a>
### Nested Schema for `items.filter_limits`

Read-Only:

* `alias` (String) - alias of filter template
* `limit` (Number)
* `qualifiers` (List of String)

