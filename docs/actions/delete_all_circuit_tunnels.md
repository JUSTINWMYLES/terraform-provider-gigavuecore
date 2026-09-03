---
page_title: "gigavuecore_delete_all_circuit_tunnels Action - gigavuecore"
subcategory: ""
description: |-
  Delete all Circuit Tunnels
---

# gigavuecore_delete_all_circuit_tunnels Action

Delete all Circuit Tunnels

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_all_circuit_tunnels" "example" {
  config {
    cluster_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID


