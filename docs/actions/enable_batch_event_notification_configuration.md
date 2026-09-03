---
page_title: "gigavuecore_enable_batch_event_notification_configuration Action - gigavuecore"
subcategory: ""
description: |-
  Enable Batch Notification Configuration
---

# gigavuecore_enable_batch_event_notification_configuration Action

Enable Batch Notification Configuration

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

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


