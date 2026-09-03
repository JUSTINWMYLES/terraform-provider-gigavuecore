---
page_title: "gigavuecore_reset_user_account_locks Action - gigavuecore"
subcategory: ""
description: |-
  Clears history of login failures, and/or unlocks user account
---

# gigavuecore_reset_user_account_locks Action

Clears history of login failures, and/or unlocks user account

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_reset_user_account_locks" "example" {
  config {
    clear_history = true
    cluster_id    = "example"
    unlock        = true
    username      = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `clear_history` (Boolean, optional) - If set to 'false', leave the history alone and only unlock the account. Therefore, one more login will be permitted, but the account could then become re-locked after another failure (if it was already over the threshold)
* `cluster_id` (String, required) - Target Cluster ID
* `unlock` (Boolean, optional) - If set to 'false', clear the history, but leave the account's lock alone. Therefore, if it was locked, it remains locked until further action is taken
* `username` (String, optional) - account to reset lockout settings for. If left out, all user accounts are affected


