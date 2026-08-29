---
page_title: "gigavuecore_modify_tunnel_logical_group Action - gigavuecore"
subcategory: ""
description: |-
  Edit Tunnel Logical Group. API handles activate/deactivate tunnel logical group and modification of Tunnel logical group name.
---

# gigavuecore_modify_tunnel_logical_group Action

Edit Tunnel Logical Group. API handles activate/deactivate tunnel logical group and modification of Tunnel logical group name.

## Example Usage

```terraform
action "gigavuecore_modify_tunnel_logical_group" "example" {
  config {
    tunnel_logical_groups = [{
      alert_policy_name = "example"
      alias             = "example"
      custom_alias      = "example"
      decap_tunnel_infos = [{
        tunnel_operation_type = "example"
        tunnel_type           = "example"
      }]
      encap_tunnel_infos = [{
        tunnel_operation_type = "example"
        tunnel_type           = "example"
      }]
      health_state = "example"
      health_state_reasons = [{
        message                               = "example"
        severity                              = "example"
        traffic_health_state_computation_type = "example"
      }]
      is_active   = true
      key         = "example"
      state       = "example"
      tunnel_id   = "example"
      tunnel_type = "example"
    }]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `tunnel_logical_groups` (Attributes List, required) (see [below for nested schema](#nestedatt--tunnel_logical_groups))

<a id="nestedatt--tunnel_logical_groups"></a>
### Nested Schema for `tunnel_logical_groups`

Required:

* `decap_tunnel_infos` (Attributes Set) (see [below for nested schema](#nestedatt--tunnel_logical_groups--decap_tunnel_infos))
* `encap_tunnel_infos` (Attributes Set) (see [below for nested schema](#nestedatt--tunnel_logical_groups--encap_tunnel_infos))

Optional:

* `alert_policy_name` (String) - Alert policy name
* `alias` (String) - FM Tunnel Logical Group Alias. FM auto generated alias cannot be edited.
* `custom_alias` (String) - FM Tunnel Logical Group custom alias
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--tunnel_logical_groups--health_state_reasons))
* `is_active` (Boolean) - User can activate or deactive Tunnel logical group. Active Tunnel logical groups of Complete state will be eligible for tunnel monitoring.
* `key` (String) - FM Tunnel Logical Group Key
* `state` (String) - Giga Tunnel Logical Group discovery state
* `tunnel_id` (Dynamic) - Tunnel logical group Id used for encapsulation and decapsulation
* `tunnel_type` (String) - Giga Tunnel Type

<a id="nestedatt--tunnel_logical_groups--decap_tunnel_infos"></a>
### Nested Schema for `tunnel_logical_groups.decap_tunnel_infos`

Optional:

* `tunnel_operation_type` (String) - Giga Tunnel Operational Type
* `tunnel_type` (String) - Giga Tunnel Type

<a id="nestedatt--tunnel_logical_groups--encap_tunnel_infos"></a>
### Nested Schema for `tunnel_logical_groups.encap_tunnel_infos`

Optional:

* `tunnel_operation_type` (String) - Giga Tunnel Operational Type
* `tunnel_type` (String) - Giga Tunnel Type

<a id="nestedatt--tunnel_logical_groups--health_state_reasons"></a>
### Nested Schema for `tunnel_logical_groups.health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

