---
page_title: "gigavuecore_tool Resource - gigavuecore"
subcategory: ""
description: |-
  Find Inline Tool by alias
---

# gigavuecore_tool Resource

Find Inline Tool by alias

## Example Usage

```terraform
resource "gigavuecore_tool" "example" {
  alias                      = "example"
  cluster_id                 = "example"
  combined_heart_beat_status = "example"
  comment                    = "example"
  enabled                    = true
  failover_action            = "example"
  flex_status                = "example"
  flex_traffic_path          = "example"
  health_state               = "example"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "example"
    traffic_health_state_computation_type = "example"
  }]
  heartbeat = {
    enabled      = true
    ip_address_a = "example"
    ip_address_b = "example"
    profile      = "example"
    status = {
      heartbeat_passing = "example"
      stats_ato_b = {
        rx_count = 0
        tx_count = 0
      }
      stats_bto_a = {
        rx_count = 0
        tx_count = 0
      }
    }
  }
  inline_tool_type = "example"
  negative_heartbeat = {
    enabled = true
    profile = "example"
    status = {
      heartbeat_passing = "example"
      stats_ato_b = {
        rx_count = 0
        tx_count = 0
      }
      stats_bto_a = {
        rx_count = 0
        tx_count = 0
      }
    }
  }
  operational_state = "example"
  port_a            = "example"
  port_a_status     = "example"
  port_b            = "example"
  port_b_status     = "example"
  recovery_mode     = "example"
  shared            = true
  timestamp         = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Inline Tool alias. Unique within a cluster
* `cluster_id` (String, required) - Target Cluster ID
* `combined_heart_beat_status` (String, optional) - combined Heartbeat status
* `comment` (String, optional)
* `enabled` (Boolean, optional)
* `failover_action` (String, optional)
* `flex_status` (String, optional)
* `flex_traffic_path` (String, optional)
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `heartbeat` (Attributes, optional) - Embedded Heartbeat configuration for an Inline Tool. Private class (see [below for nested schema](#nestedatt--heartbeat))
* `inline_tool_type` (String, optional) - inlineTool type
* `negative_heartbeat` (Attributes, optional) - Embedded Heartbeat configuration for an Inline Tool. Private class (see [below for nested schema](#nestedatt--negative_heartbeat))
* `operational_state` (String, optional) - Operational State
* `port_a` (String, required) - portId of side A inline tool port
* `port_a_status` (String, optional) - port A status
* `port_b` (String, required) - portId of side B inline tool port
* `port_b_status` (String, optional) - port B status
* `recovery_mode` (String, optional)
* `shared` (Boolean, optional) - inline tool sharing mode
* `timestamp` (String, optional) - date-time of stats collection in RFC 3339 format

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `combined_heart_beat_status` (String, computed) - combined Heartbeat status
* `comment` (String, computed)
* `enabled` (Boolean, computed)
* `failover_action` (String, computed)
* `flex_status` (String, computed)
* `flex_traffic_path` (String, computed)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `heartbeat` (Attributes, computed) - Embedded Heartbeat configuration for an Inline Tool. Private class (see [below for nested schema](#nestedatt--heartbeat))
* `inline_tool_type` (String, computed) - inlineTool type
* `negative_heartbeat` (Attributes, computed) - Embedded Heartbeat configuration for an Inline Tool. Private class (see [below for nested schema](#nestedatt--negative_heartbeat))
* `operational_state` (String, computed) - Operational State
* `port_a_status` (String, computed) - port A status
* `port_b_status` (String, computed) - port B status
* `recovery_mode` (String, computed)
* `shared` (Boolean, computed) - inline tool sharing mode
* `timestamp` (String, computed) - date-time of stats collection in RFC 3339 format

<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--heartbeat"></a>
### Nested Schema for `heartbeat`

Required:

* `enabled` (Boolean)
* `ip_address_a` (String) - the destination IP address to be used in heartbeat packets send from side A to side B (the default is N.N.N.N where N is the port number within the chassis as shown on the face plate)
* `ip_address_b` (String) - the destination IP address to be used in heartbeat packets send from side B to side A (the default is N.N.N.N where N is the port number within the chassis as shown on the face plate)
* `profile` (String) - Alias of referenced Heartbeat Profile. the default is the heartbeat profile named 'default'
Optional:

* `status` (Attributes) - Embedded Heartbeat configuration for an Inline Tool. Private class (see [below for nested schema](#nestedatt--heartbeat--status))
<a id="nestedatt--heartbeat--status"></a>
### Nested Schema for `heartbeat.status`

Required:

* `stats_ato_b` (Attributes) - Embedded Heartbeat configuration for an Inline Tool. Private class (see [below for nested schema](#nestedatt--heartbeat--status--stats_ato_b))
* `stats_bto_a` (Attributes) - Embedded Heartbeat configuration for an Inline Tool. Private class (see [below for nested schema](#nestedatt--heartbeat--status--stats_bto_a))
Optional:

* `heartbeat_passing` (String) - indicates whether heartbeat packets from portA are reaching portB and vice versa
<a id="nestedatt--heartbeat--status--stats_ato_b"></a>
### Nested Schema for `heartbeat.status.stats_ato_b`

Required:

* `rx_count` (Number) - number of heartbeat packets received
* `tx_count` (Number) - number of heartbeat packets sent
<a id="nestedatt--heartbeat--status--stats_bto_a"></a>
### Nested Schema for `heartbeat.status.stats_bto_a`

Required:

* `rx_count` (Number) - number of heartbeat packets received
* `tx_count` (Number) - number of heartbeat packets sent
<a id="nestedatt--negative_heartbeat"></a>
### Nested Schema for `negative_heartbeat`

Required:

* `enabled` (Boolean)
Optional:

* `profile` (String) - Alias of referenced Negative Heartbeat Profile
* `status` (Attributes) - Embedded Heartbeat configuration for an Inline Tool. Private class (see [below for nested schema](#nestedatt--negative_heartbeat--status))
<a id="nestedatt--negative_heartbeat--status"></a>
### Nested Schema for `negative_heartbeat.status`

Required:

* `stats_ato_b` (Attributes) - Embedded Heartbeat configuration for an Inline Tool. Private class (see [below for nested schema](#nestedatt--negative_heartbeat--status--stats_ato_b))
* `stats_bto_a` (Attributes) - Embedded Heartbeat configuration for an Inline Tool. Private class (see [below for nested schema](#nestedatt--negative_heartbeat--status--stats_bto_a))
Optional:

* `heartbeat_passing` (String) - indicates whether heartbeat packets from portA are reaching portB and vice versa
<a id="nestedatt--negative_heartbeat--status--stats_ato_b"></a>
### Nested Schema for `negative_heartbeat.status.stats_ato_b`

Required:

* `rx_count` (Number) - number of heartbeat packets received
* `tx_count` (Number) - number of heartbeat packets sent
<a id="nestedatt--negative_heartbeat--status--stats_bto_a"></a>
### Nested Schema for `negative_heartbeat.status.stats_bto_a`

Required:

* `rx_count` (Number) - number of heartbeat packets received
* `tx_count` (Number) - number of heartbeat packets sent

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_tool.example {alias}
```
