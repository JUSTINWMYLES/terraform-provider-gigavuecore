---
page_title: "gigavuecore_remove_map_ap_rule Action - gigavuecore"
subcategory: ""
description: |-
  Remove a apRule from a 'secondlevel/byRule' map
---

# gigavuecore_remove_map_ap_rule Action

Remove a apRule from a 'secondlevel/byRule' map

## Example Usage

```terraform
action "gigavuecore_remove_map_ap_rule" "example" {
  config {
    alias   = "example"
    rule_id = 0
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `rule_id` (Number, required) - apRule id


