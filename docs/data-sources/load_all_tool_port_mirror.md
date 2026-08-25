---
page_title: "gigavuecore_load_all_tool_port_mirror Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all port mirrors
---

# gigavuecore_load_all_tool_port_mirror Data Source

Load all port mirrors

## Example Usage

```terraform
data "gigavuecore_load_all_tool_port_mirror" "example" {
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
* `tool_port_mirrors` (Attributes List, computed) (see [below for nested schema](#nestedatt--tool_port_mirrors))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--tool_port_mirrors"></a>
### Nested Schema for `tool_port_mirrors`

Read-Only:

* `alias` (String) - unique tool port mirror alias
* `cluster_id` (String) - id of the defining cluster
* `comment` (String)
* `dst_ports` (List of String) - List of the 'to' tool ports
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--tool_port_mirrors--health_state_reasons))
* `src_ports` (List of String) - list of the 'from' tool ports
<a id="nestedatt--tool_port_mirrors--health_state_reasons"></a>
### Nested Schema for `tool_port_mirrors.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

