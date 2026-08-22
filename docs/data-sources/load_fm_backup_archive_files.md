---
page_title: "gigavuecore_load_fm_backup_archive_files Data Source - gigavuecore"
subcategory: ""
description: |-
  List FM Backup archive file(s) on archive server
---

# gigavuecore_load_fm_backup_archive_files Data Source

List FM Backup archive file(s) on archive server

## Example Usage

```terraform
data "gigavuecore_load_fm_backup_archive_files" "example" {
  page = null
  server_alias = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `server_alias` (String, required) - Alias of the target Archive Server
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `fm_backup_archive_files` (List(Object({file_path, last_modified, size, version})), computed)

