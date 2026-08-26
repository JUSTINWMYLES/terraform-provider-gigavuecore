---
page_title: "gigavuecore_get_flex_inline_map_config_status Data Source - gigavuecore"
subcategory: ""
description: |-
  Gets all flexInline maps config status under the given solution
---

# gigavuecore_get_flex_inline_map_config_status Data Source

Gets all flexInline maps config status under the given solution

## Example Usage

```terraform
data "gigavuecore_get_flex_inline_map_config_status" "example" {
  alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Solution alias for which map config status is needed

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `cluster_id` (String)
* `config_alias` (String) - Alias of the corresponding config object
* `config_data` (List of Dynamic) - Data holding configuration details
* `config_type` (String) - Type of the configuration object
* `deployment_status` (String) - Deployment status of the object
* `deployment_status_reasons` (List of String)
* `operation_type` (String) - Operation performed on the config object
* `solution_alias` (String) - Alias of the solution

