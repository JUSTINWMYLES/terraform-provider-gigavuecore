---
page_title: "gigavuecore_load_nrt_stats_fabric_map Data Source - gigavuecore"
subcategory: ""
description: |-
  Get NRT stats for the user-defined fabric map
---

# gigavuecore_load_nrt_stats_fabric_map Data Source

Get NRT stats for the user-defined fabric map

## Example Usage

```terraform
data "gigavuecore_load_nrt_stats_fabric_map" "example" {
  fabric_map_alias = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `fabric_map_alias` (String, required) - alias of the fabric map

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - unique nrt alias
* `cluster_name` (String) - cluster name
* `component_type` (String)
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state; 'grey' indicates undeployed state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--health_state_reasons))
* `nrt_alias` (String) - unique nrt alias
* `reason` (String)
* `stats_data` (Attributes) (see [below for nested schema](#nestedatt--items--stats_data))
* `traffic_health_state` (String) - Read-only. 'green' indicates traffic healthy state; 'red'  indicates traffic critical state; 'grey' indicates traffic not monitored state;
* `traffic_health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--traffic_health_state_reasons))
* `update_time` (String) - Last Updated time of the fabric map

<a id="nestedatt--items--health_state_reasons"></a>
### Nested Schema for `items.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

<a id="nestedatt--items--stats_data"></a>
### Nested Schema for `items.stats_data`

Read-Only:

* `alias` (String)
* `current_interval_start` (String)
* `current_offset_amount` (Number)
* `fabric_map_alias` (String) - unique fabric map alias
* `flex_dir` (String) - only applicable for flexInline maps
* `fstype` (String)
* `interval_duration` (String)
* `next_interval_start` (String)
* `octets` (String)
* `offset_shift_per_interval` (Number)
* `packets` (String)
* `rotational_time_rem` (String)
* `rules` (Attributes List) (see [below for nested schema](#nestedatt--items--stats_data--rules))
* `sub_type` (String) - 'byRule' is applicable to all map types; 'collector' is applicable to 'regular', 'inline', 'flexInline' and 'secondLevel' maps; 'passAll' is applicable to 'regular' and 'inline' maps; 'flowFilter', 'flowSample', 'flowWhitelist', 'flowSampleSip', 'flowWhitelistSip', 'flowSampleDiameter', 'flowWhitelistDiameter', 'flowSampleOverlap', 'flowWhitelistOverlap', 'flowSample5g','flowWhitelist5g','flowSample5gOverlap' and 'flowWhitelist5gOverlap' are applicable to 'secondLevel' maps
* `tags` (Map of List of String) - A label with key and list of values that can be tagged to a resource
* `timestamp` (String) - Counter collection date and time in ISO 8601 format
* `total_drop_octets` (Number)
* `total_drop_packets` (Number)
* `total_pass_octets` (Number)
* `total_pass_packets` (Number)
* `type` (String) - 'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'flexInline' is from inline network to tools; 'transitLevel' is from vport to vport

<a id="nestedatt--items--stats_data--rules"></a>
### Nested Schema for `items.stats_data.rules`

Read-Only:

* `accepted` (Number) - SecondLevel/flowSample only
* `accepted_bytes` (Number) - SecondLevel/flowSample only
* `accepted_pkts` (Number) - SecondLevel/flowSample only
* `entries` (Number) - SecondLevel/whitelist only
* `fs_interval_type` (Attributes) (see [below for nested schema](#nestedatt--items--stats_data--rules--fs_interval_type))
* `ip_can_bearer` (Number) - SecondLevel/flowFilter, flowWhitelist only
* `map_alias` (String) - Map Alias
* `matched` (Number) - SecondLevel/flowSample only
* `matched_bytes` (Number) - SecondLevel/whitelist only
* `matched_pkts` (Number) - SecondLevel/whitelist only
* `octets` (Number)
* `packets` (Number)
* `rejected` (Number) - SecondLevel/flowSample only
* `rejected_bytes` (Number) - SecondLevel/flowSample only
* `rejected_pkts` (Number) - SecondLevel/flowSample only
* `rule_id` (Number) - Map Rule Id
* `rule_type` (String)
* `rules` (Number) - SecondLevel/flowSample only
* `tags` (Map of List of String)

<a id="nestedatt--items--stats_data--rules--fs_interval_type"></a>
### Nested Schema for `items.stats_data.rules.fs_interval_type`

Read-Only:

* `current_interval` (Attributes) (see [below for nested schema](#nestedatt--items--stats_data--rules--fs_interval_type--current_interval))
* `future_intervals` (Attributes List) (see [below for nested schema](#nestedatt--items--stats_data--rules--fs_interval_type--future_intervals))
* `prior_intervals` (Attributes List) (see [below for nested schema](#nestedatt--items--stats_data--rules--fs_interval_type--prior_intervals))

<a id="nestedatt--items--stats_data--rules--fs_interval_type--current_interval"></a>
### Nested Schema for `items.stats_data.rules.fs_interval_type.current_interval`

Read-Only:

* `interval_start` (String)
* `sample_range` (String)

<a id="nestedatt--items--stats_data--rules--fs_interval_type--future_intervals"></a>
### Nested Schema for `items.stats_data.rules.fs_interval_type.future_intervals`

Read-Only:

* `interval_start` (String)
* `sample_range` (String)

<a id="nestedatt--items--stats_data--rules--fs_interval_type--prior_intervals"></a>
### Nested Schema for `items.stats_data.rules.fs_interval_type.prior_intervals`

Read-Only:

* `interval_start` (String)
* `sample_range` (String)

<a id="nestedatt--items--traffic_health_state_reasons"></a>
### Nested Schema for `items.traffic_health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

