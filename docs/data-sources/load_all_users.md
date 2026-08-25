---
page_title: "gigavuecore_load_all_users Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Users
---

# gigavuecore_load_all_users Data Source

Load all Users

## Example Usage

```terraform
data "gigavuecore_load_all_users" "example" {
  page     = null
  sort     = null
  username = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `username` (String, optional) - username of the target user

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
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

