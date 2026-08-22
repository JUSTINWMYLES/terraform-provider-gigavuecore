---
page_title: "gigavuecore_get_upgrade_job_info Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Upgrade Job Info
---

# gigavuecore_get_upgrade_job_info Data Source

Get Upgrade Job Info

## Example Usage

```terraform
data "gigavuecore_get_upgrade_job_info" "example" {
  task_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `task_id` (String, required) - Task ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `created_by` (String, computed) - Created By
* `created_time` (String, computed) - Created Time
* `description` (String, computed) - Description
* `device_upgrade_context` (List(Object({cluster_name, config_backup, device_ip, file_path, file_type, hostname, model})), computed)
* `end_date` (String, computed) - End Date
* `image_server` (String, computed) - Image Server
* `last_update_time` (String, computed) - Last Update Time
* `scheduled_time` (String, computed) - Scheduled Time
* `start_date` (String, computed) - Start Date
* `task_name` (String, computed) - Task Name
* `task_status` (String, computed) - Task Status
* `task_type` (String, computed) - Task Type
* `time_left` (String, computed) - Time Left

