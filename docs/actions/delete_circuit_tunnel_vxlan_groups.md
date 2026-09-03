---
page_title: "gigavuecore_delete_circuit_tunnel_vxlan_groups Action - gigavuecore"
subcategory: ""
description: |-
  Delete all Circuit Tunnel VxLan Groups
---

# gigavuecore_delete_circuit_tunnel_vxlan_groups Action

Delete all Circuit Tunnel VxLan Groups

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

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


