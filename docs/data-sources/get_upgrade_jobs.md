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
  created_by     = null
  end_date       = null
  page           = null
  scheduled_time = null
  sort           = null
  start_date     = null
  task_id        = null
  task_name      = null
  task_status    = null
  time_left      = null
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

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `device_upgrade_task_info_record` (Attributes List, computed) (see [below for nested schema](#nestedatt--device_upgrade_task_info_record))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--device_upgrade_task_info_record"></a>
### Nested Schema for `device_upgrade_task_info_record`

Read-Only:

* `created_by` (String) - Created By
* `created_time` (String) - Created Time
* `description` (String) - Description
* `end_date` (String) - End Date
* `scheduled_time` (String) - Scheduled Time
* `start_date` (String) - Start Date
* `task_id` (String) - Task ID
* `task_name` (String) - Task Name
* `task_status` (String) - Task Status
* `task_type` (String) - Task Type
* `time_left` (String) - Time Left

