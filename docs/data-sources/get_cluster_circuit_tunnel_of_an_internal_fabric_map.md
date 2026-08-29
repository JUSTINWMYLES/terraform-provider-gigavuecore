---
page_title: "gigavuecore_get_cluster_circuit_tunnel_of_an_internal_fabric_map Data Source - gigavuecore"
subcategory: ""
description: |-
  Get a cluster circuit tunnel endpoint of an internal fabric map
---

# gigavuecore_get_cluster_circuit_tunnel_of_an_internal_fabric_map Data Source

Get a cluster circuit tunnel endpoint of an internal fabric map

## Example Usage

```terraform
data "gigavuecore_get_cluster_circuit_tunnel_of_an_internal_fabric_map" "example" {
  alias     = "example"
  cct_alias = "example"
  ifm_alias = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the fabric map
* `cct_alias` (String, required) - alias of the cluster circuit tunnel endpoint
* `ifm_alias` (String, required) - alias of the internal fabric map

### Attributes

In addition to all arguments above, the following attributes are exported:

* `circuit_tunnel` (Attributes, computed) (see [below for nested schema](#nestedatt--circuit_tunnel))
* `cluster_id` (String, computed) - id of the cluster in which the circuit tunnel endpoint is created.
* `config_status` (String, computed) - Configuration status of this circuit tunnel endpoint.
* `error_message` (String, computed) - In case of configuration failure, this message provides details about the possible cause of the failure.
* `fabric_map_aliases` (List of String, computed) - Aliases of the fabric maps that this circuit tunnel endpoint supports.

<a id="nestedatt--circuit_tunnel"></a>
### Nested Schema for `circuit_tunnel`

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

