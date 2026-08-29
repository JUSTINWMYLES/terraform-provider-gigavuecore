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
  event_details           = "example"
  external_trap_receivers = "example"
  recipients              = "example"
  tags                    = "example"
  task_name               = "example"
  time_interval           = "example"
  type                    = "example"
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

* `event_notif_tasks` (Attributes List, computed) (see [below for nested schema](#nestedatt--event_notif_tasks))
* `instant_rate_limit` (Number, computed) - Maximum number of instant emails that can be sent per minute

<a id="nestedatt--event_notif_tasks"></a>
### Nested Schema for `event_notif_tasks`

Read-Only:

* `allow_attachment` (Boolean) - Allow attachment in notification email
* `attachment_limit` (Number) - Maximum number of events to be included in the attachment
* `comment` (String) - Comments
* `email_subject_prefix` (String) - Subject of the notification email
* `enabled` (Boolean) - Status of the notification task
* `event_details` (Attributes List) - Event Details (see [below for nested schema](#nestedatt--event_notif_tasks--event_details))
* `external_trap_receivers` (List of String) - List of External Trap Receiver aliases
* `instant_rate_limit` (Number) - Maximum number of instant emails that can be sent per minute
* `recipients` (List of String) - List of email recipients
* `recurring_schedule` (String) - Cron expression
* `send_mail_if_empty` (Boolean) - Send email if no events are generated within time interval
* `severity` (List of String) - Event Severities
* `tags` (Attributes List) - Tags (see [below for nested schema](#nestedatt--event_notif_tasks--tags))
* `task_id` (String) - ID of the notification task
* `task_name` (String) - Name of the task
* `template_details` (List of String) - List of templates used
* `time_interval` (Number) - Time interval between two batch emails
* `time_left` (String) - Time left for the next batch task to execute
* `type` (String) - Type of the event notification task
<a id="nestedatt--event_notif_tasks--event_details"></a>
### Nested Schema for `event_notif_tasks.event_details`

Read-Only:

* `description` (String) - Event Description
* `display_name` (String) - Event Display Name
* `event_type` (String) - Event Type
* `name` (String) - Event Name
* `scope` (String) - Event Scope
* `severity` (List of String) - Event Severity
* `severity_type` (String) - Event Severity Type
* `sub_type` (String) - Event Subtype
<a id="nestedatt--event_notif_tasks--tags"></a>
### Nested Schema for `event_notif_tasks.tags`

Read-Only:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag

