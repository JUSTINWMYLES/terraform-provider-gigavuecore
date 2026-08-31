---
page_title: "gigavuecore_reset_topology_viz Action - gigavuecore"
subcategory: ""
description: |-
  resets topology data by reconstructing links and node neighbour details
---

# gigavuecore_reset_topology_viz Action

resets topology data by reconstructing links and node neighbour details

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_reset_topology_viz" "example" {
  config {
  }
}
```
