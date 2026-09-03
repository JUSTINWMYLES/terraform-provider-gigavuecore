---
page_title: "gigavuecore_create_manual_topology_node Action - gigavuecore"
subcategory: ""
description: |-
  Returns generated Topology Node Id for future reference
---

# gigavuecore_create_manual_topology_node Action

Returns generated Topology Node Id for future reference

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_create_manual_topology_node" "example" {
  config {
    comment      = "example"
    model        = "example"
    node_alias   = "example"
    topo_node_id = "example"
    type         = "example"
    vendor       = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `comment` (String, optional)
* `model` (String, optional) - Node model, if available
* `node_alias` (String, required) - Node alias. User-assigned
* `topo_node_id` (String, optional) - Node's unique identifier. Auto-assigned. This is required while updating existing node
* `type` (String, optional) - Annotation: type of a node. E.g. switch, router, etc...
* `vendor` (String, optional) - Node vendor, if available


