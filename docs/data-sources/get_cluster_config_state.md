---
page_title: "gigavuecore_get_cluster_config_state Data Source - gigavuecore"
subcategory: ""
description: |-
  get the cluster configuration state
---

# gigavuecore_get_cluster_config_state Data Source

get the cluster configuration state

## Example Usage

```terraform
data "gigavuecore_get_cluster_config_state" "example" {
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - if provided, the state and the configurations of the requested cluster will be returned

### Attributes

In addition to all arguments above, the following attributes are exported:

* `cluster_enabled` (Boolean, computed) - cluster enable or disable
* `cluster_leader_box_id` (Number, computed) - cluster leader box id
* `cluster_master_box_id` (Number, computed) - cluster master box id. (deprecated: use clusterLeaderBoxId)
* `cluster_members` (Attributes List, computed) (see [below for nested schema](#nestedatt--cluster_members))
* `cluster_params` (Attributes, computed) - provides clustering parameters (see [below for nested schema](#nestedatt--cluster_params))
* `global_cluster_id` (String, computed) - global cluster id
* `leader_address` (String, computed) - cluster leader management address
* `local_cluster_role` (String, computed) - standalone: not a member of any cluster; master: cluster master; standby: stand-by master; normal: non-master member; unknown: error, cannot find master. (deprecated: use localClusterRoleAlias)
* `local_cluster_role_alias` (String, computed) - standalone: not a member of any cluster; leader: cluster leader; standby: stand-by leader; normal: non-leader member; unknown: error, cannot find leader
* `localcluster_state` (String, computed) - operational status of the cluster member
* `master_address` (String, computed) - cluster master management address. (deprecated: use leaderAddress)

<a id="nestedatt--cluster_members"></a>
### Nested Schema for `cluster_members`

Read-Only:

* `boot_time` (String) - the time node booted up. In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'
* `box_id` (Number) - box id of the cluster member
* `cc_sync_status` (String) - CC1/CC2 dynamic sync status
* `cluster_intf` (String) - cluster service interface
* `cluster_name` (String) - cluster name
* `cluster_node_id` (Number) - cluster node id
* `external_address` (String) - node external address
* `global_node_id` (String) - global node id
* `host_id` (String) - host id
* `hostname` (String) - host name of the member
* `internal_address` (String) - node internal address
* `internal_port` (Number) - node internal port
* `leader_discovery` (Attributes) - provides leader discovery information (see [below for nested schema](#nestedatt--cluster_members--leader_discovery))
* `leader_preference` (Number) - leader election preference rank
* `master_discovery` (Attributes) - provides master discovery information. (deprecated: use leaderDiscovery) (see [below for nested schema](#nestedatt--cluster_members--master_discovery))
* `master_preference` (Number) - master election preference rank. (deprecated: use leaderPreference)
* `mgmt_address` (String) - device management address
* `model` (String) - Device model
* `node_cluster_role` (String) - standalone: not a member of any cluster; master: cluster master; standby: stand-by master; normal: non-master member; unknown: error, cannot find master. (deprecated: use nodeClusterRoleAlias)
* `node_cluster_role_alias` (String) - standalone: not a member of any cluster; leader: cluster leader; standby: stand-by leader; normal: non-leader member; unknown: error, cannot find leader
* `node_cluster_state` (String) - Operational status of the cluster member
* `oper_status` (String) - 'operational': fully configurable; 'safe': no configuration possible; 'limited': limited configuration possible
* `recv_hb_from` (Number) - node id from which heartbeats received
* `send_hb_to` (Number) - node id to which heartbeats sent
* `serial_number` (String) - unique serial number for node chassis
* `sw_build_number` (String) - Device software build number
* `sw_version` (String) - Device software version
* `vip_intf` (String) - cluster master interface

<a id="nestedatt--cluster_members--leader_discovery"></a>
### Nested Schema for `cluster_members.leader_discovery`

Read-Only:

* `auto_discovery` (Boolean) - Cluster auto-discovery. Nodes with auto-discovery disabled cannot become leader
* `cluster_formation_timeout` (Number) - maximum expected time for cluster startup in seconds
* `primary_ip` (String) - Leader primary ip address. Only valid if 'autoDiscovery' is disabled
* `primary_port` (Number) - Leader primary port. Only valid if 'autoDiscovery' is disabled
* `secondary_ip` (String) - Leader secondary ip address. Only valid if 'autoDiscovery' is disabled
* `secondary_port` (Number) - Leader secondary port. Only valid if 'autoDiscovery' is disabled
* `static_discovery_timeout` (Number) - Set the cluster leader connection timeout in seconds. Maps to CLI 'connectTimeout'. Controls how long 'primaryIp' should be tried before trying to connect to 'secondaryIp' and vice versa. Only valid if 'autoDiscovery' is disabled

<a id="nestedatt--cluster_members--master_discovery"></a>
### Nested Schema for `cluster_members.master_discovery`

Read-Only:

* `auto_discovery` (Boolean) - Cluster auto-discovery. Nodes with auto-discovery disabled cannot become master
* `cluster_formation_timeout` (Number) - maximum expected time for cluster startup in seconds
* `primary_ip` (String) - Master primary ip address. Only valid if 'autoDiscovery' is disabled
* `primary_port` (Number) - Master primary port. Only valid if 'autoDiscovery' is disabled
* `secondary_ip` (String) - Master secondary ip address. Only valid if 'autoDiscovery' is disabled
* `secondary_port` (Number) - Master secondary port. Only valid if 'autoDiscovery' is disabled
* `static_discovery_timeout` (Number) - Set the cluster master connection timeout in seconds. Maps to CLI 'connectTimeout'. Controls how long 'primaryIp' should be tried before trying to connect to 'secondaryIp' and vice versa. Only valid if 'autoDiscovery' is disabled

<a id="nestedatt--cluster_params"></a>
### Nested Schema for `cluster_params`

Read-Only:

* `cluster_id` (String) - cluster id
* `cluster_vip` (String) - cluster leader virtual ip address
* `cluster_vip_mask_len` (Number) - cluster leader virtual ip mask length, valid and required when 'clusterVip' is specified
* `ip_protocol` (String) - Supported since 6.2
* `shared_secret` (String) - shared-secret used for message authentication, length 16 to 64
* `stacking_mode` (String) - stacking mode

