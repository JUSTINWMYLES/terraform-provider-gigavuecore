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
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - if provided, the state and the configurations of the requested cluster will be returned

### Attributes

In addition to all arguments above, the following attributes are exported:

* `cluster_enabled` (Bool, computed) - cluster enable or disable
* `cluster_leader_box_id` (Number, computed) - cluster leader box id
* `cluster_master_box_id` (Number, computed) - cluster master box id. (deprecated: use clusterLeaderBoxId)
* `cluster_members` (List(Object({boot_time, box_id, cc_sync_status, cluster_intf, cluster_name, cluster_node_id, external_address, global_node_id, host_id, hostname, internal_address, internal_port, leader_discovery, leader_preference, master_discovery, master_preference, mgmt_address, model, node_cluster_role, node_cluster_role_alias, node_cluster_state, oper_status, recv_hb_from, send_hb_to, serial_number, sw_build_number, sw_version, vip_intf})), computed)
* `cluster_params` (Object({cluster_id, cluster_vip, cluster_vip_mask_len, ip_protocol, shared_secret, stacking_mode}), computed) - provides clustering parameters
  * `cluster_id` (String, computed) - cluster id
  * `cluster_vip` (String, computed) - cluster leader virtual ip address
  * `cluster_vip_mask_len` (Number, computed) - cluster leader virtual ip mask length, valid and required when 'clusterVip' is specified
  * `ip_protocol` (String, computed) - Supported since 6.2
  * `shared_secret` (String, computed) - shared-secret used for message authentication, length 16 to 64
  * `stacking_mode` (String, computed) - stacking mode
* `global_cluster_id` (String, computed) - global cluster id
* `leader_address` (String, computed) - cluster leader management address
* `local_cluster_role` (String, computed) - standalone: not a member of any cluster; master: cluster master; standby: stand-by master; normal: non-master member; unknown: error, cannot find master. (deprecated: use localClusterRoleAlias)
* `local_cluster_role_alias` (String, computed) - standalone: not a member of any cluster; leader: cluster leader; standby: stand-by leader; normal: non-leader member; unknown: error, cannot find leader
* `localcluster_state` (String, computed) - operational status of the cluster member
* `master_address` (String, computed) - cluster master management address. (deprecated: use leaderAddress)

