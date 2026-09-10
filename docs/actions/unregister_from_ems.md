---
page_title: "gigavuecore_unregister_from_ems Action - gigavuecore"
subcategory: ""
description: |-
  Unregister FM from EMS
---

# gigavuecore_unregister_from_ems Action

Unregister FM from EMS

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_unregister_from_ems" "example" {
  config {
  }
}
```
