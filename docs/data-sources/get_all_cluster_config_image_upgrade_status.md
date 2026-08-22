---
page_title: "gigavuecore_get_all_cluster_config_image_upgrade_status Data Source - gigavuecore"
subcategory: ""
description: |-
  get all the cluster configuration imageUpgrade status
---

# gigavuecore_get_all_cluster_config_image_upgrade_status Data Source

get all the cluster configuration imageUpgrade status

## Example Usage

```terraform
data "gigavuecore_get_all_cluster_config_image_upgrade_status" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `clusters_status` (List(Object({cluster_id, cluster_log, end_time, node_upgrade_status, object_counts, object_counts_expected, object_diff_report, start_time, status, task_id, task_name, upgrade_summary})), computed)
* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type

