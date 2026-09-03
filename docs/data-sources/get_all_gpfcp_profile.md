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

* `alias` (String) - Alias of the Gpfcp Profile
* `comment` (String) - Description of the Gpfcp Profile
* `g_profiles` (Attributes List) (see [below for nested schema](#nestedatt--items--g_profiles))

<a id="nestedatt--items--g_profiles"></a>
### Nested Schema for `items.g_profiles`

Read-Only:

* `comment` (String) - Description of the Gpfcp Profile Rule
* `g_interface` (Attributes) (see [below for nested schema](#nestedatt--items--g_profiles--g_interface))
* `ip_interface` (String)
* `node_type` (String)
* `port_list` (List of Number)

<a id="nestedatt--items--g_profiles--g_interface"></a>
### Nested Schema for `items.g_profiles.g_interface`

Read-Only:

* `ip_addresses` (List of String)

