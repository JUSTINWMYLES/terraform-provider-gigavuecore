---
page_title: "gigavuecore_get_all_cluster_config_image_upgrade_status_by_task_group_id List Resource - gigavuecore"
subcategory: ""
description: |-
  get all the cluster configuration imageUpgrade status by TaskGroupId
---

# gigavuecore_get_all_cluster_config_image_upgrade_status_by_task_group_id List Resource

get all the cluster configuration imageUpgrade status by TaskGroupId

## Example Usage

```terraform
list "gigavuecore_get_all_cluster_config_image_upgrade_status_by_task_group_id" "example" {
  provider = gigavuecore
  limit    = 100
}

```
## Schema

### Identity Attributes

The following identity attributes are exported for each matching result:

* `task_group_id` (String, computed) - system generated id


