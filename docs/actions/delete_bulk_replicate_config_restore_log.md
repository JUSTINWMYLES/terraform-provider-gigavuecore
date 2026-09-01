---
page_title: "gigavuecore_delete_bulk_replicate_config_restore_log Action - gigavuecore"
subcategory: ""
description: |-
  Delete bulk replicate config restore log
---

# gigavuecore_delete_bulk_replicate_config_restore_log Action

Delete bulk replicate config restore log

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_bulk_replicate_config_restore_log" "example" {
  config {
    filename = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `filename` (String, required) - Target restore log file name


