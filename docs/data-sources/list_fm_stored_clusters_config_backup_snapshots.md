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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `cluster_id` (String) - Id of the cluster
* `cluster_name` (String) - Name of the cluster
* `entries` (Attributes List) (see [below for nested schema](#nestedatt--items--entries))
<a id="nestedatt--items--entries"></a>
### Nested Schema for `items.entries`

Read-Only:

* `alias` (String) - User-entered alias for this snapshot
* `backup_id` (String) - Unique identifier of the cluster config snapshot for a given cluster
* `backup_time` (String) - Configuration file backup time
* `cluster_name` (String) - Name of the cluster
* `cluster_vip` (String) - Virtual Ip of the cluster
* `comment` (String) - User-entered comments for this snapshot
* `device_global_id` (String) - Device Global Id of the hosting node
* `device_host_name` (String) - Device HostName of the hosting node
* `device_ip` (String) - Device Ip of the hosting node
* `do_not_purge` (Boolean) - Indicates whether this snapshot should not be auto-aged/purged
* `file_meta_data` (Attributes) (see [below for nested schema](#nestedatt--items--entries--file_meta_data))
* `format` (String) - format of the Backup file
* `restore_logs` (List of String) - restoreLogs associated with this file
* `sw_version` (String) - Software version of the device this snapshot was taken from
* `tags` (List of String) - Tags associated with this file
<a id="nestedatt--items--entries--file_meta_data"></a>
### Nested Schema for `items.entries.file_meta_data`

Read-Only:

* `device_global_id` (String) - Device Global Id of the hosting node
* `device_host_name` (String) - Device HostName of the hosting node
* `device_ip` (String) - Device Ip of the hosting node
* `device_model` (String) - Gigamon physical device models
* `device_version` (String) - Device's current Software version
* `do_not_purge` (Boolean) - Indicates whether this snapshot should not be auto-aged/purged
* `file_name` (String) - Config file name
* `member_info` (Attributes List) (see [below for nested schema](#nestedatt--items--entries--file_meta_data--member_info))
<a id="nestedatt--items--entries--file_meta_data--member_info"></a>
### Nested Schema for `items.entries.file_meta_data.member_info`

Read-Only:

* `box_id` (String) - Cluster BoxId of the node
* `host_name` (String) - Device HostName of the hosting node
* `role` (String) - Role of the node. (deprecated: use roleAlias)
* `role_alias` (String) - Role of the node

