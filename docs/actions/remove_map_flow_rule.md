---
page_title: "gigavuecore_remove_map_flow_rule Action - gigavuecore"
subcategory: ""
description: |-
  Remove a flowRule from a 'secondLevel/flowFilter' map
---

# gigavuecore_remove_map_flow_rule Action

Remove a flowRule from a 'secondLevel/flowFilter' map

## Example Usage

```terraform
action "gigavuecore_remove_map_flow_rule" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    rule_id    = 1
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `cluster_id` (String, required) - Target Cluster ID
* `rule_id` (Number, required) - flowRule id


