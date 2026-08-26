---
page_title: "gigavuecore_delete_by_scroll_id Action - gigavuecore"
subcategory: ""
description: |-
  Delete by ScrollId
---

# gigavuecore_delete_by_scroll_id Action

Delete by ScrollId

## Example Usage

```terraform
action "gigavuecore_delete_by_scroll_id" "example" {
  config {
    scroll_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `scroll_id` (String, required) - scrollId to clear the scroll


