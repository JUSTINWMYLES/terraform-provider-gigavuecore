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
  alias = "example"
  cluster_configs = [{
    cluster_id = "example"
    export_criteria = {
      lsb = 0
    }
    export_type = "ipBased"
    ib_pathway  = "example"
    source = {
      alias = "example"
      type  = "IN"
    }
  }]
  health_state = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  keep_on_device = true
  resilient_config = {
    side_a = "source"
    side_b = "source"
  }
  target_traffic_path = "drop"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional) - Alias of the solution
* `cluster_configs` (Attributes List, optional) - FlexInline solution source details (see [below for nested schema](#nestedatt--cluster_configs))
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `keep_on_device` (Boolean, optional) - if true, deletes the map associated to solution from device as well as from FM
* `resilient_config` (Attributes, optional) (see [below for nested schema](#nestedatt--resilient_config))
* `target_traffic_path` (String, optional) - Inline network target traffic path

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `config_status` (String, computed)
* `config_status_reasons` (List of String, computed)

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--cluster_configs"></a>
### Nested Schema for `cluster_configs`

Optional:

* `cluster_id` (String)
* `export_criteria` (Attributes) - Based on this criteria and source on which this criteria is defined, traffic is guided to the other node (see [below for nested schema](#nestedatt--cluster_configs--export_criteria))
* `export_type` (String)
* `ib_pathway` (String) - Inter-broker pathway to guide traffic from one node to other
* `source` (Attributes) - Object holding flexInline source details (see [below for nested schema](#nestedatt--cluster_configs--source))

<a id="nestedatt--cluster_configs--export_criteria"></a>
### Nested Schema for `cluster_configs.export_criteria`

Optional:

* `lsb` (Number)

<a id="nestedatt--cluster_configs--source"></a>
### Nested Schema for `cluster_configs.source`

Optional:

* `alias` (String) - Alias of the source
* `type` (String) - Source type

<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

<a id="nestedatt--resilient_config"></a>
### Nested Schema for `resilient_config`

Optional:

* `side_a` (String) - Side A of the resilient map
* `side_b` (String) - Side B of the resilient map
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_flex_inline.example {alias}
```
