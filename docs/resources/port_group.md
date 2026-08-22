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
  alias = null
  cluster_id = null
  comment = null
  gigastreams = []
  health_state = null
  health_state_reasons = []
  port_weights = []
  ports = []
  smart_lb = null
  tunnel_lb_endpoints = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Port Group alias. Uniquely identifies a PortGroup within a cluster
* `cluster_id` (String, optional) - id of the defining cluster
* `comment` (String, optional)
* `gigastreams` (List(String), optional) - Gigastream aliases
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), optional)
* `port_weights` (List(Number), optional) - load balancing weights for the ports in the port list. If included, the list size must match the size of the 'ports' list
* `ports` (List(String), optional)
* `smart_lb` (Bool, optional) - Enable or disable GigaSMART load balancing
* `tunnel_lb_endpoints` (List(Object({tunnel_endpoint, weight})), optional) - Tunnel Endpoint id with weight

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `cluster_id` (String, computed) - id of the defining cluster
* `comment` (String, computed)
* `gigastreams` (List(String), computed) - Gigastream aliases
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `port_weights` (List(Number), computed) - load balancing weights for the ports in the port list. If included, the list size must match the size of the 'ports' list
* `ports` (List(String), computed)
* `smart_lb` (Bool, computed) - Enable or disable GigaSMART load balancing
* `tunnel_lb_endpoints` (List(Object({tunnel_endpoint, weight})), computed) - Tunnel Endpoint id with weight

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_port_group.example {alias}
```
