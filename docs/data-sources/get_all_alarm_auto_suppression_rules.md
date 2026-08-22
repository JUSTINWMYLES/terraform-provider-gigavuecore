---
page_title: "gigavuecore_get_all_alarm_auto_suppression_rules Data Source - gigavuecore"
subcategory: ""
description: |-
  Get All Alarm Auto Supppression Rules
---

# gigavuecore_get_all_alarm_auto_suppression_rules Data Source

Get All Alarm Auto Supppression Rules

## Example Usage

```terraform
data "gigavuecore_get_all_alarm_auto_suppression_rules" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `auto_suppression_rules` (List(Object({enable, type})), computed) - Alarm Auto Suppression Rules
* `enable` (Bool, computed) - Enable/Disable auto alarm suppression rules for operation Image Upgrade,Config Restore,Device Reboot and Cluster Operation

