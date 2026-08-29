---
page_title: "gigavuecore_get_all_giga_flex_inline_network_group Data Source - gigavuecore"
subcategory: ""
description: |-
  Get All FM Inline Network Groups
---

# gigavuecore_get_all_giga_flex_inline_network_group Data Source

Get All FM Inline Network Groups

## Example Usage

```terraform
data "gigavuecore_get_all_giga_flex_inline_network_group" "example" {
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - if provided, Network groups only for that cluster is returned

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - Alias of the network group
* `cluster_id` (String)
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--health_state_reasons))
* `members` (Attributes List) - Array holding inline network group members (see [below for nested schema](#nestedatt--items--members))
<a id="nestedatt--items--health_state_reasons"></a>
### Nested Schema for `items.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--items--members"></a>
### Nested Schema for `items.members`

Read-Only:

* `alias` (String) - Alias of the inline network
* `cluster_name` (String)
* `type` (String) - Type of inline construct

