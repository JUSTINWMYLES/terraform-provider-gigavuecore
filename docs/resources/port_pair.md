---
page_title: "gigavuecore_port_pair Resource - gigavuecore"
subcategory: ""
description: |-
  Find PortPair by alias
---

# gigavuecore_port_pair Resource

Find PortPair by alias

## Example Usage

```terraform
resource "gigavuecore_port_pair" "example" {
  alias = null
  cluster_id = null
  comment = null
  config_mismatch = null
  health_state = null
  health_state_reasons = []
  lfp_enabled = null
  port1 = null
  port2 = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Port Pair alias. Uniquely identifies a PortPair within a cluster
* `cluster_id` (String, optional) - id of the defining cluster
* `comment` (String, optional)
* `config_mismatch` (Bool, optional) - Indicates speed/duplex of port1 and port2 are different
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), optional)
* `lfp_enabled` (Bool, optional)
* `port1` (String, required)
* `port2` (String, required)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `cluster_id` (String, computed) - id of the defining cluster
* `comment` (String, computed)
* `config_mismatch` (Bool, computed) - Indicates speed/duplex of port1 and port2 are different
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `lfp_enabled` (Bool, computed)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_port_pair.example {alias}
```
