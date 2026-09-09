---
page_title: "gigavuecore_load_fabric_path Data Source - gigavuecore"
subcategory: ""
description: |-
  Load fabricPath details for the given srcClusterID and dstClusterID
---

# gigavuecore_load_fabric_path Data Source

Load fabricPath details for the given srcClusterID and dstClusterID

## Example Usage

```terraform
data "gigavuecore_load_fabric_path" "example" {
  details        = true
  dst_cluster_id = "example"
  src_cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `details` (Boolean, optional) - load comprehensive response on Fabric Path, including detailed information on topology link endpoints
* `dst_cluster_id` (String, required) - destination Cluster ID
* `src_cluster_id` (String, required) - source Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - Composing from srcClusterUUID and dstClusterUUID strings, i.e., srcClusterUUID-to-dstClusterUUID
* `config_state` (String) - Configuration state of the fabric path
* `dst_cluster_uuid` (String) - All dstVertices must be from same dstClusterUUID
* `dst_vertices` (Attributes List) (see [below for nested schema](#nestedatt--items--dst_vertices))
* `gfp_id` (Number) - Can be a hash number from alias string
* `operation_state` (String) - Operation state of the fabric path
* `path_type` (String) - Type for Fabric Path
* `referenced` (String) - Reference boolean of fabric path
* `src_cluster_uuid` (String) - All srcVertices must be from same srcClusterUUID
* `src_vertices` (Attributes List) (see [below for nested schema](#nestedatt--items--src_vertices))
* `topology_links` (Attributes List) (see [below for nested schema](#nestedatt--items--topology_links))
* `version` (Number) - Incremental version field to indicate the fabric path format

<a id="nestedatt--items--dst_vertices"></a>
### Nested Schema for `items.dst_vertices`

Read-Only:

* `alias` (String) - gigastream alias. Uniquely identifies a gigastream within a cluster
* `cluster_id` (String) - id of the defining cluster
* `comment` (String)
* `drop_weight` (Number) - relative weight for dropping the traffic
* `failover_status` (String) - Failover Status
* `hash_size` (Number) - Hash bucket size
* `hash_tool_port` (Attributes List) - Hash bucket id to tool port mapping (see [below for nested schema](#nestedatt--items--dst_vertices--hash_tool_port))
* `hash_type` (String)
* `hash_weights` (List of Number) - hashWeights for 'ports'.If included, the list size must match the size of the 'ports' list
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--dst_vertices--health_state_reasons))
* `ports` (List of String) - list of the ports to combine into a gigastream
* `threshold_level` (String) - Threshold level
* `variance_threshold` (String) - Variance threshold percentage

<a id="nestedatt--items--dst_vertices--hash_tool_port"></a>
### Nested Schema for `items.dst_vertices.hash_tool_port`

Read-Only:

* `hash_bucket_ids` (List of Number) - hash bucket id or range
* `tool_ports` (List of String) - tool port(s) mapped to hashBucketIds

<a id="nestedatt--items--dst_vertices--health_state_reasons"></a>
### Nested Schema for `items.dst_vertices.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

<a id="nestedatt--items--src_vertices"></a>
### Nested Schema for `items.src_vertices`

Read-Only:

* `alias` (String) - gigastream alias. Uniquely identifies a gigastream within a cluster
* `cluster_id` (String) - id of the defining cluster
* `comment` (String)
* `drop_weight` (Number) - relative weight for dropping the traffic
* `failover_status` (String) - Failover Status
* `hash_size` (Number) - Hash bucket size
* `hash_tool_port` (Attributes List) - Hash bucket id to tool port mapping (see [below for nested schema](#nestedatt--items--src_vertices--hash_tool_port))
* `hash_type` (String)
* `hash_weights` (List of Number) - hashWeights for 'ports'.If included, the list size must match the size of the 'ports' list
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--src_vertices--health_state_reasons))
* `ports` (List of String) - list of the ports to combine into a gigastream
* `threshold_level` (String) - Threshold level
* `variance_threshold` (String) - Variance threshold percentage

<a id="nestedatt--items--src_vertices--hash_tool_port"></a>
### Nested Schema for `items.src_vertices.hash_tool_port`

Read-Only:

* `hash_bucket_ids` (List of Number) - hash bucket id or range
* `tool_ports` (List of String) - tool port(s) mapped to hashBucketIds

<a id="nestedatt--items--src_vertices--health_state_reasons"></a>
### Nested Schema for `items.src_vertices.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

<a id="nestedatt--items--topology_links"></a>
### Nested Schema for `items.topology_links`

Read-Only:

* `aggregated_link` (Boolean)
* `comment` (String)
* `discovery_state` (String) - specifies the discovery state of the topo link
* `endpoint1` (Attributes) - node link endpoint (see [below for nested schema](#nestedatt--items--topology_links--endpoint1))
* `endpoint2` (Attributes) - node link endpoint (see [below for nested schema](#nestedatt--items--topology_links--endpoint2))
* `giga_link_type` (String) - specifies whether this is a gigastream-based or port-based link. Only applicable for topoNodeType of 'gigamon'
* `link_alias` (String) - alias of the link
* `link_state` (String) - specifies health state
* `manual_override` (Boolean) - Specifies whether manual augmentation exist. Only applicable for topoLinkType of 'cdp' or 'lldp'
* `topo_link_id` (String) - unique Id representing a topology graph link. Generated by FM server
* `topo_link_type` (String) - indicates how link became a pert of this topology graph

<a id="nestedatt--items--topology_links--endpoint1"></a>
### Nested Schema for `items.topology_links.endpoint1`

Read-Only:

* `comment` (String)
* `component_alias` (String) - For 'gigamon': this is either portId (e.g. '1/1/x1') or Gigastream alias (e.g. 'gs1'). For 'lldp'/'cdp': this is the 'portId' as reported by neighbor discovery. For 'manual': this is user provided or may be omitted
* `component_ports` (List of String) - For the component type gigastream, this specifies the array of port Ids that forms the logical group
* `component_type` (String) - Type of the endpoint component. This is either 'Port' or 'GigaStream'
* `node_alias` (String) - Node alias. References user-assigned topology node alias. Internal use only
* `port` (String) - For 'gigamon': this is the portId (e.g. '1/1/x1'). For 'lldp'/'cdp': this is the 'portId' as reported by neighbor discovery. For 'manual': this is user provided or may be omitted
* `topo_node_id` (String) - Topology graph node Id. References one of the Topology Nodes in the graph
* `topo_node_type` (String) - topology node type

<a id="nestedatt--items--topology_links--endpoint2"></a>
### Nested Schema for `items.topology_links.endpoint2`

Read-Only:

* `comment` (String)
* `component_alias` (String) - For 'gigamon': this is either portId (e.g. '1/1/x1') or Gigastream alias (e.g. 'gs1'). For 'lldp'/'cdp': this is the 'portId' as reported by neighbor discovery. For 'manual': this is user provided or may be omitted
* `component_ports` (List of String) - For the component type gigastream, this specifies the array of port Ids that forms the logical group
* `component_type` (String) - Type of the endpoint component. This is either 'Port' or 'GigaStream'
* `node_alias` (String) - Node alias. References user-assigned topology node alias. Internal use only
* `port` (String) - For 'gigamon': this is the portId (e.g. '1/1/x1'). For 'lldp'/'cdp': this is the 'portId' as reported by neighbor discovery. For 'manual': this is user provided or may be omitted
* `topo_node_id` (String) - Topology graph node Id. References one of the Topology Nodes in the graph
* `topo_node_type` (String) - topology node type

