---
page_title: "gigavuecore_clear_task_status Action - gigavuecore"
subcategory: ""
description: |-
  Clear task status by cluster IDs.
---

# gigavuecore_clear_task_status Action

Clear task status by cluster IDs.

## Example Usage

```terraform
action "gigavuecore_clear_task_status" "example" {
  config {
    cluster_ids = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_ids` (String, required) - The requested cluster IDs task status will be cleared
