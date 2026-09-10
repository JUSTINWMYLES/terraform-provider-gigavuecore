---
page_title: "gigavuecore_delete_all_gta_profiles Action - gigavuecore"
subcategory: ""
description: |-
  delete all GTA profiles
---

# gigavuecore_delete_all_gta_profiles Action

delete all GTA profiles

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_all_gta_profiles" "example" {
  config {
  }
}
```
