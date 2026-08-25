---
page_title: "gigavuecore_load_fm_backup_archive_servers Data Source - gigavuecore"
subcategory: ""
description: |-
  List FM backup archive servers information
---

# gigavuecore_load_fm_backup_archive_servers Data Source

List FM backup archive servers information

## Example Usage

```terraform
data "gigavuecore_load_fm_backup_archive_servers" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `fm_backup_archive_servers` (Attributes List, computed) (see [below for nested schema](#nestedatt--fm_backup_archive_servers))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--fm_backup_archive_servers"></a>
### Nested Schema for `fm_backup_archive_servers`

Read-Only:

* `address` (String)
* `alias` (String) - archive server unique alias
* `remote_base_path` (String) - remote file staging location where FM copies/lists the archived file(s)
* `type` (String)
* `user_name` (String) - username to use for server login.
* `user_pwd` (String) - user password to use for server login.

