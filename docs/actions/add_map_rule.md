---
page_title: "gigavuecore_add_map_rule Action - gigavuecore"
subcategory: ""
description: |-
  Add new rule to a 'regular/byRule', 'inline/byRule' or 'firstLevel/byRule' map
---

# gigavuecore_add_map_rule Action

Add new rule to a 'regular/byRule', 'inline/byRule' or 'firstLevel/byRule' map

## Example Usage

```terraform
action "gigavuecore_add_map_rule" "example" {
  config {
    alias      = "example"
    bidi       = true
    cluster_id = "example"
    comment    = "example"
    ip_rewrite = null
    matches    = "example"
    rewrite    = null
    rule_id    = 1
    rule_type  = "example"
    vlan_tag   = null
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `bidi` (Boolean, optional)
* `cluster_id` (String, required) - Target Cluster ID
* `comment` (String, optional)
* `ip_rewrite` (Dynamic, optional) - IpRewrite options on the packets
* `matches` (Set of Dynamic, required) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rewrite` (Dynamic, optional) - Rewrite options on the packets
* `rule_id` (Number, required)
* `rule_type` (String, required) - map rule type
* `vlan_tag` (Dynamic, optional)


