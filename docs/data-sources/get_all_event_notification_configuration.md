---
page_title: "gigavuecore_get_all_event_notification_configuration Data Source - gigavuecore"
subcategory: ""
description: |-
  Get All Event Notification Configurations
---

# gigavuecore_get_all_event_notification_configuration Data Source

Get All Event Notification Configurations

## Example Usage

```terraform
data "gigavuecore_get_all_event_notification_configuration" "example" {
  event_details = null
  external_trap_receivers = null
  recipients = null
  tags = null
  task_name = null
  time_interval = null
  type = null
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

### Attributes

In addition to all arguments above, the following attributes are exported:

* `event_notif_tasks` (List(Object({allow_attachment, attachment_limit, comment, email_subject_prefix, enabled, event_details, external_trap_receivers, instant_rate_limit, recipients, recurring_schedule, send_mail_if_empty, severity, tags, task_id, task_name, template_details, time_interval, time_left, type})), computed)
* `instant_rate_limit` (Number, computed) - Maximum number of instant emails that can be sent per minute

