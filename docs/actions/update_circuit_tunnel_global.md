---
page_title: "gigavuecore_update_circuit_tunnel_global Action - gigavuecore"
subcategory: ""
description: |-
  Update Circuit Tunnel global configurations
---

# gigavuecore_update_circuit_tunnel_global Action

Update Circuit Tunnel global configurations

## Example Usage

```terraform
action "gigavuecore_update_circuit_tunnel_global" "example" {
  config {
    box_id            = "example"
    cluster_id        = "example"
    vxlan_l4_dst_port = 0
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, optional) - device box id. valid range 1 - 64.
* `cluster_id` (String, required) - Target Cluster ID
* `vxlan_l4_dst_port` (Number, optional) - l4 destination port for tunnel terminating on the box


