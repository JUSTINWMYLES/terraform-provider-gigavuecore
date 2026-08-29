---
page_title: "gigavuecore_get_map_group Data Source - gigavuecore"
subcategory: ""
description: |-
  Retrieve map group for a given policy and source rule
---

# gigavuecore_get_map_group Data Source

Retrieve map group for a given policy and source rule

## Example Usage

```terraform
data "gigavuecore_get_map_group" "example" {
  policy_alias       = "example"
  source_rules_alias = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `policy_alias` (String, required) - Policy alias
* `source_rules_alias` (String, required) - Source rules alias

### Attributes

In addition to all arguments above, the following attributes are exported:

* `alias` (String, computed)
* `map_group_config` (Attributes, computed) (see [below for nested schema](#nestedatt--map_group_config))
* `source_alias` (String, computed)

<a id="nestedatt--map_group_config"></a>
### Nested Schema for `map_group_config`

Read-Only:

* `cluster_name` (String)
* `comment` (String)
* `components` (Attributes List) (see [below for nested schema](#nestedatt--map_group_config--components))
* `map_group_alias` (String)
<a id="nestedatt--map_group_config--components"></a>
### Nested Schema for `map_group_config.components`

Read-Only:

* `alias` (String) - components alias is mandatory, it cannot be empty
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--map_group_config--components--health_state_reasons))
* `source_alias` (String)
* `traffic_health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--map_group_config--components--traffic_health_state_reasons))
* `type` (String)
<a id="nestedatt--map_group_config--components--health_state_reasons"></a>
### Nested Schema for `map_group_config.components.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--map_group_config--components--traffic_health_state_reasons"></a>
### Nested Schema for `map_group_config.components.traffic_health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

