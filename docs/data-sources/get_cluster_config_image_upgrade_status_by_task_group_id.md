---
page_title: "gigavuecore_get_cluster_config_image_upgrade_status_by_task_group_id Data Source - gigavuecore"
subcategory: ""
description: |-
  get the cluster configuration imageUpgrade status by taskGroupId
---

# gigavuecore_get_cluster_config_image_upgrade_status_by_task_group_id Data Source

get the cluster configuration imageUpgrade status by taskGroupId

## Example Usage

```terraform
data "gigavuecore_get_cluster_config_image_upgrade_status_by_task_group_id" "example" {
  task_group_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `task_group_id` (String, required) - system generated id

### Attributes

In addition to all arguments above, the following attributes are exported:

* `end_time` (String, computed) - End Time in UTC format
* `multi_upgrade_status` (List(Object({cluster_id, cluster_log, end_time, node_upgrade_status, object_counts, object_counts_expected, object_diff_report, start_time, status, task_id, task_name, upgrade_summary})), computed) - list of the upgrade Status
* `num_nodes` (Number, computed) - total number of nodes
* `num_nodes_success` (Number, computed) - total number of nodes success state
* `start_time` (String, computed) - Start Time in UTC format
* `status` (String, computed) - Image Upgrade Status
* `task_name` (String, computed) - name of the task

