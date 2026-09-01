---
page_title: "gigavuecore_edit_user_details Action - gigavuecore"
subcategory: ""
description: |-
  Edit User details
---

# gigavuecore_edit_user_details Action

Edit User details

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Warning:** This action accepts attributes whose names indicate secrets (password), but action schemas cannot mark attributes Sensitive, so their values are stored and displayed in plain text. Avoid passing real secrets through these attributes.

## Example Usage

```terraform
action "gigavuecore_edit_user_details" "example" {
  config {
    body_username = "example"
    email_id      = "example"
    enabled       = true
    full_name     = "example"
    groups        = ["example"]
    password      = "example-value"
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


