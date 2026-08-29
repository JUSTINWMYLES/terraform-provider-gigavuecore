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
  cluster_id     = "example"
  order          = "example"
  page           = "example"
  server_address = "example"
  sort           = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `order` (String, optional) - The order in which the server is to be reached. 1 means server will be contacted first
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `server_address` (String, required) - ipv4 or ipv6 or hostname. Specifies address of the LDAP server where authentication requests will be sent
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `order` (String, computed) - The order in which the server is to be reached. 1 means server will be contacted first


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_ldap_server.example {server_address}
```
