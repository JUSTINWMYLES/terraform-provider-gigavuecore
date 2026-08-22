---
page_title: "gigavuecore_fabric_port_group Resource - gigavuecore"
subcategory: ""
description: |-
  Get a Fabric Port Group by alias
---

# gigavuecore_fabric_port_group Resource

Get a Fabric Port Group by alias

## Example Usage

```terraform
resource "gigavuecore_fabric_port_group" "example" {
  alias = null
  comment = null
  port_list = []
  port_weights = []
  smart_lb = null
  tags = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the Fabric Port Group
* `comment` (String, optional) - Description of the Fabric Port Group
* `port_list` (List(String), required) - Port List
* `port_weights` (List(Number), required) - Port Weights (1 to 100)
* `smart_lb` (Bool, optional) - Smart Load Balancing, it is always enabled otherwise Fabric Port Group cannot be used in the Fabric Map
* `tags` (List(Object({tag_key, tag_values})), optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `comment` (String, computed) - Description of the Fabric Port Group
* `health_state` (String, computed) - Health State
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `smart_lb` (Bool, computed) - Smart Load Balancing, it is always enabled otherwise Fabric Port Group cannot be used in the Fabric Map
* `tags` (List(Object({tag_key, tag_values})), computed)
* `traffic_health_state` (String, computed) - Traffic Health State
* `traffic_health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_fabric_port_group.example {alias}
```
