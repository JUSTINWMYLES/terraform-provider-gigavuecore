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
  alias                = null
  comment              = null
  gigastreams          = []
  health_state         = null
  health_state_reasons = []
  port_weights         = []
  ports                = []
  smart_lb             = null
  tunnel_lb_endpoints  = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Port Group alias. Uniquely identifies a PortGroup within a cluster
* `comment` (String, optional)
* `gigastreams` (List of String, optional) - Gigastream aliases
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `port_weights` (List of Number, optional) - load balancing weights for the ports in the port list. If included, the list size must match the size of the 'ports' list
* `ports` (List of String, optional)
* `smart_lb` (Boolean, optional) - Enable or disable GigaSMART load balancing
* `tunnel_lb_endpoints` (Attributes List, optional) - Tunnel Endpoint id with weight (see [below for nested schema](#nestedatt--tunnel_lb_endpoints))

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `cluster_id` (String, computed) - id of the defining cluster
* `comment` (String, computed)
* `gigastreams` (List of String, computed) - Gigastream aliases
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `port_weights` (List of Number, computed) - load balancing weights for the ports in the port list. If included, the list size must match the size of the 'ports' list
* `ports` (List of String, computed)
* `smart_lb` (Boolean, computed) - Enable or disable GigaSMART load balancing
* `tunnel_lb_endpoints` (Attributes List, computed) - Tunnel Endpoint id with weight (see [below for nested schema](#nestedatt--tunnel_lb_endpoints))

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

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_port_group.example {alias}
```
