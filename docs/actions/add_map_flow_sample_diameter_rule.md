---
page_title: "gigavuecore_add_map_flow_sample_diameter_rule Action - gigavuecore"
subcategory: ""
description: |-
  Add new flowSampleDiameterRule to a 'secondLevel/flowSampleDiameter' map
---

# gigavuecore_add_map_flow_sample_diameter_rule Action

Add new flowSampleDiameterRule to a 'secondLevel/flowSampleDiameter' map

## Example Usage

```terraform
action "gigavuecore_add_map_flow_sample_diameter_rule" "example" {
  config {
    alias      = "example"
    diameter   = "example"
    interface  = "example"
    percentage = 1
    rule_id    = 1
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `diameter` (Dynamic, required) - Map Flow Sample Diameter Rule Definition
* `interface` (String, required) - interface type
* `percentage` (Number, required)
* `rule_id` (Number, required)


