---
page_title: "gigavuecore_bulk_policy_enable Action - gigavuecore"
subcategory: ""
description: |-
  Enable all policies or by policy Ids
---

# gigavuecore_bulk_policy_enable Action

Enable all policies or by policy Ids

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_bulk_policy_enable" "example" {
  config {
    enabled    = true
    policy_ids = ["example"]
    type       = "all"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `enabled` (Boolean, required) - enable/disable policies
* `policy_ids` (List of String, optional) - list of policy Ids
* `type` (String, required) - Enable all policies or by policy Ids


