---
page_title: "gigavuecore_network_group Resource - gigavuecore"
subcategory: ""
description: |-
  Get FM Network group details by alias
---

# gigavuecore_network_group Resource

Get FM Network group details by alias

## Example Usage

```terraform
resource "gigavuecore_network_group" "example" {
  alias = null
  cluster_id = null
  health_state = null
  health_state_reasons = []
  members = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional) - Alias of the network group
* `cluster_id` (String, optional)
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), optional)
* `members` (List(Object({alias, cluster_name, type})), optional) - Array holding inline network group members

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `alias` (String, computed) - Alias of the network group
* `cluster_id` (String, computed)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `members` (List(Object({alias, cluster_name, type})), computed) - Array holding inline network group members

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_network_group.example {alias}
```
