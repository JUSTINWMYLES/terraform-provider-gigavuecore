---
page_title: "gigavuecore_remove_map_gs_rule Action - gigavuecore"
subcategory: ""
description: |-
  Remove a gsRule from a 'secondLevel/byRule' map
---

# gigavuecore_remove_map_gs_rule Action

Remove a gsRule from a 'secondLevel/byRule' map

## Example Usage

```terraform
action "gigavuecore_remove_map_gs_rule" "example" {
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
* `rule_id` (Number, required) - GsRule id


