---
page_title: "gigavuecore_delete_all_enhanced_slicing Action - gigavuecore"
subcategory: ""
description: |-
  new in version H 5.7
---

# gigavuecore_delete_all_enhanced_slicing Action

new in version H 5.7

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_all_enhanced_slicing" "example" {
  config {
  }
}
```
