---
page_title: "gigavuecore_update_flex_inline_nrt_stats_config Action - gigavuecore"
subcategory: ""
description: |-
  Register FlexInline solution to NRT stats
---

# gigavuecore_update_flex_inline_nrt_stats_config Action

Register FlexInline solution to NRT stats

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_update_flex_inline_nrt_stats_config" "example" {
  config {
    operation_type = "example"
    solution_alias = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `operation_type` (String, required) - adds flexInline solution to NRT
* `solution_alias` (String, required) - Register flexinline solution to NRT


