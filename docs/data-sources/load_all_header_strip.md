---
page_title: "gigavuecore_load_all_header_strip Data Source - gigavuecore"
subcategory: ""
description: |-
  Load header strip for all boxes
---

# gigavuecore_load_all_header_strip Data Source

Load header strip for all boxes

## Example Usage

```terraform
data "gigavuecore_load_all_header_strip" "example" {
  cluster_id = "example"
  page       = "example"
  sort       = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target cluster ID.
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `box_id` (String) - device box id. valid range 1 - 64.
* `mpls_labels` (List of String) - mpls ids, valid and required. Range can be specified. Example:1..200

