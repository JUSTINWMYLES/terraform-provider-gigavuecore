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
  node_id    = null
  page       = null
  ports      = null
  sort       = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - Target Cluster ID. Either 'clusterId' or 'nodeId' is required
* `node_id` (String, optional) - ID of the target device. Either 'clusterId' or 'nodeId' is required
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `ports` (Boolean, optional) - Indicates whether per-chassis port list should be included in the response
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `chassis_list` (Attributes List, computed) (see [below for nested schema](#nestedatt--chassis_list))
* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))

<a id="nestedatt--chassis_list"></a>
### Nested Schema for `chassis_list`

Read-Only:

* `boot_time` (String) - the time system booted up. In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'
* `box_id` (String)
* `build_id` (String)
* `cards` (Attributes List) (see [below for nested schema](#nestedatt--chassis_list--cards))
* `cc_sync_status` (String)
* `chassis_id` (String)
* `cpld_version` (String)
* `cpu_load_avg` (Attributes) - CPU Load Average (see [below for nested schema](#nestedatt--chassis_list--cpu_load_avg))
* `device_name` (String) - maps to hostname
* `fan_trays` (Attributes List) (see [below for nested schema](#nestedatt--chassis_list--fan_trays))
* `file_system_statistics` (Attributes) - File System Statistics (see [below for nested schema](#nestedatt--chassis_list--file_system_statistics))
* `gdp` (Boolean) - enable/disable GDP for chassis
* `git_full_hash` (String)
* `global_node_id` (String)
* `hardware_info` (Attributes) - hardware info  of dell box (see [below for nested schema](#nestedatt--chassis_list--hardware_info))
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--chassis_list--health_state_reasons))
* `host_id` (String)
* `hw_revision` (String)
* `hw_type` (String)
* `l2_gre_id` (Number) - Global l2gre id configured, Value of 0 disables the vxlanId
* `leaf_config` (Attributes) - Configuration of Leaf node for Spine-Link (see [below for nested schema](#nestedatt--chassis_list--leaf_config))
* `memory` (Attributes) - Device Memory (see [below for nested schema](#nestedatt--chassis_list--memory))
* `mode` (String) - Use of 100G ports requires chassis mode value to be 100G. 100G applicable to HC2-v2 only. 100GLeft, Right refers to left, right side chassis bank.
* `model` (String) - Gigamon physical device models
* `oper_status` (String)
* `ports` (Attributes List) - details of the ports on this chassis (see [below for nested schema](#nestedatt--chassis_list--ports))
* `power_management` (Attributes) - Power Manager Information (see [below for nested schema](#nestedatt--chassis_list--power_management))
* `power_modules` (Attributes List) (see [below for nested schema](#nestedatt--chassis_list--power_modules))
* `product_build_date` (String)
* `product_code` (String)
* `product_name` (String)
* `product_version` (String)
* `serial_number` (String)
* `sys_contact` (String)
* `sys_descr` (String)
* `sys_location` (String)
* `system_mode` (String) - 'operational': fully configurable; 'safe': no configuration possible; 'limited': limited configuration possible
* `ts_version` (Number)
* `uboot_version` (String)
* `version_summary` (String)
* `vxlan_id` (Number) - Global vxlan id configured , Value of 0 disables the vxlanId
<a id="nestedatt--chassis_list--cards"></a>
### Nested Schema for `chassis_list.cards`

Read-Only:

* `admin_status` (String)
* `alarm_buffer_threshold` (Number) - card micro burst threshold
* `config_status` (String)
* `fabric_hash_adv` (Boolean) - Advanced Fabric Hash. Supported only for Q02X32/Q08 cards
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--chassis_list--cards--health_state_reasons))
* `hw_revision` (String)
* `hw_type` (String)
* `mode` (String)
* `oper_status` (String)
* `pld_info` (Attributes) (see [below for nested schema](#nestedatt--chassis_list--cards--pld_info))
* `power_priority` (Number) - Power slot Priority
* `power_req` (Number) - watt
* `product_code` (String)
* `serial_number` (String)
* `slot_id` (String) - device card slot id
* `temperatures` (Attributes) (see [below for nested schema](#nestedatt--chassis_list--cards--temperatures))
* `voltages` (Attributes List) (see [below for nested schema](#nestedatt--chassis_list--cards--voltages))
<a id="nestedatt--chassis_list--cards--health_state_reasons"></a>
### Nested Schema for `chassis_list.cards.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--chassis_list--cards--pld_info"></a>
### Nested Schema for `chassis_list.cards.pld_info`

Read-Only:

* `need_upgrade` (Boolean)
* `pld_revision` (String) - PLD revision
<a id="nestedatt--chassis_list--cards--temperatures"></a>
### Nested Schema for `chassis_list.cards.temperatures`

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
<a id="nestedatt--chassis_list--cards--voltages"></a>
### Nested Schema for `chassis_list.cards.voltages`

Read-Only:

* `time_stamp` (String) - date-time of stats collection in RFC 3339 format
* `value` (Number) - volt
* `voltage` (String)
<a id="nestedatt--chassis_list--cpu_load_avg"></a>
### Nested Schema for `chassis_list.cpu_load_avg`

Read-Only:

* `load1` (Number) - cpu load averages for past 1 minute
* `load15` (Number) - cpu load averages for past 15 minutes
* `load5` (Number) - cpu load averages for past 5 minutes
* `time_stamp` (String) - date-time of stats collection in RFC 3339 format
<a id="nestedatt--chassis_list--fan_trays"></a>
### Nested Schema for `chassis_list.fan_trays`

Read-Only:

* `fan_details` (Attributes List) (see [below for nested schema](#nestedatt--chassis_list--fan_trays--fan_details))
* `fantray_id` (String) - chassis Fan Tray id
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--chassis_list--fan_trays--health_state_reasons))
* `hw_revision` (String)
* `hw_type` (String)
* `oper_status` (String)
* `product_code` (String)
* `serial_number` (String)
<a id="nestedatt--chassis_list--fan_trays--fan_details"></a>
### Nested Schema for `chassis_list.fan_trays.fan_details`

Read-Only:

* `fan_id` (String) - id of a fan in a FanTray
* `fan_speed_rpm` (Number) - fan speed in rpm
* `time_stamp` (String) - date-time of stats collection in RFC 3339 format
<a id="nestedatt--chassis_list--fan_trays--health_state_reasons"></a>
### Nested Schema for `chassis_list.fan_trays.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--chassis_list--file_system_statistics"></a>
### Nested Schema for `chassis_list.file_system_statistics`

Read-Only:

* `config_statistics` (Attributes) - File System Statistics (see [below for nested schema](#nestedatt--chassis_list--file_system_statistics--config_statistics))
* `time_stamp` (String) - date-time of stats collection in RFC 3339 format
* `var_statistics` (Attributes) - File System Statistics (see [below for nested schema](#nestedatt--chassis_list--file_system_statistics--var_statistics))
<a id="nestedatt--chassis_list--file_system_statistics--config_statistics"></a>
### Nested Schema for `chassis_list.file_system_statistics.config_statistics`

Read-Only:

* `inodes_percent_free` (Number) - percent
* `space_available` (Number) - MB
* `space_free` (Number) - MB
* `space_percent_free` (Number) - percent
* `space_total` (Number) - MB
* `space_used` (Number) - MB
<a id="nestedatt--chassis_list--file_system_statistics--var_statistics"></a>
### Nested Schema for `chassis_list.file_system_statistics.var_statistics`

Read-Only:

* `inodes_percent_free` (Number) - percent
* `space_available` (Number) - MB
* `space_free` (Number) - MB
* `space_percent_free` (Number) - percent
* `space_total` (Number) - MB
* `space_used` (Number) - MB
<a id="nestedatt--chassis_list--hardware_info"></a>
### Nested Schema for `chassis_list.hardware_info`

Read-Only:

* `base_mac_address` (String) - macAddress
* `device_version` (String) - device version
* `diag_version` (String) - diag version
* `label_revision` (String) - labelRevision of dell box
* `manufacture_date` (String) - loader version
* `part_number` (String) - part number
* `platform_name` (String) - platformName
* `product_name` (String) - product name
* `service_tag` (String) - service tag
<a id="nestedatt--chassis_list--health_state_reasons"></a>
### Nested Schema for `chassis_list.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--chassis_list--leaf_config"></a>
### Nested Schema for `chassis_list.leaf_config`

Read-Only:

* `mode` (String)
<a id="nestedatt--chassis_list--memory"></a>
### Nested Schema for `chassis_list.memory`

Read-Only:

* `free` (Number) - MB
* `time_stamp` (String) - date-time of stats collection in RFC 3339 format
* `total` (Number) - MB
* `used` (Number) - MB
<a id="nestedatt--chassis_list--ports"></a>
### Nested Schema for `chassis_list.ports`

Read-Only:

* `admin_status` (String)
* `alias` (String) - device port alias
* `auto_neg` (Boolean) - When auto-negotiation is enabled, duplex and speed settings are ignored (they are set via auto-negotiation). Auto-negotiation is always disabled for 40Gb and 100Gb ports. For 1Gb speeds over copper, auto-negotiation must be enabled, per the IEEE 802.3 specification.
* `breakout_mode` (String) - none = port not broken out; na = port not capable of breakout; 4x = 4x10G; 2q = 2x40G
* `cable_length` (String) - Attached cable length in meter
* `comment` (String)
* `config_speed` (String) - The configured line speed of a port. Only applicable for copper ports. Only applicable if auto-negotiation is off
* `duplex` (String) - Only applicable for 10M/100M operations. Only applicable if auto-negotiation is off. Duplex is always set to full for 10G ports. If this parameter is set explicitly on one end of the connection, it must also be set explicitly on the other end. Duplex mismatches will occur if the duplex setting are forced on one end of the connection while the other end attempts to auto-negotiate settings
* `force_link_up` (Boolean) - Forces connection on an optical port. Use this option when an optical GigaPORT tool port is connected to a legacy optical tool that does not transmit light; Available for optical 1Gb/10Gb tool ports; Not available for 10Gb-capable ports with a 1Gb SFP installed. 10Gb-capable optical tool ports only support force-linkup when a 10Gb SFP+ is installed.
* `force_link_up_status` (String) - Forcelinkup state for the port configured
* `gsparams` (Attributes) - Gsparams Resource HW limit for engine port (see [below for nested schema](#nestedatt--chassis_list--ports--gsparams))
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--chassis_list--ports--health_state_reasons))
* `host_name` (String) - device host name
* `is_signal_detected` (String)
* `last_state_transition_time` (Number) - Port state change timestamp in milliseconds
* `mtu` (Number)
* `oper_speed` (String) - Port operational runtime speed
* `oper_status` (String)
* `port_id` (String) - device port id
* `port_role` (String) - master/slave role of port. (deprecated: use portRoleAlias)
* `port_role_alias` (String) - source/receiver role of port
* `port_type` (String) - Configures port type for eligible ports based on the underlying card/chassis hardware type. 'gigasmart' port type can never be explicitly assigned
* `sfp` (Attributes) - Port SFP details (see [below for nested schema](#nestedatt--chassis_list--ports--sfp))
* `ude` (Attributes) - Unidirectional Ethernet (see [below for nested schema](#nestedatt--chassis_list--ports--ude))
<a id="nestedatt--chassis_list--ports--gsparams"></a>
### Nested Schema for `chassis_list.ports.gsparams`

Read-Only:

* `bufferasf` (Attributes) (see [below for nested schema](#nestedatt--chassis_list--ports--gsparams--bufferasf))
* `metadata` (Attributes) (see [below for nested schema](#nestedatt--chassis_list--ports--gsparams--metadata))
<a id="nestedatt--chassis_list--ports--gsparams--bufferasf"></a>
### Nested Schema for `chassis_list.ports.gsparams.bufferasf`

Read-Only:

* `initial` (Number)
* `max` (Number)
* `min` (Number)
<a id="nestedatt--chassis_list--ports--gsparams--metadata"></a>
### Nested Schema for `chassis_list.ports.gsparams.metadata`

Read-Only:

* `initial` (Number)
* `max` (Number)
* `min` (Number)
<a id="nestedatt--chassis_list--ports--health_state_reasons"></a>
### Nested Schema for `chassis_list.ports.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--chassis_list--ports--sfp"></a>
### Nested Schema for `chassis_list.ports.sfp`

Read-Only:

* `sfp_type` (String)
* `vendor_name` (String)
* `vendor_pn` (String)
* `vendor_sn` (String)
<a id="nestedatt--chassis_list--ports--ude"></a>
### Nested Schema for `chassis_list.ports.ude`

Read-Only:

* `enabled` (Boolean) - Only applicable if 100g-bidi is detected
<a id="nestedatt--chassis_list--power_management"></a>
### Nested Schema for `chassis_list.power_management`

Read-Only:

* `power_allocated` (Number) - watts
* `power_available` (Number) - watts
* `power_redundancy` (String)
* `total_base_system_power` (Number) - watts
* `total_power` (Number) - watts
* `total_redundant_power` (Number) - watts
<a id="nestedatt--chassis_list--power_modules"></a>
### Nested Schema for `chassis_list.power_modules`

Read-Only:

* `capacity` (Number) - watts
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--chassis_list--power_modules--health_state_reasons))
* `hw_revision` (String)
* `hw_type` (String)
* `oper_status` (String) - applicable for non-HC2 devices
* `oper_status_bottom` (String) - applicable for HC2 devices
* `oper_status_top` (String) - applicable for HC2 devices
* `power_module_id` (String) - chassis Power Module id
* `power_voltages` (Attributes List) (see [below for nested schema](#nestedatt--chassis_list--power_modules--power_voltages))
* `product_code` (String)
* `serial_number` (String)
<a id="nestedatt--chassis_list--power_modules--health_state_reasons"></a>
### Nested Schema for `chassis_list.power_modules.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--chassis_list--power_modules--power_voltages"></a>
### Nested Schema for `chassis_list.power_modules.power_voltages`

Read-Only:

* `time_stamp` (String) - date-time of stats collection in RFC 3339 format
* `value` (Number) - volt
* `voltage` (String)
<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type

