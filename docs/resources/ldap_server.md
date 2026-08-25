---
page_title: "gigavuecore_ldap_server Resource - gigavuecore"
subcategory: ""
description: |-
  Load all LDAP Servers
---

# gigavuecore_ldap_server Resource

Load all LDAP Servers

## Example Usage

```terraform
resource "gigavuecore_ldap_server" "example" {
  order          = null
  server_address = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `order` (String, optional) - The order in which the server is to be reached. 1 means server will be contacted first
* `server_address` (String, required) - ipv4 or ipv6 or hostname. Specifies address of the LDAP server where authentication requests will be sent

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `id` (String, computed)
* `ldap_servers` (Attributes List, computed) (see [below for nested schema](#nestedatt--ldap_servers))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--ldap_servers"></a>
### Nested Schema for `ldap_servers`

Read-Only:

* `order` (String) - The order in which the server is to be reached. 1 means server will be contacted first
* `server_address` (String) - ipv4 or ipv6 or hostname. Specifies address of the LDAP server where authentication requests will be sent

