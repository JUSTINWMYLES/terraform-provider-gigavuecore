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

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `giga_fabric_port_groups` (Attributes List, computed) (see [below for nested schema](#nestedatt--giga_fabric_port_groups))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--giga_fabric_port_groups"></a>
### Nested Schema for `giga_fabric_port_groups`

Read-Only:

* `alias` (String) - Alias of the Fabric Port Group
* `comment` (String) - Description of the Fabric Port Group
* `health_state` (String) - Health State
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--giga_fabric_port_groups--health_state_reasons))
* `port_list` (List of String) - Port List
* `port_weights` (List of Number) - Port Weights (1 to 100)
* `smart_lb` (Boolean) - Smart Load Balancing, it is always enabled otherwise Fabric Port Group cannot be used in the Fabric Map
* `tags` (Attributes List) (see [below for nested schema](#nestedatt--giga_fabric_port_groups--tags))
* `traffic_health_state` (String) - Traffic Health State
* `traffic_health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--giga_fabric_port_groups--traffic_health_state_reasons))
<a id="nestedatt--giga_fabric_port_groups--health_state_reasons"></a>
### Nested Schema for `giga_fabric_port_groups.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--giga_fabric_port_groups--tags"></a>
### Nested Schema for `giga_fabric_port_groups.tags`

Read-Only:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--giga_fabric_port_groups--traffic_health_state_reasons"></a>
### Nested Schema for `giga_fabric_port_groups.traffic_health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

