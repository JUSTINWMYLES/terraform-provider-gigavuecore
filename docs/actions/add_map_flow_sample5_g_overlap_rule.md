---
page_title: "gigavuecore_add_map_flow_sample5_g_overlap_rule Action - gigavuecore"
subcategory: ""
description: |-
  Add new flowSample5gOverlapRule to a 'secondLevel/flowSample5gOverlap' map
---

# gigavuecore_add_map_flow_sample5_g_overlap_rule Action

Add new flowSample5gOverlapRule to a 'secondLevel/flowSample5gOverlap' map

## Example Usage

```terraform
action "gigavuecore_add_map_flow_sample5_g_overlap_rule" "example" {
  config {
    alias      = "example"
    comment    = "example"
    flow5_g    = "example"
    percentage = 1
    rule_id    = 1
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `comment` (String, optional)
* `flow5_g` (Dynamic, required) - Map Flow Sample Overlap Rule 5g match Definition. Private class
* `percentage` (Number, required)
* `rule_id` (Number, required)


