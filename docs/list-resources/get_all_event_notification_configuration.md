---
page_title: "gigavuecore_get_all_event_notification_configuration List Resource - gigavuecore"
subcategory: ""
description: |-
  Get All Event Notification Configurations
---

# gigavuecore_get_all_event_notification_configuration List Resource

Get All Event Notification Configurations

## Example Usage

```terraform
list "gigavuecore_get_all_event_notification_configuration" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    event_details           = "example"
    external_trap_receivers = "example"
    recipients              = "example"
    tags                    = "example"
    task_name               = "example"
    time_interval           = "example"
    type                    = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `event_details` (String, optional) - eventDetails
* `external_trap_receivers` (String, optional) - Alias of External Trap Receivers
* `recipients` (String, optional) - Email Recipeints
* `tags` (String, optional) - tags
* `task_name` (String, optional) - Name of the task
* `time_interval` (String, optional) - timeInterval
* `type` (String, optional) - Type of the task


