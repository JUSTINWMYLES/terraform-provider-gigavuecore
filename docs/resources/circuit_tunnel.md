---
page_title: "gigavuecore_circuit_tunnel Resource - gigavuecore"
subcategory: ""
description: |-
  Create a Circuit Tunnel
---

# gigavuecore_circuit_tunnel Resource

Create a Circuit Tunnel

## Example Usage

```terraform
resource "gigavuecore_circuit_tunnel" "example" {
  alias       = "example"
  attach      = ["example"]
  circuit_ids = [2]
  cluster_id  = "example"
  comment     = "example"
  dip_address = "example"
  l4_src_port = 1
  mode        = "encap"
  type        = "vxlan"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - tunnel name
* `attach` (List of String, optional) - ipInterface if type is vxlan. ports or gigastreams if type is circuit and is valid only with decap mode.
* `circuit_ids` (List of Number, optional) - circuit ids, valid and required if type is circuit
* `cluster_id` (String, required) - id of the cluster in which this circuit tunnel is created
* `comment` (String, optional)
* `dip_address` (String, optional)
* `l4_src_port` (Number, optional)
* `mode` (String, optional) - tunnel mode
* `type` (String, optional)

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_circuit_tunnel.example {alias}/{cluster_id}
```
