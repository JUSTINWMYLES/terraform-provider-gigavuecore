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
  cluster_id = null
  node_id = null
  slot_id = null
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
* `fabric_hash_adv` (Bool, computed) - Advanced Fabric Hash. Supported only for Q02X32/Q08 cards
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `hw_revision` (String, computed)
* `hw_type` (String, computed)
* `mode` (String, computed)
* `oper_status` (String, computed)
* `pld_info` (Object({need_upgrade, pld_revision}), computed)
  * `need_upgrade` (Bool, computed)
  * `pld_revision` (String, computed) - PLD revision
* `power_priority` (Number, computed) - Power slot Priority
* `power_req` (Number, computed) - watt
* `product_code` (String, computed)
* `serial_number` (String, computed)
* `temperatures` (Object({board, bottom_switch, bottom_switch_major, bottom_switch_minor, bottom_switch_shut, cav_cpu, cpu, cpu_major, cpu_minor, cpu_shut, e1_cpu, e1_port, e2_cpu, exhaust, intake, near_cpu, near_cpu_major, near_cpu_minor, near_cpu_shut, qsfp_cage, qsfp_cage_major, qsfp_cage_minor, qsfp_cage_shut, rear_panel, rear_panel_major, rear_panel_minor, rear_panel_shut, sfp_cage, sfp_cage_major, sfp_cage_minor, sfp_cage_shut, switch, switch_major, switch_minor, switch_shut, time_stamp, top_switch, top_switch_major, top_switch_minor, top_switch_shut}), computed)
  * `board` (Number, computed) - celsius
  * `bottom_switch` (Number, computed) - celsius
  * `bottom_switch_major` (Number, computed) - celsius
  * `bottom_switch_minor` (Number, computed) - celsius
  * `bottom_switch_shut` (Number, computed) - celsius
  * `cav_cpu` (Number, computed) - celsius
  * `cpu` (Number, computed) - celsius
  * `cpu_major` (Number, computed) - celsius
  * `cpu_minor` (Number, computed) - celsius
  * `cpu_shut` (Number, computed) - celsius
  * `e1_cpu` (Number, computed) - celsius
  * `e1_port` (Number, computed) - celsius
  * `e2_cpu` (Number, computed) - celsius
  * `exhaust` (Number, computed) - celsius
  * `intake` (Number, computed) - celsius
  * `near_cpu` (Number, computed) - celsius
  * `near_cpu_major` (Number, computed) - celsius
  * `near_cpu_minor` (Number, computed) - celsius
  * `near_cpu_shut` (Number, computed) - celsius
  * `qsfp_cage` (Number, computed) - celsius
  * `qsfp_cage_major` (Number, computed) - celsius
  * `qsfp_cage_minor` (Number, computed) - celsius
  * `qsfp_cage_shut` (Number, computed) - celsius
  * `rear_panel` (Number, computed) - celsius
  * `rear_panel_major` (Number, computed) - celsius
  * `rear_panel_minor` (Number, computed) - celsius
  * `rear_panel_shut` (Number, computed) - celsius
  * `sfp_cage` (Number, computed) - celsius
  * `sfp_cage_major` (Number, computed) - celsius
  * `sfp_cage_minor` (Number, computed) - celsius
  * `sfp_cage_shut` (Number, computed) - celsius
  * `switch` (Number, computed) - celsius
  * `switch_major` (Number, computed) - celsius
  * `switch_minor` (Number, computed) - celsius
  * `switch_shut` (Number, computed) - celsius
  * `time_stamp` (String, computed) - date-time of stats collection in RFC 3339 format
  * `top_switch` (Number, computed) - celsius
  * `top_switch_major` (Number, computed) - celsius
  * `top_switch_minor` (Number, computed) - celsius
  * `top_switch_shut` (Number, computed) - celsius
* `voltages` (List(Object({time_stamp, value, voltage})), computed)

