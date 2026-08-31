---
page_title: "gigavuecore_action_rfs_sync_now Action - gigavuecore"
subcategory: ""
description: |-
  Force Sync Now
---

# gigavuecore_action_rfs_sync_now Action

Force Sync Now

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_action_rfs_sync_now" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of target HSM Group
* `cluster_id` (String, required) - Target cluster ID.


