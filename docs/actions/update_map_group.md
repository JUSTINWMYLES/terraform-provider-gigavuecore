---
page_title: "gigavuecore_update_map_group Action - gigavuecore"
subcategory: ""
description: |-
  Update map group for a given policy and source rule
---

# gigavuecore_update_map_group Action

Update map group for a given policy and source rule

## Example Usage

```terraform
action "gigavuecore_update_map_group" "example" {
  config {
    alias = "example"
    map_group_config = {
      cluster_name = "example"
      comment      = "example"
      components = [{
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
      map_group_alias = "example"
    }
    policy_alias       = "example"
    source_alias       = "example"
    source_rules_alias = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `map_group_config` (Attributes, required) (see [below for nested schema](#nestedatt--map_group_config))
* `policy_alias` (String, required) - Policy alias
* `source_alias` (String, required)
* `source_rules_alias` (String, required) - Source rules alias

<a id="nestedatt--map_group_config"></a>
### Nested Schema for `map_group_config`

Required:

* `cluster_name` (String)
* `components` (Attributes List) (see [below for nested schema](#nestedatt--map_group_config--components))
* `map_group_alias` (String)

Optional:

* `comment` (String)

<a id="nestedatt--map_group_config--components"></a>
### Nested Schema for `map_group_config.components`

Required:

* `alias` (String) - components alias is mandatory, it cannot be empty

Optional:

* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--map_group_config--components--health_state_reasons))
* `source_alias` (String)
* `traffic_health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--map_group_config--components--traffic_health_state_reasons))
* `type` (String)

<a id="nestedatt--map_group_config--components--health_state_reasons"></a>
### Nested Schema for `map_group_config.components.health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

<a id="nestedatt--map_group_config--components--traffic_health_state_reasons"></a>
### Nested Schema for `map_group_config.components.traffic_health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

