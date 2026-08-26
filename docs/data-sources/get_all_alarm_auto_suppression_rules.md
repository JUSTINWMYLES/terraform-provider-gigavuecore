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

### Attributes

In addition to all arguments above, the following attributes are exported:

* `auto_suppression_rules` (Attributes List, computed) - Alarm Auto Suppression Rules (see [below for nested schema](#nestedatt--auto_suppression_rules))
* `enable` (Boolean, computed) - Enable/Disable auto alarm suppression rules for operation Image Upgrade,Config Restore,Device Reboot and Cluster Operation

<a id="nestedatt--auto_suppression_rules"></a>
### Nested Schema for `auto_suppression_rules`

Read-Only:

* `enable` (Boolean)
* `type` (String) - Type of operation on Cluster/Node

