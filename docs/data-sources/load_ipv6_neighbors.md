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

* `arp_entries` (Attributes List, computed) (see [below for nested schema](#nestedatt--arp_entries))
* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))

<a id="nestedatt--arp_entries"></a>
### Nested Schema for `arp_entries`

Read-Only:

* `address` (String) - ip address
* `age` (String) - Age in ISO-8601 partial-time format (time-hour : time-minute : time-second)
* `hw_address` (String) - hardware address
* `iface` (String) - interface name
* `state` (String)
<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type

