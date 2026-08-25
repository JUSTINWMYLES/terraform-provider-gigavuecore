---
page_title: "gigavuecore_load_all_tacacs_servers Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all TACACS+ Servers
---

# gigavuecore_load_all_tacacs_servers Data Source

Load all TACACS+ Servers

## Example Usage

```terraform
data "gigavuecore_load_all_tacacs_servers" "example" {
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
* `tacacs_servers` (Attributes List, computed) (see [below for nested schema](#nestedatt--tacacs_servers))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--tacacs_servers"></a>
### Nested Schema for `tacacs_servers`

Read-Only:

* `auth_type` (String) - Specify whether this TACACS+ server uses ASCII or PAP authentication
* `enabled` (Boolean)
* `port` (Number)
* `retries` (Number) - value of 0 disables retries. Defaults to the value defined in the TacacsServerDefaults
* `secret_key` (String) - if not included, defaults to the value defined in the TacacsServerDefaults
* `server_address` (String) - ipv4 or ipv6 or hostname
* `timeout` (Number) - in seconds. Defaults to the value defined in the TacacsServerDefaults

