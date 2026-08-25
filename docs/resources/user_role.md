---
page_title: "gigavuecore_user_role Resource - gigavuecore"
subcategory: ""
description: |-
  Load all User Roles
---

# gigavuecore_user_role Resource

Load all User Roles

## Example Usage

```terraform
resource "gigavuecore_user_role" "example" {
  description = null
  name        = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `description` (String, optional)
* `name` (String, required)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `id` (String, computed)
* `user_roles` (Attributes List, computed) (see [below for nested schema](#nestedatt--user_roles))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--user_roles"></a>
### Nested Schema for `user_roles`

Read-Only:

* `description` (String)
* `name` (String)

