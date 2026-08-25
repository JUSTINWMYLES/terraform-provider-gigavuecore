---
page_title: "gigavuecore_update_map_flow_sample_diameter_rule Action - gigavuecore"
subcategory: ""
description: |-
  update flowSampleDiameterRule of a 'secondLevel/flowSampleDiameter' map
---

# gigavuecore_update_map_flow_sample_diameter_rule Action

update flowSampleDiameterRule of a 'secondLevel/flowSampleDiameter' map

## Example Usage

```terraform
action "gigavuecore_update_map_flow_sample_diameter_rule" "example" {
  config {
    alias        = "example"
    body_rule_id = 1
    diameter     = "example"
    interface    = "example"
    percentage   = 1
    rule_id      = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `body_rule_id` (Number, required)
* `diameter` (Dynamic, required) - Map Flow Sample Diameter Rule Definition
* `interface` (String, required) - interface type
* `percentage` (Number, required)
* `rule_id` (String, required) - id of the rule to update


