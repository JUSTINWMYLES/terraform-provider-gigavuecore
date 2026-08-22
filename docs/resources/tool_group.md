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
  alias = null
  comment = null
  current_state = {}
  enabled = null
  failover_action = null
  failover_mode = null
  flex_status = null
  flex_traffic_path = null
  hash = null
  health_state = null
  health_state_reasons = []
  inline_tools = []
  min_group_size = null
  oper_status = null
  release_spare_if_possible = null
  spare_inline_tool = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Inline Tool Group alias. Unique within a cluster
* `comment` (String, optional)
* `current_state` (Object({inline_tools, spare_tool, spare_tool_status, switched_inline_tool}), optional) - Inline Tool Group Smart Load Balancing definition
  * `inline_tools` (List(String), optional) - Current list of inlineTools
  * `spare_tool` (String, optional) - Current spare tool
  * `spare_tool_status` (String, optional) - Current spare tool status
  * `switched_inline_tool` (String, optional) - Current switched inlineTool
* `enabled` (Bool, optional) - setting to false is equivalent to forcing the inline tool group failure (useful for taking the inline tool group out of commission for maintenance or other purposes)
* `failover_action` (String, required)
* `failover_mode` (String, required) - the way of handling a failure of an individual member of the inline tool list when no spare inline tool is configured or if the spare inline tool is failed
* `flex_status` (String, optional)
* `flex_traffic_path` (String, optional)
* `hash` (String, optional)
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), optional)
* `inline_tools` (List(String), required) - If no spare inline tool configured, list of aliases of inline tools participating in hash-based traffic distribution. If the spare inline tool is configured, list of aliases of primary inline tools to which traffic is forwarded as long as all of them are healthy. The number of inline tools in the list must be between 1 and 64 if the spare inline tool is configured or between 2 and 64 otherwise
* `min_group_size` (Number, optional) - the minimum number of inline tools (the inline tools in the list plus spare if configured) that must be up so that the entire inline-tool-group is considered up
* `oper_status` (String, optional)
* `release_spare_if_possible` (Bool, optional) - when set to true, if the spare inline tool became active it remains active regardless of the health state of the originally failed primary inline tool
* `spare_inline_tool` (String, optional) - alias of an inline tool to which traffic is forwarded when the first failure occurs in the set of primary inline tools

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `comment` (String, computed)
* `current_state` (Object({inline_tools, spare_tool, spare_tool_status, switched_inline_tool}), computed) - Inline Tool Group Smart Load Balancing definition
  * `inline_tools` (List(String), optional) - Current list of inlineTools
  * `spare_tool` (String, optional) - Current spare tool
  * `spare_tool_status` (String, optional) - Current spare tool status
  * `switched_inline_tool` (String, optional) - Current switched inlineTool
* `enabled` (Bool, computed) - setting to false is equivalent to forcing the inline tool group failure (useful for taking the inline tool group out of commission for maintenance or other purposes)
* `flex_status` (String, computed)
* `flex_traffic_path` (String, computed)
* `hash` (String, computed)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `min_group_size` (Number, computed) - the minimum number of inline tools (the inline tools in the list plus spare if configured) that must be up so that the entire inline-tool-group is considered up
* `oper_status` (String, computed)
* `release_spare_if_possible` (Bool, computed) - when set to true, if the spare inline tool became active it remains active regardless of the health state of the originally failed primary inline tool
* `spare_inline_tool` (String, computed) - alias of an inline tool to which traffic is forwarded when the first failure occurs in the set of primary inline tools

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_tool_group.example {alias}
```
