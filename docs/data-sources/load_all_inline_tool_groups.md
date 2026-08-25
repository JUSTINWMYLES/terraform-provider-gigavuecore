---
page_title: "gigavuecore_load_all_inline_tool_groups Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Inline Tool Groups
---

# gigavuecore_load_all_inline_tool_groups Data Source

Load all Inline Tool Groups

## Example Usage

```terraform
data "gigavuecore_load_all_inline_tool_groups" "example" {
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
* `inline_tool_groups` (Attributes List, computed) (see [below for nested schema](#nestedatt--inline_tool_groups))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--inline_tool_groups"></a>
### Nested Schema for `inline_tool_groups`

Read-Only:

* `alias` (String) - Inline Tool Group alias. Unique within a cluster
* `comment` (String)
* `current_state` (Attributes) - Inline Tool Group Smart Load Balancing definition (see [below for nested schema](#nestedatt--inline_tool_groups--current_state))
* `enabled` (Boolean) - setting to false is equivalent to forcing the inline tool group failure (useful for taking the inline tool group out of commission for maintenance or other purposes)
* `failover_action` (String)
* `failover_mode` (String) - the way of handling a failure of an individual member of the inline tool list when no spare inline tool is configured or if the spare inline tool is failed
* `flex_status` (String)
* `flex_traffic_path` (String)
* `hash` (String)
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--inline_tool_groups--health_state_reasons))
* `inline_tools` (List of String) - If no spare inline tool configured, list of aliases of inline tools participating in hash-based traffic distribution. If the spare inline tool is configured, list of aliases of primary inline tools to which traffic is forwarded as long as all of them are healthy. The number of inline tools in the list must be between 1 and 64 if the spare inline tool is configured or between 2 and 64 otherwise
* `min_group_size` (Number) - the minimum number of inline tools (the inline tools in the list plus spare if configured) that must be up so that the entire inline-tool-group is considered up
* `oper_status` (String)
* `release_spare_if_possible` (Boolean) - when set to true, if the spare inline tool became active it remains active regardless of the health state of the originally failed primary inline tool
* `spare_inline_tool` (String) - alias of an inline tool to which traffic is forwarded when the first failure occurs in the set of primary inline tools
<a id="nestedatt--inline_tool_groups--current_state"></a>
### Nested Schema for `inline_tool_groups.current_state`

Read-Only:

* `inline_tools` (List of String) - Current list of inlineTools
* `spare_tool` (String) - Current spare tool
* `spare_tool_status` (String) - Current spare tool status
* `switched_inline_tool` (String) - Current switched inlineTool
<a id="nestedatt--inline_tool_groups--health_state_reasons"></a>
### Nested Schema for `inline_tool_groups.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

