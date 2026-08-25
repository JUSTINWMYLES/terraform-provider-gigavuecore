---
page_title: "gigavuecore_enable_batch_event_notification_configuration Action - gigavuecore"
subcategory: ""
description: |-
  Enable Batch Notification Configuration
---

# gigavuecore_enable_batch_event_notification_configuration Action

Enable Batch Notification Configuration

## Example Usage

```terraform
action "gigavuecore_enable_batch_event_notification_configuration" "example" {
  config {
    task_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `task_id` (String, required) - ID of the Task


