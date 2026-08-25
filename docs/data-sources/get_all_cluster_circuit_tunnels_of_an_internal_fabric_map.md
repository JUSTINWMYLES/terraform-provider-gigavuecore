---
page_title: "gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all cluster circuit tunnel endpoints of an internal fabric map
---

# gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map Data Source

Get all cluster circuit tunnel endpoints of an internal fabric map

## Example Usage

```terraform
data "gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map" "example" {
  alias     = null
  ifm_alias = null
  mode      = null
  type      = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the fabric map
* `ifm_alias` (String, required) - alias of the internal fabric map
* `mode` (String, optional) - tunnel mode: 'encap' or 'decap'.
* `type` (String, optional) - tunnel type: 'circuit' or 'vxlan'.

### Attributes

In addition to all arguments above, the following attributes are exported:

* `cluster_maps` (Attributes List, computed) (see [below for nested schema](#nestedatt--cluster_maps))
* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))

<a id="nestedatt--cluster_maps"></a>
### Nested Schema for `cluster_maps`

Read-Only:

* `circuit_tunnel` (Attributes) (see [below for nested schema](#nestedatt--cluster_maps--circuit_tunnel))
* `cluster_id` (String) - id of the cluster in which the circuit tunnel endpoint is created.
* `config_status` (String) - Configuration status of this circuit tunnel endpoint.
* `error_message` (String) - In case of configuration failure, this message provides details about the possible cause of the failure.
* `fabric_map_aliases` (List of String) - Aliases of the fabric maps that this circuit tunnel endpoint supports.
<a id="nestedatt--cluster_maps--circuit_tunnel"></a>
### Nested Schema for `cluster_maps.circuit_tunnel`

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

