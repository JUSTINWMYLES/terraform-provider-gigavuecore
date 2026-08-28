---
page_title: "gigavuecore_get_aggregate_links Action - gigavuecore"
subcategory: ""
description: |-
  Link speed and count information between 2 set of nodes having different tag values for a given key
---

# gigavuecore_get_aggregate_links Action

Link speed and count information between 2 set of nodes having different tag values for a given key

## Example Usage

```terraform
action "gigavuecore_get_aggregate_links" "example" {
  config {
    tags_filter = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `tags_filter` (Dynamic, required)


