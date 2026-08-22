---
page_title: "gigavuecore_get_flex_inline_map_config Data Source - gigavuecore"
subcategory: ""
description: |-
  Get maps related to the given solution alias
---

# gigavuecore_get_flex_inline_map_config Data Source

Get maps related to the given solution alias

## Example Usage

```terraform
data "gigavuecore_get_flex_inline_map_config" "example" {
  alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - if provided, returns maps configured for the given solution alias

### Attributes

In addition to all arguments above, the following attributes are exported:

* `cluster_id` (String, computed)
* `config_data` (List(Dynamic), computed) - Data holding configuration details
* `config_status` (String, computed) - Status of the created config object
* `config_status_reasons` (List(String), computed)
* `config_type` (String, computed) - Type of the configuration object
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `solution_alias` (String, computed) - Alias of the solution

