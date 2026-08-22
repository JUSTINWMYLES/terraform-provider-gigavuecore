---
page_title: "gigavuecore_recover_inline_tool Action - gigavuecore"
subcategory: ""
description: |-
  Recover Inline Tool
---

# gigavuecore_recover_inline_tool Action

Recover Inline Tool

## Example Usage

```terraform
action "gigavuecore_recover_inline_tool" "example" {
  config {
    alias = "example"
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of Inline Tool to redefine
* `cluster_id` (String, required) - Target Cluster ID
