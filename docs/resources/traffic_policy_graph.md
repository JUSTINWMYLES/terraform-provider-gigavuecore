---
page_title: "gigavuecore_traffic_policy_graph Resource - gigavuecore"
subcategory: ""
description: |-
  Find TrafficPolicyGraphs by alias
---

# gigavuecore_traffic_policy_graph Resource

Find TrafficPolicyGraphs by alias

## Example Usage

```terraform
resource "gigavuecore_traffic_policy_graph" "example" {
  alias       = "example"
  description = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional) - Traffic Policy Graph alias. Uniquely identifies a TrafficPolicyGraph
* `description` (String, optional) - Traffic Policy Graph description

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `alias` (String, computed) - Traffic Policy Graph alias. Uniquely identifies a TrafficPolicyGraph
* `comment` (String, computed)
* `description` (String, computed) - Traffic Policy Graph description
* `end_points` (Attributes, computed) - End Points Container. (see [below for nested schema](#nestedatt--end_points))
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `tags` (Attributes List, computed) (see [below for nested schema](#nestedatt--tags))
* `traffic_flows` (Attributes, computed) - Traffic Flows Container (see [below for nested schema](#nestedatt--traffic_flows))
* `traffic_health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--traffic_health_state_reasons))

<a id="nestedatt--end_points"></a>
### Nested Schema for `end_points`

Read-Only:

* `applications` (Attributes Set) (see [below for nested schema](#nestedatt--end_points--applications))
* `destinations` (Attributes Set) (see [below for nested schema](#nestedatt--end_points--destinations))
* `maps` (Attributes Set) (see [below for nested schema](#nestedatt--end_points--maps))
* `sources` (Attributes Set) (see [below for nested schema](#nestedatt--end_points--sources))
* `tunnels` (Attributes Set) (see [below for nested schema](#nestedatt--end_points--tunnels))
<a id="nestedatt--end_points--applications"></a>
### Nested Schema for `end_points.applications`

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
<a id="nestedatt--end_points--destinations"></a>
### Nested Schema for `end_points.destinations`

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
<a id="nestedatt--end_points--maps"></a>
### Nested Schema for `end_points.maps`

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
<a id="nestedatt--end_points--sources"></a>
### Nested Schema for `end_points.sources`

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
<a id="nestedatt--end_points--tunnels"></a>
### Nested Schema for `end_points.tunnels`

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
<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--traffic_flows"></a>
### Nested Schema for `traffic_flows`

Read-Only:

* `traffic_flows` (Attributes Set) (see [below for nested schema](#nestedatt--traffic_flows--traffic_flows))
<a id="nestedatt--traffic_flows--traffic_flows"></a>
### Nested Schema for `traffic_flows.traffic_flows`

Read-Only:

* `flow_destination` (String) - Traffic Flow Destination Endpoint
* `flow_destination_reference` (String) - Reference point at Traffic Flow Source
* `flow_id` (String) - Unique id for Traffic Flow
* `flow_source` (String) - Traffic Flow Source Endpoint
* `flow_source_reference` (String) - Reference point at Traffic Flow Source
<a id="nestedatt--traffic_health_state_reasons"></a>
### Nested Schema for `traffic_health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_traffic_policy_graph.example {alias}
```
