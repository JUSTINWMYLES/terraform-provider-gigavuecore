---
page_title: "gigavuecore_cluster_config_switch_i_pprotocol_spec Action - gigavuecore"
subcategory: ""
description: |-
  switch ip protocol for physical cluster nodes
---

# gigavuecore_cluster_config_switch_i_pprotocol_spec Action

switch ip protocol for physical cluster nodes

## Example Usage

```terraform
action "gigavuecore_cluster_config_switch_i_pprotocol_spec" "example" {
  config {
    cluster_id           = "example"
    cluster_primary_ip   = "example"
    cluster_secondary_ip = "example"
    cluster_vip          = "example"
    cluster_vip_mask_len = 1
    ip_protocol          = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - The requested cluster ID
* `cluster_primary_ip` (String, optional) - Cluster primary IP
* `cluster_secondary_ip` (String, optional) - Cluster secondary ip
* `cluster_vip` (String, optional) - Cluster VIP
* `cluster_vip_mask_len` (Number, optional) - cluster leader virtual ip mask length, valid and required when 'clusterVip' is specified
* `ip_protocol` (String, optional) - ip protocol


