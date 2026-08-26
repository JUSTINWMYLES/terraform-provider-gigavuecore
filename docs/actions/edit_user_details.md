---
page_title: "gigavuecore_edit_user_details Action - gigavuecore"
subcategory: ""
description: |-
  Edit User details
---

# gigavuecore_edit_user_details Action

Edit User details

## Example Usage

```terraform
action "gigavuecore_edit_user_details" "example" {
  config {
    body_username = "example"
    email_id      = "example"
    enabled       = true
    full_name     = "example"
    groups        = [ "example" ]
    password      = "example"
    username      = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `body_username` (String, required) - username
* `email_id` (String, required) - email ID
* `enabled` (Boolean, optional)
* `full_name` (String, optional) - user's full name
* `groups` (List of String, optional)
* `password` (String, required) - password
* `username` (String, required) - Target Username


