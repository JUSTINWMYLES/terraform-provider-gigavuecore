---
page_title: "gigavuecore_get_event_notification_template Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Event Notification Template
---

# gigavuecore_get_event_notification_template Data Source

Get Event Notification Template

## Example Usage

```terraform
data "gigavuecore_get_event_notification_template" "example" {
  template_type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `template_type` (String, required) - Type of tempalte

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `event_details` (Attributes List) (see [below for nested schema](#nestedatt--items--event_details))
* `template_name` (String)
* `template_type` (String) - Template Type
<a id="nestedatt--items--event_details"></a>
### Nested Schema for `items.event_details`

Read-Only:

* `description` (String) - Event Description
* `display_name` (String) - Event Display Name
* `event_type` (String) - Event Type
* `name` (String) - Event Name
* `scope` (String) - Event Scope
* `severity` (List of String) - Event Severity
* `severity_type` (String) - Event Severity Type
* `sub_type` (String) - Event Subtype

