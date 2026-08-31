---
page_title: "gigavuecore_tool_group Resource - gigavuecore"
subcategory: ""
description: |-
  Find Inline Tool Group by alias
---

# gigavuecore_tool_group Resource

Find Inline Tool Group by alias

## Example Usage

```terraform
resource "gigavuecore_tool_group" "example" {
  alias      = "example"
  cluster_id = "example"
  comment    = "example"
  current_state = {
    inline_tools         = ["example"]
    spare_tool           = "example"
    spare_tool_status    = "down"
    switched_inline_tool = "example"
  }
  enabled           = true
  failover_action   = "toolBypass"
  failover_mode     = "disabled"
  flex_status       = "forwarding"
  flex_traffic_path = "drop"
  hash              = "advanced"
  health_state      = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  inline_tools              = ["example"]
  min_group_size            = 1
  oper_status               = "up"
  release_spare_if_possible = true
  spare_inline_tool         = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Inline Tool Group alias. Unique within a cluster
* `cluster_id` (String, required) - Target Cluster ID
* `comment` (String, optional)
* `current_state` (Attributes, optional) - Inline Tool Group Smart Load Balancing definition (see [below for nested schema](#nestedatt--current_state))
* `enabled` (Boolean, optional) - setting to false is equivalent to forcing the inline tool group failure (useful for taking the inline tool group out of commission for maintenance or other purposes)
* `failover_action` (String, required)
* `failover_mode` (String, required) - the way of handling a failure of an individual member of the inline tool list when no spare inline tool is configured or if the spare inline tool is failed
* `flex_status` (String, optional)
* `flex_traffic_path` (String, optional)
* `hash` (String, optional)
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `inline_tools` (List of String, required) - If no spare inline tool configured, list of aliases of inline tools participating in hash-based traffic distribution. If the spare inline tool is configured, list of aliases of primary inline tools to which traffic is forwarded as long as all of them are healthy. The number of inline tools in the list must be between 1 and 64 if the spare inline tool is configured or between 2 and 64 otherwise
* `min_group_size` (Number, optional) - the minimum number of inline tools (the inline tools in the list plus spare if configured) that must be up so that the entire inline-tool-group is considered up
* `oper_status` (String, optional)
* `release_spare_if_possible` (Boolean, optional) - when set to true, if the spare inline tool became active it remains active regardless of the health state of the originally failed primary inline tool
* `spare_inline_tool` (String, optional) - alias of an inline tool to which traffic is forwarded when the first failure occurs in the set of primary inline tools

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--current_state"></a>
### Nested Schema for `current_state`

Optional:

* `inline_tools` (List of String) - Current list of inlineTools
* `spare_tool` (String) - Current spare tool
* `spare_tool_status` (String) - Current spare tool status
* `switched_inline_tool` (String) - Current switched inlineTool

<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_tool_group.example {alias}/{cluster_id}
```
