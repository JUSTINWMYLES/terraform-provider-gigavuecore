---
page_title: "gigavuecore_add_map_flow_sample5_g_rule Action - gigavuecore"
subcategory: ""
description: |-
  Add new flowSample5gRule to a 'secondLevel/flowSample5g' map
---

# gigavuecore_add_map_flow_sample5_g_rule Action

Add new flowSample5gRule to a 'secondLevel/flowSample5g' map

## Example Usage

```terraform
action "gigavuecore_add_map_flow_sample5_g_rule" "example" {
  config {
    alias = "example"
    comment = "example"
    flow5_g = "example"
    percentage = 1
    priority = 1
    rule_id = 1
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `comment` (String, optional)
* `flow5_g` (Dynamic, required) - Map Flow Sample Rule 5g match Definition. Private class
* `percentage` (Number, required)
* `priority` (Number, optional)
* `rule_id` (Number, required)
