---
page_title: "gigavuecore_delete_manual_topology_entities Action - gigavuecore"
subcategory: ""
description: |-
  Delete Manual Topology Nodes ans Links
---

# gigavuecore_delete_manual_topology_entities Action

Delete Manual Topology Nodes ans Links

## Example Usage

```terraform
action "gigavuecore_delete_manual_topology_entities" "example" {
  config {
    links = [ "example" ]
    nodes = [ "example" ]
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `links` (List(String), optional) - Array of ids of links to be deleted
* `nodes` (List(String), optional) - Array of ids of nodes to be deleted
