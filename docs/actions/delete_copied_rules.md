---
page_title: "gigavuecore_delete_copied_rules Action - gigavuecore"
subcategory: ""
description: |-
  Remove all previously copied rules from the user-specific clipboard. This operation clears the entire clipboard for the authenticated user.
---

# gigavuecore_delete_copied_rules Action

Remove all previously copied rules from the user-specific clipboard.
This operation clears the entire clipboard for the authenticated user.

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_copied_rules" "example" {
  config {
  }
}
```
