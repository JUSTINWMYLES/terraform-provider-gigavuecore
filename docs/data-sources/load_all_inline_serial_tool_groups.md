---
page_title: "gigavuecore_load_all_inline_serial_tool_groups Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Inline Serial Tool Groups
---

# gigavuecore_load_all_inline_serial_tool_groups Data Source

Load all Inline Serial Tool Groups

## Example Usage

```terraform
data "gigavuecore_load_all_inline_serial_tool_groups" "example" {
  cluster_id = null
  page       = null
  sort       = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `inline_serial_tool_groups` (Attributes List, computed) (see [below for nested schema](#nestedatt--inline_serial_tool_groups))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--inline_serial_tool_groups"></a>
### Nested Schema for `inline_serial_tool_groups`

Read-Only:

* `alias` (String) - Inline Tool Group alias. Unique within a cluster
* `comment` (String)
* `enabled` (Boolean) - setting to false is equivalent to forcing the inline serial tool failure (useful for taking the inline serial tool out of commission for maintenance or other purposes)
* `failover_action` (String)
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--inline_serial_tool_groups--health_state_reasons))
* `inline_tools` (List of String) - a list of aliases of the inline tools or inline tool groups that participate in the series
* `per_direction_order` (String) - direction of traffic flow in reference to the order that the inline-tools are configured
<a id="nestedatt--inline_serial_tool_groups--health_state_reasons"></a>
### Nested Schema for `inline_serial_tool_groups.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

