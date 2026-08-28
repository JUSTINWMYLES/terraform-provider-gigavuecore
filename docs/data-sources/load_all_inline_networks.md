---
page_title: "gigavuecore_load_all_inline_networks Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Inline Networks
---

# gigavuecore_load_all_inline_networks Data Source

Load all Inline Networks

## Example Usage

```terraform
data "gigavuecore_load_all_inline_networks" "example" {
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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - Inline Network alias. Unique within a cluster
* `comment` (String)
* `forwarding_state` (String)
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--health_state_reasons))
* `heartbeat` (Attributes) - Embedded Heartbeat configuration for an Inline Network. (see [below for nested schema](#nestedatt--items--heartbeat))
* `lfp` (Boolean) - Link Failure Propagation. Controls whether the inline network link failure on one side of the inline network gets propagated to the other side
* `physical_bypass` (Boolean) - only applicable for 'protected' inline networks
* `port_a` (String) - portId of side A inline network port
* `port_b` (String) - portId of side B inline network port
* `redundancy_control_state` (String) - Redundancy Control State. For 'protected' inline networks
* `redundancy_profile` (String) - Alias of one of the pre-defined Redundancy Profiles. only applicable for 'protected' inline networks
* `traffic_path` (String)
* `type` (String)
<a id="nestedatt--items--health_state_reasons"></a>
### Nested Schema for `items.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--items--heartbeat"></a>
### Nested Schema for `items.heartbeat`

Read-Only:

* `enabled` (Boolean)

