---
page_title: "gigavuecore_circuit_tunnel Resource - gigavuecore"
subcategory: ""
description: |-
  Load Circuit Tunnel by alias
---

# gigavuecore_circuit_tunnel Resource

Load Circuit Tunnel by alias

## Example Usage

```terraform
resource "gigavuecore_circuit_tunnel" "example" {
  alias       = "example"
  attach      = [ "example" ]
  circuit_ids = [ 0 ]
  cluster_id  = "example"
  comment     = "example"
  dip_address = "example"
  l4_src_port = 0
  mode        = "example"
  type        = "example"
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

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `attach` (List of String, computed) - ipInterface if type is vxlan. ports or gigastreams if type is circuit and is valid only with decap mode.
* `circuit_ids` (List of Number, computed) - circuit ids, valid and required if type is circuit
* `comment` (String, computed)
* `dip_address` (String, computed)
* `l4_src_port` (Number, computed)
* `mode` (String, computed) - tunnel mode
* `type` (String, computed)


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_circuit_tunnel.example {alias}
```
