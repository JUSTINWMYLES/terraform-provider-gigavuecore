---
page_title: "gigavuecore_delete_by_scroll_id Action - gigavuecore"
subcategory: ""
description: |-
  Delete by ScrollId
---

# gigavuecore_delete_by_scroll_id Action

Delete by ScrollId

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

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


