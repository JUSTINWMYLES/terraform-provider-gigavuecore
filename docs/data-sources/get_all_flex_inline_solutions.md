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

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - Alias of the solution
* `cluster_configs` (Attributes List) - FlexInline solution source details (see [below for nested schema](#nestedatt--items--cluster_configs))
* `config_status` (String)
* `config_status_reasons` (List of String)
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--health_state_reasons))
* `resilient_config` (Attributes) (see [below for nested schema](#nestedatt--items--resilient_config))
* `target_traffic_path` (String) - Inline network target traffic path
<a id="nestedatt--items--cluster_configs"></a>
### Nested Schema for `items.cluster_configs`

Read-Only:

* `cluster_id` (String)
* `export_criteria` (Attributes) - Based on this criteria and source on which this criteria is defined, traffic is guided to the other node (see [below for nested schema](#nestedatt--items--cluster_configs--export_criteria))
* `export_type` (String)
* `ib_pathway` (String) - Inter-broker pathway to guide traffic from one node to other
* `source` (Attributes) (see [below for nested schema](#nestedatt--items--cluster_configs--source))
<a id="nestedatt--items--cluster_configs--export_criteria"></a>
### Nested Schema for `items.cluster_configs.export_criteria`

Read-Only:

* `lsb` (Number)
<a id="nestedatt--items--cluster_configs--source"></a>
### Nested Schema for `items.cluster_configs.source`

Read-Only:

* `alias` (String) - Alias of the source
* `type` (String) - Source type
<a id="nestedatt--items--health_state_reasons"></a>
### Nested Schema for `items.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--items--resilient_config"></a>
### Nested Schema for `items.resilient_config`

Read-Only:

* `side_a` (String) - Side A of the resilient map
* `side_b` (String) - Side B of the resilient map

