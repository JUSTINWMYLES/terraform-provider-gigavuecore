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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `address` (String)
* `alias` (String) - archive server unique alias
* `remote_base_path` (String) - remote file staging location where FM copies/lists the archived file(s)
* `type` (String)
* `user_name` (String) - username to use for server login.
* `user_pwd` (String) - user password to use for server login.

