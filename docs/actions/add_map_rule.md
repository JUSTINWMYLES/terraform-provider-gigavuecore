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
    ip_rewrite = {
      dst_ip = "example"
      src_ip = "example"
    }
    matches = [ "example" ]
    rewrite = {
      dst_mac = "example"
      src_mac = "example"
    }
    rule_id   = 0
    rule_type = "example"
    vlan_tag = {
      tag_protocol_id = "example"
      vlan_action     = "example"
      vlan_id         = 0
    }
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
* `ip_rewrite` (Attributes, optional) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--ip_rewrite))
* `matches` (Set of Dynamic, required) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rewrite` (Attributes, optional) - Rewrite options on the packets (see [below for nested schema](#nestedatt--rewrite))
* `rule_id` (Number, required)
* `rule_type` (String, required) - map rule type
* `vlan_tag` (Attributes, optional) - This field is only  applicable for pass rules (see [below for nested schema](#nestedatt--vlan_tag))

<a id="nestedatt--ip_rewrite"></a>
### Nested Schema for `ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)

<a id="nestedatt--rewrite"></a>
### Nested Schema for `rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)

<a id="nestedatt--vlan_tag"></a>
### Nested Schema for `vlan_tag`

Required:

* `vlan_action` (String)

Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)

