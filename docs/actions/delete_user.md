---
page_title: "gigavuecore_delete_user Action - gigavuecore"
subcategory: ""
description: |-
  Delete a user
---

# gigavuecore_delete_user Action

Delete a user

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_user" "example" {
  config {
    username = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `username` (String, required) - username of target user


