---
page_title: "gigavuecore_get_policy_nrt_stats Data Source - gigavuecore"
subcategory: ""
description: |-
  Get near real-time statistics for a traffic flow
---

# gigavuecore_get_policy_nrt_stats Data Source

Get near real-time statistics for a traffic flow

## Example Usage

```terraform
data "gigavuecore_get_policy_nrt_stats" "example" {
  alias = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Traffic flow alias

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String)
* `cluster_name` (String)
* `component_type` (String)
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--health_state_reasons))
* `nrt_alias` (String)
* `reason` (String)
* `stats_data` (Dynamic)
* `traffic_health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--traffic_health_state_reasons))
* `update_time` (String)
<a id="nestedatt--items--health_state_reasons"></a>
### Nested Schema for `items.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--items--traffic_health_state_reasons"></a>
### Nested Schema for `items.traffic_health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

