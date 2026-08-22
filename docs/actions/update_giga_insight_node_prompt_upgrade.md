---
page_title: "gigavuecore_update_giga_insight_node_prompt_upgrade Action - gigavuecore"
subcategory: ""
description: |-
  Upload prompt bundle upgrade for a GigaInsight Node
---

# gigavuecore_update_giga_insight_node_prompt_upgrade Action

Upload prompt bundle upgrade for a GigaInsight Node

## Example Usage

```terraform
action "gigavuecore_update_giga_insight_node_prompt_upgrade" "example" {
  config {
    bundle = "example"
    node_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `bundle` (String, optional)
* `node_id` (String, required)
