---
page_title: "gigavuecore_load_all_critical_notifications Data Source - gigavuecore"
subcategory: ""
description: |-
  Load All Critical Notifications
---

# gigavuecore_load_all_critical_notifications Data Source

Load All Critical Notifications

## Example Usage

```terraform
data "gigavuecore_load_all_critical_notifications" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({banner_config_data, description, severity, type})), computed)

