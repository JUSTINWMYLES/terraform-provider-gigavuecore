---
page_title: "gigavuecore_update_map_chain_priority Action - gigavuecore"
subcategory: ""
description: |-
  Update map chain priority for a given policy and source rule
---

# gigavuecore_update_map_chain_priority Action

Update map chain priority for a given policy and source rule

## Example Usage

```terraform
action "gigavuecore_update_map_chain_priority" "example" {
  config {
    alias        = "example"
    policy_alias = "example"
    priority_configs = [{
      cluster_name = "example"
      component = {
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
      }
      map_chain_id  = "example"
      priority_type = "example"
      ref_component = {
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
      }
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

* `alias` (String, required) - Alias is mandatory, cannot be null or blank
* `policy_alias` (String, required) - Policy alias
* `priority_configs` (Attributes List, optional) - List of clustered priority specifications (see [below for nested schema](#nestedatt--priority_configs))
* `source_alias` (String, required) - SourceAlias is mandatory, cannot be null or blank
* `source_rules_alias` (String, required) - Source rules alias

<a id="nestedatt--priority_configs"></a>
### Nested Schema for `priority_configs`

Required:

* `cluster_name` (String) - Cluster name, cannot be empty
* `component` (Attributes) - Schema for PriorityComponent representing a component with health and traffic state. (see [below for nested schema](#nestedatt--priority_configs--component))
* `priority_type` (String) - Priority type, mandatory
* `src_ports_as_id` (String) - Source ports as ID, cannot be empty

Optional:

* `map_chain_id` (String) - Map chain ID
* `ref_component` (Attributes) - Schema for PriorityComponent representing a component with health and traffic state. (see [below for nested schema](#nestedatt--priority_configs--ref_component))

<a id="nestedatt--priority_configs--component"></a>
### Nested Schema for `priority_configs.component`

Required:

* `alias` (String) - components alias is mandatory, it cannot be empty

Optional:

* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--priority_configs--component--health_state_reasons))
* `source_alias` (String)
* `traffic_health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--priority_configs--component--traffic_health_state_reasons))
* `type` (String)

<a id="nestedatt--priority_configs--component--health_state_reasons"></a>
### Nested Schema for `priority_configs.component.health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

<a id="nestedatt--priority_configs--component--traffic_health_state_reasons"></a>
### Nested Schema for `priority_configs.component.traffic_health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

<a id="nestedatt--priority_configs--ref_component"></a>
### Nested Schema for `priority_configs.ref_component`

Required:

* `alias` (String) - components alias is mandatory, it cannot be empty

Optional:

* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--priority_configs--ref_component--health_state_reasons))
* `source_alias` (String)
* `traffic_health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--priority_configs--ref_component--traffic_health_state_reasons))
* `type` (String)

<a id="nestedatt--priority_configs--ref_component--health_state_reasons"></a>
### Nested Schema for `priority_configs.ref_component.health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

<a id="nestedatt--priority_configs--ref_component--traffic_health_state_reasons"></a>
### Nested Schema for `priority_configs.ref_component.traffic_health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

