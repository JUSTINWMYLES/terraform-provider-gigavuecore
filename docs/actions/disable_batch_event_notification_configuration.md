---
page_title: "gigavuecore_disable_batch_event_notification_configuration Action - gigavuecore"
subcategory: ""
description: |-
  Disable Batch Notification Configuration
---

# gigavuecore_disable_batch_event_notification_configuration Action

Disable Batch Notification Configuration

## Example Usage

```terraform
action "gigavuecore_disable_batch_event_notification_configuration" "example" {
  config {
    task_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `task_id` (String, required) - ID of the Task
