---
page_title: "gigavuecore_get_all_cluster_circuit_tunnels_of_an_user_fabric_map List Resource - gigavuecore"
subcategory: ""
description: |-
  Get all cluster circuit tunnel endpoints of a user-defined fabric map
---

# gigavuecore_get_all_cluster_circuit_tunnels_of_an_user_fabric_map List Resource

Get all cluster circuit tunnel endpoints of a user-defined fabric map

## Example Usage

```terraform
list "gigavuecore_get_all_cluster_circuit_tunnels_of_an_user_fabric_map" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    alias = "example"
    mode  = "example"
    type  = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the fabric map
* `mode` (String, optional) - tunnel mode: 'encap' or 'decap'.
* `type` (String, optional) - tunnel type: 'circuit' or 'vxlan'.


### Identity Attributes

The following identity attributes are exported for each matching result:

* `alias` (String, computed)
* `cct_alias` (String, computed)


