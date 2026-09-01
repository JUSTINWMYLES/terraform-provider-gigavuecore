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
  page = "example"
  sort = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - network profile alias
* `description` (String)
* `ipv4_subnets` (List of String) - ipv4 subnets, to identify client
* `ipv6_masks` (List of String) - ipv6 subnets, to identify client

