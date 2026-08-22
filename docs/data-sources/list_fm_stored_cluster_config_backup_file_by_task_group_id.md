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
  task_group_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `task_group_id` (String, required) - TaskGroupId of the Config Backup File

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({device_global_id, device_host_name, device_ip, device_model, device_version, do_not_purge, file_name, member_info})), computed)

