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
  task_id = "example"
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
* `device_upgrade_context` (Attributes List, computed) (see [below for nested schema](#nestedatt--device_upgrade_context))
* `end_date` (String, computed) - End Date
* `image_server` (String, computed) - Image Server
* `last_update_time` (String, computed) - Last Update Time
* `scheduled_time` (String, computed) - Scheduled Time
* `start_date` (String, computed) - Start Date
* `task_name` (String, computed) - Task Name
* `task_status` (String, computed) - Task Status
* `task_type` (String, computed) - Task Type
* `time_left` (String, computed) - Time Left

<a id="nestedatt--device_upgrade_context"></a>
### Nested Schema for `device_upgrade_context`

Read-Only:

* `cluster_name` (String) - Cluster Name
* `config_backup` (Boolean) - Config Backup
* `device_ip` (String) - Device IP
* `file_path` (String) - File Path
* `file_type` (String) - File Type
* `hostname` (String) - Host Name
* `model` (String) - Model

