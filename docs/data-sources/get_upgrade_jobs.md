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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

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

