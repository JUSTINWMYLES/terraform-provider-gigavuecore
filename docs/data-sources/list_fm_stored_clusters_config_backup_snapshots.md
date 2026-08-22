---
page_title: "gigavuecore_list_fm_stored_clusters_config_backup_snapshots Data Source - gigavuecore"
subcategory: ""
description: |-
  Lists config backup snapshots for multiple clusters
---

# gigavuecore_list_fm_stored_clusters_config_backup_snapshots Data Source

Lists config backup snapshots for multiple clusters

## Example Usage

```terraform
data "gigavuecore_list_fm_stored_clusters_config_backup_snapshots" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - if provided, only backups associated with this cluster will be returned

### Attributes

In addition to all arguments above, the following attributes are exported:

* `clusters_config_backups` (List(Object({cluster_id, cluster_name, entries})), computed)
* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type

