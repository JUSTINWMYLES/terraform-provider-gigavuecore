---
page_title: "gigavuecore_get_ssl_client_trust_store_list Data Source - gigavuecore"
subcategory: ""
description: |-
  Get information about the SSL Client trust-stores
---

# gigavuecore_get_ssl_client_trust_store_list Data Source

Get information about the SSL Client trust-stores

## Example Usage

```terraform
data "gigavuecore_get_ssl_client_trust_store_list" "example" {
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

* `client_trust_stores` (Attributes List, computed) (see [below for nested schema](#nestedatt--client_trust_stores))
* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))

<a id="nestedatt--client_trust_stores"></a>
### Nested Schema for `client_trust_stores`

Read-Only:

* `alias` (String)
* `cluster_id` (String) - id of the defining cluster
* `ssl_trust_store_entries` (Attributes List) (see [below for nested schema](#nestedatt--client_trust_stores--ssl_trust_store_entries))
* `type` (String)
<a id="nestedatt--client_trust_stores--ssl_trust_store_entries"></a>
### Nested Schema for `client_trust_stores.ssl_trust_store_entries`

Read-Only:

* `expire_date` (String)
* `fingerprint` (String)
* `issuer_common_name` (String)
* `issuer_name` (String)
* `subject_common_name` (String)
* `subject_name` (String)
<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type

