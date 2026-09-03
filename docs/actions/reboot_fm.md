---
page_title: "gigavuecore_reboot_fm Action - gigavuecore"
subcategory: ""
description: |-
  Reboot FM
---

# gigavuecore_reboot_fm Action

Reboot FM

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_reboot_fm" "example" {
  config {
  }
}
```
