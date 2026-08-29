---
page_title: "gigavuecore_list_bulk_replicate_config_restore_logs Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all bulk replicate config restore logs
---

# gigavuecore_list_bulk_replicate_config_restore_logs Data Source

Load all bulk replicate config restore logs

## Example Usage

```terraform
data "gigavuecore_list_bulk_replicate_config_restore_logs" "example" {
  page = "example"
  sort = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `created_ts` (Number) - File creation timestamp in UTC milliseconds
* `created_ts_utc` (String) - File creation timestamp in ISO 8601 format
* `filename` (String) - Bulk Replicate restore log file name

