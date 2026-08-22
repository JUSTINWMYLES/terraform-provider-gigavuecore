---
page_title: "gigavuecore_get_event_notification_category Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Event Notification Category
---

# gigavuecore_get_event_notification_category Data Source

Get Event Notification Category

## Example Usage

```terraform
data "gigavuecore_get_event_notification_category" "example" {
  category_type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `category_type` (String, required) - Type of category

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({name, scope, sub_types})), computed)

