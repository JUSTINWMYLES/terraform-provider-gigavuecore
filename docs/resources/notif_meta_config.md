---
page_title: "gigavuecore_notif_meta_config Resource - gigavuecore"
subcategory: ""
description: |-
  Get Notification Configuration
---

# gigavuecore_notif_meta_config Resource

Get Notification Configuration

## Example Usage

```terraform
resource "gigavuecore_notif_meta_config" "example" {
  allow_attachment = null
  attachment_limit = null
  comment = null
  email_subject_prefix = null
  enabled = null
  event_details = []
  external_trap_receivers = []
  instant_rate_limit = null
  recipients = []
  recurring_schedule = null
  send_mail_if_empty = null
  severity = []
  tags = []
  task_id = null
  task_name = null
  template_details = []
  time_interval = null
  time_left = null
  type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `allow_attachment` (Bool, optional) - Allow attachment in notification email
* `attachment_limit` (Number, optional) - Maximum number of events to be included in the attachment
* `comment` (String, optional) - Comments
* `email_subject_prefix` (String, optional) - Subject of the notification email
* `enabled` (Bool, optional) - Status of the notification task
* `event_details` (List(Object({description, display_name, event_type, name, scope, severity, severity_type, sub_type})), required) - Event Details
* `external_trap_receivers` (List(String), optional) - List of External Trap Receiver aliases
* `instant_rate_limit` (Number, optional) - Maximum number of instant emails that can be sent per minute
* `recipients` (List(String), optional) - List of email recipients
* `recurring_schedule` (String, optional) - Cron expression
* `send_mail_if_empty` (Bool, optional) - Send email if no events are generated within time interval
* `severity` (List(String), optional) - Event Severities
* `tags` (List(Object({tag_key, tag_values})), optional) - Tags
* `task_id` (String, required) - ID of the notification task
* `task_name` (String, required) - Name of the task
* `template_details` (List(String), optional) - List of templates used
* `time_interval` (Number, optional) - Time interval between two batch emails
* `time_left` (String, optional) - Time left for the next batch task to execute
* `type` (String, required) - Type of the event notification task

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `allow_attachment` (Bool, computed) - Allow attachment in notification email
* `attachment_limit` (Number, computed) - Maximum number of events to be included in the attachment
* `comment` (String, computed) - Comments
* `email_subject_prefix` (String, computed) - Subject of the notification email
* `enabled` (Bool, computed) - Status of the notification task
* `external_trap_receivers` (List(String), computed) - List of External Trap Receiver aliases
* `id` (String, computed)
* `instant_rate_limit` (Number, computed) - Maximum number of instant emails that can be sent per minute
* `recipients` (List(String), computed) - List of email recipients
* `recurring_schedule` (String, computed) - Cron expression
* `send_mail_if_empty` (Bool, computed) - Send email if no events are generated within time interval
* `severity` (List(String), computed) - Event Severities
* `tags` (List(Object({tag_key, tag_values})), computed) - Tags
* `template_details` (List(String), computed) - List of templates used
* `time_interval` (Number, computed) - Time interval between two batch emails
* `time_left` (String, computed) - Time left for the next batch task to execute

