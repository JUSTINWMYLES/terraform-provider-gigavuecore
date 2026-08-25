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
  email_id  = null
  enabled   = null
  full_name = null
  groups    = []
  password  = null
  username  = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `email_id` (String, required) - email ID
* `enabled` (Boolean, optional)
* `full_name` (String, optional) - user's full name
* `groups` (List of String, optional)
* `password` (String, required) - password
* `username` (String, required) - username

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `id` (String, computed)
* `users` (Attributes List, computed) (see [below for nested schema](#nestedatt--users))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--users"></a>
### Nested Schema for `users`

Read-Only:

* `email_id` (String) - email ID
* `enabled` (Boolean)
* `full_name` (String) - user's full name
* `groups` (List of String)
* `password` (String) - password
* `username` (String) - username

