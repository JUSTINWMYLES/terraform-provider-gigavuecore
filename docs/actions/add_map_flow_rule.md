---
page_title: "gigavuecore_add_map_flow_rule Action - gigavuecore"
subcategory: ""
description: |-
  Add new flowRule to a 'secondLevel/flowFilter' map
---

# gigavuecore_add_map_flow_rule Action

Add new flowRule to a 'secondLevel/flowFilter' map

## Example Usage

```terraform
action "gigavuecore_add_map_flow_rule" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    gtp        = "example"
    rule_id    = 1
    rule_type  = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `cluster_id` (String, required) - Target Cluster ID
* `gtp` (Dynamic, required) - Map Flow Rule GTP match Definition. Private class
* `rule_id` (Number, required)
* `rule_type` (String, required) - map rule type


