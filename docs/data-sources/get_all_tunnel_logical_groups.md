---
page_title: "gigavuecore_get_all_tunnel_logical_groups Data Source - gigavuecore"
subcategory: ""
description: |-
  Get All Tunnel Logical Groups
---

# gigavuecore_get_all_tunnel_logical_groups Data Source

Get All Tunnel Logical Groups

## Example Usage

```terraform
data "gigavuecore_get_all_tunnel_logical_groups" "example" {
  alias          = "example"
  destination_id = "example"
  is_active      = "example"
  page           = "example"
  sort           = "example"
  source_id      = "example"
  state          = "example"
  tunnel_id      = "example"
  tunnel_type    = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional) - Tunnel Logical Group Alias
* `destination_id` (String, optional) - Destination Id(ClusterName/HostName) of Decapsulation Tunnel in Tunnel Logical Group
* `is_active` (String, optional) - Tunnel Logical Group Active/Inactive State. 'true' or 'false'
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:30) to prevent reading entire DB by accident
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is DESC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `source_id` (String, optional) - Source Id(ClusterName/HostName) of Encapsulation Tunnel in Tunnel Logical Group
* `state` (String, optional) - Tunnel Discovery State
* `tunnel_id` (String, optional) - Tunnel logical group Id used for encapsulation and decapsulation
* `tunnel_type` (String, optional) - Tunnel Type

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alert_policy_name` (String) - Alert policy name
* `alias` (String) - FM Tunnel Logical Group Alias. FM auto generated alias cannot be edited.
* `custom_alias` (String) - FM Tunnel Logical Group custom alias
* `decap_tunnel_infos` (Set of Dynamic)
* `encap_tunnel_infos` (Set of Dynamic)
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--health_state_reasons))
* `is_active` (Boolean) - User can activate or deactive Tunnel logical group. Active Tunnel logical groups of Complete state will be eligible for tunnel monitoring.
* `key` (String) - FM Tunnel Logical Group Key
* `state` (String) - Giga Tunnel Logical Group discovery state
* `tunnel_id` (Number) - Tunnel logical group Id used for encapsulation and decapsulation
* `tunnel_type` (String) - Giga Tunnel Type

<a id="nestedatt--items--health_state_reasons"></a>
### Nested Schema for `items.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

