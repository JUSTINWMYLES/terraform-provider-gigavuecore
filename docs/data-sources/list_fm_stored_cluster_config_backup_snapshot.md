---
page_title: "gigavuecore_list_fm_stored_cluster_config_backup_snapshot Data Source - gigavuecore"
subcategory: ""
description: |-
  Lists a labeled config backup snapshot
---

# gigavuecore_list_fm_stored_cluster_config_backup_snapshot Data Source

Lists a labeled config backup snapshot

## Example Usage

```terraform
data "gigavuecore_list_fm_stored_cluster_config_backup_snapshot" "example" {
  backup_id  = "example"
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `backup_id` (String, required) - Unique identifier of the cluster config snapshot for a given cluster
* `cluster_id` (String, required) - Target Cluster Id

### Attributes

In addition to all arguments above, the following attributes are exported:

* `alias` (String, computed) - User-entered alias for this snapshot
* `backup_time` (String, computed) - Configuration file backup time
* `cluster_name` (String, computed) - Name of the cluster
* `cluster_vip` (String, computed) - Virtual Ip of the cluster
* `comment` (String, computed) - User-entered comments for this snapshot
* `device_global_id` (String, computed) - Device Global Id of the hosting node
* `device_host_name` (String, computed) - Device HostName of the hosting node
* `device_ip` (String, computed) - Device Ip of the hosting node
* `do_not_purge` (Boolean, computed) - Indicates whether this snapshot should not be auto-aged/purged
* `file_meta_data` (Attributes, computed) (see [below for nested schema](#nestedatt--file_meta_data))
* `format` (String, computed) - format of the Backup file
* `restore_logs` (List of String, computed) - restoreLogs associated with this file
* `sw_version` (String, computed) - Software version of the device this snapshot was taken from
* `tags` (List of String, computed) - Tags associated with this file

<a id="nestedatt--file_meta_data"></a>
### Nested Schema for `file_meta_data`

Read-Only:

* `device_global_id` (String) - Device Global Id of the hosting node
* `device_host_name` (String) - Device HostName of the hosting node
* `device_ip` (String) - Device Ip of the hosting node
* `device_model` (String) - Gigamon physical device models
* `device_version` (String) - Device's current Software version
* `do_not_purge` (Boolean) - Indicates whether this snapshot should not be auto-aged/purged
* `file_name` (String) - Config file name
* `member_info` (Attributes List) (see [below for nested schema](#nestedatt--file_meta_data--member_info))

<a id="nestedatt--file_meta_data--member_info"></a>
### Nested Schema for `file_meta_data.member_info`

Read-Only:

* `box_id` (String) - Cluster BoxId of the node
* `host_name` (String) - Device HostName of the hosting node
* `role` (String) - Role of the node. (deprecated: use roleAlias)
* `role_alias` (String) - Role of the node

