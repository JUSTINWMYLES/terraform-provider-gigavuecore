---
page_title: "gigavuecore_port_pair Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new PortPair
---

# gigavuecore_port_pair Resource

Create a new PortPair

## Example Usage

```terraform
resource "gigavuecore_port_pair" "example" {
  alias           = "example"
  cluster_id      = "example"
  comment         = "example"
  config_mismatch = true
  health_state    = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  lfp_enabled = true
  port1       = "example"
  port2       = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Port Pair alias. Uniquely identifies a PortPair within a cluster
* `cluster_id` (String, required) - id of the defining cluster
* `comment` (String, optional)
* `config_mismatch` (Boolean, optional) - Indicates speed/duplex of port1 and port2 are different
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `lfp_enabled` (Boolean, optional)
* `port1` (String, required)
* `port2` (String, required)

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
terraform import gigavuecore_port_pair.example {alias}/{cluster_id}
```
