---
page_title: "gigavuecore_load_all_fabric_port_groups Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all the Fabric Port Groups
---

# gigavuecore_load_all_fabric_port_groups Data Source

Get all the Fabric Port Groups

## Example Usage

```terraform
data "gigavuecore_load_all_fabric_port_groups" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - Alias of the Fabric Port Group
* `comment` (String) - Description of the Fabric Port Group
* `health_state` (String) - Health State
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--health_state_reasons))
* `port_list` (List of String) - Port List
* `port_weights` (List of Number) - Port Weights (1 to 100)
* `smart_lb` (Boolean) - Smart Load Balancing, it is always enabled otherwise Fabric Port Group cannot be used in the Fabric Map
* `tags` (Attributes List) (see [below for nested schema](#nestedatt--items--tags))
* `traffic_health_state` (String) - Traffic Health State
* `traffic_health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--traffic_health_state_reasons))
<a id="nestedatt--items--health_state_reasons"></a>
### Nested Schema for `items.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--items--tags"></a>
### Nested Schema for `items.tags`

Read-Only:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--items--traffic_health_state_reasons"></a>
### Nested Schema for `items.traffic_health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

