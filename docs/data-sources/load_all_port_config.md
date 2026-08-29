---
page_title: "gigavuecore_load_all_port_config Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all PortConfigs
---

# gigavuecore_load_all_port_config Data Source

Load all PortConfigs

## Example Usage

```terraform
data "gigavuecore_load_all_port_config" "example" {
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

* `access_roles` (Attributes) - Port Access RBAC definitions (see [below for nested schema](#nestedatt--items--access_roles))
* `alarm_thresholds` (Attributes) - Port Alarm Thresholds definitions (see [below for nested schema](#nestedatt--items--alarm_thresholds))
* `fec` (String) - enable/disable forward error correction
* `gdp` (Boolean) - enable/disable GDP packets on port
* `header_strip` (String) - protocol type
* `ingress_vlan_tag` (Number) - Ingress port VLAN tag, with valid numbers between 2 and 4000 that is added to a packet. Setting the value to 0 disables VLAN tagging. The port must be a network port
* `l2_gre_id` (Number) - Value of 0 disables the l2greId
* `licensed` (Boolean) - for TA series indicates whether a port is licensed. Defaults to 'true'
* `lock` (Attributes) - Port Locking definitions (see [below for nested schema](#nestedatt--items--lock))
* `mpls_advanced` (Attributes) - Select a combination of Mpls-Advanced options (see [below for nested schema](#nestedatt--items--mpls_advanced))
* `neighbor_discovery` (String) - Configures port neighbor discovery options
* `port_id` (String) - device port id
* `ptp` (Attributes) - Port PTP configurations (see [below for nested schema](#nestedatt--items--ptp))
* `share` (Attributes) - Port Sharing definitions (see [below for nested schema](#nestedatt--items--share))
* `tag_protocol_id` (String) - When Ingress port VLAN tag is added , this protocol Id will be added which egress out the traffic of tool port
* `taptx` (String) - Opens or closes the copper tap port relay. Active closes the port relay. Passive opens the port relay
* `timestamp` (Attributes) - Timestamping definitions for GigaPORT-X12-TS ports (see [below for nested schema](#nestedatt--items--timestamp))
* `vxlan_id` (Number) - Value of 0 disables the vxlanId
<a id="nestedatt--items--access_roles"></a>
### Nested Schema for `items.access_roles`

Read-Only:

* `level1` (List of String) - list of roles with 'Level 1' access privileges. Level 1 - Read-only access. Can view port configuration and statistics
* `level2` (List of String) - list of roles with 'Level 2' access privileges. Level 2 - Level 1 plus the capability to configure port-lock, lock-share, and all traffic objects except port-pair
* `level3` (List of String) - list of roles with 'Level 3' access privileges. Level 3 - Level 2 plus the capability to configure port params (such as administrative status of the port, speed, duplex, and auto-negotiation), as well as port-pair
* `level4` (List of String) - list of roles with 'Level 4' access privileges. Level 4 - Level 3 plus the capability to change the port type
<a id="nestedatt--items--alarm_thresholds"></a>
### Nested Schema for `items.alarm_thresholds`

Read-Only:

* `alarm_buffer_threshold_rx` (Number) - in percent's. Alarm is reported every time threshold is crossed. Value of 0 disables the threshold
* `alarm_buffer_threshold_tx` (Number) - in percent's. Alarm is reported every time threshold is crossed. Value of 0 disables the threshold
* `alarm_threshold` (Number) - in percent's. Alarm is reported every time threshold is crossed. The threshold must be exceeded for at least six consecutive seconds. Value of 0 disables the threshold
* `alarm_threshold_low` (Number) - in percents. Alarm is reported every time this low threshold is crossed. It must be under this threshold for at least six consecutive seconds. Value of 0 disables the threshold
<a id="nestedatt--items--lock"></a>
### Nested Schema for `items.lock`

Read-Only:

* `description` (String) - Optional lock description string
* `locking_user` (String) - User account name this lock belongs to. Lock is used to restrict use of the port by only the specified user account. Users with Default/Operator privileges can only lock ports to which their account has been granted access under their own user account name only. Administrators can lock any port in the system for any existing user account name
* `shared_with` (List of String) - Used to share a locked port with other user accounts. Users with Default/Operator privileges can only share a lock if the port is locked under their account name. Administrators can lock a port for another user account
<a id="nestedatt--items--mpls_advanced"></a>
### Nested Schema for `items.mpls_advanced`

Read-Only:

* `mpls_advanced_opt` (Attributes List) - List of Mpls-Advanced Options (see [below for nested schema](#nestedatt--items--mpls_advanced--mpls_advanced_opt))
<a id="nestedatt--items--mpls_advanced--mpls_advanced_opt"></a>
### Nested Schema for `items.mpls_advanced.mpls_advanced_opt`

Read-Only:

* `adv_opt` (String) - Mpls-Advanced Options
<a id="nestedatt--items--ptp"></a>
### Nested Schema for `items.ptp`

Read-Only:

* `announce_interval` (Number) - Configures the interval between PTP announce messages on an interface. The range for the PTP announcement interval is from -2 to 4 log seconds. For the domain 0 as well as domain 24 to 43 the interval limit will be of -3 to 4
* `delay_request_interval` (Number) - Configures the minimum interval allowed between PTP delay messages when the port is in the source state. The range is from log(-7) to log(5) seconds, where log(-1) = 1 frame per second. For the domain 0 the interval limit is -4 to 5 and for the range of domain 24 to 43 the interval limit will be of -7 to 4.
* `enable` (Boolean) - Enable PTP on the port
* `local_priority` (Number)
* `role` (String) - deprecated: use roleAlias
* `role_alias` (String)
* `sync_interval` (Number) - Configures the interval between PTP synchronization messages on an interface. The range is from log(-7) to log(1) seconds. For domain 0 the interval limit is -4 to 1 and for the range of domain 24 to 43 the interval limit will be of -7 to 4.
* `timestamp` (Attributes) - Port PTP timestamp configurations (see [below for nested schema](#nestedatt--items--ptp--timestamp))
* `vlan` (Number) - configures vlan on the PTP configured port
<a id="nestedatt--items--ptp--timestamp"></a>
### Nested Schema for `items.ptp.timestamp`

Read-Only:

* `egress` (Attributes) (see [below for nested schema](#nestedatt--items--ptp--timestamp--egress))
* `ingress` (Attributes) (see [below for nested schema](#nestedatt--items--ptp--timestamp--ingress))
<a id="nestedatt--items--ptp--timestamp--egress"></a>
### Nested Schema for `items.ptp.timestamp.egress`

Read-Only:

* `insert` (Boolean) - When set, will insert timestamp for packets on the port
* `source_id` (Number) - Source identifier of Packet Time Stamp record header
<a id="nestedatt--items--ptp--timestamp--ingress"></a>
### Nested Schema for `items.ptp.timestamp.ingress`

Read-Only:

* `insert` (Boolean) - When set, will insert timestamp for packets on the port
* `source_id` (Number) - Source identifier of Packet Time Stamp record header
<a id="nestedatt--items--share"></a>
### Nested Schema for `items.share`

Read-Only:

* `tool_share_roles` (List of String) - Only applicable to tool ports. Used to designate a tool port as available for tool-to-tool pass-alls (tool-mirrors) with the specified roles
<a id="nestedatt--items--timestamp"></a>
### Nested Schema for `items.timestamp`

Read-Only:

* `append_ingress` (Boolean) - Used to add a timestamp to ingress packets. Applicable to GigaPORT-X12-TS card ports x1..x12 when configured as network ports
* `source_id` (Number) - Used to specify a custom source-id to be included in the timestamp appended by the GigaPORT-X12-TS. Applicable to GigaPORT-X12-TS card ports x1..x12 when configured as network ports. The timestamp always includes a source-id field. If this custom value is not explicitly specified, the GigaPORT-X12-TS generates one automatically using the following formula: (<box-id> \* 2048) + (<slot-id> \* 256) + <port-number>
* `strip_egress` (Boolean) - Used to strip timestamps from egress packets. Use this argument to strip timestamps from egress packets. Important: this option should only be enabled to packets with time stamps appended. This function will strip the last 14 bytes of each packet regardless of whether a timestamp has been added. Applicable to GigaPORT-X12-TS card ports x9..x12 when configured as tool ports

