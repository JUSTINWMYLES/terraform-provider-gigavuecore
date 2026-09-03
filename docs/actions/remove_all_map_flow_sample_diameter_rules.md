---
page_title: "gigavuecore_remove_all_map_flow_sample_diameter_rules Action - gigavuecore"
subcategory: ""
description: |-
  Remove all flowSampleDiameterRules from a 'secondLevel/flowSampleDiameter' map
---

# gigavuecore_remove_all_map_flow_sample_diameter_rules Action

Remove all flowSampleDiameterRules from a 'secondLevel/flowSampleDiameter' map

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

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


