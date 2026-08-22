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
  name = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `description` (String, optional)
* `name` (String, required)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `id` (String, computed)
* `user_roles` (List(Object({description, name})), computed)

