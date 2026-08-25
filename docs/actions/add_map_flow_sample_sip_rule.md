---
page_title: "gigavuecore_add_map_flow_sample_sip_rule Action - gigavuecore"
subcategory: ""
description: |-
  Add new flowSampleSipRule to a 'secondLevel/flowSampleSip' map
---

# gigavuecore_add_map_flow_sample_sip_rule Action

Add new flowSampleSipRule to a 'secondLevel/flowSampleSip' map

## Example Usage

```terraform
action "gigavuecore_add_map_flow_sample_sip_rule" "example" {
  config {
    alias      = "example"
    percentage = 1
    rule_id    = 1
    sip        = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `percentage` (Number, required)
* `rule_id` (Number, required)
* `sip` (Dynamic, required)


