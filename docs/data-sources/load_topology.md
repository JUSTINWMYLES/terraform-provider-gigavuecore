---
page_title: "gigavuecore_load_topology Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Classic Topology
---

# gigavuecore_load_topology Data Source

Load Classic Topology

## Example Usage

```terraform
data "gigavuecore_load_topology" "example" {
  cluster_id = null
  topo_node_type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - Cluster ID to filter by
* `topo_node_type` (String, optional) - Comma-separated list of topoNodeTypes to filter by. Valid values are \['gigamon', 'manual', 'lldp', 'cdp'\]

### Attributes

In addition to all arguments above, the following attributes are exported:

* `links` (List(Object({comment, connections, endpoint1, endpoint2, giga_link_type, link_state, manual_override, topo_link_id, topo_link_type})), computed)
* `nodes` (List(Dynamic), computed)

