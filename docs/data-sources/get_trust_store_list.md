---
page_title: "gigavuecore_get_trust_store_list Data Source - gigavuecore"
subcategory: ""
description: |-
  Get information about the trust-store
---

# gigavuecore_get_trust_store_list Data Source

Get information about the trust-store

## Example Usage

```terraform
data "gigavuecore_get_trust_store_list" "example" {
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
* `trust_store_list` (Attributes List, computed) (see [below for nested schema](#nestedatt--trust_store_list))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--trust_store_list"></a>
### Nested Schema for `trust_store_list`

Read-Only:

* `cluster_id` (String) - id of the defining cluster
* `expire_date` (String)
* `fingerprint` (String)
* `issuer_common_name` (String)
* `issuer_name` (String)
* `subject_common_name` (String)
* `subject_name` (String)

