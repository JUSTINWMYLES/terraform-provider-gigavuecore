---
page_title: "gigavuecore_load_all_circuit_tunnels Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Circuit Tunnels
---

# gigavuecore_load_all_circuit_tunnels Data Source

Load all Circuit Tunnels

## Example Usage

```terraform
data "gigavuecore_load_all_circuit_tunnels" "example" {
  cluster_id = null
  mode       = null
  page       = null
  sort       = null
  type       = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `mode` (String, optional) - Filter circuit tunnels by mode
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `type` (String, optional) - Filter circuit tunnels by type

### Attributes

In addition to all arguments above, the following attributes are exported:

* `circuit_tunnels` (Attributes List, computed) (see [below for nested schema](#nestedatt--circuit_tunnels))
* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))

<a id="nestedatt--circuit_tunnels"></a>
### Nested Schema for `circuit_tunnels`

Read-Only:

* `alias` (String) - tunnel name
* `attach` (List of String) - ipInterface if type is vxlan. ports or gigastreams if type is circuit and is valid only with decap mode.
* `circuit_ids` (List of Number) - circuit ids, valid and required if type is circuit
* `cluster_id` (String) - id of the cluster in which this circuit tunnel is created
* `comment` (String)
* `dip_address` (String)
* `l4_src_port` (Number)
* `mode` (String) - tunnel mode
* `type` (String)
<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type

