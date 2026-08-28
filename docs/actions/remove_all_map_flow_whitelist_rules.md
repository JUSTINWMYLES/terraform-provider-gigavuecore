---
page_title: "gigavuecore_remove_all_map_flow_whitelist_rules Action - gigavuecore"
subcategory: ""
description: |-
  Remove all flowRules from a 'secondLevel/flowWhitelist' map
---

# gigavuecore_remove_all_map_flow_whitelist_rules Action

Remove all flowRules from a 'secondLevel/flowWhitelist' map

## Example Usage

```terraform
action "gigavuecore_remove_all_map_flow_whitelist_rules" "example" {
  config {
    alias = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map


