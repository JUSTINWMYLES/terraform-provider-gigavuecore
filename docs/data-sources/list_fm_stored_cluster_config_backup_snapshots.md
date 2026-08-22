---
page_title: "gigavuecore_list_fm_stored_cluster_config_backup_snapshots Data Source - gigavuecore"
subcategory: ""
description: |-
  Lists config backup snapshots a given cluster
---

# gigavuecore_list_fm_stored_cluster_config_backup_snapshots Data Source

Lists config backup snapshots a given cluster

## Example Usage

```terraform
data "gigavuecore_list_fm_stored_cluster_config_backup_snapshots" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Id of the cluster

### Attributes

In addition to all arguments above, the following attributes are exported:

* `cluster_name` (String, computed) - Name of the cluster
* `entries` (List(Object({alias, backup_id, backup_time, cluster_name, cluster_vip, comment, device_global_id, device_host_name, device_ip, do_not_purge, file_meta_data, format, restore_logs, sw_version, tags})), computed)

