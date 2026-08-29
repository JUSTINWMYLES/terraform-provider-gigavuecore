---
page_title: "gigavuecore_get_system_arp_refresh_interval Data Source - gigavuecore"
subcategory: ""
description: |-
  Get System Arp Refresh Interval
---

# gigavuecore_get_system_arp_refresh_interval Data Source

Get System Arp Refresh Interval

## Example Usage

```terraform
data "gigavuecore_get_system_arp_refresh_interval" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `arp_refresh_interval` (Number, computed) - System Arp Refresh Interval in Seconds


