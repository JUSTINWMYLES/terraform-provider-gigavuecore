---
page_title: "gigavuecore_update_map_flow_sample_sip_rule Action - gigavuecore"
subcategory: ""
description: |-
  update flowSampleSipRule of a 'secondLevel/flowSampleSip' map
---

# gigavuecore_update_map_flow_sample_sip_rule Action

update flowSampleSipRule of a 'secondLevel/flowSampleSip' map

## Example Usage

```terraform
action "gigavuecore_update_map_flow_sample_sip_rule" "example" {
  config {
    alias = "example"
    body_rule_id = 1
    percentage = 1
    rule_id = "example"
    sip = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `body_rule_id` (Number, required)
* `percentage` (Number, required)
* `rule_id` (String, required) - id of the rule to update
* `sip` (Dynamic, required)
