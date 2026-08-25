---
page_title: "gigavuecore_load_all_vports Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all vPorts
---

# gigavuecore_load_all_vports Data Source

Load all vPorts

## Example Usage

```terraform
data "gigavuecore_load_all_vports" "example" {
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
* `vports` (Attributes List, computed) (see [below for nested schema](#nestedatt--vports))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--vports"></a>
### Nested Schema for `vports`

Read-Only:

* `alias` (String)
* `deferred_binding` (Boolean) - enable/disable deferred-binding
* `fail_over_action` (String)
* `gs_group` (String) - Alias of referenced managing GsGroup
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--vports--health_state_reasons))
* `inline_status` (String)
* `inner_traffic_path` (String) - Similar to inline-network traffic-path, applicable for inner map
* `metadata_monitoring` (Attributes) (see [below for nested schema](#nestedatt--vports--metadata_monitoring))
* `mode` (String)
* `outer_traffic_path` (String) - Similar to inline-tool/inline-tool-group flex-traffic path, applicable for outer map
* `sa_apf_profile` (String) - ASF session profile
<a id="nestedatt--vports--health_state_reasons"></a>
### Nested Schema for `vports.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--vports--metadata_monitoring"></a>
### Nested Schema for `vports.metadata_monitoring`

Read-Only:

* `action` (String) - metadata monitoring action
* `exporters` (List of String)

