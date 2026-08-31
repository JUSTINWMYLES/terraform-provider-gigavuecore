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
    allow_attachment     = true
    attachment_limit     = 0
    comment              = "example"
    email_subject_prefix = "example"
    enabled              = true
    event_details = [{
      description   = "example"
      display_name  = "example"
      event_type    = "example"
      name          = "example"
      scope         = "example"
      severity      = [ "Clear" ]
      severity_type = "critical"
      sub_type      = "example"
    }]
    external_trap_receivers = [ "example" ]
    instant_rate_limit      = 0
    notif_type              = "example"
    recipients              = [ "example" ]
    recurring_schedule      = "example"
    send_mail_if_empty      = true
    severity                = [ "Clear" ]
    tags = [{
      tag_key    = "example"
      tag_values = [ "example" ]
    }]
    task_id          = "example"
    task_name        = "example"
    template_details = [ "example" ]
    time_interval    = 1
    time_left        = "example"
    type             = "instant"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `allow_attachment` (Boolean, optional) - Allow attachment in notification email
* `attachment_limit` (Number, optional) - Maximum number of events to be included in the attachment
* `comment` (String, optional) - Comments
* `email_subject_prefix` (String, optional) - Subject of the notification email
* `enabled` (Boolean, optional) - Status of the notification task
* `event_details` (Attributes List, required) - Event Details (see [below for nested schema](#nestedatt--event_details))
* `external_trap_receivers` (List of String, optional) - List of External Trap Receiver aliases
* `instant_rate_limit` (Number, optional) - Maximum number of instant emails that can be sent per minute
* `notif_type` (String, required) - Type of Notification
* `recipients` (List of String, optional) - List of email recipients
* `recurring_schedule` (String, optional) - Cron expression
* `send_mail_if_empty` (Boolean, optional) - Send email if no events are generated within time interval
* `severity` (List of String, optional) - Event Severities
* `tags` (Attributes List, optional) - Tags (see [below for nested schema](#nestedatt--tags))
* `task_id` (String, optional) - ID of the notification task
* `task_name` (String, required) - Name of the task
* `template_details` (List of String, optional) - List of templates used
* `time_interval` (Number, optional) - Time interval between two batch emails
* `time_left` (String, optional) - Time left for the next batch task to execute
* `type` (String, required) - Type of the event notification task

<a id="nestedatt--event_details"></a>
### Nested Schema for `event_details`

Optional:

* `description` (String) - Event Description
* `display_name` (String) - Event Display Name
* `event_type` (String) - Event Type
* `name` (String) - Event Name
* `scope` (String) - Event Scope
* `severity` (List of String) - Event Severity
* `severity_type` (String) - Event Severity Type
* `sub_type` (String) - Event Subtype

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag

