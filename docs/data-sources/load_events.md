---
page_title: "gigavuecore_load_events Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Events
---

# gigavuecore_load_events Data Source

Load Events

## Example Usage

```terraform
data "gigavuecore_load_events" "example" {
  end_time = null
  page = null
  resource_id = null
  resource_type = null
  scope = null
  severity = null
  sort = null
  source = null
  start_time = null
  type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `end_time` (String, optional) - End Time to filter by. In ISO 8601 format
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:100) to prevent reading entire DB by accident
* `resource_id` (String, optional) - Affected Entity to filter by
* `resource_type` (String, optional) - Affected Entity Type to filter by
* `scope` (String, optional) - Event Scope to filter by
* `severity` (String, optional) - Event Scope to filter by
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `source` (String, optional) - Event Source identifier to filter by
* `start_time` (String, optional) - Start Time to filter by. In ISO 8601 format
* `type` (String, optional) - Event Type to filter by

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `events` (List(Object({description, resource_id, resource_type, scope, severity, source, ts, ts_utc, type})), computed)

