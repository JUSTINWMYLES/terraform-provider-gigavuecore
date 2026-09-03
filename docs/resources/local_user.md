---
page_title: "gigavuecore_local_user Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new Local User
---

# gigavuecore_local_user Resource

Create a new Local User

## Example Usage

```terraform
resource "gigavuecore_local_user" "example" {
  account_status   = "accountDisabled"
  capability       = "admin"
  cluster_id       = "example"
  current_password = "example"
  enabled          = true
  full_name        = "example"
  roles            = ["example"]
  user_pwd         = "example"
  username         = "example"
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

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_local_user.example {username}/{cluster_id}
```
