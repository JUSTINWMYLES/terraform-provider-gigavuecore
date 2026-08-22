---
page_title: "gigavuecore_remove_all_map_flow_sample5_g_overlap_rules Action - gigavuecore"
subcategory: ""
description: |-
  Remove all flowSample5gOverlapRules from a 'secondLevel/flowSample5gOverlap' map
---

# gigavuecore_remove_all_map_flow_sample5_g_overlap_rules Action

Remove all flowSample5gOverlapRules from a 'secondLevel/flowSample5gOverlap' map

## Example Usage

```terraform
action "gigavuecore_remove_all_map_flow_sample5_g_overlap_rules" "example" {
  config {
    alias = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
