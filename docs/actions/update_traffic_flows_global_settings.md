---
page_title: "gigavuecore_update_traffic_flows_global_settings Action - gigavuecore"
subcategory: ""
description: |-
  Update global settings for Traffic Flows
---

# gigavuecore_update_traffic_flows_global_settings Action

Update global settings for Traffic Flows

## Example Usage

```terraform
action "gigavuecore_update_traffic_flows_global_settings" "example" {
  config {
    body = {
      auto_migrate = true
      fabric_resource = {
        mode  = "example"
        scope = "example"
        type  = "example"
      }
      l2_circuit = {
        vlan_ids = "example"
      }
    }
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `body` (Attributes, required) - Settings for traffic flows, including migration and fabric resource configuration. (see [below for nested schema](#nestedatt--body))

<a id="nestedatt--body"></a>
### Nested Schema for `body`

Optional:

* `auto_migrate` (Boolean)
* `fabric_resource` (Attributes) (see [below for nested schema](#nestedatt--body--fabric_resource))
* `l2_circuit` (Attributes) (see [below for nested schema](#nestedatt--body--l2_circuit))

<a id="nestedatt--body--fabric_resource"></a>
### Nested Schema for `body.fabric_resource`

Optional:

* `mode` (String) - \['SHARE' or 'NOT\_SHARE'\]: Resource sharing mode. SHARE: same resource can be shared by different fabric maps; NOT\_SHARE: not shared. Default: NOT\_SHARE.
* `scope` (String) - \['GLOBAL'\]: Scope of resource pool. GLOBAL: only one global resource pool.
* `type` (String) - \['L2CIRCUIT'\]: Resource type.

<a id="nestedatt--body--l2_circuit"></a>
### Nested Schema for `body.l2_circuit`

Required:

* `vlan_ids` (String) - VLAN id ranges used for L2CIRCUIT resource type.

