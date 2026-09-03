---
page_title: "gigavuecore_clear_policy_status Action - gigavuecore"
subcategory: ""
description: |-
  Clear policy deployment status
---

# gigavuecore_clear_policy_status Action

Clear policy deployment status

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_clear_policy_status" "example" {
  config {
    name = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `name` (String, required) - policy name


