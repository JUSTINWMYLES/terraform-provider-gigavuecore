---
page_title: "gigavuecore_user Resource - gigavuecore"
subcategory: ""
description: |-
  Load User by Username
---

# gigavuecore_user Resource

Load User by Username

## Example Usage

```terraform
resource "gigavuecore_user" "example" {
  email_id = null
  enabled = null
  full_name = null
  groups = []
  password = null
  username = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `email_id` (String, required) - email ID
* `enabled` (Bool, optional)
* `full_name` (String, optional) - user's full name
* `groups` (List(String), optional)
* `password` (String, required) - password
* `username` (String, required) - username

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `id` (String, computed)
* `users` (List(Object({email_id, enabled, full_name, groups, password, username})), computed)

