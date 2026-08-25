---
page_title: "gigavuecore_update_map_template_rule Action - gigavuecore"
subcategory: ""
description: |-
  Update a map template rule
---

# gigavuecore_update_map_template_rule Action

Update a map template rule

## Example Usage

```terraform
action "gigavuecore_update_map_template_rule" "example" {
  config {
    alias        = "example"
    bidi         = true
    body_rule_id = 1
    cluster_id   = "example"
    comment      = "example"
    ip_rewrite   = null
    matches      = "example"
    rewrite      = null
    rule_id      = "example"
    rule_type    = "example"
    vlan_tag     = null
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map template
* `bidi` (Boolean, optional)
* `body_rule_id` (Number, required)
* `cluster_id` (String, required) - Target Cluster ID
* `comment` (String, optional)
* `ip_rewrite` (Dynamic, optional) - IpRewrite options on the packets
* `matches` (Set of Dynamic, required) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rewrite` (Dynamic, optional) - Rewrite options on the packets
* `rule_id` (String, required) - id of the rule to update
* `rule_type` (String, required) - map template rule type
* `vlan_tag` (Dynamic, optional)


