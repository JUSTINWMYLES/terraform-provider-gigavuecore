---
page_title: "gigavuecore_delete_avisi_policy_condition Action - gigavuecore"
subcategory: ""
description: |-
  Delete Active Visibility Policy Condition
---

# gigavuecore_delete_avisi_policy_condition Action

Delete Active Visibility Policy Condition

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_avisi_policy_condition" "example" {
  config {
    condition_id = "example"
    policy_id    = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `condition_id` (String, required) - Policy Condition Id
* `policy_id` (String, required) - Policy Id


