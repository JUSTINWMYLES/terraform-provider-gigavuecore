---
page_title: "gigavuecore_load_tunnel_endpoint_entries Data Source - gigavuecore"
subcategory: ""
description: |-
  Load known Tunnel Endpoints
---

# gigavuecore_load_tunnel_endpoint_entries Data Source

Load known Tunnel Endpoints

## Example Usage

```terraform
data "gigavuecore_load_tunnel_endpoint_entries" "example" {
  cluster_id  = null
  page        = null
  sort        = null
  tunnel_type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Cluster ID to filter by
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:100) to prevent reading entire DB by accident
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `tunnel_type` (String, optional) - Tunnel type to filter by

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `tunnel_endpoints` (Attributes List, computed) (see [below for nested schema](#nestedatt--tunnel_endpoints))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--tunnel_endpoints"></a>
### Nested Schema for `tunnel_endpoints`

Read-Only:

* `tunnel_endpoint` (Attributes) - Tunnel Endpoint configuration (see [below for nested schema](#nestedatt--tunnel_endpoints--tunnel_endpoint))
* `tunneled_port` (Attributes) - Tunnel Endpoint's servicing Tunneled Port information (see [below for nested schema](#nestedatt--tunnel_endpoints--tunneled_port))
<a id="nestedatt--tunnel_endpoints--tunnel_endpoint"></a>
### Nested Schema for `tunnel_endpoints.tunnel_endpoint`

Read-Only:

* `address` (String) - Tunnel Endpoint IP Address on the TunneledPort
* `gre_key` (Number) - only applicable for 'l2gre' type
* `port` (Number) - Tunnel Endpoint Port on the TunneledPort. Only applicable and required for 'gmip' type
* `src_port` (Number) - Tunnel Source Port (the port on the remote/client side). Only applicable and required for 'gmip' type
* `type` (String)
<a id="nestedatt--tunnel_endpoints--tunneled_port"></a>
### Nested Schema for `tunnel_endpoints.tunneled_port`

Read-Only:

* `cluster_id` (String) - ID of a cluster the tunneled port is a member of
* `device_box_id` (String) - BoxId of the device the tunneled port is a member of
* `device_id` (String) - ID of the device the tunneled port is a member of
* `device_model` (String) - model of the device hosting the tunneled port
* `port_id` (String) - Tunneled port ID on a Gigamon Chassis

