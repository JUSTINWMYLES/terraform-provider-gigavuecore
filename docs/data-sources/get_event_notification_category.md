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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `name` (String) - Scope Name
* `scope` (String) - Event Scope
* `sub_types` (Attributes List) (see [below for nested schema](#nestedatt--items--sub_types))
<a id="nestedatt--items--sub_types"></a>
### Nested Schema for `items.sub_types`

Read-Only:

* `event_types` (Attributes List) (see [below for nested schema](#nestedatt--items--sub_types--event_types))
* `name` (String) - Subtype Name
* `sub_type` (String) - Event Subtype
<a id="nestedatt--items--sub_types--event_types"></a>
### Nested Schema for `items.sub_types.event_types`

Read-Only:

* `description` (String) - Event Description
* `display_name` (String) - Event Display Name
* `event_type` (String) - Event Type
* `name` (String) - Event Name

