---
page_title: "gigavuecore_update_map_flow_sample_overlap_rule Action - gigavuecore"
subcategory: ""
description: |-
  update flowRule of a 'secondLevel/flowSampleOverlap' map
---

# gigavuecore_update_map_flow_sample_overlap_rule Action

update flowRule of a 'secondLevel/flowSampleOverlap' map

## Example Usage

```terraform
action "gigavuecore_update_map_flow_sample_overlap_rule" "example" {
  config {
    alias           = "example"
    body_rule_id    = 1
    comment         = "example"
    gtp             = "example"
    percentage      = 1
    periodic_recalc = true
    priority        = 1
    rule_id         = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `body_rule_id` (Number, required)
* `comment` (String, optional)
* `gtp` (Dynamic, required) - Map Flow Sample Rule GTP match Definition. Private class
* `percentage` (Number, required)
* `periodic_recalc` (Boolean, optional) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number, optional) - maximum value is equal to the number of rules upon completion of the request
* `rule_id` (String, required) - id of the rule to update


