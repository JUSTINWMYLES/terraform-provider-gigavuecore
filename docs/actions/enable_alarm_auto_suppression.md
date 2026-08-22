---
page_title: "gigavuecore_enable_alarm_auto_suppression Action - gigavuecore"
subcategory: ""
description: |-
  Edit auto alarm suppression rules
---

# gigavuecore_enable_alarm_auto_suppression Action

Edit auto alarm suppression rules

## Example Usage

```terraform
action "gigavuecore_enable_alarm_auto_suppression" "example" {
  config {
    auto_suppression_rules = "example"
    enable = true
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `auto_suppression_rules` (List(Dynamic), required) - Alarm Auto Suppression Rules
* `enable` (Bool, required) - Enable/Disable auto alarm suppression rules for operation Image Upgrade,Config Restore,Device Reboot and Cluster Operation
