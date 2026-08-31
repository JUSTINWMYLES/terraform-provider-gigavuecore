---
page_title: "gigavuecore_enable_alarm_auto_suppression Action - gigavuecore"
subcategory: ""
description: |-
  Edit auto alarm suppression rules
---

# gigavuecore_enable_alarm_auto_suppression Action

Edit auto alarm suppression rules

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_enable_alarm_auto_suppression" "example" {
  config {
    auto_suppression_rules = [{
      enable = true
      type   = "image_upgrade"
    }]
    enable = true
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `auto_suppression_rules` (Attributes List, required) - Alarm Auto Suppression Rules (see [below for nested schema](#nestedatt--auto_suppression_rules))
* `enable` (Boolean, required) - Enable/Disable auto alarm suppression rules for operation Image Upgrade,Config Restore,Device Reboot and Cluster Operation

<a id="nestedatt--auto_suppression_rules"></a>
### Nested Schema for `auto_suppression_rules`

Required:

* `enable` (Boolean)
* `type` (String) - Type of operation on Cluster/Node

