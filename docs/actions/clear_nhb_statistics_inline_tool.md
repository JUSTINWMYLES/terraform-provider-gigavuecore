---
page_title: "gigavuecore_clear_nhb_statistics_inline_tool Action - gigavuecore"
subcategory: ""
description: |-
  Clear negative heartbeat statistics of all Inline Tools
---

# gigavuecore_clear_nhb_statistics_inline_tool Action

Clear negative heartbeat statistics of all Inline Tools

## Example Usage

```terraform
action "gigavuecore_clear_nhb_statistics_inline_tool" "example" {
  config {
    cluster_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID


