---
page_title: "gigavuecore_load_all_roles Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all roles
---

# gigavuecore_load_all_roles Data Source

Load all roles

## Example Usage

```terraform
data "gigavuecore_load_all_roles" "example" {
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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `description` (String) - description
* `name` (String) - name
* `scope` (Attributes List) (see [below for nested schema](#nestedatt--items--scope))
<a id="nestedatt--items--scope"></a>
### Nested Schema for `items.scope`

Read-Only:

* `actions` (List of String)
* `hierarchy` (Boolean)
* `type` (String)

