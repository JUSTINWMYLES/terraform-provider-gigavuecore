---
page_title: "gigavuecore_get_system_arp_entries Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all ARP Entries
---

# gigavuecore_get_system_arp_entries Data Source

Get all ARP Entries

## Example Usage

```terraform
data "gigavuecore_get_system_arp_entries" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `address` (String) - ip address
* `age` (String) - Age in ISO-8601 partial-time format (time-hour : time-minute : time-second)
* `hw_address` (String) - hardware address
* `iface` (String) - interface name
* `state` (String)

