---
page_title: "gigavuecore_backup_fm_config Action - gigavuecore"
subcategory: ""
description: |-
  Backup FM configuration state
---

# gigavuecore_backup_fm_config Action

Backup FM configuration state

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_backup_fm_config" "example" {
  config {
    archive_server = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `archive_server` (String, required) - Alias of the target Archive Server


