---
page_title: "gigavuecore_replace_flex_inline_map_config Action - gigavuecore"
subcategory: ""
description: |-
  Replace the existing map config with the provided configs
---

# gigavuecore_replace_flex_inline_map_config Action

Replace the existing map config with the provided configs

## Example Usage

```terraform
action "gigavuecore_replace_flex_inline_map_config" "example" {
  config {
    alias                 = "example"
    bypass_strategy       = "example"
    cluster_id            = "example"
    config_data           = [ "example" ]
    config_status         = "example"
    config_status_reasons = [ "example" ]
    config_type           = "example"
    health_state          = "example"
    health_state_reasons = [{
      message                               = "example"
      severity                              = "example"
      traffic_health_state_computation_type = "example"
    }]
    solution_alias = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Updates maps for the given solution
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

