---
page_title: "gigavuecore_get_tools Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Tools
---

# gigavuecore_get_tools Data Source

Get Tools

## Example Usage

```terraform
data "gigavuecore_get_tools" "example" {
  node_alias = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `node_alias` (String, optional) - manual node alias

### Attributes

In addition to all arguments above, the following attributes are exported:

* `alias` (String, computed)
* `tools_info` (Attributes List, computed) (see [below for nested schema](#nestedatt--tools_info))
* `type` (String, computed)

<a id="nestedatt--tools_info"></a>
### Nested Schema for `tools_info`

Read-Only:

* `comment` (String)
* `compression_ratio` (Number)
* `giga_streams` (Attributes List) (see [below for nested schema](#nestedatt--tools_info--giga_streams))
* `is_tool` (Boolean)
* `is_used_in_deployed_policy` (Boolean)
* `max_throughput` (Number)
* `model` (String)
* `node_alias` (String)
* `ports` (Attributes List) (see [below for nested schema](#nestedatt--tools_info--ports))
* `topo_node_id` (String)
* `total_storage` (Number)
* `type` (String)
* `vendor` (String)

<a id="nestedatt--tools_info--giga_streams"></a>
### Nested Schema for `tools_info.giga_streams`

Read-Only:

* `alias` (String) - gigastream alias. Uniquely identifies a gigastream within a cluster
* `cluster_id` (String) - id of the defining cluster
* `comment` (String)
* `drop_weight` (Number) - relative weight for dropping the traffic
* `failover_status` (String) - Failover Status
* `hash_size` (Number) - Hash bucket size
* `hash_tool_port` (Attributes List) - Hash bucket id to tool port mapping (see [below for nested schema](#nestedatt--tools_info--giga_streams--hash_tool_port))
* `hash_type` (String)
* `hash_weights` (List of Number) - hashWeights for 'ports'.If included, the list size must match the size of the 'ports' list
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--tools_info--giga_streams--health_state_reasons))
* `ports` (List of String) - list of the ports to combine into a gigastream
* `threshold_level` (String) - Threshold level
* `variance_threshold` (String) - Variance threshold percentage

<a id="nestedatt--tools_info--giga_streams--hash_tool_port"></a>
### Nested Schema for `tools_info.giga_streams.hash_tool_port`

Read-Only:

* `hash_bucket_ids` (List of Number) - hash bucket id or range
* `tool_ports` (List of String) - tool port(s) mapped to hashBucketIds

<a id="nestedatt--tools_info--giga_streams--health_state_reasons"></a>
### Nested Schema for `tools_info.giga_streams.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

<a id="nestedatt--tools_info--ports"></a>
### Nested Schema for `tools_info.ports`

Read-Only:

* `alias` (String)
* `cluster_id` (String)
* `device_ip` (String)
* `hostname` (String)
* `link_id` (String)
* `port_id` (String)
* `port_type` (String)

