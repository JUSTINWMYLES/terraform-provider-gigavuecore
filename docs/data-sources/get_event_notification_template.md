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

* `items` (List(Object({event_details, template_name, template_type})), computed)

