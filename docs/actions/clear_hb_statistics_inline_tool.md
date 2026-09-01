---
page_title: "gigavuecore_clear_hb_statistics_inline_tool Action - gigavuecore"
subcategory: ""
description: |-
  Clear heartbeat statistics of all Inline Tools
---

# gigavuecore_clear_hb_statistics_inline_tool Action

Clear heartbeat statistics of all Inline Tools

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_clear_hb_statistics_inline_tool" "example" {
  config {
    cluster_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID


