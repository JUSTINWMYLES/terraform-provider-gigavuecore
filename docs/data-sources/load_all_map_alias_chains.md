---
page_title: "gigavuecore_load_all_map_alias_chains Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all aliased map chains
---

# gigavuecore_load_all_map_alias_chains Data Source

Load all aliased map chains

## Example Usage

```terraform
data "gigavuecore_load_all_map_alias_chains" "example" {
  cluster_id = "example"
  map_alias  = "example"
  page       = "example"
  sort       = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `map_alias` (String, optional) - Filters the response to include only the mapChain that includes the map of the given alias. Note that the map can be either cluster map or fabric map.
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `cluster_id` (String) - id of the defining cluster
* `collector` (String) - alias of the chain collector map. must reference a map of matching type and  'collector' subtype
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--health_state_reasons))
* `map_chain_id` (String) - mapChain ID - replaces srcPortsAsId - should be used anywhere that requires srcPortsAsId
* `ordered_map_aliases` (Set of String) - aliases of the maps in the chain
* `src_ports` (List of String) - network ports for the maps chain
* `src_ports_as_id` (String) - (Deprecated - use mapChainId instead) mapChain ID
<a id="nestedatt--items--health_state_reasons"></a>
### Nested Schema for `items.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

