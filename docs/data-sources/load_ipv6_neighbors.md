---
page_title: "gigavuecore_load_ipv6_neighbors Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Ipv6 Neighbors
---

# gigavuecore_load_ipv6_neighbors Data Source

Load Ipv6 Neighbors

## Example Usage

```terraform
data "gigavuecore_load_ipv6_neighbors" "example" {
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

