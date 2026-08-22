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
  backup_id = null
  cluster_id = null
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
* `do_not_purge` (Bool, computed) - Indicates whether this snapshot should not be auto-aged/purged
* `file_meta_data` (Object({device_global_id, device_host_name, device_ip, device_model, device_version, do_not_purge, file_name, member_info}), computed)
  * `device_global_id` (String, computed) - Device Global Id of the hosting node
  * `device_host_name` (String, computed) - Device HostName of the hosting node
  * `device_ip` (String, computed) - Device Ip of the hosting node
  * `device_model` (String, computed) - Gigamon physical device models
  * `device_version` (String, computed) - Device's current Software version
  * `do_not_purge` (Bool, computed) - Indicates whether this snapshot should not be auto-aged/purged
  * `file_name` (String, computed) - Config file name
  * `member_info` (List(Object({box_id, host_name, role, role_alias})), computed)
* `format` (String, computed) - format of the Backup file
* `restore_logs` (List(String), computed) - restoreLogs associated with this file
* `sw_version` (String, computed) - Software version of the device this snapshot was taken from
* `tags` (List(String), computed) - Tags associated with this file

