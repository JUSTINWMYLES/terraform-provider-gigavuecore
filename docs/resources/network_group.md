---
page_title: "gigavuecore_network_group Resource - gigavuecore"
subcategory: ""
description: |-
  Get FM Network group details by alias
---

# gigavuecore_network_group Resource

Get FM Network group details by alias

## Example Usage

```terraform
resource "gigavuecore_network_group" "example" {
  alias                = null
  cluster_id           = null
  health_state         = null
  health_state_reasons = []
  members              = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional) - Alias of the network group
* `cluster_id` (String, optional)
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `members` (Attributes List, optional) - Array holding inline network group members (see [below for nested schema](#nestedatt--members))

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `alias` (String, computed) - Alias of the network group
* `cluster_id` (String, computed)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `members` (Attributes List, computed) - Array holding inline network group members (see [below for nested schema](#nestedatt--members))

<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--members"></a>
### Nested Schema for `members`

Optional:

* `alias` (String) - Alias of the inline network
* `cluster_name` (String)
* `type` (String) - Type of inline construct

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_network_group.example {alias}
```
