---
page_title: "gigavuecore_add_map_gs_rule Action - gigavuecore"
subcategory: ""
description: |-
  Add new gsRule to a 'secondLevel/byRule' map
---

# gigavuecore_add_map_gs_rule Action

Add new gsRule to a 'secondLevel/byRule' map

## Example Usage

```terraform
action "gigavuecore_add_map_gs_rule" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    comment    = "example"
    matches    = [ "example" ]
    rule_id    = 1
    rule_type  = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `cluster_id` (String, required) - Target Cluster ID
* `comment` (String, optional)
* `matches` (Set of Dynamic, required) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number, required)
* `rule_type` (String, required) - map rule type


