---
page_title: "gigavuecore_fm_server_aaa_config Action - gigavuecore"
subcategory: ""
description: |-
  fmServer AAA Config
---

# gigavuecore_fm_server_aaa_config Action

fmServer AAA Config

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_fm_server_aaa_config" "example" {
  config {
    accounting_method      = "example"
    auth_method            = "local"
    default_login_attempts = 3
    default_user_group     = "example"
    user_lock_enabled      = true
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `accounting_method` (String, optional) - Accounting Method
* `auth_method` (String, required) - Auth Method
* `default_login_attempts` (Number, optional) - DefaultLoginAttempts will only be applicable for local auth type
* `default_user_group` (String, required) - Default UserGroup
* `user_lock_enabled` (Boolean, optional) - UserLockEnabled will only be applicable for local auth type


