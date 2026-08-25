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
  alias                 = null
  cluster_configs       = []
  config_status         = null
  config_status_reasons = []
  resilient_config      = {}
  target_traffic_path   = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional) - Alias of the solution
* `cluster_configs` (Attributes List, optional) - FlexInline solution source details (see [below for nested schema](#nestedatt--cluster_configs))
* `config_status` (String, optional)
* `config_status_reasons` (List of String, optional)
* `resilient_config` (Attributes, optional) (see [below for nested schema](#nestedatt--resilient_config))
* `target_traffic_path` (String, optional) - Inline network target traffic path

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `alias` (String, computed) - Alias of the solution
* `cluster_configs` (Attributes List, computed) - FlexInline solution source details (see [below for nested schema](#nestedatt--cluster_configs))
* `config_status` (String, computed)
* `config_status_reasons` (List of String, computed)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `resilient_config` (Attributes, computed) (see [below for nested schema](#nestedatt--resilient_config))
* `target_traffic_path` (String, computed) - Inline network target traffic path

<a id="nestedatt--cluster_configs"></a>
### Nested Schema for `cluster_configs`

Optional:

* `cluster_id` (String)
* `export_criteria` (Attributes) - Based on this criteria and source on which this criteria is defined, traffic is guided to the other node (see [below for nested schema](#nestedatt--cluster_configs--export_criteria))
* `export_type` (String)
* `ib_pathway` (String) - Inter-broker pathway to guide traffic from one node to other
* `source` (Attributes) (see [below for nested schema](#nestedatt--cluster_configs--source))
<a id="nestedatt--cluster_configs--export_criteria"></a>
### Nested Schema for `cluster_configs.export_criteria`

Optional:

* `lsb` (Number)
<a id="nestedatt--cluster_configs--source"></a>
### Nested Schema for `cluster_configs.source`

Optional:

* `alias` (String) - Alias of the source
* `type` (String) - Source type
<a id="nestedatt--resilient_config"></a>
### Nested Schema for `resilient_config`

Optional:

* `side_a` (String) - Side A of the resilient map
* `side_b` (String) - Side B of the resilient map
<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_flex_inline.example {alias}
```
