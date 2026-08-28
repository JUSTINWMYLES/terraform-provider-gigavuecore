---
page_title: "gigavuecore_get_map_chain Data Source - gigavuecore"
subcategory: ""
description: |-
  Retrieve map chain for a given policy and source rule
---

# gigavuecore_get_map_chain Data Source

Retrieve map chain for a given policy and source rule

## Example Usage

```terraform
data "gigavuecore_get_map_chain" "example" {
  policy_alias       = null
  source_rules_alias = null
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
* `priority_configs` (Attributes List, computed) (see [below for nested schema](#nestedatt--priority_configs))
* `source_alias` (String, computed)

<a id="nestedatt--priority_configs"></a>
### Nested Schema for `priority_configs`

Read-Only:

* `cluster_name` (String)
* `collector_map_alias` (String)
* `map_chain_id` (String)
* `ordered_components` (Attributes List) (see [below for nested schema](#nestedatt--priority_configs--ordered_components))
* `src_ports_as_id` (String)
<a id="nestedatt--priority_configs--ordered_components"></a>
### Nested Schema for `priority_configs.ordered_components`

Read-Only:

* `alias` (String) - components alias is mandatory, it cannot be empty
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--priority_configs--ordered_components--health_state_reasons))
* `source_alias` (String)
* `traffic_health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--priority_configs--ordered_components--traffic_health_state_reasons))
* `type` (String)
<a id="nestedatt--priority_configs--ordered_components--health_state_reasons"></a>
### Nested Schema for `priority_configs.ordered_components.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--priority_configs--ordered_components--traffic_health_state_reasons"></a>
### Nested Schema for `priority_configs.ordered_components.traffic_health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

