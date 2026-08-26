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
  node_id    = null
  page       = null
  sort       = null
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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `admin_status` (String)
* `alarm_buffer_threshold` (Number) - card micro burst threshold
* `config_status` (String)
* `fabric_hash_adv` (Boolean) - Advanced Fabric Hash. Supported only for Q02X32/Q08 cards
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--health_state_reasons))
* `hw_revision` (String)
* `hw_type` (String)
* `mode` (String)
* `oper_status` (String)
* `pld_info` (Attributes) (see [below for nested schema](#nestedatt--items--pld_info))
* `power_priority` (Number) - Power slot Priority
* `power_req` (Number) - watt
* `product_code` (String)
* `serial_number` (String)
* `slot_id` (String) - device card slot id
* `temperatures` (Attributes) (see [below for nested schema](#nestedatt--items--temperatures))
* `voltages` (Attributes List) (see [below for nested schema](#nestedatt--items--voltages))
<a id="nestedatt--items--health_state_reasons"></a>
### Nested Schema for `items.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--items--pld_info"></a>
### Nested Schema for `items.pld_info`

Read-Only:

* `need_upgrade` (Boolean)
* `pld_revision` (String) - PLD revision
<a id="nestedatt--items--temperatures"></a>
### Nested Schema for `items.temperatures`

Read-Only:

* `board` (Number) - celsius
* `bottom_switch` (Number) - celsius
* `bottom_switch_major` (Number) - celsius
* `bottom_switch_minor` (Number) - celsius
* `bottom_switch_shut` (Number) - celsius
* `cav_cpu` (Number) - celsius
* `cpu` (Number) - celsius
* `cpu_major` (Number) - celsius
* `cpu_minor` (Number) - celsius
* `cpu_shut` (Number) - celsius
* `e1_cpu` (Number) - celsius
* `e1_port` (Number) - celsius
* `e2_cpu` (Number) - celsius
* `exhaust` (Number) - celsius
* `intake` (Number) - celsius
* `near_cpu` (Number) - celsius
* `near_cpu_major` (Number) - celsius
* `near_cpu_minor` (Number) - celsius
* `near_cpu_shut` (Number) - celsius
* `qsfp_cage` (Number) - celsius
* `qsfp_cage_major` (Number) - celsius
* `qsfp_cage_minor` (Number) - celsius
* `qsfp_cage_shut` (Number) - celsius
* `rear_panel` (Number) - celsius
* `rear_panel_major` (Number) - celsius
* `rear_panel_minor` (Number) - celsius
* `rear_panel_shut` (Number) - celsius
* `sfp_cage` (Number) - celsius
* `sfp_cage_major` (Number) - celsius
* `sfp_cage_minor` (Number) - celsius
* `sfp_cage_shut` (Number) - celsius
* `switch` (Number) - celsius
* `switch_major` (Number) - celsius
* `switch_minor` (Number) - celsius
* `switch_shut` (Number) - celsius
* `time_stamp` (String) - date-time of stats collection in RFC 3339 format
* `top_switch` (Number) - celsius
* `top_switch_major` (Number) - celsius
* `top_switch_minor` (Number) - celsius
* `top_switch_shut` (Number) - celsius
<a id="nestedatt--items--voltages"></a>
### Nested Schema for `items.voltages`

Read-Only:

* `time_stamp` (String) - date-time of stats collection in RFC 3339 format
* `value` (Number) - volt
* `voltage` (String)

