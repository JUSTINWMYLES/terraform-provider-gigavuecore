---
page_title: "gigavuecore_create_event_notification_configuration Action - gigavuecore"
subcategory: ""
description: |-
  Create Notification Configuration
---

# gigavuecore_create_event_notification_configuration Action

Create Notification Configuration

## Example Usage

```terraform
action "gigavuecore_create_event_notification_configuration" "example" {
  config {
    allow_attachment = true
    attachment_limit = 1
    comment = "example"
    email_subject_prefix = "example"
    enabled = true
    event_details = "example"
    external_trap_receivers = [ "example" ]
    instant_rate_limit = 1
    notif_type = "example"
    recipients = [ "example" ]
    recurring_schedule = "example"
    send_mail_if_empty = true
    severity = [ "example" ]
    tags = null
    task_id = "example"
    task_name = "example"
    template_details = [ "example" ]
    time_interval = 1
    time_left = "example"
    type = "example"
  }
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
* `event_details` (List(Dynamic), required) - Event Details
* `external_trap_receivers` (List(String), optional) - List of External Trap Receiver aliases
* `instant_rate_limit` (Number, optional) - Maximum number of instant emails that can be sent per minute
* `notif_type` (String, required) - Type of Notification
* `recipients` (List(String), optional) - List of email recipients
* `recurring_schedule` (String, optional) - Cron expression
* `send_mail_if_empty` (Bool, optional) - Send email if no events are generated within time interval
* `severity` (List(String), optional) - Event Severities
* `tags` (List(Dynamic), optional) - Tags
* `task_id` (String, optional) - ID of the notification task
* `task_name` (String, required) - Name of the task
* `template_details` (List(String), optional) - List of templates used
* `time_interval` (Number, optional) - Time interval between two batch emails
* `time_left` (String, optional) - Time left for the next batch task to execute
* `type` (String, required) - Type of the event notification task
