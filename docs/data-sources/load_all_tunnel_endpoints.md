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

* `alias` (String) - tunnel endpoint alias
* `ip_address` (String) - tunnel endpoint remote ip address. IPv4 or IPv6
* `te_id` (String) - tunnel endpoint alias, format: teN , where 1 <= N <= 128

