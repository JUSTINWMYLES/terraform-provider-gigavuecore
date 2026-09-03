---
page_title: "gigavuecore_redefine_system_arp_refresh_interval Action - gigavuecore"
subcategory: ""
description: |-
  Redefine System Arp's Refresh Interval
---

# gigavuecore_redefine_system_arp_refresh_interval Action

Redefine System Arp's Refresh Interval

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_system_arp_refresh_interval" "example" {
  config {
    arp_refresh_interval = 3
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `arp_refresh_interval` (Number, optional) - System Arp Refresh Interval in Seconds


