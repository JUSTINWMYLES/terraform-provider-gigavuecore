---
page_title: "gigavuecore_get_giga_port_by_port_id Data Source - gigavuecore"
subcategory: ""
description: |-
  Find Device Port by portId
---

# gigavuecore_get_giga_port_by_port_id Data Source

Find Device Port by portId

## Example Usage

```terraform
data "gigavuecore_get_giga_port_by_port_id" "example" {
  cluster_id = null
  port_id    = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `port_id` (String, required) - device port id

### Attributes

In addition to all arguments above, the following attributes are exported:

* `admin_status` (String, computed)
* `alias` (String, computed) - device port alias
* `auto_neg` (Boolean, computed) - When auto-negotiation is enabled, duplex and speed settings are ignored (they are set via auto-negotiation). Auto-negotiation is always disabled for 40Gb and 100Gb ports. For 1Gb speeds over copper, auto-negotiation must be enabled, per the IEEE 802.3 specification.
* `breakout_mode` (String, computed) - none = port not broken out; na = port not capable of breakout; 4x = 4x10G; 2q = 2x40G
* `cable_length` (String, computed) - Attached cable length in meter
* `comment` (String, computed)
* `config_speed` (String, computed) - The configured line speed of a port. Only applicable for copper ports. Only applicable if auto-negotiation is off
* `duplex` (String, computed) - Only applicable for 10M/100M operations. Only applicable if auto-negotiation is off. Duplex is always set to full for 10G ports. If this parameter is set explicitly on one end of the connection, it must also be set explicitly on the other end. Duplex mismatches will occur if the duplex setting are forced on one end of the connection while the other end attempts to auto-negotiate settings
* `force_link_up` (Boolean, computed) - Forces connection on an optical port. Use this option when an optical GigaPORT tool port is connected to a legacy optical tool that does not transmit light; Available for optical 1Gb/10Gb tool ports; Not available for 10Gb-capable ports with a 1Gb SFP installed. 10Gb-capable optical tool ports only support force-linkup when a 10Gb SFP+ is installed.
* `force_link_up_status` (String, computed) - Forcelinkup state for the port configured
* `gsparams` (Attributes, computed) - Gsparams Resource HW limit for engine port (see [below for nested schema](#nestedatt--gsparams))
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `host_name` (String, computed) - device host name
* `is_signal_detected` (String, computed)
* `last_state_transition_time` (Number, computed) - Port state change timestamp in milliseconds
* `mtu` (Number, computed)
* `oper_speed` (String, computed) - Port operational runtime speed
* `oper_status` (String, computed)
* `port_role` (String, computed) - master/slave role of port. (deprecated: use portRoleAlias)
* `port_role_alias` (String, computed) - source/receiver role of port
* `port_type` (String, computed) - Configures port type for eligible ports based on the underlying card/chassis hardware type. 'gigasmart' port type can never be explicitly assigned
* `sfp` (Attributes, computed) - Port SFP details (see [below for nested schema](#nestedatt--sfp))
* `ude` (Attributes, computed) - Unidirectional Ethernet (see [below for nested schema](#nestedatt--ude))

<a id="nestedatt--gsparams"></a>
### Nested Schema for `gsparams`

Read-Only:

* `bufferasf` (Attributes) (see [below for nested schema](#nestedatt--gsparams--bufferasf))
* `metadata` (Attributes) (see [below for nested schema](#nestedatt--gsparams--metadata))
<a id="nestedatt--gsparams--bufferasf"></a>
### Nested Schema for `gsparams.bufferasf`

Read-Only:

* `initial` (Number)
* `max` (Number)
* `min` (Number)
<a id="nestedatt--gsparams--metadata"></a>
### Nested Schema for `gsparams.metadata`

Read-Only:

* `initial` (Number)
* `max` (Number)
* `min` (Number)
<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--sfp"></a>
### Nested Schema for `sfp`

Read-Only:

* `sfp_type` (String)
* `vendor_name` (String)
* `vendor_pn` (String)
* `vendor_sn` (String)
<a id="nestedatt--ude"></a>
### Nested Schema for `ude`

Read-Only:

* `enabled` (Boolean) - Only applicable if 100g-bidi is detected

