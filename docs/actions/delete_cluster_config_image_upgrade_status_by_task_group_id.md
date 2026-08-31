---
page_title: "gigavuecore_delete_cluster_config_image_upgrade_status_by_task_group_id Action - gigavuecore"
subcategory: ""
description: |-
  Delete cluster configuration imageUpgrade status by taskGroupId
---

# gigavuecore_delete_cluster_config_image_upgrade_status_by_task_group_id Action

Delete cluster configuration imageUpgrade status by taskGroupId

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_cluster_config_image_upgrade_status_by_task_group_id" "example" {
  config {
    task_group_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `task_group_id` (String, required) - ID of the taskGroup


