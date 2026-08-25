---
page_title: "gigavuecore_load_all_gtp_persistence_backup_restore Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Gtp Persistence backup restore information
---

# gigavuecore_load_all_gtp_persistence_backup_restore Data Source

Load all Gtp Persistence backup restore information

## Example Usage

```terraform
data "gigavuecore_load_all_gtp_persistence_backup_restore" "example" {
  cluster_id = null
  page       = null
  sort       = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `gtp_persistant_backup_restore_infos` (Attributes List, computed) (see [below for nested schema](#nestedatt--gtp_persistant_backup_restore_infos))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--gtp_persistant_backup_restore_infos"></a>
### Nested Schema for `gtp_persistant_backup_restore_infos`

Read-Only:

* `alias` (String) - Gs Group alias
* `backup_info` (Attributes) (see [below for nested schema](#nestedatt--gtp_persistant_backup_restore_infos--backup_info))
* `cluster_id` (String) - id of the defining cluster
* `restore_info` (Attributes) (see [below for nested schema](#nestedatt--gtp_persistant_backup_restore_infos--restore_info))
<a id="nestedatt--gtp_persistant_backup_restore_infos--backup_info"></a>
### Nested Schema for `gtp_persistant_backup_restore_infos.backup_info`

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
<a id="nestedatt--gtp_persistant_backup_restore_infos--restore_info"></a>
### Nested Schema for `gtp_persistant_backup_restore_infos.restore_info`

Read-Only:

* `restore_time` (String) - The timestamp of the last restore. In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'
* `sessions` (Number) - The number of sessions restored
* `tunnels` (Number) - The number of tunnels restored

