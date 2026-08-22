---
page_title: "gigavuecore_get_upgrade_jobs Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Upgrade Jobs
---

# gigavuecore_get_upgrade_jobs Data Source

Get Upgrade Jobs

## Example Usage

```terraform
data "gigavuecore_get_upgrade_jobs" "example" {
  created_by = null
  end_date = null
  page = null
  scheduled_time = null
  sort = null
  start_date = null
  task_id = null
  task_name = null
  task_status = null
  time_left = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `created_by` (String, optional)
* `end_date` (String, optional)
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `scheduled_time` (String, optional)
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `start_date` (String, optional)
* `task_id` (String, optional)
* `task_name` (String, optional)
* `task_status` (String, optional)
* `time_left` (String, optional)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `device_upgrade_task_info_record` (List(Object({created_by, created_time, description, end_date, scheduled_time, start_date, task_id, task_name, task_status, task_type, time_left})), computed)

