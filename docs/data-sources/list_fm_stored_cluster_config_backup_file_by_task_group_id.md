---
page_title: "gigavuecore_list_fm_stored_cluster_config_backup_file_by_task_group_id Data Source - gigavuecore"
subcategory: ""
description: |-
  Lists a labeled config backup file by TaskGroupId
---

# gigavuecore_list_fm_stored_cluster_config_backup_file_by_task_group_id Data Source

Lists a labeled config backup file by TaskGroupId

## Example Usage

```terraform
data "gigavuecore_list_fm_stored_cluster_config_backup_file_by_task_group_id" "example" {
  task_group_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `task_group_id` (String, required) - TaskGroupId of the Config Backup File

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `device_global_id` (String) - Device Global Id of the hosting node
* `device_host_name` (String) - Device HostName of the hosting node
* `device_ip` (String) - Device Ip of the hosting node
* `device_model` (String) - Gigamon physical device models
* `device_version` (String) - Device's current Software version
* `do_not_purge` (Boolean) - Indicates whether this snapshot should not be auto-aged/purged
* `file_name` (String) - Config file name
* `member_info` (Attributes List) (see [below for nested schema](#nestedatt--items--member_info))
<a id="nestedatt--items--member_info"></a>
### Nested Schema for `items.member_info`

Read-Only:

* `box_id` (String) - Cluster BoxId of the node
* `host_name` (String) - Device HostName of the hosting node
* `role` (String) - Role of the node. (deprecated: use roleAlias)
* `role_alias` (String) - Role of the node

