---
page_title: "gigavuecore_remove_map_rule Action - gigavuecore"
subcategory: ""
description: |-
  Remove a rule from a 'regular/byRule', 'inline/byRule' or 'firstLevel/byRule' map
---

# gigavuecore_remove_map_rule Action

Remove a rule from a 'regular/byRule', 'inline/byRule' or 'firstLevel/byRule' map

## Example Usage

```terraform
action "gigavuecore_remove_map_rule" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    rule_id    = 1
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `cluster_id` (String, required) - Target Cluster ID
* `rule_id` (Number, required) - map rule id


