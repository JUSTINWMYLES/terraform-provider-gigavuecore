---
page_title: "gigavuecore_remove_map_flow_whitelist_overlap_rule Action - gigavuecore"
subcategory: ""
description: |-
  Remove a flowRule from a 'secondLevel/flowWhitelistOverlap' map
---

# gigavuecore_remove_map_flow_whitelist_overlap_rule Action

Remove a flowRule from a 'secondLevel/flowWhitelistOverlap' map

## Example Usage

```terraform
action "gigavuecore_remove_map_flow_whitelist_overlap_rule" "example" {
  config {
    alias = "example"
    rule_id = 1
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `rule_id` (Number, required) - flowWhitelistOverlapRule id
