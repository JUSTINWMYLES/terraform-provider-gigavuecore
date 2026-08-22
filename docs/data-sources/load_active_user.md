---
page_title: "gigavuecore_load_active_user Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Active User Details
---

# gigavuecore_load_active_user Data Source

Load Active User Details

## Example Usage

```terraform
data "gigavuecore_load_active_user" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `auth_method` (String, computed) - authentication method
* `enabled` (Bool, computed)
* `fm_authorized_roles` (List(String), computed)
* `full_name` (String, computed) - user's full name
* `groups` (List(String), computed)
* `password` (String, computed) - password
* `username` (String, computed) - username

