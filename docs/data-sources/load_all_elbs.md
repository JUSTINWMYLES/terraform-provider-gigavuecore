---
page_title: "gigavuecore_load_all_elbs Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Elbs
---

# gigavuecore_load_all_elbs Data Source

Load all Elbs

## Example Usage

```terraform
data "gigavuecore_load_all_elbs" "example" {
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
* `elbs` (Attributes List, computed) (see [below for nested schema](#nestedatt--elbs))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--elbs"></a>
### Nested Schema for `elbs`

Read-Only:

* `alias` (String) - app elb alias
* `hash_fields` (Attributes List) (see [below for nested schema](#nestedatt--elbs--hash_fields))
<a id="nestedatt--elbs--hash_fields"></a>
### Nested Schema for `elbs.hash_fields`

Read-Only:

* `hash_field` (String)
* `hash_location` (String) - Ignored when 'hashField' == 'gtpuTeid'. required otherwise

