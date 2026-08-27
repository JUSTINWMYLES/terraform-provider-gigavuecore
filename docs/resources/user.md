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

* `enabled` (Boolean, computed)
* `full_name` (String, computed) - user's full name
* `groups` (List of String, computed)
* `id` (String, computed)


