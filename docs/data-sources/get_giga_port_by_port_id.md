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
  port_id = null
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
* `auto_neg` (Bool, computed) - When auto-negotiation is enabled, duplex and speed settings are ignored (they are set via auto-negotiation). Auto-negotiation is always disabled for 40Gb and 100Gb ports. For 1Gb speeds over copper, auto-negotiation must be enabled, per the IEEE 802.3 specification.
* `breakout_mode` (String, computed) - none = port not broken out; na = port not capable of breakout; 4x = 4x10G; 2q = 2x40G
* `cable_length` (String, computed) - Attached cable length in meter
* `comment` (String, computed)
* `config_speed` (String, computed) - The configured line speed of a port. Only applicable for copper ports. Only applicable if auto-negotiation is off
* `duplex` (String, computed) - Only applicable for 10M/100M operations. Only applicable if auto-negotiation is off. Duplex is always set to full for 10G ports. If this parameter is set explicitly on one end of the connection, it must also be set explicitly on the other end. Duplex mismatches will occur if the duplex setting are forced on one end of the connection while the other end attempts to auto-negotiate settings
* `force_link_up` (Bool, computed) - Forces connection on an optical port. Use this option when an optical GigaPORT tool port is connected to a legacy optical tool that does not transmit light; Available for optical 1Gb/10Gb tool ports; Not available for 10Gb-capable ports with a 1Gb SFP installed. 10Gb-capable optical tool ports only support force-linkup when a 10Gb SFP+ is installed.
* `force_link_up_status` (String, computed) - Forcelinkup state for the port configured
* `gsparams` (Object({bufferasf, metadata}), computed) - Gsparams Resource HW limit for engine port
  * `bufferasf` (Object({initial, max, min}), computed)
    * `initial` (Number, computed)
    * `max` (Number, computed)
    * `min` (Number, computed)
  * `metadata` (Object({initial, max, min}), computed)
    * `initial` (Number, computed)
    * `max` (Number, computed)
    * `min` (Number, computed)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `host_name` (String, computed) - device host name
* `is_signal_detected` (String, computed)
* `last_state_transition_time` (Number, computed) - Port state change timestamp in milliseconds
* `mtu` (Number, computed)
* `oper_speed` (String, computed) - Port operational runtime speed
* `oper_status` (String, computed)
* `port_role` (String, computed) - master/slave role of port. (deprecated: use portRoleAlias)
* `port_role_alias` (String, computed) - source/receiver role of port
* `port_type` (String, computed) - Configures port type for eligible ports based on the underlying card/chassis hardware type. 'gigasmart' port type can never be explicitly assigned
* `sfp` (Object({sfp_type, vendor_name, vendor_pn, vendor_sn}), computed) - Port SFP details
  * `sfp_type` (String, computed)
  * `vendor_name` (String, computed)
  * `vendor_pn` (String, computed)
  * `vendor_sn` (String, computed)
* `ude` (Object({enabled}), computed) - Unidirectional Ethernet
  * `enabled` (Bool, computed) - Only applicable if 100g-bidi is detected

