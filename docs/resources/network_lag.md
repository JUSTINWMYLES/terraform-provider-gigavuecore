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
  alias            = "example"
  cdp              = true
  cluster_id       = "example"
  comment          = "example"
  forwarding_state = "example"
  health_state     = "example"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "example"
    traffic_health_state_computation_type = "example"
  }]
  inline_networks          = [ "example" ]
  lacp                     = true
  lfp                      = true
  physical_bypass          = true
  redundancy_control_state = "example"
  redundancy_profile       = "example"
  traffic_path             = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Inline Network LAG alias. Unique within a cluster
* `cdp` (Boolean, optional) - Enable or Disable Cisco Discovery Protocol
* `cluster_id` (String, required) - Target Cluster ID
* `comment` (String, optional)
* `forwarding_state` (String, optional)
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `inline_networks` (Set of String, required) - list of inline-network-lag aliases
* `lacp` (Boolean, optional) - Enable or Disable Link Aggregation Control Protocol
* `lfp` (Boolean, optional) - Link Failure Propagation. Controls whether the inline network link failure on one side of the inline network gets propagated to the other side. Changes will be pushed down to all inline-network members
* `physical_bypass` (Boolean, optional) - changes will be pushed down to all inline-network members
* `redundancy_control_state` (String, optional) - For Redundancy Control State. From 'protected' inline networks
* `redundancy_profile` (String, optional) - Alias of one of the pre-defined Redundancy Profiles. only applicable for 'protected' inline networks
* `traffic_path` (String, optional)

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
terraform import gigavuecore_network_lag.example {alias}/{cluster_id}
```
