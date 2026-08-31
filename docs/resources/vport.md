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
  alias            = "example"
  cluster_id       = "example"
  deferred_binding = true
  fail_over_action = "vport-bypass"
  gs_group         = "example"
  health_state     = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  inline_status      = "up"
  inner_traffic_path = "to-inline-tool"
  metadata_monitoring = {
    action    = "enable"
    exporters = ["example"]
  }
  mode               = "none"
  outer_traffic_path = "to-inline-tool"
  sa_apf_profile     = "example"
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

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

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
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_vport.example {alias}/{cluster_id}
```
