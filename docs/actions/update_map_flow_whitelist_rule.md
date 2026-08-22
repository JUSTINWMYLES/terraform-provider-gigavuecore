---
page_title: "gigavuecore_update_map_flow_whitelist_rule Action - gigavuecore"
subcategory: ""
description: |-
  update flowRule of a 'secondLevel/flowWhitelist' map
---

# gigavuecore_update_map_flow_whitelist_rule Action

update flowRule of a 'secondLevel/flowWhitelist' map

## Example Usage

```terraform
action "gigavuecore_update_map_flow_whitelist_rule" "example" {
  config {
    alias = "example"
    body_rule_id = 1
    flow5_g = null
    gtp = null
    rule_id = "example"
    sip = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `body_rule_id` (Number, required)
* `flow5_g` (Dynamic, optional) - Map Flow Whitelist 5g Rule GTP match Definition. Private class
* `gtp` (Dynamic, optional) - Map Flow Whitelist Rule GTP match Definition. Private class
* `rule_id` (String, required) - id of the rule to update
* `sip` (Dynamic, optional) - Map Flow Whitelist Rule Sip match definition
