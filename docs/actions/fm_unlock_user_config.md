---
page_title: "gigavuecore_fm_unlock_user_config Action - gigavuecore"
subcategory: ""
description: |-
  FM Unlock User
---

# gigavuecore_fm_unlock_user_config Action

FM Unlock User

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_fm_unlock_user_config" "example" {
  config {
    username = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `username` (String, required) - User Name


