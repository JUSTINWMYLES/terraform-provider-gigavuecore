---
page_title: "gigavuecore_update_topology_link Action - gigavuecore"
subcategory: ""
description: |-
  Update Manual Topology Link
---

# gigavuecore_update_topology_link Action

Update Manual Topology Link

## Example Usage

```terraform
action "gigavuecore_update_topology_link" "example" {
  config {
    body_topo_link_id = "example"
    comment           = "example"
    connections       = null
    endpoint1         = "example"
    endpoint2         = "example"
    topo_link_id      = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `body_topo_link_id` (String, optional) - Link's unique identifier. Auto-assigned. This is required while updating existing links
* `comment` (String, optional)
* `connections` (List of Dynamic, optional) - Actual port ids which are part of the physical connection
* `endpoint1` (Dynamic, required) - Node link endpoint Create/Update Spec
* `endpoint2` (Dynamic, required) - Node link endpoint Create/Update Spec
* `topo_link_id` (String, required) - Topology Link Id


