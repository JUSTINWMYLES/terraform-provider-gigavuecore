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
  page   = null
  sort   = null
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

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `header_strips_aging_def` (Attributes List, computed) (see [below for nested schema](#nestedatt--header_strips_aging_def))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--header_strips_aging_def"></a>
### Nested Schema for `header_strips_aging_def`

Read-Only:

* `aging_interval` (Number) - Interval in sec. Valid range 300-1000000. Enter 0 to disable.
* `box_id` (String) - device box id. valid range 1 - 64. all is applicable only for post request.
* `dst_port` (Number) - L4 destination port number.Valid value is between 0 to 65535.
* `protocol_type` (String) - protocol type

