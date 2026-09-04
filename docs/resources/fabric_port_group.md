---
page_title: "gigavuecore_fabric_port_group Resource - gigavuecore"
subcategory: ""
description: |-
  Create a Fabric Port Group
---

# gigavuecore_fabric_port_group Resource

Create a Fabric Port Group

## Example Usage

```terraform
resource "gigavuecore_fabric_port_group" "example" {
  alias        = "example"
  comment      = "example"
  port_list    = ["example"]
  port_weights = [0]
  smart_lb     = true
  tags = [{
    tag_key    = "example"
    tag_values = ["example"]
  }]
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the Fabric Port Group
* `comment` (String, optional) - Description of the Fabric Port Group
* `port_list` (List of String, required) - Port List
* `port_weights` (List of Number, required) - Port Weights (1 to 100)
* `smart_lb` (Boolean, optional) - Smart Load Balancing, it is always enabled otherwise Fabric Port Group cannot be used in the Fabric Map
* `tags` (Attributes List, optional) (see [below for nested schema](#nestedatt--tags))

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `health_state` (String, computed) - Health State
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `traffic_health_state` (String, computed) - Traffic Health State
* `traffic_health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--traffic_health_state_reasons))

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag

<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

<a id="nestedatt--traffic_health_state_reasons"></a>
### Nested Schema for `traffic_health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_fabric_port_group.example {alias}
```
