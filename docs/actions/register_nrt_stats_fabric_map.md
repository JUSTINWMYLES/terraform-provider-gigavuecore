---
page_title: "gigavuecore_register_nrt_stats_fabric_map Action - gigavuecore"
subcategory: ""
description: |-
  Register / de-register user-defined fabric map to NRT stats
---

# gigavuecore_register_nrt_stats_fabric_map Action

Register / de-register user-defined fabric map to NRT stats

## Example Usage

```terraform
action "gigavuecore_register_nrt_stats_fabric_map" "example" {
  config {
    fabric_map_alias = "example"
    operation        = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `fabric_map_alias` (String, required) - alias of the fabric map
* `operation` (String, optional) - operation type (add/delete) to update near real time config for fabric map


