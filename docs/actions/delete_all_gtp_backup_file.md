---
page_title: "gigavuecore_delete_all_gtp_backup_file Action - gigavuecore"
subcategory: ""
description: |-
  Delete all Gtp backup files
---

# gigavuecore_delete_all_gtp_backup_file Action

Delete all Gtp backup files

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_all_gtp_backup_file" "example" {
  config {
  }
}
```
