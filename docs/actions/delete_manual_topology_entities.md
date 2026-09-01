---
page_title: "gigavuecore_delete_manual_topology_entities Action - gigavuecore"
subcategory: ""
description: |-
  Delete Manual Topology Nodes ans Links
---

# gigavuecore_delete_manual_topology_entities Action

Delete Manual Topology Nodes ans Links

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_manual_topology_entities" "example" {
  config {
    links = ["example"]
    nodes = ["example"]
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `links` (List of String, optional) - Array of ids of links to be deleted
* `nodes` (List of String, optional) - Array of ids of nodes to be deleted


