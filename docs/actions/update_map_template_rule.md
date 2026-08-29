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
    body_rule_id = 0
    cluster_id   = "example"
    comment      = "example"
    ip_rewrite = {
      dst_ip = "example"
      src_ip = "example"
    }
    matches = [ "example" ]
    rewrite = {
      dst_mac = "example"
      src_mac = "example"
    }
    rule_id   = "example"
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

* `alias` (String, required) - alias of the target map template
* `bidi` (Boolean, optional)
* `body_rule_id` (Number, required)
* `cluster_id` (String, required) - Target Cluster ID
* `comment` (String, optional)
* `ip_rewrite` (Attributes, optional) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--ip_rewrite))
* `matches` (Set of Dynamic, required) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rewrite` (Attributes, optional) - Rewrite options on the packets (see [below for nested schema](#nestedatt--rewrite))
* `rule_id` (String, required) - id of the rule to update
* `rule_type` (String, required) - map template rule type
* `vlan_tag` (Attributes, optional) (see [below for nested schema](#nestedatt--vlan_tag))

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

