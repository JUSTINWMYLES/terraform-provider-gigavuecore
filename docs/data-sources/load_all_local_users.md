---
page_title: "gigavuecore_load_all_local_users Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Local Users
---

# gigavuecore_load_all_local_users Data Source

Load all Local Users

## Example Usage

```terraform
data "gigavuecore_load_all_local_users" "example" {
  cluster_id = null
  page       = null
  sort       = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `account_status` (String)
* `capability` (String)
* `current_password` (String) - Specifies the logged in password
* `enabled` (Boolean) - Temporarily enable/disable logins for the specified account. Disabling an account closes any currently open sessions for the specified account
* `full_name` (String) - Full name for the account (referred to sometimes as the gecos)
* `roles` (List of String) - References to the Roles defined for the User
* `user_pwd` (String)
* `username` (String) - Specifies the username of the local user

