---
page_title: "gigavuecore_get_traffic_flows_global_settings Data Source - gigavuecore"
subcategory: ""
description: |-
  Get global settings for Traffic Flows
---

# gigavuecore_get_traffic_flows_global_settings Data Source

Get global settings for Traffic Flows

## Example Usage

```terraform
data "gigavuecore_get_traffic_flows_global_settings" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `auto_migrate` (Bool, computed)
* `fabric_resource` (Object({mode, scope, type}), computed)
  * `mode` (String, computed) - \['SHARE' or 'NOT\_SHARE'\]: Resource sharing mode. SHARE: same resource can be shared by different fabric maps; NOT\_SHARE: not shared. Default: NOT\_SHARE.
  * `scope` (String, computed) - \['GLOBAL'\]: Scope of resource pool. GLOBAL: only one global resource pool.
  * `type` (String, computed) - \['L2CIRCUIT'\]: Resource type.
* `l2_circuit` (Object({vlan_ids}), computed)
  * `vlan_ids` (String, computed) - VLAN id ranges used for L2CIRCUIT resource type.

