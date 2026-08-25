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
  page         = null
  server_alias = null
  sort         = null
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

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `fm_backup_archive_files` (Attributes List, computed) (see [below for nested schema](#nestedatt--fm_backup_archive_files))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--fm_backup_archive_files"></a>
### Nested Schema for `fm_backup_archive_files`

Read-Only:

* `file_path` (String) - Path of the file
* `last_modified` (String) - last modified time of the file
* `size` (Number) - size of the file
* `version` (String) - version of the archive file

