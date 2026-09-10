---
page_title: "gigavuecore_tool_port_mirror Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new ToolPortMirror
---

# gigavuecore_tool_port_mirror Resource

Create a new ToolPortMirror

## Example Usage

```terraform
resource "gigavuecore_tool_port_mirror" "example" {
  alias        = "example"
  cluster_id   = "example"
  comment      = "example"
  dst_ports    = ["example"]
  health_state = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  src_ports = ["example"]
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - unique tool port mirror alias
* `cluster_id` (String, required) - id of the defining cluster
* `comment` (String, optional)
* `dst_ports` (List of String, required) - List of the 'to' tool ports
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `src_ports` (List of String, required) - list of the 'from' tool ports

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

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_tool_port_mirror.example {alias}/{cluster_id}
```
