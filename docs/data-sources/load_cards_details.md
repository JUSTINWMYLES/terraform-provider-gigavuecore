---
page_title: "gigavuecore_load_cards_details Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Device Cards details
---

# gigavuecore_load_cards_details Data Source

Load all Device Cards details

## Example Usage

```terraform
data "gigavuecore_load_cards_details" "example" {
  cluster_id = null
  node_id = null
  page = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - Target Cluster ID. Either 'clusterId' or 'nodeId' is required
* `node_id` (String, optional) - ID of the target device. Either 'clusterId' or 'nodeId' is required
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `cards` (List(Object({admin_status, alarm_buffer_threshold, config_status, fabric_hash_adv, health_state, health_state_reasons, hw_revision, hw_type, mode, oper_status, pld_info, power_priority, power_req, product_code, serial_number, slot_id, temperatures, voltages})), computed)
* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type

