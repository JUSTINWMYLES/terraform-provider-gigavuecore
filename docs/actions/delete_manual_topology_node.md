---
page_title: "gigavuecore_delete_manual_topology_node Action - gigavuecore"
subcategory: ""
description: |-
  Delete Manual Topology Node
---

# gigavuecore_delete_manual_topology_node Action

Delete Manual Topology Node

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_manual_topology_node" "example" {
  config {
    topo_node_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `topo_node_id` (String, required) - Topology Node Id


