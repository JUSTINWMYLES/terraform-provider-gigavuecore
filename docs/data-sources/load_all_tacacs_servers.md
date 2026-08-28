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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `auth_type` (String) - Specify whether this TACACS+ server uses ASCII or PAP authentication
* `enabled` (Boolean)
* `port` (Number)
* `retries` (Number) - value of 0 disables retries. Defaults to the value defined in the TacacsServerDefaults
* `secret_key` (String) - if not included, defaults to the value defined in the TacacsServerDefaults
* `server_address` (String) - ipv4 or ipv6 or hostname
* `timeout` (Number) - in seconds. Defaults to the value defined in the TacacsServerDefaults

