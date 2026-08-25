---
page_title: "gigavuecore_remove_map_flow_sample_sip_rule Action - gigavuecore"
subcategory: ""
description: |-
  Remove a flowSampleSipRule from a 'secondLevel/flowSampleSip' map
---

# gigavuecore_remove_map_flow_sample_sip_rule Action

Remove a flowSampleSipRule from a 'secondLevel/flowSampleSip' map

## Example Usage

```terraform
action "gigavuecore_remove_map_flow_sample_sip_rule" "example" {
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
* `rule_id` (Number, required) - flowSampleSipRule id


