---
page_title: "gigavuecore_update_map_chain Action - gigavuecore"
subcategory: ""
description: |-
  Update map chain for a given policy and source rule
---

# gigavuecore_update_map_chain Action

Update map chain for a given policy and source rule

## Example Usage

```terraform
action "gigavuecore_update_map_chain" "example" {
  config {
    alias        = "example"
    policy_alias = "example"
    priority_configs = [{
      cluster_name        = "example"
      collector_map_alias = "example"
      map_chain_id        = "example"
      ordered_components = [{
        alias        = "example"
        health_state = "example"
        health_state_reasons = [{
          message                               = "example"
          severity                              = "example"
          traffic_health_state_computation_type = "example"
        }]
        source_alias         = "example"
        traffic_health_state = "example"
        traffic_health_state_reasons = [{
          message                               = "example"
          severity                              = "example"
          traffic_health_state_computation_type = "example"
        }]
        type = "example"
      }]
      src_ports_as_id = "example"
    }]
    source_alias       = "example"
    source_rules_alias = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `policy_alias` (String, required) - Policy alias
* `priority_configs` (Attributes List, optional) (see [below for nested schema](#nestedatt--priority_configs))
* `source_alias` (String, required)
* `source_rules_alias` (String, required) - Source rules alias

<a id="nestedatt--priority_configs"></a>
### Nested Schema for `priority_configs`

Required:

* `cluster_name` (String)
* `ordered_components` (Attributes List) (see [below for nested schema](#nestedatt--priority_configs--ordered_components))
* `src_ports_as_id` (String)
Optional:

* `collector_map_alias` (String)
* `map_chain_id` (String)
<a id="nestedatt--priority_configs--ordered_components"></a>
### Nested Schema for `priority_configs.ordered_components`

Required:

* `alias` (String) - components alias is mandatory, it cannot be empty
Optional:

* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--priority_configs--ordered_components--health_state_reasons))
* `source_alias` (String)
* `traffic_health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--priority_configs--ordered_components--traffic_health_state_reasons))
* `type` (String)
<a id="nestedatt--priority_configs--ordered_components--health_state_reasons"></a>
### Nested Schema for `priority_configs.ordered_components.health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--priority_configs--ordered_components--traffic_health_state_reasons"></a>
### Nested Schema for `priority_configs.ordered_components.traffic_health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

