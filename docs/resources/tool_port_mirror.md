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
  alias        = "example"
  cluster_id   = "example"
  comment      = "example"
  dst_ports    = [ "example" ]
  health_state = "example"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "example"
    traffic_health_state_computation_type = "example"
  }]
  src_ports = [ "example" ]
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

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_tool_port_mirror.example {alias}/{cluster_id}
```
