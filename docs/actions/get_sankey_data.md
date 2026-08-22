---
page_title: "gigavuecore_get_sankey_data Action - gigavuecore"
subcategory: ""
description: |-
  Response for sankey representation
---

# gigavuecore_get_sankey_data Action

Response for sankey representation

## Example Usage

```terraform
action "gigavuecore_get_sankey_data" "example" {
  config {
    nodes_filter = null
    tags_filter = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `nodes_filter` (List(Dynamic), optional)
* `tags_filter` (Dynamic, required)
