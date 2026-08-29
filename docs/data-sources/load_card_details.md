---
page_title: "gigavuecore_load_card_details Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Device Card details
---

# gigavuecore_load_card_details Data Source

Load Device Card details

## Example Usage

```terraform
data "gigavuecore_load_card_details" "example" {
  cluster_id = "example"
  node_id    = "example"
  slot_id    = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - Target Cluster ID. Either 'clusterId' or 'nodeId' is required
* `node_id` (String, optional) - ID of the target device. Either 'clusterId' or 'nodeId' is required
* `slot_id` (String, required) - device card slot id

### Attributes

In addition to all arguments above, the following attributes are exported:

* `admin_status` (String, computed)
* `alarm_buffer_threshold` (Number, computed) - card micro burst threshold
* `config_status` (String, computed)
* `fabric_hash_adv` (Boolean, computed) - Advanced Fabric Hash. Supported only for Q02X32/Q08 cards
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `hw_revision` (String, computed)
* `hw_type` (String, computed)
* `mode` (String, computed)
* `oper_status` (String, computed)
* `pld_info` (Attributes, computed) (see [below for nested schema](#nestedatt--pld_info))
* `power_priority` (Number, computed) - Power slot Priority
* `power_req` (Number, computed) - watt
* `product_code` (String, computed)
* `serial_number` (String, computed)
* `temperatures` (Attributes, computed) (see [below for nested schema](#nestedatt--temperatures))
* `voltages` (Attributes List, computed) (see [below for nested schema](#nestedatt--voltages))

<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--pld_info"></a>
### Nested Schema for `pld_info`

Read-Only:

* `need_upgrade` (Boolean)
* `pld_revision` (String) - PLD revision
<a id="nestedatt--temperatures"></a>
### Nested Schema for `temperatures`

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
<a id="nestedatt--voltages"></a>
### Nested Schema for `voltages`

Read-Only:

* `time_stamp` (String) - date-time of stats collection in RFC 3339 format
* `value` (Number) - volt
* `voltage` (String)

