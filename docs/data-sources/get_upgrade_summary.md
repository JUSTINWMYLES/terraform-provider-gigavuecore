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
  activate_info = null
  activate_stage = null
  chassis_oper_status = null
  cluster_id = null
  fetch_info = null
  fetch_stage = null
  health_state = null
  hostname = null
  install_info = null
  install_stage = null
  last_executed_task = null
  licensed = null
  model = null
  page = null
  role = null
  serial_number = null
  sort = null
  sw_version = null
  tags = null
  task_status = null
  uboot_version = null
  verify_info = null
  verify_stage = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `activate_info` (String, optional)
* `activate_stage` (String, optional)
* `chassis_oper_status` (String, optional)
* `cluster_id` (String, optional)
* `fetch_info` (String, optional)
* `fetch_stage` (String, optional)
* `health_state` (String, optional)
* `hostname` (String, optional)
* `install_info` (String, optional)
* `install_stage` (String, optional)
* `last_executed_task` (String, optional)
* `licensed` (String, optional)
* `model` (String, optional)
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `role` (String, optional)
* `serial_number` (String, optional)
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `sw_version` (String, optional)
* `tags` (String, optional)
* `task_status` (String, optional)
* `uboot_version` (String, optional)
* `verify_info` (String, optional)
* `verify_stage` (String, optional)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `device_upgrade_summary` (List(Object({activate_info, activate_stage, chassis_oper_status, cluster_id, device_ip, discovery_outcome, fetch_info, fetch_stage, health_state, health_state_reasons, hostname, install_info, install_stage, last_executed_task, licensed, maintenance, model, role, serial_number, sw_version, tags, task_status, uboot_version, verify_info, verify_stage})), computed)

