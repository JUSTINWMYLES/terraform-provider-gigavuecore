---
page_title: "gigavuecore_network Resource - gigavuecore"
subcategory: ""
description: |-
  Find Inline Network by alias
---

# gigavuecore_network Resource

Find Inline Network by alias

## Example Usage

```terraform
resource "gigavuecore_network" "example" {
  alias = null
  comment = null
  forwarding_state = null
  health_state = null
  health_state_reasons = []
  heartbeat = {}
  lfp = null
  physical_bypass = null
  port_a = null
  port_b = null
  redundancy_control_state = null
  redundancy_profile = null
  traffic_path = null
  type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Inline Network alias. Unique within a cluster
* `comment` (String, optional)
* `forwarding_state` (String, optional)
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), optional)
* `heartbeat` (Object({enabled}), optional) - Embedded Heartbeat configuration for an Inline Network.
  * `enabled` (Bool, required)
* `lfp` (Bool, optional) - Link Failure Propagation. Controls whether the inline network link failure on one side of the inline network gets propagated to the other side
* `physical_bypass` (Bool, optional) - only applicable for 'protected' inline networks
* `port_a` (String, required) - portId of side A inline network port
* `port_b` (String, required) - portId of side B inline network port
* `redundancy_control_state` (String, optional) - Redundancy Control State. For 'protected' inline networks
* `redundancy_profile` (String, optional) - Alias of one of the pre-defined Redundancy Profiles. only applicable for 'protected' inline networks
* `traffic_path` (String, optional)
* `type` (String, optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `comment` (String, computed)
* `forwarding_state` (String, computed)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `heartbeat` (Object({enabled}), computed) - Embedded Heartbeat configuration for an Inline Network.
  * `enabled` (Bool, required)
* `lfp` (Bool, computed) - Link Failure Propagation. Controls whether the inline network link failure on one side of the inline network gets propagated to the other side
* `physical_bypass` (Bool, computed) - only applicable for 'protected' inline networks
* `redundancy_control_state` (String, computed) - Redundancy Control State. For 'protected' inline networks
* `redundancy_profile` (String, computed) - Alias of one of the pre-defined Redundancy Profiles. only applicable for 'protected' inline networks
* `traffic_path` (String, computed)
* `type` (String, computed)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_network.example {alias}
```
