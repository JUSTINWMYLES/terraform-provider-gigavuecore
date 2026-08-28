---
page_title: "gigavuecore_load_all_inline_tools Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Inline Tools
---

# gigavuecore_load_all_inline_tools Data Source

Load all Inline Tools

## Example Usage

```terraform
data "gigavuecore_load_all_inline_tools" "example" {
  cluster_id = null
  page       = null
  sort       = null
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

* `alias` (String) - Inline Tool alias. Unique within a cluster
* `combined_heart_beat_status` (String) - combined Heartbeat status
* `comment` (String)
* `enabled` (Boolean)
* `failover_action` (String)
* `flex_status` (String)
* `flex_traffic_path` (String)
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--health_state_reasons))
* `heartbeat` (Attributes) - Embedded Heartbeat configuration for an Inline Tool. Private class (see [below for nested schema](#nestedatt--items--heartbeat))
* `inline_tool_type` (String) - inlineTool type
* `negative_heartbeat` (Attributes) - Embedded Heartbeat configuration for an Inline Tool. Private class (see [below for nested schema](#nestedatt--items--negative_heartbeat))
* `operational_state` (String) - Operational State
* `port_a` (String) - portId of side A inline tool port
* `port_a_status` (String) - port A status
* `port_b` (String) - portId of side B inline tool port
* `port_b_status` (String) - port B status
* `recovery_mode` (String)
* `shared` (Boolean) - inline tool sharing mode
* `timestamp` (String) - date-time of stats collection in RFC 3339 format
<a id="nestedatt--items--health_state_reasons"></a>
### Nested Schema for `items.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--items--heartbeat"></a>
### Nested Schema for `items.heartbeat`

Read-Only:

* `enabled` (Boolean)
* `ip_address_a` (String) - the destination IP address to be used in heartbeat packets send from side A to side B (the default is N.N.N.N where N is the port number within the chassis as shown on the face plate)
* `ip_address_b` (String) - the destination IP address to be used in heartbeat packets send from side B to side A (the default is N.N.N.N where N is the port number within the chassis as shown on the face plate)
* `profile` (String) - Alias of referenced Heartbeat Profile. the default is the heartbeat profile named 'default'
* `status` (Attributes) - Embedded Heartbeat configuration for an Inline Tool. Private class (see [below for nested schema](#nestedatt--items--heartbeat--status))
<a id="nestedatt--items--heartbeat--status"></a>
### Nested Schema for `items.heartbeat.status`

Read-Only:

* `heartbeat_passing` (String) - indicates whether heartbeat packets from portA are reaching portB and vice versa
* `stats_ato_b` (Attributes) - Embedded Heartbeat configuration for an Inline Tool. Private class (see [below for nested schema](#nestedatt--items--heartbeat--status--stats_ato_b))
* `stats_bto_a` (Attributes) - Embedded Heartbeat configuration for an Inline Tool. Private class (see [below for nested schema](#nestedatt--items--heartbeat--status--stats_bto_a))
<a id="nestedatt--items--heartbeat--status--stats_ato_b"></a>
### Nested Schema for `items.heartbeat.status.stats_ato_b`

Read-Only:

* `rx_count` (Number) - number of heartbeat packets received
* `tx_count` (Number) - number of heartbeat packets sent
<a id="nestedatt--items--heartbeat--status--stats_bto_a"></a>
### Nested Schema for `items.heartbeat.status.stats_bto_a`

Read-Only:

* `rx_count` (Number) - number of heartbeat packets received
* `tx_count` (Number) - number of heartbeat packets sent
<a id="nestedatt--items--negative_heartbeat"></a>
### Nested Schema for `items.negative_heartbeat`

Read-Only:

* `enabled` (Boolean)
* `profile` (String) - Alias of referenced Negative Heartbeat Profile
* `status` (Attributes) - Embedded Heartbeat configuration for an Inline Tool. Private class (see [below for nested schema](#nestedatt--items--negative_heartbeat--status))
<a id="nestedatt--items--negative_heartbeat--status"></a>
### Nested Schema for `items.negative_heartbeat.status`

Read-Only:

* `heartbeat_passing` (String) - indicates whether heartbeat packets from portA are reaching portB and vice versa
* `stats_ato_b` (Attributes) - Embedded Heartbeat configuration for an Inline Tool. Private class (see [below for nested schema](#nestedatt--items--negative_heartbeat--status--stats_ato_b))
* `stats_bto_a` (Attributes) - Embedded Heartbeat configuration for an Inline Tool. Private class (see [below for nested schema](#nestedatt--items--negative_heartbeat--status--stats_bto_a))
<a id="nestedatt--items--negative_heartbeat--status--stats_ato_b"></a>
### Nested Schema for `items.negative_heartbeat.status.stats_ato_b`

Read-Only:

* `rx_count` (Number) - number of heartbeat packets received
* `tx_count` (Number) - number of heartbeat packets sent
<a id="nestedatt--items--negative_heartbeat--status--stats_bto_a"></a>
### Nested Schema for `items.negative_heartbeat.status.stats_bto_a`

Read-Only:

* `rx_count` (Number) - number of heartbeat packets received
* `tx_count` (Number) - number of heartbeat packets sent

