---
page_title: "gigavuecore_create_flex_inline_config Action - gigavuecore"
subcategory: ""
description: |-
  Creates flexInline maps under the given solution alias
---

# gigavuecore_create_flex_inline_config Action

Creates flexInline maps under the given solution alias

## Example Usage

```terraform
action "gigavuecore_create_flex_inline_config" "example" {
  config {
    alias                 = "example"
    bypass_strategy       = "example"
    cluster_id            = "example"
    config_data           = [ "example" ]
    config_status         = "OPEN"
    config_status_reasons = [ "example" ]
    config_type           = "FLEXINLINE_MAP"
    health_state          = "green"
    health_state_reasons = [{
      message                               = "example"
      severity                              = "green"
      traffic_health_state_computation_type = "PORT_LOW_UTIL"
    }]
    solution_alias = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Creates maps for the given solution
* `bypass_strategy` (String, optional) - if bypass, Inline network's traffic path would be bypass during the solution deployment and would return to the original state on successful deployment, If skipBypass, Inline network's traffic path would not be changed during the deployment
* `cluster_id` (String, optional)
* `config_data` (List of Dynamic, optional) - Data holding configuration details
* `config_status` (String, optional) - Status of the created config object
* `config_status_reasons` (List of String, optional)
* `config_type` (String, optional) - Type of the configuration object
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `solution_alias` (String, optional) - Alias of the solution

<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

