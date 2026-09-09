---
page_title: "gigavuecore_network Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new Inline Network
---

# gigavuecore_network Resource

Create a new Inline Network

## Example Usage

```terraform
resource "gigavuecore_network" "example" {
  alias            = "example"
  cluster_id       = "example"
  comment          = "example"
  forwarding_state = "physicalBypass"
  health_state     = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  heartbeat = {
    enabled = true
  }
  lfp                      = true
  physical_bypass          = true
  port_a                   = "example"
  port_b                   = "example"
  redundancy_control_state = "neutral"
  redundancy_profile       = "example"
  traffic_path             = "drop"
  type                     = "protected"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Inline Network alias. Unique within a cluster
* `cluster_id` (String, required) - Target Cluster ID
* `comment` (String, optional)
* `forwarding_state` (String, optional)
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `heartbeat` (Attributes, optional) - Embedded Heartbeat configuration for an Inline Network. (see [below for nested schema](#nestedatt--heartbeat))
* `lfp` (Boolean, optional) - Link Failure Propagation. Controls whether the inline network link failure on one side of the inline network gets propagated to the other side
* `physical_bypass` (Boolean, optional) - only applicable for 'protected' inline networks
* `port_a` (String, required) - portId of side A inline network port
* `port_b` (String, required) - portId of side B inline network port
* `redundancy_control_state` (String, optional) - Redundancy Control State. For 'protected' inline networks
* `redundancy_profile` (String, optional) - Alias of one of the pre-defined Redundancy Profiles. only applicable for 'protected' inline networks
* `traffic_path` (String, optional)
* `type` (String, optional)

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

<a id="nestedatt--heartbeat"></a>
### Nested Schema for `heartbeat`

Required:

* `enabled` (Boolean)
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
terraform import gigavuecore_network.example {alias}/{cluster_id}
```
