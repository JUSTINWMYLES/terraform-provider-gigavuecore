---
page_title: "gigavuecore_load_topology_viz_nodes Data Source - gigavuecore"
subcategory: ""
description: |-
  Load manual and gigamon nodes
---

# gigavuecore_load_topology_viz_nodes Data Source

Load manual and gigamon nodes

## Example Usage

```terraform
data "gigavuecore_load_topology_viz_nodes" "example" {
  node_type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `node_type` (String, optional) - Topology node type

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List of Dynamic, computed)


