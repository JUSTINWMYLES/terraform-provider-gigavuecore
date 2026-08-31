---
page_title: "gigavuecore_port_group Resource - gigavuecore"
subcategory: ""
description: |-
  Find PortGroup by alias
---

# gigavuecore_port_group Resource

Find PortGroup by alias

## Example Usage

```terraform
resource "gigavuecore_port_group" "example" {
  alias        = "example"
  cluster_id   = "example"
  comment      = "example"
  gigastreams  = ["example"]
  health_state = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  port_weights = [1]
  ports        = ["example"]
  smart_lb     = true
  tunnel_lb_endpoints = [{
    tunnel_endpoint = "example"
    weight          = 1
  }]
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Port Group alias. Uniquely identifies a PortGroup within a cluster
* `cluster_id` (String, required) - id of the defining cluster
* `comment` (String, optional)
* `gigastreams` (List of String, optional) - Gigastream aliases
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `port_weights` (List of Number, optional) - load balancing weights for the ports in the port list. If included, the list size must match the size of the 'ports' list
* `ports` (List of String, optional)
* `smart_lb` (Boolean, optional) - Enable or disable GigaSMART load balancing
* `tunnel_lb_endpoints` (Attributes List, optional) - Tunnel Endpoint id with weight (see [below for nested schema](#nestedatt--tunnel_lb_endpoints))

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

<a id="nestedatt--tunnel_lb_endpoints"></a>
### Nested Schema for `tunnel_lb_endpoints`

Required:

* `tunnel_endpoint` (String) - tunnel endpoint id

Optional:

* `weight` (Number) - tunnel endpoint weight
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
terraform import gigavuecore_port_group.example {alias}/{cluster_id}
```
