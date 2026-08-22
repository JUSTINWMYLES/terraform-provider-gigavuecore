---
page_title: "gigavuecore_get_all_alarm_suppression_metadata Data Source - gigavuecore"
subcategory: ""
description: |-
  Get All Alarm Supppression Rules
---

# gigavuecore_get_all_alarm_suppression_metadata Data Source

Get All Alarm Supppression Rules

## Example Usage

```terraform
data "gigavuecore_get_all_alarm_suppression_metadata" "example" {
  alarm_types = null
  alias = null
  cluster_id = null
  hostname = null
  page = null
  resource_id = null
  selected_reason = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alarm_types` (String, optional) - comma-separated list of Alarm Type's to filter
* `alias` (String, optional)
* `cluster_id` (String, optional)
* `hostname` (String, optional)
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:30) to prevent reading entire DB by accident
* `resource_id` (String, optional)
* `selected_reason` (String, optional)
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is DESC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `suppressed_entities` (List(Object({alarm_types, alias, cluster_id, created_by, created_ts, enable, expiry_time, expiry_time_unit, expiry_ts, hostname, resource_id, resource_type, selected_reason, tags, updated_by, updated_ts})), computed)

