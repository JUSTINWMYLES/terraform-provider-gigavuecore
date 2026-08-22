---
page_title: "gigavuecore_add_map_flow_sample_overlap_rule Action - gigavuecore"
subcategory: ""
description: |-
  Add new flowSampleOverlapRule to a 'secondLevel/flowSampleOverlap' map
---

# gigavuecore_add_map_flow_sample_overlap_rule Action

Add new flowSampleOverlapRule to a 'secondLevel/flowSampleOverlap' map

## Example Usage

```terraform
action "gigavuecore_add_map_flow_sample_overlap_rule" "example" {
  config {
    alias = "example"
    comment = "example"
    gtp = "example"
    percentage = 1
    periodic_recalc = true
    priority = 1
    rule_id = 1
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `comment` (String, optional)
* `gtp` (Dynamic, required) - Map Flow Sample Rule GTP match Definition. Private class
* `percentage` (Number, required)
* `periodic_recalc` (Bool, optional) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number, optional) - maximum value is equal to the number of rules upon completion of the request
* `rule_id` (Number, required)
