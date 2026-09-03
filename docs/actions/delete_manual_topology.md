---
page_title: "gigavuecore_delete_manual_topology Action - gigavuecore"
subcategory: ""
description: |-
  Delete all manually-augmented topology elements
---

# gigavuecore_delete_manual_topology Action

Delete all manually-augmented topology elements

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_manual_topology" "example" {
  config {
  }
}
```
