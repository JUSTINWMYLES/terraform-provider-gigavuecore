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
  page = "example"
  sort = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - app elb alias
* `hash_fields` (Attributes List) (see [below for nested schema](#nestedatt--items--hash_fields))
<a id="nestedatt--items--hash_fields"></a>
### Nested Schema for `items.hash_fields`

Read-Only:

* `hash_field` (String)
* `hash_location` (String) - Ignored when 'hashField' == 'gtpuTeid'. required otherwise

