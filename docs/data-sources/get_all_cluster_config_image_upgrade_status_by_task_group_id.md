---
page_title: "gigavuecore_get_all_cluster_config_image_upgrade_status_by_task_group_id Data Source - gigavuecore"
subcategory: ""
description: |-
  get all the cluster configuration imageUpgrade status by TaskGroupId
---

# gigavuecore_get_all_cluster_config_image_upgrade_status_by_task_group_id Data Source

get all the cluster configuration imageUpgrade status by TaskGroupId

## Example Usage

```terraform
data "gigavuecore_get_all_cluster_config_image_upgrade_status_by_task_group_id" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `upgrade_task_status` (List(Object({end_time, multi_upgrade_status, num_nodes, num_nodes_success, start_time, status, task_group_id, task_name})), computed)

