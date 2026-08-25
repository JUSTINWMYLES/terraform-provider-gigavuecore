---
page_title: "gigavuecore_update_manual_topology_entity Action - gigavuecore"
subcategory: ""
description: |-
  Update Manual Topology Nodes and Links
---

# gigavuecore_update_manual_topology_entity Action

Update Manual Topology Nodes and Links

## Example Usage

```terraform
action "gigavuecore_update_manual_topology_entity" "example" {
  config {
    links = null
    nodes = null
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `links` (List of Dynamic, optional)
* `nodes` (List of Dynamic, optional)


