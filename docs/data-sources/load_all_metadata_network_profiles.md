---
page_title: "gigavuecore_load_all_metadata_network_profiles Data Source - gigavuecore"
subcategory: ""
description: |-
  Load All Metadata Network Profiles
---

# gigavuecore_load_all_metadata_network_profiles Data Source

Load All Metadata Network Profiles

## Example Usage

```terraform
data "gigavuecore_load_all_metadata_network_profiles" "example" {
  page = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `network_profiles` (Attributes List, computed) (see [below for nested schema](#nestedatt--network_profiles))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--network_profiles"></a>
### Nested Schema for `network_profiles`

Read-Only:

* `alias` (String) - network profile alias
* `description` (String)
* `ipv4_subnets` (List of String) - ipv4 subnets, to identify client
* `ipv6_masks` (List of String) - ipv6 subnets, to identify client

