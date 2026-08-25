---
page_title: "gigavuecore_serial_tool_group Resource - gigavuecore"
subcategory: ""
description: |-
  Find Inline Serial Tool Group by alias
---

# gigavuecore_serial_tool_group Resource

Find Inline Serial Tool Group by alias

## Example Usage

```terraform
resource "gigavuecore_serial_tool_group" "example" {
  alias                = null
  comment              = null
  enabled              = null
  failover_action      = null
  health_state         = null
  health_state_reasons = []
  inline_tools         = []
  per_direction_order  = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Inline Tool Group alias. Unique within a cluster
* `comment` (String, optional)
* `enabled` (Boolean, optional) - setting to false is equivalent to forcing the inline serial tool failure (useful for taking the inline serial tool out of commission for maintenance or other purposes)
* `failover_action` (String, required)
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `inline_tools` (List of String, required) - a list of aliases of the inline tools or inline tool groups that participate in the series
* `per_direction_order` (String, optional) - direction of traffic flow in reference to the order that the inline-tools are configured

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `comment` (String, computed)
* `enabled` (Boolean, computed) - setting to false is equivalent to forcing the inline serial tool failure (useful for taking the inline serial tool out of commission for maintenance or other purposes)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `per_direction_order` (String, computed) - direction of traffic flow in reference to the order that the inline-tools are configured

<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_serial_tool_group.example {alias}
```
