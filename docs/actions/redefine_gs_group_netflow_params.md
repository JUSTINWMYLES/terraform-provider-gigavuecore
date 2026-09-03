---
page_title: "gigavuecore_redefine_gs_group_netflow_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Netflow Params
---

# gigavuecore_redefine_gs_group_netflow_params Action

Redefine GS Group's Netflow Params

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_netflow_params" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    monitor    = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `monitor` (String, optional) - Alias of referenced Netflow Monitor


