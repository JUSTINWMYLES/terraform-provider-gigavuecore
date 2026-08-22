---
page_title: "gigavuecore_remove_all_map_flow_sample_diameter_rules Action - gigavuecore"
subcategory: ""
description: |-
  Remove all flowSampleDiameterRules from a 'secondLevel/flowSampleDiameter' map
---

# gigavuecore_remove_all_map_flow_sample_diameter_rules Action

Remove all flowSampleDiameterRules from a 'secondLevel/flowSampleDiameter' map

## Example Usage

```terraform
action "gigavuecore_remove_all_map_flow_sample_diameter_rules" "example" {
  config {
    alias = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
