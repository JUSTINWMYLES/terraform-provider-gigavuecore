---
page_title: "gigavuecore_clear_nrt_stats_fabric_map Action - gigavuecore"
subcategory: ""
description: |-
  Clear registered Fabric Map from NRT stats
---

# gigavuecore_clear_nrt_stats_fabric_map Action

Clear registered Fabric Map from NRT stats

## Example Usage

```terraform
action "gigavuecore_clear_nrt_stats_fabric_map" "example" {
  config {
    fabric_map_alias = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `fabric_map_alias` (String, required) - alias of the fabric map


