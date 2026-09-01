---
page_title: "gigavuecore_load_all_stack_links Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Stack Links
---

# gigavuecore_load_all_stack_links Data Source

Load all Stack Links

## Example Usage

```terraform
data "gigavuecore_load_all_stack_links" "example" {
  cluster_id = "example"
  page       = "example"
  sort       = "example"
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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - Port Pair alias. Uniquely identifies a StackLink within a cluster
* `cluster_id` (String) - id of the defining cluster
* `comment` (String)
* `endpoint1` (String) - for 'gigastream' type, this is the gigastream alias of node1. for 'port' type, this is the port id on node1
* `endpoint2` (String) - for 'gigastream' type, this is the gigastream alias of node2. for 'port' type, this is the port id on node2
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--health_state_reasons))
* `port_list1` (List of String) - only applicable for type 'gigastream'. port-list corresponding to endpoint1
* `port_list2` (List of String) - only applicable for type 'gigastream'. port-list corresponding to endpoint2
* `type` (String) - specifies whether this is a gigastream-based or port-based stack link

<a id="nestedatt--items--health_state_reasons"></a>
### Nested Schema for `items.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

