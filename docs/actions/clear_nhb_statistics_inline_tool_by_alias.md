---
page_title: "gigavuecore_clear_nhb_statistics_inline_tool_by_alias Action - gigavuecore"
subcategory: ""
description: |-
  Clear negative heartbeat statistics of Inline Tool by alias
---

# gigavuecore_clear_nhb_statistics_inline_tool_by_alias Action

Clear negative heartbeat statistics of Inline Tool by alias

## Example Usage

```terraform
action "gigavuecore_clear_nhb_statistics_inline_tool_by_alias" "example" {
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


