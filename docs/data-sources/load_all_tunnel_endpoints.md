---
page_title: "gigavuecore_load_all_tunnel_endpoints Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Tunnel Endpoints
---

# gigavuecore_load_all_tunnel_endpoints Data Source

Load all Tunnel Endpoints

## Example Usage

```terraform
data "gigavuecore_load_all_tunnel_endpoints" "example" {
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
* `tunnel_lb_endpoints` (Attributes List, computed) (see [below for nested schema](#nestedatt--tunnel_lb_endpoints))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--tunnel_lb_endpoints"></a>
### Nested Schema for `tunnel_lb_endpoints`

Read-Only:

* `alias` (String) - tunnel endpoint alias
* `ip_address` (String) - tunnel endpoint remote ip address. IPv4 or IPv6
* `te_id` (String) - tunnel endpoint alias, format: teN , where 1 <= N <= 128

