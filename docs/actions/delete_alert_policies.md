---
page_title: "gigavuecore_delete_alert_policies Action - gigavuecore"
subcategory: ""
description: |-
  Delete all or given list of alert policies
---

# gigavuecore_delete_alert_policies Action

Delete all or given list of alert policies

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_alert_policies" "example" {
  config {
    drop_all     = true
    policy_names = ["example"]
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `drop_all` (Boolean, optional) - If list of policies are not provided this is required
* `policy_names` (List of String, required) - List of alert policies that needs to be deleted


