---
page_title: "gigavuecore_get_upgrade_summary Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Upgrade Summary
---

# gigavuecore_get_upgrade_summary Data Source

Get Upgrade Summary

## Example Usage

```terraform
data "gigavuecore_get_upgrade_summary" "example" {
  activate_info       = "example"
  activate_stage      = "example"
  chassis_oper_status = "example"
  cluster_id          = "example"
  fetch_info          = "example"
  fetch_stage         = "example"
  health_state        = "example"
  hostname            = "example"
  install_info        = "example"
  install_stage       = "example"
  last_executed_task  = "example"
  licensed            = "example"
  model               = "example"
  page                = "example"
  role                = "example"
  serial_number       = "example"
  sort                = "example"
  sw_version          = "example"
  tags                = "example"
  task_status         = "example"
  uboot_version       = "example"
  verify_info         = "example"
  verify_stage        = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `activate_info` (String, optional) - Activate Info
* `activate_stage` (String, optional) - Activate Stage
* `chassis_oper_status` (String, optional) - Chassis OperStatus
* `cluster_id` (String, optional) - Cluster ID
* `fetch_info` (String, optional) - Fetch Info
* `fetch_stage` (String, optional) - Fetch Stage
* `health_state` (String, optional) - Health State
* `hostname` (String, optional) - Host Name
* `install_info` (String, optional) - Install Info
* `install_stage` (String, optional) - Install Stage
* `last_executed_task` (String, optional) - Last Executed Task
* `licensed` (String, optional) - Licensed
* `model` (String, optional) - Model
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `role` (String, optional) - Role
* `serial_number` (String, optional) - Serial Number
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `sw_version` (String, optional) - Software Version
* `tags` (String, optional) - Tags
* `task_status` (String, optional) - Task Status
* `uboot_version` (String, optional) - UBoot Version
* `verify_info` (String, optional) - verifyInfo
* `verify_stage` (String, optional) - Verify Stage

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `activate_info` (String) - Activate Info
* `activate_stage` (String) - Activate Stage
* `chassis_oper_status` (String) - Chassis OperStatus
* `cluster_id` (String) - Cluster ID
* `device_ip` (String) - Device IP
* `discovery_outcome` (String) - Discovery Outcome
* `fetch_info` (String) - Fetch Info
* `fetch_stage` (String) - Fetch Stage
* `health_state` (String) - Health State
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--health_state_reasons))
* `hostname` (String) - Host Name
* `install_info` (String) - Install Info
* `install_stage` (String) - Install Stage
* `last_executed_task` (String) - Last Executed Task
* `licensed` (Boolean) - Licensed
* `maintenance` (Boolean) - Maintenance
* `model` (String) - Model
* `role` (String) - Role
* `serial_number` (String) - Serial Number
* `sw_version` (String) - Software Version
* `tags` (Attributes List) (see [below for nested schema](#nestedatt--items--tags))
* `task_status` (String) - Task Status
* `uboot_version` (String) - UBoot Version
* `verify_info` (String) - Verify Info
* `verify_stage` (String) - Verify Stage

<a id="nestedatt--items--health_state_reasons"></a>
### Nested Schema for `items.health_state_reasons`

Read-Only:

* `message` (String) - Message
* `severity` (String) - Severity
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

<a id="nestedatt--items--tags"></a>
### Nested Schema for `items.tags`

Read-Only:

* `tag_key` (String) - Tag Key
* `tag_values` (List of Dynamic) - Tag Values

