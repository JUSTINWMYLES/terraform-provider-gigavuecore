---
page_title: "gigavuecore_delete_topology_link Action - gigavuecore"
subcategory: ""
description: |-
  Delete Manual Topology Link
---

# gigavuecore_delete_topology_link Action

Delete Manual Topology Link

## Example Usage

```terraform
action "gigavuecore_delete_topology_link" "example" {
  config {
    topo_link_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `topo_link_id` (String, required) - Topology Link Id


