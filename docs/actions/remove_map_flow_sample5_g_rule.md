---
page_title: "gigavuecore_remove_map_flow_sample5_g_rule Action - gigavuecore"
subcategory: ""
description: |-
  Remove a flowSample5gRule from a 'secondLevel/flowSample5g' map
---

# gigavuecore_remove_map_flow_sample5_g_rule Action

Remove a flowSample5gRule from a 'secondLevel/flowSample5g' map

## Example Usage

```terraform
action "gigavuecore_remove_map_flow_sample5_g_rule" "example" {
  config {
    alias   = "example"
    rule_id = 1
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `rule_id` (Number, required) - flowSample5gRule id


