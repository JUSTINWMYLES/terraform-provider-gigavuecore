---
page_title: "gigavuecore_add_map_flow_whitelist_overlap_rule Action - gigavuecore"
subcategory: ""
description: |-
  Add new flowWhitelistOverlapRule to a 'secondLevel/flowWhitelistOverlap' map
---

# gigavuecore_add_map_flow_whitelist_overlap_rule Action

Add new flowWhitelistOverlapRule to a 'secondLevel/flowWhitelistOverlap' map

## Example Usage

```terraform
action "gigavuecore_add_map_flow_whitelist_overlap_rule" "example" {
  config {
    alias = "example"
    flow5_g = null
    gtp = null
    rule_id = 1
    sip = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `flow5_g` (Dynamic, optional) - Map Flow Whitelist 5g Rule GTP match Definition. Private class
* `gtp` (Dynamic, optional) - Map Flow Whitelist Rule GTP match Definition. Private class
* `rule_id` (Number, required)
* `sip` (Dynamic, optional) - Map Flow Whitelist Rule Sip match definition
