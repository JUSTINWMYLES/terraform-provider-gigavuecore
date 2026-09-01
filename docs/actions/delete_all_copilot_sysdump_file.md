---
page_title: "gigavuecore_delete_all_copilot_sysdump_file Action - gigavuecore"
subcategory: ""
description: |-
  Delete all Copilot Sysdump Files
---

# gigavuecore_delete_all_copilot_sysdump_file Action

Delete all Copilot Sysdump Files

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_all_copilot_sysdump_file" "example" {
  config {
  }
}
```
