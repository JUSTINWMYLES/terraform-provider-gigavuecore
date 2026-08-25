---
page_title: "gigavuecore_load_all_v3_users Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all SNMPv3 Users
---

# gigavuecore_load_all_v3_users Data Source

Load all SNMPv3 Users

## Example Usage

```terraform
data "gigavuecore_load_all_v3_users" "example" {
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
* `v3_users` (Attributes List, computed) (see [below for nested schema](#nestedatt--v3_users))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--v3_users"></a>
### Nested Schema for `v3_users`

Read-Only:

* `auth_key` (String)
* `auth_protocol` (String)
* `enabled` (Boolean)
* `priv_key` (String)
* `priv_protocol` (String)
* `read_only` (Boolean) - This property is introduced for GUI's purpose to restrict the edit/delete on the SNMPv3 user used by FM. If this property doesn't exist, consider it false. And this property is ignored if it exists in the create/update spec.
* `username` (String)

