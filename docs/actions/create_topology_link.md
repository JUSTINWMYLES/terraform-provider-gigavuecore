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
    connections = null
    endpoint1 = "example"
    endpoint2 = "example"
    topo_link_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `comment` (String, optional)
* `connections` (List(Dynamic), optional) - Actual port ids which are part of the physical connection
* `endpoint1` (Dynamic, required) - Node link endpoint Create/Update Spec
* `endpoint2` (Dynamic, required) - Node link endpoint Create/Update Spec
* `topo_link_id` (String, optional) - Link's unique identifier. Auto-assigned. This is required while updating existing links
