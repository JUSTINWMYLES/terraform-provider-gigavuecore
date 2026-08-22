---
page_title: "gigavuecore_flex_inline Resource - gigavuecore"
subcategory: ""
description: |-
  Get a flexInlineSolution details by alias
---

# gigavuecore_flex_inline Resource

Get a flexInlineSolution details by alias

## Example Usage

```terraform
resource "gigavuecore_flex_inline" "example" {
  alias = null
  cluster_configs = []
  config_status = null
  config_status_reasons = []
  resilient_config = {}
  target_traffic_path = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional) - Alias of the solution
* `cluster_configs` (List(Object({cluster_id, export_criteria, export_type, ib_pathway, source})), optional) - FlexInline solution source details
* `config_status` (String, optional)
* `config_status_reasons` (List(String), optional)
* `resilient_config` (Object({side_a, side_b}), optional)
  * `side_a` (String, optional) - Side A of the resilient map
  * `side_b` (String, optional) - Side B of the resilient map
* `target_traffic_path` (String, optional) - Inline network target traffic path

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `alias` (String, computed) - Alias of the solution
* `cluster_configs` (List(Object({cluster_id, export_criteria, export_type, ib_pathway, source})), computed) - FlexInline solution source details
* `config_status` (String, computed)
* `config_status_reasons` (List(String), computed)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `resilient_config` (Object({side_a, side_b}), computed)
  * `side_a` (String, optional) - Side A of the resilient map
  * `side_b` (String, optional) - Side B of the resilient map
* `target_traffic_path` (String, computed) - Inline network target traffic path

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_flex_inline.example {alias}
```
