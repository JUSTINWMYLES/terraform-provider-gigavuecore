---
page_title: "gigavuecore_load_all_user_tags Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all tags
---

# gigavuecore_load_all_user_tags Data Source

Load all tags

## Example Usage

```terraform
data "gigavuecore_load_all_user_tags" "example" {
  page    = null
  sort    = null
  tag_key = null
  type    = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `tag_key` (String, optional) - Filter based on provided tag keys
* `type` (String, optional) - If provided, returns all tags which matches the tag type

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `description` (String)
* `hierarchical` (Boolean)
* `multi_valued` (Boolean)
* `tag_key` (String) - tag key
* `tag_type` (String)
* `tag_values` (List of String)

