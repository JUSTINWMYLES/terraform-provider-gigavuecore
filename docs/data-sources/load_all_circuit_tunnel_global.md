---
page_title: "gigavuecore_load_all_circuit_tunnel_global Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Circuit Tunnel global configurations
---

# gigavuecore_load_all_circuit_tunnel_global Data Source

Load all Circuit Tunnel global configurations

## Example Usage

```terraform
data "gigavuecore_load_all_circuit_tunnel_global" "example" {
  box_id     = null
  cluster_id = null
  page       = null
  sort       = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, optional) - device box id. valid range 1 - 64.
* `cluster_id` (String, required) - Target Cluster ID
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `box_id` (String) - device box id. valid range 1 - 64.
* `l2_gre_entropy` (Number) - entropy size
* `l2_gre_protocol_type` (String) - l2gre protocol type
* `vxlan_l4_dst_port` (Number) - l4 destination port for tunnel terminating on the box

