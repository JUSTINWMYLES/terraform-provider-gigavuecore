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

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `roles` (Attributes List, computed) (see [below for nested schema](#nestedatt--roles))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--roles"></a>
### Nested Schema for `roles`

Read-Only:

* `description` (String) - description
* `name` (String) - name
* `scope` (Attributes List) (see [below for nested schema](#nestedatt--roles--scope))
<a id="nestedatt--roles--scope"></a>
### Nested Schema for `roles.scope`

Read-Only:

* `actions` (List of String)
* `hierarchy` (Boolean)
* `type` (String)

