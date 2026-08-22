---
page_title: "gigavuecore_delete_manual_topology_node Action - gigavuecore"
subcategory: ""
description: |-
  Delete Manual Topology Node
---

# gigavuecore_delete_manual_topology_node Action

Delete Manual Topology Node

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
