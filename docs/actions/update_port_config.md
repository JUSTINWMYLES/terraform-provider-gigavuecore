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
    access_roles = null
    alarm_thresholds = null
    body_port_id = "example"
    cluster_id = "example"
    fec = "example"
    gdp = true
    header_strip = "example"
    ingress_vlan_tag = 1
    l2_gre_id = 1
    licensed = true
    lock = null
    mpls_advanced = null
    neighbor_discovery = "example"
    port_id = "example"
    ptp = null
    share = null
    tag_protocol_id = "example"
    taptx = "example"
    timestamp = null
    vxlan_id = 1
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `access_roles` (Dynamic, optional) - Port Access RBAC definitions
* `alarm_thresholds` (Dynamic, optional) - Port Alarm Thresholds definitions
* `body_port_id` (String, required) - device port id
* `cluster_id` (String, required) - Target Cluster ID
* `fec` (String, optional) - enable/disable forward error correction
* `gdp` (Bool, optional) - enable/disable GDP packets on port
* `header_strip` (String, optional) - protocol type
* `ingress_vlan_tag` (Number, optional) - Ingress port VLAN tag, with valid numbers between 2 and 4000 that is added to a packet. Setting the value to 0 disables VLAN tagging. The port must be a network port
* `l2_gre_id` (Number, optional) - Value of 0 disables the l2greId
* `licensed` (Bool, optional) - for TA series indicates whether a port is licensed. Defaults to 'true'
* `lock` (Dynamic, optional) - Port Locking definitions
* `mpls_advanced` (Dynamic, optional) - Select a combination of Mpls-Advanced options
* `neighbor_discovery` (String, optional) - Configures port neighbor discovery options
* `port_id` (String, required) - Device Port ID (format: boxId\_slotId\_port, example: 1\_1\_c1)
* `ptp` (Dynamic, optional) - Port PTP configurations
* `share` (Dynamic, optional) - Port Sharing definitions
* `tag_protocol_id` (String, optional) - When Ingress port VLAN tag is added , this protocol Id will be added which egress out the traffic of tool port
* `taptx` (String, optional) - Opens or closes the copper tap port relay. Active closes the port relay. Passive opens the port relay
* `timestamp` (Dynamic, optional) - Timestamping definitions for GigaPORT-X12-TS ports
* `vxlan_id` (Number, optional) - Value of 0 disables the vxlanId
