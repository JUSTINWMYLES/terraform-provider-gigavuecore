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

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `arp_entries` (List(Object({address, age, hw_address, iface, state})), computed)
* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type

