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

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `banner_config_data` (Attributes) (see [below for nested schema](#nestedatt--items--banner_config_data))
* `description` (String) - Notification description
* `severity` (String) - Notification Severity
* `type` (String) - Notification type
<a id="nestedatt--items--banner_config_data"></a>
### Nested Schema for `items.banner_config_data`

Read-Only:

* `current_usage_percent` (Number)
* `disk_partition` (String) - Disk Partition
* `host_name` (String) - Host Name
* `role` (String) - Node Role
* `threshold_percent` (Number)
* `total_gb` (Number)

