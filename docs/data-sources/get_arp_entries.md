---
page_title: "gigavuecore_get_arp_entries Data Source - gigavuecore"
subcategory: ""
description: |-
  Load GigaSMART port ARP entries
---

# gigavuecore_get_arp_entries Data Source

Load GigaSMART port ARP entries

## Example Usage

```terraform
data "gigavuecore_get_arp_entries" "example" {
  cluster_id = null
  eport = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - id of the defining cluster
* `eport` (String, required) - GigaSMART engine port

### Attributes

In addition to all arguments above, the following attributes are exported:

* `arp` (List(Object({hw_address, ip_address})), computed)

