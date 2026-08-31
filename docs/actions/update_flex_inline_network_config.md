---
page_title: "gigavuecore_update_flex_inline_network_config Action - gigavuecore"
subcategory: ""
description: |-
  Updates inline network details used in the given solution
---

# gigavuecore_update_flex_inline_network_config Action

Updates inline network details used in the given solution

## Example Usage

```terraform
action "gigavuecore_update_flex_inline_network_config" "example" {
  config {
    alias                 = "example"
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

* `alias` (String, required) - Solution alias for which inline network is to be updated
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

