---
page_title: "gigavuecore_load_all_gigastream Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all gigastreams
---

# gigavuecore_load_all_gigastream Data Source

Load all gigastreams

## Example Usage

```terraform
data "gigavuecore_load_all_gigastream" "example" {
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
* `gigastreams` (Attributes List, computed) (see [below for nested schema](#nestedatt--gigastreams))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--gigastreams"></a>
### Nested Schema for `gigastreams`

Read-Only:

* `alias` (String) - gigastream alias. Uniquely identifies a gigastream within a cluster
* `cluster_id` (String) - id of the defining cluster
* `comment` (String)
* `drop_weight` (Number) - relative weight for dropping the traffic
* `failover_status` (String) - Failover Status
* `hash_size` (Number) - Hash bucket size
* `hash_tool_port` (Attributes List) - Hash bucket id to tool port mapping (see [below for nested schema](#nestedatt--gigastreams--hash_tool_port))
* `hash_type` (String)
* `hash_weights` (List of Number) - hashWeights for 'ports'.If included, the list size must match the size of the 'ports' list
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--gigastreams--health_state_reasons))
* `ports` (List of String) - list of the ports to combine into a gigastream
* `threshold_level` (String) - Threshold level
* `variance_threshold` (String) - Variance threshold percentage
<a id="nestedatt--gigastreams--hash_tool_port"></a>
### Nested Schema for `gigastreams.hash_tool_port`

Read-Only:

* `hash_bucket_ids` (List of Number) - hash bucket id or range
* `tool_ports` (List of String) - tool port(s) mapped to hashBucketIds
<a id="nestedatt--gigastreams--health_state_reasons"></a>
### Nested Schema for `gigastreams.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

