---
page_title: "gigavuecore_list_bulk_replicate_config_files Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all bulk replicate config files
---

# gigavuecore_list_bulk_replicate_config_files Data Source

Load all bulk replicate config files

## Example Usage

```terraform
data "gigavuecore_list_bulk_replicate_config_files" "example" {
  page = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({comment, created_ts, created_ts_utc, family, filename})), computed)

