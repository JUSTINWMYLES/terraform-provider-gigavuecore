---
page_title: "gigavuecore_delete_circuit_tunnel_vxlan_groups Action - gigavuecore"
subcategory: ""
description: |-
  Delete all Circuit Tunnel VxLan Groups
---

# gigavuecore_delete_circuit_tunnel_vxlan_groups Action

Delete all Circuit Tunnel VxLan Groups

## Example Usage

```terraform
action "gigavuecore_delete_circuit_tunnel_vxlan_groups" "example" {
  config {
    box_id     = "example"
    cluster_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, optional) - device box id. valid range 1 - 64.
* `cluster_id` (String, required) - Target Cluster ID


