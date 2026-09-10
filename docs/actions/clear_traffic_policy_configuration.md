---
page_title: "gigavuecore_clear_traffic_policy_configuration Action - gigavuecore"
subcategory: ""
description: |-
  Clear Prefiltering Policy configuration
---

# gigavuecore_clear_traffic_policy_configuration Action

Clear Prefiltering Policy configuration

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_clear_traffic_policy_configuration" "example" {
  config {
    id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `id` (String, required) - policy graph id


