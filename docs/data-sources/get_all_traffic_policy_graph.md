---
page_title: "gigavuecore_get_all_traffic_policy_graph Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all TrafficPolicyGraphs
---

# gigavuecore_get_all_traffic_policy_graph Data Source

Load all TrafficPolicyGraphs

## Example Usage

```terraform
data "gigavuecore_get_all_traffic_policy_graph" "example" {
  page = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `traffic_policy_graphs` (Attributes List, computed) (see [below for nested schema](#nestedatt--traffic_policy_graphs))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--traffic_policy_graphs"></a>
### Nested Schema for `traffic_policy_graphs`

Read-Only:

* `alias` (String) - Traffic Policy Graph alias. Uniquely identifies a TrafficPolicyGraph
* `comment` (String)
* `end_points` (Attributes) - End Points Container. (see [below for nested schema](#nestedatt--traffic_policy_graphs--end_points))
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--traffic_policy_graphs--health_state_reasons))
* `tags` (Attributes List) (see [below for nested schema](#nestedatt--traffic_policy_graphs--tags))
* `traffic_flows` (Attributes) - Traffic Flows Container (see [below for nested schema](#nestedatt--traffic_policy_graphs--traffic_flows))
* `traffic_health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--traffic_policy_graphs--traffic_health_state_reasons))
<a id="nestedatt--traffic_policy_graphs--end_points"></a>
### Nested Schema for `traffic_policy_graphs.end_points`

Read-Only:

* `applications` (Attributes Set) (see [below for nested schema](#nestedatt--traffic_policy_graphs--end_points--applications))
* `destinations` (Attributes Set) (see [below for nested schema](#nestedatt--traffic_policy_graphs--end_points--destinations))
* `maps` (Attributes Set) (see [below for nested schema](#nestedatt--traffic_policy_graphs--end_points--maps))
* `sources` (Attributes Set) (see [below for nested schema](#nestedatt--traffic_policy_graphs--end_points--sources))
* `tunnels` (Attributes Set) (see [below for nested schema](#nestedatt--traffic_policy_graphs--end_points--tunnels))
<a id="nestedatt--traffic_policy_graphs--end_points--applications"></a>
### Nested Schema for `traffic_policy_graphs.end_points.applications`

Read-Only:

* `comment` (String)
* `end_point_entry_references` (List of String) - End Point entry references
* `end_point_exit_references` (List of String) - End Point exit references
* `end_point_id` (String)
* `end_point_node` (String) - End Point Node
* `end_point_properties` (Dynamic) - End Point Properties
* `end_point_resource` (String) - End Point Resource by Resource Manager
* `end_point_template` (String) - End Point Template
* `end_point_type` (String) - End Point type
* `end_point_universe` (String) - End Point Universe
<a id="nestedatt--traffic_policy_graphs--end_points--destinations"></a>
### Nested Schema for `traffic_policy_graphs.end_points.destinations`

Read-Only:

* `comment` (String)
* `end_point_entry_references` (List of String) - End Point entry references
* `end_point_exit_references` (List of String) - End Point exit references
* `end_point_id` (String)
* `end_point_node` (String) - End Point Node
* `end_point_properties` (Dynamic) - End Point Properties
* `end_point_resource` (String) - End Point Resource by Resource Manager
* `end_point_template` (String) - End Point Template
* `end_point_type` (String) - End Point type
* `end_point_universe` (String) - End Point Universe
<a id="nestedatt--traffic_policy_graphs--end_points--maps"></a>
### Nested Schema for `traffic_policy_graphs.end_points.maps`

Read-Only:

* `comment` (String)
* `end_point_entry_references` (List of String) - End Point entry references
* `end_point_exit_references` (List of String) - End Point exit references
* `end_point_id` (String)
* `end_point_node` (String) - End Point Node
* `end_point_properties` (Dynamic) - End Point Properties
* `end_point_resource` (String) - End Point Resource by Resource Manager
* `end_point_template` (String) - End Point Template
* `end_point_type` (String) - End Point type
* `end_point_universe` (String) - End Point Universe
<a id="nestedatt--traffic_policy_graphs--end_points--sources"></a>
### Nested Schema for `traffic_policy_graphs.end_points.sources`

Read-Only:

* `comment` (String)
* `end_point_entry_references` (List of String) - End Point entry references
* `end_point_exit_references` (List of String) - End Point exit references
* `end_point_id` (String)
* `end_point_node` (String) - End Point Node
* `end_point_properties` (Dynamic) - End Point Properties
* `end_point_resource` (String) - End Point Resource by Resource Manager
* `end_point_template` (String) - End Point Template
* `end_point_type` (String) - End Point type
* `end_point_universe` (String) - End Point Universe
<a id="nestedatt--traffic_policy_graphs--end_points--tunnels"></a>
### Nested Schema for `traffic_policy_graphs.end_points.tunnels`

Read-Only:

* `comment` (String)
* `end_point_entry_references` (List of String) - End Point entry references
* `end_point_exit_references` (List of String) - End Point exit references
* `end_point_id` (String)
* `end_point_node` (String) - End Point Node
* `end_point_properties` (Dynamic) - End Point Properties
* `end_point_resource` (String) - End Point Resource by Resource Manager
* `end_point_template` (String) - End Point Template
* `end_point_type` (String) - End Point type
* `end_point_universe` (String) - End Point Universe
<a id="nestedatt--traffic_policy_graphs--health_state_reasons"></a>
### Nested Schema for `traffic_policy_graphs.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--traffic_policy_graphs--tags"></a>
### Nested Schema for `traffic_policy_graphs.tags`

Read-Only:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--traffic_policy_graphs--traffic_flows"></a>
### Nested Schema for `traffic_policy_graphs.traffic_flows`

Read-Only:

* `traffic_flows` (Attributes Set) (see [below for nested schema](#nestedatt--traffic_policy_graphs--traffic_flows--traffic_flows))
<a id="nestedatt--traffic_policy_graphs--traffic_flows--traffic_flows"></a>
### Nested Schema for `traffic_policy_graphs.traffic_flows.traffic_flows`

Read-Only:

* `flow_destination` (String) - Traffic Flow Destination Endpoint
* `flow_destination_reference` (String) - Reference point at Traffic Flow Source
* `flow_id` (String) - Unique id for Traffic Flow
* `flow_source` (String) - Traffic Flow Source Endpoint
* `flow_source_reference` (String) - Reference point at Traffic Flow Source
<a id="nestedatt--traffic_policy_graphs--traffic_health_state_reasons"></a>
### Nested Schema for `traffic_policy_graphs.traffic_health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

