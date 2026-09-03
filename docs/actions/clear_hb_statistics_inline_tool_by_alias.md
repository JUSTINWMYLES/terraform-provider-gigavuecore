---
page_title: "gigavuecore_clear_hb_statistics_inline_tool_by_alias Action - gigavuecore"
subcategory: ""
description: |-
  Clear heartbeat statistics of Inline Tool by alias
---

# gigavuecore_clear_hb_statistics_inline_tool_by_alias Action

Clear heartbeat statistics of Inline Tool by alias

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_clear_hb_statistics_inline_tool_by_alias" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target Inline Tool
* `cluster_id` (String, required) - Target Cluster ID


