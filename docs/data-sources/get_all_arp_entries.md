---
page_title: "gigavuecore_get_all_arp_entries Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all GigaSMART ports ARP entries
---

# gigavuecore_get_all_arp_entries Data Source

Load all GigaSMART ports ARP entries

## Example Usage

```terraform
data "gigavuecore_get_all_arp_entries" "example" {
  cluster_id = "example"
  page       = "example"
  sort       = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `arp` (Attributes List) (see [below for nested schema](#nestedatt--items--arp))
* `cluster_id` (String) - id of the defining cluster
* `eport` (String) - GigaSMART engine port

<a id="nestedatt--items--arp"></a>
### Nested Schema for `items.arp`

Read-Only:

* `hw_address` (String)
* `ip_address` (String)

