---
page_title: "gigavuecore_serial_tool_group Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new Inline Serial Tool Group
---

# gigavuecore_serial_tool_group Resource

Create a new Inline Serial Tool Group

## Example Usage

```terraform
resource "gigavuecore_serial_tool_group" "example" {
  alias           = "example"
  cluster_id      = "example"
  comment         = "example"
  enabled         = true
  failover_action = "toolBypass"
  health_state    = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  inline_tools        = ["example"]
  per_direction_order = "reverse"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Inline Tool Group alias. Unique within a cluster
* `cluster_id` (String, required) - Target Cluster ID
* `comment` (String, optional)
* `enabled` (Boolean, optional) - setting to false is equivalent to forcing the inline serial tool failure (useful for taking the inline serial tool out of commission for maintenance or other purposes)
* `failover_action` (String, required)
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `inline_tools` (List of String, required) - a list of aliases of the inline tools or inline tool groups that participate in the series
* `per_direction_order` (String, optional) - direction of traffic flow in reference to the order that the inline-tools are configured

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
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
terraform import gigavuecore_serial_tool_group.example {alias}/{cluster_id}
```
