---
page_title: "gigavuecore_load_giga_ports Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Device Ports
---

# gigavuecore_load_giga_ports Data Source

Load all Device Ports

## Example Usage

```terraform
data "gigavuecore_load_giga_ports" "example" {
  cluster_id = "example"
  page       = "example"
  sort       = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

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
* `gsparams` (Attributes) - Gsparams Resource HW limit for engine port (see [below for nested schema](#nestedatt--items--gsparams))
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--health_state_reasons))
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
* `sfp` (Attributes) - Port SFP details (see [below for nested schema](#nestedatt--items--sfp))
* `ude` (Attributes) - Unidirectional Ethernet (see [below for nested schema](#nestedatt--items--ude))
<a id="nestedatt--items--gsparams"></a>
### Nested Schema for `items.gsparams`

Read-Only:

* `bufferasf` (Attributes) (see [below for nested schema](#nestedatt--items--gsparams--bufferasf))
* `metadata` (Attributes) (see [below for nested schema](#nestedatt--items--gsparams--metadata))
<a id="nestedatt--items--gsparams--bufferasf"></a>
### Nested Schema for `items.gsparams.bufferasf`

Read-Only:

* `initial` (Number)
* `max` (Number)
* `min` (Number)
<a id="nestedatt--items--gsparams--metadata"></a>
### Nested Schema for `items.gsparams.metadata`

Read-Only:

* `initial` (Number)
* `max` (Number)
* `min` (Number)
<a id="nestedatt--items--health_state_reasons"></a>
### Nested Schema for `items.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--items--sfp"></a>
### Nested Schema for `items.sfp`

Read-Only:

* `sfp_type` (String)
* `vendor_name` (String)
* `vendor_pn` (String)
* `vendor_sn` (String)
<a id="nestedatt--items--ude"></a>
### Nested Schema for `items.ude`

Read-Only:

* `enabled` (Boolean) - Only applicable if 100g-bidi is detected

