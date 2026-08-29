---
page_title: "gigavuecore_create_manual_topology_entity Action - gigavuecore"
subcategory: ""
description: |-
  Create Manual Topology Nodes and Links
---

# gigavuecore_create_manual_topology_entity Action

Create Manual Topology Nodes and Links

## Example Usage

```terraform
action "gigavuecore_create_manual_topology_entity" "example" {
  config {
    links = [{
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
    }]
    nodes = [{
      comment      = "example"
      model        = "example"
      node_alias   = "example"
      topo_node_id = "example"
      type         = "example"
      vendor       = "example"
    }]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `links` (Attributes List, optional) (see [below for nested schema](#nestedatt--links))
* `nodes` (Attributes List, optional) (see [below for nested schema](#nestedatt--nodes))

<a id="nestedatt--links"></a>
### Nested Schema for `links`

Required:

* `endpoint1` (Attributes) - Node link endpoint Create/Update Spec (see [below for nested schema](#nestedatt--links--endpoint1))
* `endpoint2` (Attributes) - Node link endpoint Create/Update Spec (see [below for nested schema](#nestedatt--links--endpoint2))

Optional:

* `comment` (String)
* `connections` (Attributes List) - Actual port ids which are part of the physical connection (see [below for nested schema](#nestedatt--links--connections))
* `topo_link_id` (String) - Link's unique identifier. Auto-assigned. This is required while updating existing links

<a id="nestedatt--links--endpoint1"></a>
### Nested Schema for `links.endpoint1`

Required:

* `component_type` (String) - Type of the endpoint component. This is either 'Port' or 'GigaStream'
* `topo_node_id` (String) - Topology graph node Id. References one of the Topology Nodes in the graph

Optional:

* `comment` (String)
* `component_alias` (String) - For 'gigamon': this is either portId (e.g. '1/1/x1') or Gigastream alias (e.g. 'gs1').  For 'manual': this is user provided or may be omitted
* `port` (String) - For 'gigamon': this is the portId (e.g. '1/1/x1'). For 'manual': this is user provided or may be omitted
* `topo_node_type` (String) - topology node type

<a id="nestedatt--links--endpoint2"></a>
### Nested Schema for `links.endpoint2`

Required:

* `component_type` (String) - Type of the endpoint component. This is either 'Port' or 'GigaStream'
* `topo_node_id` (String) - Topology graph node Id. References one of the Topology Nodes in the graph

Optional:

* `comment` (String)
* `component_alias` (String) - For 'gigamon': this is either portId (e.g. '1/1/x1') or Gigastream alias (e.g. 'gs1').  For 'manual': this is user provided or may be omitted
* `port` (String) - For 'gigamon': this is the portId (e.g. '1/1/x1'). For 'manual': this is user provided or may be omitted
* `topo_node_type` (String) - topology node type

<a id="nestedatt--links--connections"></a>
### Nested Schema for `links.connections`

Optional:

* `port1` (String)
* `port2` (String)

<a id="nestedatt--nodes"></a>
### Nested Schema for `nodes`

Required:

* `node_alias` (String) - Node alias. User-assigned

Optional:

* `comment` (String)
* `model` (String) - Node model, if available
* `topo_node_id` (String) - Node's unique identifier. Auto-assigned. This is required while updating existing node
* `type` (String) - Annotation: type of a node. E.g. switch, router, etc...
* `vendor` (String) - Node vendor, if available

