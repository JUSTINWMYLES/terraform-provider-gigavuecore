---
page_title: "gigavuecore_load_all_port_group Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all PortGroups
---

# gigavuecore_load_all_port_group Data Source

Load all PortGroups

## Example Usage

```terraform
data "gigavuecore_load_all_port_group" "example" {
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
* `port_groups` (Attributes List, computed) (see [below for nested schema](#nestedatt--port_groups))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--port_groups"></a>
### Nested Schema for `port_groups`

Read-Only:

* `alias` (String) - Port Group alias. Uniquely identifies a PortGroup within a cluster
* `cluster_id` (String) - id of the defining cluster
* `comment` (String)
* `gigastreams` (List of String) - Gigastream aliases
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--port_groups--health_state_reasons))
* `port_weights` (List of Number) - load balancing weights for the ports in the port list. If included, the list size must match the size of the 'ports' list
* `ports` (List of String)
* `smart_lb` (Boolean) - Enable or disable GigaSMART load balancing
* `tunnel_lb_endpoints` (Attributes List) - Tunnel Endpoint id with weight (see [below for nested schema](#nestedatt--port_groups--tunnel_lb_endpoints))
<a id="nestedatt--port_groups--health_state_reasons"></a>
### Nested Schema for `port_groups.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--port_groups--tunnel_lb_endpoints"></a>
### Nested Schema for `port_groups.tunnel_lb_endpoints`

Read-Only:

* `tunnel_endpoint` (String) - tunnel endpoint id
* `weight` (Number) - tunnel endpoint weight

