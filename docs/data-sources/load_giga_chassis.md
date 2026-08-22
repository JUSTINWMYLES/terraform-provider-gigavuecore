---
page_title: "gigavuecore_load_giga_chassis Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Device Chassis
---

# gigavuecore_load_giga_chassis Data Source

Load all Device Chassis

## Example Usage

```terraform
data "gigavuecore_load_giga_chassis" "example" {
  cluster_id = null
  node_id = null
  page = null
  ports = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - Target Cluster ID. Either 'clusterId' or 'nodeId' is required
* `node_id` (String, optional) - ID of the target device. Either 'clusterId' or 'nodeId' is required
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `ports` (Bool, optional) - Indicates whether per-chassis port list should be included in the response
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `chassis_list` (List(Object({boot_time, box_id, build_id, cards, cc_sync_status, chassis_id, cpld_version, cpu_load_avg, device_name, fan_trays, file_system_statistics, gdp, git_full_hash, global_node_id, hardware_info, health_state, health_state_reasons, host_id, hw_revision, hw_type, l2_gre_id, leaf_config, memory, mode, model, oper_status, ports, power_management, power_modules, product_build_date, product_code, product_name, product_version, serial_number, sys_contact, sys_descr, sys_location, system_mode, ts_version, uboot_version, version_summary, vxlan_id})), computed)
* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type

