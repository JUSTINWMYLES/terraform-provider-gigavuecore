---
page_title: "gigavuecore_delete_flex_inline_map_config Action - gigavuecore"
subcategory: ""
description: |-
  Deletes the map configs from the given solution
---

# gigavuecore_delete_flex_inline_map_config Action

Deletes the map configs from the given solution

## Example Usage

```terraform
action "gigavuecore_delete_flex_inline_map_config" "example" {
  config {
    alias = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Deletes maps for the given solution


