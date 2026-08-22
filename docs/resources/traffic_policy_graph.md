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
  alias = null
  description = null
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
* `end_points` (Object({applications, destinations, maps, sources, tunnels}), computed) - End Points Container.
  * `applications` (Set(Object({comment, end_point_entry_references, end_point_exit_references, end_point_id, end_point_node, end_point_properties, end_point_resource, end_point_template, end_point_type, end_point_universe})), computed)
  * `destinations` (Set(Object({comment, end_point_entry_references, end_point_exit_references, end_point_id, end_point_node, end_point_properties, end_point_resource, end_point_template, end_point_type, end_point_universe})), computed)
  * `maps` (Set(Object({comment, end_point_entry_references, end_point_exit_references, end_point_id, end_point_node, end_point_properties, end_point_resource, end_point_template, end_point_type, end_point_universe})), computed)
  * `sources` (Set(Object({comment, end_point_entry_references, end_point_exit_references, end_point_id, end_point_node, end_point_properties, end_point_resource, end_point_template, end_point_type, end_point_universe})), computed)
  * `tunnels` (Set(Object({comment, end_point_entry_references, end_point_exit_references, end_point_id, end_point_node, end_point_properties, end_point_resource, end_point_template, end_point_type, end_point_universe})), computed)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `tags` (List(Object({tag_key, tag_values})), computed)
* `traffic_flows` (Object({traffic_flows}), computed) - Traffic Flows Container
  * `traffic_flows` (Set(Object({flow_destination, flow_destination_reference, flow_id, flow_source, flow_source_reference})), computed)
* `traffic_health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_traffic_policy_graph.example {alias}
```
