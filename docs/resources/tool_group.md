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
  alias                     = null
  cluster_id                = null
  comment                   = null
  current_state             = {}
  enabled                   = null
  failover_action           = null
  failover_mode             = null
  flex_status               = null
  flex_traffic_path         = null
  hash                      = null
  health_state              = null
  health_state_reasons      = []
  inline_tools              = []
  min_group_size            = null
  oper_status               = null
  release_spare_if_possible = null
  spare_inline_tool         = null
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

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `comment` (String, computed)
* `current_state` (Attributes, computed) - Inline Tool Group Smart Load Balancing definition (see [below for nested schema](#nestedatt--current_state))
* `enabled` (Boolean, computed) - setting to false is equivalent to forcing the inline tool group failure (useful for taking the inline tool group out of commission for maintenance or other purposes)
* `flex_status` (String, computed)
* `flex_traffic_path` (String, computed)
* `hash` (String, computed)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `min_group_size` (Number, computed) - the minimum number of inline tools (the inline tools in the list plus spare if configured) that must be up so that the entire inline-tool-group is considered up
* `oper_status` (String, computed)
* `release_spare_if_possible` (Boolean, computed) - when set to true, if the spare inline tool became active it remains active regardless of the health state of the originally failed primary inline tool
* `spare_inline_tool` (String, computed) - alias of an inline tool to which traffic is forwarded when the first failure occurs in the set of primary inline tools

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

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_tool_group.example {alias}
```
