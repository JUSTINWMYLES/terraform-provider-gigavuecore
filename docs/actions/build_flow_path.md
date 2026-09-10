---
page_title: "gigavuecore_build_flow_path Action - gigavuecore"
subcategory: ""
description: |-
  Build traffic flow path for a cluster
---

# gigavuecore_build_flow_path Action

Build traffic flow path for a cluster

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_build_flow_path" "example" {
  config {
    cluster_name = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_name` (String, required) - Target Cluster Name


