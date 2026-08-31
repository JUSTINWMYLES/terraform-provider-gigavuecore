---
page_title: "gigavuecore_create_cluster_config Action - gigavuecore"
subcategory: ""
description: |-
  creates the cluster with the given device members list and cluster parameters
---

# gigavuecore_create_cluster_config Action

creates the cluster with the given device members list and cluster parameters

## Example Usage

```terraform
action "gigavuecore_create_cluster_config" "example" {
  config {
    cluster_member_specs = [{
      box_id       = 0
      cluster_intf = "example"
      cluster_port = 0
      leader_discovery = {
        auto_discovery            = true
        cluster_formation_timeout = 0
        primary_ip                = "example"
        primary_port              = 0
        secondary_ip              = "example"
        secondary_port            = 0
        static_discovery_timeout  = 10
      }
      leader_preference = 0
      master_discovery = {
        auto_discovery            = true
        cluster_formation_timeout = 0
        primary_ip                = "example"
        primary_port              = 0
        secondary_ip              = "example"
        secondary_port            = 0
        static_discovery_timeout  = 10
      }
      master_preference = 0
      mgmt_address      = "example"
      vip_intf          = "example"
    }]
    cluster_params = {
      cluster_id           = "example"
      cluster_vip          = "example"
      cluster_vip_mask_len = 0
      ip_protocol          = "ipv4"
      shared_secret        = "example"
      stacking_mode        = "example"
    }
    cluster_type           = "example"
    seed_node_mgmt_address = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_member_specs` (Attributes List, optional) (see [below for nested schema](#nestedatt--cluster_member_specs))
* `cluster_params` (Attributes, optional) - provides clustering parameters (see [below for nested schema](#nestedatt--cluster_params))
* `cluster_type` (String, optional) - Type of the cluster eg. oob. It cannot be null
* `seed_node_mgmt_address` (String, optional) - seed node management address

<a id="nestedatt--cluster_member_specs"></a>
### Nested Schema for `cluster_member_specs`

Optional:

* `box_id` (Number) - box id of the cluster member
* `cluster_intf` (String) - cluster service interface
* `cluster_port` (Number) - cluster service port
* `leader_discovery` (Attributes) - provides leader discovery information (see [below for nested schema](#nestedatt--cluster_member_specs--leader_discovery))
* `leader_preference` (Number) - cluster leader preference. Larger value means more preferred, set 1-9 to exclude leadership
* `master_discovery` (Attributes) - provides master discovery information. (deprecated: use leaderDiscovery) (see [below for nested schema](#nestedatt--cluster_member_specs--master_discovery))
* `master_preference` (Number) - cluster master preference. Larger value means more preferred, set 1-9 to exclude mastership. (deprecated: use leaderPreference)
* `mgmt_address` (String) - Management Address
* `vip_intf` (String) - cluster master interface

<a id="nestedatt--cluster_member_specs--leader_discovery"></a>
### Nested Schema for `cluster_member_specs.leader_discovery`

Optional:

* `auto_discovery` (Boolean) - Cluster auto-discovery. Nodes with auto-discovery disabled cannot become leader
* `cluster_formation_timeout` (Number) - maximum expected time for cluster startup in seconds
* `primary_ip` (String) - Leader primary ip address. Only valid if 'autoDiscovery' is disabled
* `primary_port` (Number) - Leader primary port. Only valid if 'autoDiscovery' is disabled
* `secondary_ip` (String) - Leader secondary ip address. Only valid if 'autoDiscovery' is disabled
* `secondary_port` (Number) - Leader secondary port. Only valid if 'autoDiscovery' is disabled
* `static_discovery_timeout` (Number) - Set the cluster leader connection timeout in seconds. Maps to CLI 'connectTimeout'. Controls how long 'primaryIp' should be tried before trying to connect to 'secondaryIp' and vice versa. Only valid if 'autoDiscovery' is disabled

<a id="nestedatt--cluster_member_specs--master_discovery"></a>
### Nested Schema for `cluster_member_specs.master_discovery`

Optional:

* `auto_discovery` (Boolean) - Cluster auto-discovery. Nodes with auto-discovery disabled cannot become master
* `cluster_formation_timeout` (Number) - maximum expected time for cluster startup in seconds
* `primary_ip` (String) - Master primary ip address. Only valid if 'autoDiscovery' is disabled
* `primary_port` (Number) - Master primary port. Only valid if 'autoDiscovery' is disabled
* `secondary_ip` (String) - Master secondary ip address. Only valid if 'autoDiscovery' is disabled
* `secondary_port` (Number) - Master secondary port. Only valid if 'autoDiscovery' is disabled
* `static_discovery_timeout` (Number) - Set the cluster master connection timeout in seconds. Maps to CLI 'connectTimeout'. Controls how long 'primaryIp' should be tried before trying to connect to 'secondaryIp' and vice versa. Only valid if 'autoDiscovery' is disabled

<a id="nestedatt--cluster_params"></a>
### Nested Schema for `cluster_params`

Optional:

* `cluster_id` (String) - cluster id
* `cluster_vip` (String) - cluster leader virtual ip address
* `cluster_vip_mask_len` (Number) - cluster leader virtual ip mask length, valid and required when 'clusterVip' is specified
* `ip_protocol` (String) - Supported since 6.2
* `shared_secret` (String) - shared-secret used for message authentication, length 16 to 64
* `stacking_mode` (String) - stacking mode

