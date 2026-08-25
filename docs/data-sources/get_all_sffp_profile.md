---
page_title: "gigavuecore_get_all_sffp_profile Data Source - gigavuecore"
subcategory: ""
description: |-
  new in H5.8
---

# gigavuecore_get_all_sffp_profile Data Source

new in H5.8

## Example Usage

```terraform
data "gigavuecore_get_all_sffp_profile" "example" {
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
* `sffp_profiles` (Attributes List, computed) (see [below for nested schema](#nestedatt--sffp_profiles))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--sffp_profiles"></a>
### Nested Schema for `sffp_profiles`

Read-Only:

* `alias` (String) - Alias of the SFFP Profile
* `profiles` (Attributes List) (see [below for nested schema](#nestedatt--sffp_profiles--profiles))
<a id="nestedatt--sffp_profiles--profiles"></a>
### Nested Schema for `sffp_profiles.profiles`

Read-Only:

* `ip_interface` (String)
* `node_type` (String)
* `port_list` (List of Number)
* `sx_interface` (Attributes) (see [below for nested schema](#nestedatt--sffp_profiles--profiles--sx_interface))
<a id="nestedatt--sffp_profiles--profiles--sx_interface"></a>
### Nested Schema for `sffp_profiles.profiles.sx_interface`

Read-Only:

* `ip_addresses` (List of String)

