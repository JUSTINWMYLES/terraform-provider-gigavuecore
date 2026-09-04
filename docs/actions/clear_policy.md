---
page_title: "gigavuecore_clear_policy Action - gigavuecore"
subcategory: ""
description: |-
  clear all policy
---

# gigavuecore_clear_policy Action

clear all policy

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_clear_policy" "example" {
  config {
  }
}
```
