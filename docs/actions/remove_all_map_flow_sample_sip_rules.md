---
page_title: "gigavuecore_remove_all_map_flow_sample_sip_rules Action - gigavuecore"
subcategory: ""
description: |-
  Remove all flowSampleSipRules from a 'secondLevel/flowSampleSip' map
---

# gigavuecore_remove_all_map_flow_sample_sip_rules Action

Remove all flowSampleSipRules from a 'secondLevel/flowSampleSip' map

## Example Usage

```terraform
action "gigavuecore_remove_all_map_flow_sample_sip_rules" "example" {
  config {
    alias = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map


