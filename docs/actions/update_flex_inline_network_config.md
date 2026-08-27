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
    config_data           = null
    config_status         = "example"
    config_status_reasons = [ "example" ]
    config_type           = "example"
    health_state          = "example"
    health_state_reasons  = null
    solution_alias        = "example"
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
* `health_state_reasons` (List of Dynamic, optional)
* `solution_alias` (String, optional) - Alias of the solution


