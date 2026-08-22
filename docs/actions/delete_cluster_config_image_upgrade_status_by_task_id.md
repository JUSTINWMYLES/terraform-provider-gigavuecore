---
page_title: "gigavuecore_delete_cluster_config_image_upgrade_status_by_task_id Action - gigavuecore"
subcategory: ""
description: |-
  Delete cluster configuration imageUpgrade status by taskId
---

# gigavuecore_delete_cluster_config_image_upgrade_status_by_task_id Action

Delete cluster configuration imageUpgrade status by taskId

## Example Usage

```terraform
action "gigavuecore_delete_cluster_config_image_upgrade_status_by_task_id" "example" {
  config {
    task_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `task_id` (String, required) - ID of the task
