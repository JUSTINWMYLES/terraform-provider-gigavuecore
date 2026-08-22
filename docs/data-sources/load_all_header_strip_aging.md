---
page_title: "gigavuecore_load_all_header_strip_aging Data Source - gigavuecore"
subcategory: ""
description: |-
  Load header strip aging for all boxes
---

# gigavuecore_load_all_header_strip_aging Data Source

Load header strip aging for all boxes

## Example Usage

```terraform
data "gigavuecore_load_all_header_strip_aging" "example" {
  box_id = null
  page = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, optional) - device box id. valid range 1 - 64.
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `header_strips_aging_def` (List(Object({aging_interval, box_id, dst_port, protocol_type})), computed)

