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
  alias = null
  cct_alias = null
  ifm_alias = null
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

* `circuit_tunnel` (Object({alias, attach, circuit_ids, cluster_id, comment, dip_address, l4_src_port, mode, type}), computed)
  * `alias` (String, computed) - tunnel name
  * `attach` (List(String), computed) - ipInterface if type is vxlan. ports or gigastreams if type is circuit and is valid only with decap mode.
  * `circuit_ids` (List(Number), computed) - circuit ids, valid and required if type is circuit
  * `cluster_id` (String, computed) - id of the cluster in which this circuit tunnel is created
  * `comment` (String, computed)
  * `dip_address` (String, computed)
  * `l4_src_port` (Number, computed)
  * `mode` (String, computed) - tunnel mode
  * `type` (String, computed)
* `cluster_id` (String, computed) - id of the cluster in which the circuit tunnel endpoint is created.
* `config_status` (String, computed) - Configuration status of this circuit tunnel endpoint.
* `error_message` (String, computed) - In case of configuration failure, this message provides details about the possible cause of the failure.
* `fabric_map_aliases` (List(String), computed) - Aliases of the fabric maps that this circuit tunnel endpoint supports.

