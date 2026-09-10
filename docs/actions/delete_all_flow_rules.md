---
page_title: "gigavuecore_delete_all_flow_rules Action - gigavuecore"
subcategory: ""
description: |-
  Delete all flow rules for a subflow
---

# gigavuecore_delete_all_flow_rules Action

Delete all flow rules for a subflow

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_all_flow_rules" "example" {
  config {
    flow_alias     = "example"
    policy_alias   = "example"
    sub_flow_alias = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `flow_alias` (String, required) - Flow identifier
* `policy_alias` (String, required) - Policy identifier
* `sub_flow_alias` (String, required) - Subflow identifier


