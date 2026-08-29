---
page_title: "gigavuecore_update_map_gs_rule Action - gigavuecore"
subcategory: ""
description: |-
  update gsRule of a 'secondLevel/byRule' map
---

# gigavuecore_update_map_gs_rule Action

update gsRule of a 'secondLevel/byRule' map

## Example Usage

```terraform
action "gigavuecore_update_map_gs_rule" "example" {
  config {
    alias        = "example"
    body_rule_id = 0
    cluster_id   = "example"
    comment      = "example"
    matches      = [ "example" ]
    rule_id      = "example"
    rule_type    = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `body_rule_id` (Number, required)
* `cluster_id` (String, required) - Target Cluster ID
* `comment` (String, optional)
* `matches` (Set of Dynamic, required) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (String, required) - id of the rule to update
* `rule_type` (String, required) - map rule type


