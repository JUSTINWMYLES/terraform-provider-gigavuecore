---
page_title: "gigavuecore_get_all_gpfcp_profile Data Source - gigavuecore"
subcategory: ""
description: |-
  new in H6.8
---

# gigavuecore_get_all_gpfcp_profile Data Source

new in H6.8

## Example Usage

```terraform
data "gigavuecore_get_all_gpfcp_profile" "example" {
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

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `gpfcp_profiles` (Attributes List, computed) (see [below for nested schema](#nestedatt--gpfcp_profiles))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--gpfcp_profiles"></a>
### Nested Schema for `gpfcp_profiles`

Read-Only:

* `alias` (String) - Alias of the Gpfcp Profile
* `comment` (String) - Description of the Gpfcp Profile
* `g_profiles` (Attributes List) (see [below for nested schema](#nestedatt--gpfcp_profiles--g_profiles))
<a id="nestedatt--gpfcp_profiles--g_profiles"></a>
### Nested Schema for `gpfcp_profiles.g_profiles`

Read-Only:

* `comment` (String) - Description of the Gpfcp Profile Rule
* `g_interface` (Attributes) (see [below for nested schema](#nestedatt--gpfcp_profiles--g_profiles--g_interface))
* `ip_interface` (String)
* `node_type` (String)
* `port_list` (List of Number)
<a id="nestedatt--gpfcp_profiles--g_profiles--g_interface"></a>
### Nested Schema for `gpfcp_profiles.g_profiles.g_interface`

Read-Only:

* `ip_addresses` (List of String)

