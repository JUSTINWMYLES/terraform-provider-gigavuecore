---
page_title: "gigavuecore_remove_map_flow_whitelist_rule Action - gigavuecore"
subcategory: ""
description: |-
  Remove a flowRule from a 'secondLevel/flowWhitelist' map
---

# gigavuecore_remove_map_flow_whitelist_rule Action

Remove a flowRule from a 'secondLevel/flowWhitelist' map

## Example Usage

```terraform
action "gigavuecore_remove_map_flow_whitelist_rule" "example" {
  config {
    alias   = "example"
    rule_id = 0
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `rule_id` (Number, required) - flowWhitelistRule id


