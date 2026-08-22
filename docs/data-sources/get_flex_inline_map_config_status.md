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

* `items` (List(Object({cluster_id, config_alias, config_data, config_type, deployment_status, deployment_status_reasons, operation_type, solution_alias})), computed)

