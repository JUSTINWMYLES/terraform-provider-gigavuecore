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

* `alias` (String) - Alias of the SFFP Profile
* `profiles` (Attributes List) (see [below for nested schema](#nestedatt--items--profiles))
<a id="nestedatt--items--profiles"></a>
### Nested Schema for `items.profiles`

Read-Only:

* `ip_interface` (String)
* `node_type` (String)
* `port_list` (List of Number)
* `sx_interface` (Attributes) (see [below for nested schema](#nestedatt--items--profiles--sx_interface))
<a id="nestedatt--items--profiles--sx_interface"></a>
### Nested Schema for `items.profiles.sx_interface`

Read-Only:

* `ip_addresses` (List of String)

