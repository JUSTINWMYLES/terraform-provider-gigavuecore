---
page_title: "gigavuecore_get_overlap_components Data Source - gigavuecore"
subcategory: ""
description: |-
  Retrieve overlap components for a cluster and vport
---

# gigavuecore_get_overlap_components Data Source

Retrieve overlap components for a cluster and vport

## Example Usage

```terraform
data "gigavuecore_get_overlap_components" "example" {
  cluster_id  = "example"
  vport_alias = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `vport_alias` (String, required) - VPort alias

### Attributes

In addition to all arguments above, the following attributes are exported:

* `cluster_name` (String, computed)
* `comment` (String, computed)
* `components` (Attributes List, computed) (see [below for nested schema](#nestedatt--components))
* `map_group_alias` (String, computed)

<a id="nestedatt--components"></a>
### Nested Schema for `components`

Read-Only:

* `alias` (String) - components alias is mandatory, it cannot be empty
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--components--health_state_reasons))
* `source_alias` (String)
* `traffic_health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--components--traffic_health_state_reasons))
* `type` (String)
<a id="nestedatt--components--health_state_reasons"></a>
### Nested Schema for `components.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--components--traffic_health_state_reasons"></a>
### Nested Schema for `components.traffic_health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

