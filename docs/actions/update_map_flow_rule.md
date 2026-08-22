---
page_title: "gigavuecore_update_map_flow_rule Action - gigavuecore"
subcategory: ""
description: |-
  update flowRule of a 'secondLevel/flowFilter' map
---

# gigavuecore_update_map_flow_rule Action

update flowRule of a 'secondLevel/flowFilter' map

## Example Usage

```terraform
action "gigavuecore_update_map_flow_rule" "example" {
  config {
    alias = "example"
    body_rule_id = 1
    cluster_id = "example"
    gtp = "example"
    rule_id = "example"
    rule_type = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `body_rule_id` (Number, required)
* `cluster_id` (String, required) - Target Cluster ID
* `gtp` (Dynamic, required) - Map Flow Rule GTP match Definition. Private class
* `rule_id` (String, required) - id of the rule to update
* `rule_type` (String, required) - map rule type
