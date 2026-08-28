---
page_title: "gigavuecore_get_managed_clusters Data Source - gigavuecore"
subcategory: ""
description: |-
  get cluster-grouped managed device list
---

# gigavuecore_get_managed_clusters Data Source

get cluster-grouped managed device list

## Example Usage

```terraform
data "gigavuecore_get_managed_clusters" "example" {
  cluster_id = null
  page       = null
  sort       = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - if provided, only requested cluster is returned
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `cluster_id` (String) - for H-series clusters only. Identifies the configured cluster id
* `cluster_vip` (String) - for H-series clusters only. Identifies the configured cluster VIP address
* `family` (String) - identifies whether this is an H-series cluster or a G-series stack
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--health_state_reasons))
* `leader_id` (String) - id of stack's leader device
* `master_id` (String) - id of stack's master device. (deprecated: use leaderId)
* `members` (Attributes List) (see [below for nested schema](#nestedatt--items--members))
<a id="nestedatt--items--health_state_reasons"></a>
### Nested Schema for `items.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--items--members"></a>
### Nested Schema for `items.members`

Read-Only:

* `box_id` (String)
* `chassis_oper_status` (String) - Describes operational state available in Chassis information. 'up': node is up; 'down': node is down; 'left': node is left from the cluster; 'not-reachable': node is not reachable from the cluster' . Applicable only for H-series devices
* `cluster_id` (String) - applicable only for clustered H-series devices
* `cluster_leader` (String) - address of stack/cluster leader device
* `cluster_master` (String) - address of stack/cluster master device. (deprecated: use clusterLeader)
* `cluster_mode` (String) - standalone: not a member of any cluster/stack; master: cluster/stack master; standby: for H-series only; slave: non-master member of an H-series cluster or a slave in a G-series stack; vfex: visibility fabric extender of an H-Series cluster. (deprecated: use clusterModeAlias)
* `cluster_mode_alias` (String) - standalone: not a member of any cluster/stack; leader: cluster/stack leader; standby: for H-series only; member: non-leader member of an H-series cluster or a member in a G-series stack; vfex: visibility fabric extender of an H-Series cluster
* `cluster_state` (String) - Describes the cluster membership state. Applicable only for clustered H-series devices
* `cluster_vip` (String) - applicable only for clustered H-series devices
* `device_id` (String) - unique ID representing this managed Gigamon device
* `device_ip` (String) - device management IP Address
* `device_ips` (List of String) - list of all the IPs on the device
* `disc_outcome` (String) - status of FM-to-device communication channel
* `disconnected` (Boolean) - Status of connection between FM and Device
* `dns_name` (String) - device DNS name, as registered in DNS server
* `failure_desc` (String) - reason for device discovery failure
* `family` (String)
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--members--health_state_reasons))
* `hostname` (String) - device configured hostname
* `leader_pref` (Number) - leader election preference rank
* `licensed` (Boolean)
* `master_pref` (Number) - master election preference rank. (deprecated: use leaderPref)
* `model` (String) - Gigamon physical device models
* `nat_ip` (String) - Device NAT IP Address. Applicable for only devices behind NAT set up
* `oper_status` (String) - Describes device operational state. 'operational': fully configurable; 'safe': no configuration possible; 'limited': limited configuration possible. Applicable only for H-series devices
* `sw_build_number` (String)
* `sw_version` (String)
* `topo_node_id` (String) - unique topology node ID. Generated by FM server
<a id="nestedatt--items--members--health_state_reasons"></a>
### Nested Schema for `items.members.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

