---
page_title: "gigavuecore_delete_upgrade_jobs Action - gigavuecore"
subcategory: ""
description: |-
  Delete Upgrade Jobs
---

# gigavuecore_delete_upgrade_jobs Action

Delete Upgrade Jobs

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_upgrade_jobs" "example" {
  config {
    task_ids = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `task_ids` (String, required) - Task IDs


