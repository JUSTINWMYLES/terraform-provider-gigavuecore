---
page_title: "gigavuecore_redefine_system_arp_refresh_interval Action - gigavuecore"
subcategory: ""
description: |-
  Redefine System Arp's Refresh Interval
---

# gigavuecore_redefine_system_arp_refresh_interval Action

Redefine System Arp's Refresh Interval

## Example Usage

```terraform
action "gigavuecore_redefine_system_arp_refresh_interval" "example" {
  config {
    arp_refresh_interval = 1
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `arp_refresh_interval` (Number, optional) - System Arp Refresh Interval in Seconds


