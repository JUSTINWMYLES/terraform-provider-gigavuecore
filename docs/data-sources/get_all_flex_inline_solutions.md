---
page_title: "gigavuecore_get_all_flex_inline_solutions Data Source - gigavuecore"
subcategory: ""
description: |-
  Get All flexInline solutions present on FM
---

# gigavuecore_get_all_flex_inline_solutions Data Source

Get All flexInline solutions present on FM

## Example Usage

```terraform
data "gigavuecore_get_all_flex_inline_solutions" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({alias, cluster_configs, config_status, config_status_reasons, health_state, health_state_reasons, resilient_config, target_traffic_path})), computed)

