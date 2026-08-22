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

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `fm_backup_archive_servers` (List(Object({address, alias, remote_base_path, type, user_name, user_pwd})), computed)

