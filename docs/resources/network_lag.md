---
page_title: "gigavuecore_network_lag Resource - gigavuecore"
subcategory: ""
description: |-
  Find Inline Network LAG by alias
---

# gigavuecore_network_lag Resource

Find Inline Network LAG by alias

## Example Usage

```terraform
resource "gigavuecore_network_lag" "example" {
  alias = null
  cdp = null
  comment = null
  forwarding_state = null
  health_state = null
  health_state_reasons = []
  inline_networks = []
  lacp = null
  lfp = null
  physical_bypass = null
  redundancy_control_state = null
  redundancy_profile = null
  traffic_path = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Inline Network LAG alias. Unique within a cluster
* `cdp` (Bool, optional) - Enable or Disable Cisco Discovery Protocol
* `comment` (String, optional)
* `forwarding_state` (String, optional)
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), optional)
* `inline_networks` (Set(String), required) - list of inline-network-lag aliases
* `lacp` (Bool, optional) - Enable or Disable Link Aggregation Control Protocol
* `lfp` (Bool, optional) - Link Failure Propagation. Controls whether the inline network link failure on one side of the inline network gets propagated to the other side. Changes will be pushed down to all inline-network members
* `physical_bypass` (Bool, optional) - changes will be pushed down to all inline-network members
* `redundancy_control_state` (String, optional) - For Redundancy Control State. From 'protected' inline networks
* `redundancy_profile` (String, optional) - Alias of one of the pre-defined Redundancy Profiles. only applicable for 'protected' inline networks
* `traffic_path` (String, optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `cdp` (Bool, computed) - Enable or Disable Cisco Discovery Protocol
* `comment` (String, computed)
* `forwarding_state` (String, computed)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `lacp` (Bool, computed) - Enable or Disable Link Aggregation Control Protocol
* `lfp` (Bool, computed) - Link Failure Propagation. Controls whether the inline network link failure on one side of the inline network gets propagated to the other side. Changes will be pushed down to all inline-network members
* `physical_bypass` (Bool, computed) - changes will be pushed down to all inline-network members
* `redundancy_control_state` (String, computed) - For Redundancy Control State. From 'protected' inline networks
* `redundancy_profile` (String, computed) - Alias of one of the pre-defined Redundancy Profiles. only applicable for 'protected' inline networks
* `traffic_path` (String, computed)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_network_lag.example {alias}
```
