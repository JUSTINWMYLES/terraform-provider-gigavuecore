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
  alias      = null
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

* `backup_info` (Attributes, computed) (see [below for nested schema](#nestedatt--backup_info))
* `restore_info` (Attributes, computed) (see [below for nested schema](#nestedatt--restore_info))

<a id="nestedatt--backup_info"></a>
### Nested Schema for `backup_info`

Read-Only:

* `config_status` (String) - The status of a backup
* `control_tunnels` (Number) - The number of control tunnels backed up
* `failed` (Number) - The number of failed backups
* `filename` (String) - The internal name of the backup file
* `in_progress` (Boolean)
* `last_fail_time` (String) - The timestamp of the last failed backup. In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'
* `last_sucessful_time` (String) - The timestamp of the last successful backup. In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'
* `sessions` (Number) - The number of sessions tunnels backed up
* `success` (Number) - The number of successful backups
* `user_tunnels` (Number) - The number of user tunnels backed up
<a id="nestedatt--restore_info"></a>
### Nested Schema for `restore_info`

Read-Only:

* `restore_time` (String) - The timestamp of the last restore. In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'
* `sessions` (Number) - The number of sessions restored
* `tunnels` (Number) - The number of tunnels restored

