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
  alias          = null
  destination_id = null
  is_active      = null
  page           = null
  sort           = null
  source_id      = null
  state          = null
  tunnel_id      = null
  tunnel_type    = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional)
* `destination_id` (String, optional)
* `is_active` (String, optional)
* `page` (String, optional)
* `sort` (String, optional)
* `source_id` (String, optional)
* `state` (String, optional)
* `tunnel_id` (String, optional)
* `tunnel_type` (String, optional)

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
* `tunnel_id` (Dynamic) - Tunnel logical group Id used for encapsulation and decapsulation
* `tunnel_type` (String) - Giga Tunnel Type
<a id="nestedatt--items--health_state_reasons"></a>
### Nested Schema for `items.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

