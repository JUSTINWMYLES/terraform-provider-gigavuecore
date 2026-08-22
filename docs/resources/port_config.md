---
page_title: "gigavuecore_port_config Resource - gigavuecore"
subcategory: ""
description: |-
  Find PortConfig by portId
---

# gigavuecore_port_config Resource

Find PortConfig by portId

## Example Usage

```terraform
resource "gigavuecore_port_config" "example" {
  access_roles = {}
  alarm_thresholds = {}
  fec = null
  gdp = null
  header_strip = null
  id = null
  ingress_vlan_tag = null
  l2_gre_id = null
  licensed = null
  lock = {}
  mpls_advanced = {}
  neighbor_discovery = null
  port_id = null
  ptp = {}
  share = {}
  tag_protocol_id = null
  taptx = null
  timestamp = {}
  vxlan_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `access_roles` (Object({level1, level2, level3, level4}), optional) - Port Access RBAC definitions
  * `level1` (List(String), optional) - list of roles with 'Level 1' access privileges. Level 1 - Read-only access. Can view port configuration and statistics
  * `level2` (List(String), optional) - list of roles with 'Level 2' access privileges. Level 2 - Level 1 plus the capability to configure port-lock, lock-share, and all traffic objects except port-pair
  * `level3` (List(String), optional) - list of roles with 'Level 3' access privileges. Level 3 - Level 2 plus the capability to configure port params (such as administrative status of the port, speed, duplex, and auto-negotiation), as well as port-pair
  * `level4` (List(String), optional) - list of roles with 'Level 4' access privileges. Level 4 - Level 3 plus the capability to change the port type
* `alarm_thresholds` (Object({alarm_buffer_threshold_rx, alarm_buffer_threshold_tx, alarm_threshold, alarm_threshold_low}), optional) - Port Alarm Thresholds definitions
  * `alarm_buffer_threshold_rx` (Number, optional) - in percent's. Alarm is reported every time threshold is crossed. Value of 0 disables the threshold
  * `alarm_buffer_threshold_tx` (Number, optional) - in percent's. Alarm is reported every time threshold is crossed. Value of 0 disables the threshold
  * `alarm_threshold` (Number, optional) - in percent's. Alarm is reported every time threshold is crossed. The threshold must be exceeded for at least six consecutive seconds. Value of 0 disables the threshold
  * `alarm_threshold_low` (Number, optional) - in percents. Alarm is reported every time this low threshold is crossed. It must be under this threshold for at least six consecutive seconds. Value of 0 disables the threshold
* `fec` (String, optional) - enable/disable forward error correction
* `gdp` (Bool, optional) - enable/disable GDP packets on port
* `header_strip` (String, optional) - protocol type
* `id` (String, required)
* `ingress_vlan_tag` (Number, optional) - Ingress port VLAN tag, with valid numbers between 2 and 4000 that is added to a packet. Setting the value to 0 disables VLAN tagging. The port must be a network port
* `l2_gre_id` (Number, optional) - Value of 0 disables the l2greId
* `licensed` (Bool, optional) - for TA series indicates whether a port is licensed. Defaults to 'true'
* `lock` (Object({description, locking_user, shared_with}), optional) - Port Locking definitions
  * `description` (String, optional) - Optional lock description string
  * `locking_user` (String, required) - User account name this lock belongs to. Lock is used to restrict use of the port by only the specified user account. Users with Default/Operator privileges can only lock ports to which their account has been granted access under their own user account name only. Administrators can lock any port in the system for any existing user account name
  * `shared_with` (List(String), optional) - Used to share a locked port with other user accounts. Users with Default/Operator privileges can only share a lock if the port is locked under their account name. Administrators can lock a port for another user account
* `mpls_advanced` (Object({mpls_advanced_opt}), optional) - Select a combination of Mpls-Advanced options
  * `mpls_advanced_opt` (List(Object({adv_opt})), optional) - List of Mpls-Advanced Options
* `neighbor_discovery` (String, optional) - Configures port neighbor discovery options
* `port_id` (String, required) - device port id
* `ptp` (Object({announce_interval, delay_request_interval, enable, local_priority, role, role_alias, sync_interval, timestamp, vlan}), optional) - Port PTP configurations
  * `announce_interval` (Number, optional) - Configures the interval between PTP announce messages on an interface. The range for the PTP announcement interval is from -2 to 4 log seconds. For the domain 0 as well as domain 24 to 43 the interval limit will be of -3 to 4
  * `delay_request_interval` (Number, optional) - Configures the minimum interval allowed between PTP delay messages when the port is in the source state. The range is from log(-7) to log(5) seconds, where log(-1) = 1 frame per second. For the domain 0 the interval limit is -4 to 5 and for the range of domain 24 to 43 the interval limit will be of -7 to 4.
  * `enable` (Bool, optional) - Enable PTP on the port
  * `local_priority` (Number, optional)
  * `role` (String, optional) - deprecated: use roleAlias
  * `role_alias` (String, optional)
  * `sync_interval` (Number, optional) - Configures the interval between PTP synchronization messages on an interface. The range is from log(-7) to log(1) seconds. For domain 0 the interval limit is -4 to 1 and for the range of domain 24 to 43 the interval limit will be of -7 to 4.
  * `timestamp` (Object({egress, ingress}), optional) - Port PTP timestamp configurations
    * `egress` (Object({insert, source_id}), optional)
      * `insert` (Bool, optional) - When set, will insert timestamp for packets on the port
      * `source_id` (Number, optional) - Source identifier of Packet Time Stamp record header
    * `ingress` (Object({insert, source_id}), optional)
      * `insert` (Bool, optional) - When set, will insert timestamp for packets on the port
      * `source_id` (Number, optional) - Source identifier of Packet Time Stamp record header
  * `vlan` (Number, optional) - configures vlan on the PTP configured port
* `share` (Object({tool_share_roles}), optional) - Port Sharing definitions
  * `tool_share_roles` (List(String), optional) - Only applicable to tool ports. Used to designate a tool port as available for tool-to-tool pass-alls (tool-mirrors) with the specified roles
* `tag_protocol_id` (String, optional) - When Ingress port VLAN tag is added , this protocol Id will be added which egress out the traffic of tool port
* `taptx` (String, optional) - Opens or closes the copper tap port relay. Active closes the port relay. Passive opens the port relay
* `timestamp` (Object({append_ingress, source_id, strip_egress}), optional) - Timestamping definitions for GigaPORT-X12-TS ports
  * `append_ingress` (Bool, optional) - Used to add a timestamp to ingress packets. Applicable to GigaPORT-X12-TS card ports x1..x12 when configured as network ports
  * `source_id` (Number, optional) - Used to specify a custom source-id to be included in the timestamp appended by the GigaPORT-X12-TS. Applicable to GigaPORT-X12-TS card ports x1..x12 when configured as network ports. The timestamp always includes a source-id field. If this custom value is not explicitly specified, the GigaPORT-X12-TS generates one automatically using the following formula: (<box-id> \* 2048) + (<slot-id> \* 256) + <port-number>
  * `strip_egress` (Bool, optional) - Used to strip timestamps from egress packets. Use this argument to strip timestamps from egress packets. Important: this option should only be enabled to packets with time stamps appended. This function will strip the last 14 bytes of each packet regardless of whether a timestamp has been added. Applicable to GigaPORT-X12-TS card ports x9..x12 when configured as tool ports
* `vxlan_id` (Number, optional) - Value of 0 disables the vxlanId

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `access_roles` (Object({level1, level2, level3, level4}), computed) - Port Access RBAC definitions
  * `level1` (List(String), optional) - list of roles with 'Level 1' access privileges. Level 1 - Read-only access. Can view port configuration and statistics
  * `level2` (List(String), optional) - list of roles with 'Level 2' access privileges. Level 2 - Level 1 plus the capability to configure port-lock, lock-share, and all traffic objects except port-pair
  * `level3` (List(String), optional) - list of roles with 'Level 3' access privileges. Level 3 - Level 2 plus the capability to configure port params (such as administrative status of the port, speed, duplex, and auto-negotiation), as well as port-pair
  * `level4` (List(String), optional) - list of roles with 'Level 4' access privileges. Level 4 - Level 3 plus the capability to change the port type
* `alarm_thresholds` (Object({alarm_buffer_threshold_rx, alarm_buffer_threshold_tx, alarm_threshold, alarm_threshold_low}), computed) - Port Alarm Thresholds definitions
  * `alarm_buffer_threshold_rx` (Number, optional) - in percent's. Alarm is reported every time threshold is crossed. Value of 0 disables the threshold
  * `alarm_buffer_threshold_tx` (Number, optional) - in percent's. Alarm is reported every time threshold is crossed. Value of 0 disables the threshold
  * `alarm_threshold` (Number, optional) - in percent's. Alarm is reported every time threshold is crossed. The threshold must be exceeded for at least six consecutive seconds. Value of 0 disables the threshold
  * `alarm_threshold_low` (Number, optional) - in percents. Alarm is reported every time this low threshold is crossed. It must be under this threshold for at least six consecutive seconds. Value of 0 disables the threshold
* `fec` (String, computed) - enable/disable forward error correction
* `gdp` (Bool, computed) - enable/disable GDP packets on port
* `header_strip` (String, computed) - protocol type
* `ingress_vlan_tag` (Number, computed) - Ingress port VLAN tag, with valid numbers between 2 and 4000 that is added to a packet. Setting the value to 0 disables VLAN tagging. The port must be a network port
* `l2_gre_id` (Number, computed) - Value of 0 disables the l2greId
* `licensed` (Bool, computed) - for TA series indicates whether a port is licensed. Defaults to 'true'
* `lock` (Object({description, locking_user, shared_with}), computed) - Port Locking definitions
  * `description` (String, optional) - Optional lock description string
  * `locking_user` (String, required) - User account name this lock belongs to. Lock is used to restrict use of the port by only the specified user account. Users with Default/Operator privileges can only lock ports to which their account has been granted access under their own user account name only. Administrators can lock any port in the system for any existing user account name
  * `shared_with` (List(String), optional) - Used to share a locked port with other user accounts. Users with Default/Operator privileges can only share a lock if the port is locked under their account name. Administrators can lock a port for another user account
* `mpls_advanced` (Object({mpls_advanced_opt}), computed) - Select a combination of Mpls-Advanced options
  * `mpls_advanced_opt` (List(Object({adv_opt})), optional) - List of Mpls-Advanced Options
* `neighbor_discovery` (String, computed) - Configures port neighbor discovery options
* `ptp` (Object({announce_interval, delay_request_interval, enable, local_priority, role, role_alias, sync_interval, timestamp, vlan}), computed) - Port PTP configurations
  * `announce_interval` (Number, optional) - Configures the interval between PTP announce messages on an interface. The range for the PTP announcement interval is from -2 to 4 log seconds. For the domain 0 as well as domain 24 to 43 the interval limit will be of -3 to 4
  * `delay_request_interval` (Number, optional) - Configures the minimum interval allowed between PTP delay messages when the port is in the source state. The range is from log(-7) to log(5) seconds, where log(-1) = 1 frame per second. For the domain 0 the interval limit is -4 to 5 and for the range of domain 24 to 43 the interval limit will be of -7 to 4.
  * `enable` (Bool, optional) - Enable PTP on the port
  * `local_priority` (Number, optional)
  * `role` (String, optional) - deprecated: use roleAlias
  * `role_alias` (String, optional)
  * `sync_interval` (Number, optional) - Configures the interval between PTP synchronization messages on an interface. The range is from log(-7) to log(1) seconds. For domain 0 the interval limit is -4 to 1 and for the range of domain 24 to 43 the interval limit will be of -7 to 4.
  * `timestamp` (Object({egress, ingress}), optional) - Port PTP timestamp configurations
    * `egress` (Object({insert, source_id}), optional)
      * `insert` (Bool, optional) - When set, will insert timestamp for packets on the port
      * `source_id` (Number, optional) - Source identifier of Packet Time Stamp record header
    * `ingress` (Object({insert, source_id}), optional)
      * `insert` (Bool, optional) - When set, will insert timestamp for packets on the port
      * `source_id` (Number, optional) - Source identifier of Packet Time Stamp record header
  * `vlan` (Number, optional) - configures vlan on the PTP configured port
* `share` (Object({tool_share_roles}), computed) - Port Sharing definitions
  * `tool_share_roles` (List(String), optional) - Only applicable to tool ports. Used to designate a tool port as available for tool-to-tool pass-alls (tool-mirrors) with the specified roles
* `tag_protocol_id` (String, computed) - When Ingress port VLAN tag is added , this protocol Id will be added which egress out the traffic of tool port
* `taptx` (String, computed) - Opens or closes the copper tap port relay. Active closes the port relay. Passive opens the port relay
* `timestamp` (Object({append_ingress, source_id, strip_egress}), computed) - Timestamping definitions for GigaPORT-X12-TS ports
  * `append_ingress` (Bool, optional) - Used to add a timestamp to ingress packets. Applicable to GigaPORT-X12-TS card ports x1..x12 when configured as network ports
  * `source_id` (Number, optional) - Used to specify a custom source-id to be included in the timestamp appended by the GigaPORT-X12-TS. Applicable to GigaPORT-X12-TS card ports x1..x12 when configured as network ports. The timestamp always includes a source-id field. If this custom value is not explicitly specified, the GigaPORT-X12-TS generates one automatically using the following formula: (<box-id> \* 2048) + (<slot-id> \* 256) + <port-number>
  * `strip_egress` (Bool, optional) - Used to strip timestamps from egress packets. Use this argument to strip timestamps from egress packets. Important: this option should only be enabled to packets with time stamps appended. This function will strip the last 14 bytes of each packet regardless of whether a timestamp has been added. Applicable to GigaPORT-X12-TS card ports x9..x12 when configured as tool ports
* `vxlan_id` (Number, computed) - Value of 0 disables the vxlanId

