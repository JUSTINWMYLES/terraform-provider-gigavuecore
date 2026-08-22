---
page_title: "gigavuecore_clear_hb_statistics_inline_tool Action - gigavuecore"
subcategory: ""
description: |-
  Clear heartbeat statistics of all Inline Tools
---

# gigavuecore_clear_hb_statistics_inline_tool Action

Clear heartbeat statistics of all Inline Tools

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
