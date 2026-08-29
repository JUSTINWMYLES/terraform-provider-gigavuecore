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

### Attributes

In addition to all arguments above, the following attributes are exported:

* `auto_migrate` (Boolean, computed)
* `fabric_resource` (Attributes, computed) (see [below for nested schema](#nestedatt--fabric_resource))
* `l2_circuit` (Attributes, computed) (see [below for nested schema](#nestedatt--l2_circuit))

<a id="nestedatt--fabric_resource"></a>
### Nested Schema for `fabric_resource`

Read-Only:

* `mode` (String) - \['SHARE' or 'NOT\_SHARE'\]: Resource sharing mode. SHARE: same resource can be shared by different fabric maps; NOT\_SHARE: not shared. Default: NOT\_SHARE.
* `scope` (String) - \['GLOBAL'\]: Scope of resource pool. GLOBAL: only one global resource pool.
* `type` (String) - \['L2CIRCUIT'\]: Resource type.
<a id="nestedatt--l2_circuit"></a>
### Nested Schema for `l2_circuit`

Read-Only:

* `vlan_ids` (String) - VLAN id ranges used for L2CIRCUIT resource type.

