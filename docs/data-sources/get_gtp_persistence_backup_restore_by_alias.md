---
page_title: "gigavuecore_get_gtp_persistence_backup_restore_by_alias Data Source - gigavuecore"
subcategory: ""
description: |-
  Find Gtp Persistence backup restore information by alias
---

# gigavuecore_get_gtp_persistence_backup_restore_by_alias Data Source

Find Gtp Persistence backup restore information by alias

## Example Usage

```terraform
data "gigavuecore_get_gtp_persistence_backup_restore_by_alias" "example" {
  alias = null
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Gs Group alias
* `cluster_id` (String, required) - id of the defining cluster

### Attributes

In addition to all arguments above, the following attributes are exported:

* `backup_info` (Object({config_status, control_tunnels, failed, filename, in_progress, last_fail_time, last_sucessful_time, sessions, success, user_tunnels}), computed)
  * `config_status` (String, computed) - The status of a backup
  * `control_tunnels` (Number, computed) - The number of control tunnels backed up
  * `failed` (Number, computed) - The number of failed backups
  * `filename` (String, computed) - The internal name of the backup file
  * `in_progress` (Bool, computed)
  * `last_fail_time` (String, computed) - The timestamp of the last failed backup. In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'
  * `last_sucessful_time` (String, computed) - The timestamp of the last successful backup. In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'
  * `sessions` (Number, computed) - The number of sessions tunnels backed up
  * `success` (Number, computed) - The number of successful backups
  * `user_tunnels` (Number, computed) - The number of user tunnels backed up
* `restore_info` (Object({restore_time, sessions, tunnels}), computed)
  * `restore_time` (String, computed) - The timestamp of the last restore. In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'
  * `sessions` (Number, computed) - The number of sessions restored
  * `tunnels` (Number, computed) - The number of tunnels restored

