---
page_title: "gigavuecore_tool_port_mirror Resource - gigavuecore"
subcategory: ""
description: |-
  Find ToolPortMirror by alias
---

# gigavuecore_tool_port_mirror Resource

Find ToolPortMirror by alias

## Example Usage

```terraform
resource "gigavuecore_tool_port_mirror" "example" {
  alias                = null
  cluster_id           = null
  comment              = null
  dst_ports            = []
  health_state         = null
  health_state_reasons = []
  src_ports            = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - unique tool port mirror alias
* `cluster_id` (String, optional) - id of the defining cluster
* `comment` (String, optional)
* `dst_ports` (List of String, required) - List of the 'to' tool ports
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `src_ports` (List of String, required) - list of the 'from' tool ports

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `cluster_id` (String, computed) - id of the defining cluster
* `comment` (String, computed)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))

<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_tool_port_mirror.example {alias}
```
