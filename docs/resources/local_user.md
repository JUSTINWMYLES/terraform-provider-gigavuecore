---
page_title: "gigavuecore_local_user Resource - gigavuecore"
subcategory: ""
description: |-
  Find Local User by username
---

# gigavuecore_local_user Resource

Find Local User by username

## Example Usage

```terraform
resource "gigavuecore_local_user" "example" {
  account_status   = null
  capability       = null
  cluster_id       = null
  current_password = null
  enabled          = null
  full_name        = null
  roles            = []
  user_pwd         = null
  username         = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `account_status` (String, optional)
* `capability` (String, optional)
* `cluster_id` (String, required) - Target Cluster ID
* `current_password` (String, required) - Specifies the logged in password
* `enabled` (Boolean, optional) - Temporarily enable/disable logins for the specified account. Disabling an account closes any currently open sessions for the specified account
* `full_name` (String, optional) - Full name for the account (referred to sometimes as the gecos)
* `roles` (List of String, optional) - References to the Roles defined for the User
* `user_pwd` (String, required)
* `username` (String, required) - Specifies the username of the local user

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `account_status` (String, computed)
* `capability` (String, computed)
* `enabled` (Boolean, computed) - Temporarily enable/disable logins for the specified account. Disabling an account closes any currently open sessions for the specified account
* `full_name` (String, computed) - Full name for the account (referred to sometimes as the gecos)
* `roles` (List of String, computed) - References to the Roles defined for the User


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_local_user.example {username}
```
