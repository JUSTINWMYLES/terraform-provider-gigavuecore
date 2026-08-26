---
page_title: "gigavuecore_update_map_flow_sample5_g_overlap_rule Action - gigavuecore"
subcategory: ""
description: |-
  update flowSample5gOverlapRule of a 'secondLevel/flowSample5gOverlap' map
---

# gigavuecore_update_map_flow_sample5_g_overlap_rule Action

update flowSample5gOverlapRule of a 'secondLevel/flowSample5gOverlap' map

## Example Usage

```terraform
action "gigavuecore_update_map_flow_sample5_g_overlap_rule" "example" {
  config {
    alias        = "example"
    body_rule_id = 1
    comment      = "example"
    flow5_g      = "example"
    percentage   = 1
    rule_id      = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `body_rule_id` (Number, required)
* `comment` (String, optional)
* `flow5_g` (Dynamic, required) - Map Flow Sample Overlap Rule 5g match Definition. Private class
* `percentage` (Number, required)
* `rule_id` (String, required) - id of the rule to update


