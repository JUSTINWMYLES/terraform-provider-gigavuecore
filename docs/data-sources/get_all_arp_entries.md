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
  cluster_id = null
  page       = null
  sort       = null
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

* `arp_responses` (Attributes List, computed) (see [below for nested schema](#nestedatt--arp_responses))
* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))

<a id="nestedatt--arp_responses"></a>
### Nested Schema for `arp_responses`

Read-Only:

* `arp` (Attributes List) (see [below for nested schema](#nestedatt--arp_responses--arp))
* `cluster_id` (String) - id of the defining cluster
* `eport` (String) - GigaSMART engine port
<a id="nestedatt--arp_responses--arp"></a>
### Nested Schema for `arp_responses.arp`

Read-Only:

* `hw_address` (String)
* `ip_address` (String)
<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type

