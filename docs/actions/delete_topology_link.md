---
page_title: "gigavuecore_delete_topology_link Action - gigavuecore"
subcategory: ""
description: |-
  Delete Manual Topology Link
---

# gigavuecore_delete_topology_link Action

Delete Manual Topology Link

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

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


