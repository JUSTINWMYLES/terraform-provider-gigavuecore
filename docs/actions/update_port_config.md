---
page_title: "gigavuecore_update_port_config Action - gigavuecore"
subcategory: ""
description: |-
  Update a PortConfig configuration
---

# gigavuecore_update_port_config Action

Update a PortConfig configuration

## Example Usage

```terraform
action "gigavuecore_update_port_config" "example" {
  config {
    access_roles = {
      level1 = [ "example" ]
      level2 = [ "example" ]
      level3 = [ "example" ]
      level4 = [ "example" ]
    }
    alarm_thresholds = {
      alarm_buffer_threshold_rx = 0
      alarm_buffer_threshold_tx = 0
      alarm_threshold           = 0
      alarm_threshold_low       = 0
    }
    body_port_id     = "example"
    cluster_id       = "example"
    fec              = "example"
    gdp              = true
    header_strip     = "example"
    ingress_vlan_tag = 0
    l2_gre_id        = 0
    licensed         = true
    lock = {
      description  = "example"
      locking_user = "example"
      shared_with  = [ "example" ]
    }
    mpls_advanced = {
      mpls_advanced_opt = [{
        adv_opt = "example"
      }]
    }
    neighbor_discovery = "example"
    port_id            = "example"
    ptp = {
      announce_interval      = 0
      delay_request_interval = 0
      enable                 = true
      local_priority         = 0
      role                   = "example"
      role_alias             = "example"
      sync_interval          = 0
      timestamp = {
        egress = {
          insert    = true
          source_id = 0
        }
        ingress = {
          insert    = true
          source_id = 0
        }
      }
      vlan = 0
    }
    share = {
      tool_share_roles = [ "example" ]
    }
    tag_protocol_id = "example"
    taptx           = "example"
    timestamp = {
      append_ingress = true
      source_id      = 0
      strip_egress   = true
    }
    vxlan_id = 0
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `access_roles` (Attributes, optional) - Port Access RBAC definitions (see [below for nested schema](#nestedatt--access_roles))
* `alarm_thresholds` (Attributes, optional) - Port Alarm Thresholds definitions (see [below for nested schema](#nestedatt--alarm_thresholds))
* `body_port_id` (String, required) - device port id
* `cluster_id` (String, required) - Target Cluster ID
* `fec` (String, optional) - enable/disable forward error correction
* `gdp` (Boolean, optional) - enable/disable GDP packets on port
* `header_strip` (String, optional) - protocol type
* `ingress_vlan_tag` (Number, optional) - Ingress port VLAN tag, with valid numbers between 2 and 4000 that is added to a packet. Setting the value to 0 disables VLAN tagging. The port must be a network port
* `l2_gre_id` (Number, optional) - Value of 0 disables the l2greId
* `licensed` (Boolean, optional) - for TA series indicates whether a port is licensed. Defaults to 'true'
* `lock` (Attributes, optional) - Port Locking definitions (see [below for nested schema](#nestedatt--lock))
* `mpls_advanced` (Attributes, optional) - Select a combination of Mpls-Advanced options (see [below for nested schema](#nestedatt--mpls_advanced))
* `neighbor_discovery` (String, optional) - Configures port neighbor discovery options
* `port_id` (String, required) - Device Port ID (format: boxId\_slotId\_port, example: 1\_1\_c1)
* `ptp` (Attributes, optional) - Port PTP configurations (see [below for nested schema](#nestedatt--ptp))
* `share` (Attributes, optional) - Port Sharing definitions (see [below for nested schema](#nestedatt--share))
* `tag_protocol_id` (String, optional) - When Ingress port VLAN tag is added , this protocol Id will be added which egress out the traffic of tool port
* `taptx` (String, optional) - Opens or closes the copper tap port relay. Active closes the port relay. Passive opens the port relay
* `timestamp` (Attributes, optional) - Timestamping definitions for GigaPORT-X12-TS ports (see [below for nested schema](#nestedatt--timestamp))
* `vxlan_id` (Number, optional) - Value of 0 disables the vxlanId

<a id="nestedatt--access_roles"></a>
### Nested Schema for `access_roles`

Optional:

* `level1` (List of String) - list of roles with 'Level 1' access privileges. Level 1 - Read-only access. Can view port configuration and statistics
* `level2` (List of String) - list of roles with 'Level 2' access privileges. Level 2 - Level 1 plus the capability to configure port-lock, lock-share, and all traffic objects except port-pair
* `level3` (List of String) - list of roles with 'Level 3' access privileges. Level 3 - Level 2 plus the capability to configure port params (such as administrative status of the port, speed, duplex, and auto-negotiation), as well as port-pair
* `level4` (List of String) - list of roles with 'Level 4' access privileges. Level 4 - Level 3 plus the capability to change the port type

<a id="nestedatt--alarm_thresholds"></a>
### Nested Schema for `alarm_thresholds`

Optional:

* `alarm_buffer_threshold_rx` (Number) - in percent's. Alarm is reported every time threshold is crossed. Value of 0 disables the threshold
* `alarm_buffer_threshold_tx` (Number) - in percent's. Alarm is reported every time threshold is crossed. Value of 0 disables the threshold
* `alarm_threshold` (Number) - in percent's. Alarm is reported every time threshold is crossed. The threshold must be exceeded for at least six consecutive seconds. Value of 0 disables the threshold
* `alarm_threshold_low` (Number) - in percents. Alarm is reported every time this low threshold is crossed. It must be under this threshold for at least six consecutive seconds. Value of 0 disables the threshold

<a id="nestedatt--lock"></a>
### Nested Schema for `lock`

Required:

* `locking_user` (String) - User account name this lock belongs to. Lock is used to restrict use of the port by only the specified user account. Users with Default/Operator privileges can only lock ports to which their account has been granted access under their own user account name only. Administrators can lock any port in the system for any existing user account name

Optional:

* `description` (String) - Optional lock description string
* `shared_with` (List of String) - Used to share a locked port with other user accounts. Users with Default/Operator privileges can only share a lock if the port is locked under their account name. Administrators can lock a port for another user account

<a id="nestedatt--mpls_advanced"></a>
### Nested Schema for `mpls_advanced`

Optional:

* `mpls_advanced_opt` (Attributes List) - List of Mpls-Advanced Options (see [below for nested schema](#nestedatt--mpls_advanced--mpls_advanced_opt))

<a id="nestedatt--mpls_advanced--mpls_advanced_opt"></a>
### Nested Schema for `mpls_advanced.mpls_advanced_opt`

Optional:

* `adv_opt` (String) - Mpls-Advanced Options

<a id="nestedatt--ptp"></a>
### Nested Schema for `ptp`

Optional:

* `announce_interval` (Number) - Configures the interval between PTP announce messages on an interface. The range for the PTP announcement interval is from -2 to 4 log seconds. For the domain 0 as well as domain 24 to 43 the interval limit will be of -3 to 4
* `delay_request_interval` (Number) - Configures the minimum interval allowed between PTP delay messages when the port is in the source state. The range is from log(-7) to log(5) seconds, where log(-1) = 1 frame per second. For the domain 0 the interval limit is -4 to 5 and for the range of domain 24 to 43 the interval limit will be of -7 to 4.
* `enable` (Boolean) - Enable PTP on the port
* `local_priority` (Number)
* `role` (String) - deprecated: use roleAlias
* `role_alias` (String)
* `sync_interval` (Number) - Configures the interval between PTP synchronization messages on an interface. The range is from log(-7) to log(1) seconds. For domain 0 the interval limit is -4 to 1 and for the range of domain 24 to 43 the interval limit will be of -7 to 4.
* `timestamp` (Attributes) - Port PTP timestamp configurations (see [below for nested schema](#nestedatt--ptp--timestamp))
* `vlan` (Number) - configures vlan on the PTP configured port

<a id="nestedatt--ptp--timestamp"></a>
### Nested Schema for `ptp.timestamp`

Optional:

* `egress` (Attributes) - Tx (Egress) settings (see [below for nested schema](#nestedatt--ptp--timestamp--egress))
* `ingress` (Attributes) - Rx (Ingress) settings (see [below for nested schema](#nestedatt--ptp--timestamp--ingress))

<a id="nestedatt--ptp--timestamp--egress"></a>
### Nested Schema for `ptp.timestamp.egress`

Optional:

* `insert` (Boolean) - When set, will insert timestamp for packets on the port
* `source_id` (Number) - Source identifier of Packet Time Stamp record header

<a id="nestedatt--ptp--timestamp--ingress"></a>
### Nested Schema for `ptp.timestamp.ingress`

Optional:

* `insert` (Boolean) - When set, will insert timestamp for packets on the port
* `source_id` (Number) - Source identifier of Packet Time Stamp record header

<a id="nestedatt--share"></a>
### Nested Schema for `share`

Optional:

* `tool_share_roles` (List of String) - Only applicable to tool ports. Used to designate a tool port as available for tool-to-tool pass-alls (tool-mirrors) with the specified roles

<a id="nestedatt--timestamp"></a>
### Nested Schema for `timestamp`

Optional:

* `append_ingress` (Boolean) - Used to add a timestamp to ingress packets. Applicable to GigaPORT-X12-TS card ports x1..x12 when configured as network ports
* `source_id` (Number) - Used to specify a custom source-id to be included in the timestamp appended by the GigaPORT-X12-TS. Applicable to GigaPORT-X12-TS card ports x1..x12 when configured as network ports. The timestamp always includes a source-id field. If this custom value is not explicitly specified, the GigaPORT-X12-TS generates one automatically using the following formula: (<box-id> \* 2048) + (<slot-id> \* 256) + <port-number>
* `strip_egress` (Boolean) - Used to strip timestamps from egress packets. Use this argument to strip timestamps from egress packets. Important: this option should only be enabled to packets with time stamps appended. This function will strip the last 14 bytes of each packet regardless of whether a timestamp has been added. Applicable to GigaPORT-X12-TS card ports x9..x12 when configured as tool ports

