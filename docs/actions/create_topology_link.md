---
page_title: "gigavuecore_create_topology_link Action - gigavuecore"
subcategory: ""
description: |-
  Returns generated Topology Link Id for future reference
---

# gigavuecore_create_topology_link Action

Returns generated Topology Link Id for future reference

## Example Usage

```terraform
action "gigavuecore_create_topology_link" "example" {
  config {
    comment = "example"
    connections = [{
      port1 = "example"
      port2 = "example"
    }]
    endpoint1 = {
      comment         = "example"
      component_alias = "example"
      component_type  = "example"
      port            = "example"
      topo_node_id    = "example"
      topo_node_type  = "example"
    }
    endpoint2 = {
      comment         = "example"
      component_alias = "example"
      component_type  = "example"
      port            = "example"
      topo_node_id    = "example"
      topo_node_type  = "example"
    }
    topo_link_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `comment` (String, optional)
* `connections` (Attributes List, optional) - Actual port ids which are part of the physical connection (see [below for nested schema](#nestedatt--connections))
* `endpoint1` (Attributes, required) - Node link endpoint Create/Update Spec (see [below for nested schema](#nestedatt--endpoint1))
* `endpoint2` (Attributes, required) - Node link endpoint Create/Update Spec (see [below for nested schema](#nestedatt--endpoint2))
* `topo_link_id` (String, optional) - Link's unique identifier. Auto-assigned. This is required while updating existing links

<a id="nestedatt--connections"></a>
### Nested Schema for `connections`

Optional:

* `port1` (String)
* `port2` (String)

<a id="nestedatt--endpoint1"></a>
### Nested Schema for `endpoint1`

Required:

* `component_type` (String) - Type of the endpoint component. This is either 'Port' or 'GigaStream'
* `topo_node_id` (String) - Topology graph node Id. References one of the Topology Nodes in the graph

Optional:

* `comment` (String)
* `component_alias` (String) - For 'gigamon': this is either portId (e.g. '1/1/x1') or Gigastream alias (e.g. 'gs1').  For 'manual': this is user provided or may be omitted
* `port` (String) - For 'gigamon': this is the portId (e.g. '1/1/x1'). For 'manual': this is user provided or may be omitted
* `topo_node_type` (String) - topology node type

<a id="nestedatt--endpoint2"></a>
### Nested Schema for `endpoint2`

Required:

* `component_type` (String) - Type of the endpoint component. This is either 'Port' or 'GigaStream'
* `topo_node_id` (String) - Topology graph node Id. References one of the Topology Nodes in the graph

Optional:

* `comment` (String)
* `component_alias` (String) - For 'gigamon': this is either portId (e.g. '1/1/x1') or Gigastream alias (e.g. 'gs1').  For 'manual': this is user provided or may be omitted
* `port` (String) - For 'gigamon': this is the portId (e.g. '1/1/x1'). For 'manual': this is user provided or may be omitted
* `topo_node_type` (String) - topology node type

