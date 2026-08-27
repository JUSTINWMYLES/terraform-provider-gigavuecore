---
page_title: "gigavuecore_vport Resource - gigavuecore"
subcategory: ""
description: |-
  Find vPort by alias
---

# gigavuecore_vport Resource

Find vPort by alias

## Example Usage

```terraform
resource "gigavuecore_vport" "example" {
  alias                = null
  cluster_id           = null
  deferred_binding     = null
  fail_over_action     = null
  gs_group             = null
  health_state         = null
  health_state_reasons = []
  inline_status        = null
  inner_traffic_path   = null
  metadata_monitoring  = {}
  mode                 = null
  outer_traffic_path   = null
  sa_apf_profile       = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `cluster_id` (String, required) - Target Cluster ID
* `deferred_binding` (Boolean, optional) - enable/disable deferred-binding
* `fail_over_action` (String, optional)
* `gs_group` (String, required) - Alias of referenced managing GsGroup
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `inline_status` (String, optional)
* `inner_traffic_path` (String, optional) - Similar to inline-network traffic-path, applicable for inner map
* `metadata_monitoring` (Attributes, optional) (see [below for nested schema](#nestedatt--metadata_monitoring))
* `mode` (String, optional)
* `outer_traffic_path` (String, optional) - Similar to inline-tool/inline-tool-group flex-traffic path, applicable for outer map
* `sa_apf_profile` (String, optional) - ASF session profile

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `deferred_binding` (Boolean, computed) - enable/disable deferred-binding
* `fail_over_action` (String, computed)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `inline_status` (String, computed)
* `inner_traffic_path` (String, computed) - Similar to inline-network traffic-path, applicable for inner map
* `metadata_monitoring` (Attributes, computed) (see [below for nested schema](#nestedatt--metadata_monitoring))
* `mode` (String, computed)
* `outer_traffic_path` (String, computed) - Similar to inline-tool/inline-tool-group flex-traffic path, applicable for outer map
* `sa_apf_profile` (String, computed) - ASF session profile

<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--metadata_monitoring"></a>
### Nested Schema for `metadata_monitoring`

Optional:

* `action` (String) - metadata monitoring action
* `exporters` (List of String)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_vport.example {alias}
```
