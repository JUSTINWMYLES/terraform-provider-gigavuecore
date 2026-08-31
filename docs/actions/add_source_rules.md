---
page_title: "gigavuecore_add_source_rules Action - gigavuecore"
subcategory: ""
description: |-
  Add source rules to a policy and sourceRulesAlias
---

# gigavuecore_add_source_rules Action

Add source rules to a policy and sourceRulesAlias

## Example Usage

```terraform
action "gigavuecore_add_source_rules" "example" {
  config {
    drop_rules = [{
      bidi    = true
      comment = "example"
      ip_rewrite = {
        dst_ip = "example"
        src_ip = "example"
      }
      matches = [ "example" ]
      rewrite = {
        dst_mac = "example"
        src_mac = "example"
      }
      rule_id = 1
      vlan_tag = {
        tag_protocol_id = "0x8100"
        vlan_action     = "add"
        vlan_id         = 0
      }
    }]
    pass_rules = [{
      bidi    = true
      comment = "example"
      ip_rewrite = {
        dst_ip = "example"
        src_ip = "example"
      }
      matches = [ "example" ]
      rewrite = {
        dst_mac = "example"
        src_mac = "example"
      }
      rule_id = 1
      vlan_tag = {
        tag_protocol_id = "0x8100"
        vlan_action     = "add"
        vlan_id         = 0
      }
    }]
    policy_alias       = "example"
    source_rules_alias = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `drop_rules` (Attributes Set, optional) (see [below for nested schema](#nestedatt--drop_rules))
* `pass_rules` (Attributes Set, optional) (see [below for nested schema](#nestedatt--pass_rules))
* `policy_alias` (String, required) - Policy alias
* `source_rules_alias` (String, required) - Source rules alias

<a id="nestedatt--drop_rules"></a>
### Nested Schema for `drop_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)

Optional:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--drop_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--drop_rules--rewrite))
* `vlan_tag` (Attributes) - This field is only  applicable for pass rules (see [below for nested schema](#nestedatt--drop_rules--vlan_tag))

<a id="nestedatt--drop_rules--ip_rewrite"></a>
### Nested Schema for `drop_rules.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)

<a id="nestedatt--drop_rules--rewrite"></a>
### Nested Schema for `drop_rules.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)

<a id="nestedatt--drop_rules--vlan_tag"></a>
### Nested Schema for `drop_rules.vlan_tag`

Required:

* `vlan_action` (String)

Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)

<a id="nestedatt--pass_rules"></a>
### Nested Schema for `pass_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)

Optional:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--pass_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--pass_rules--rewrite))
* `vlan_tag` (Attributes) - This field is only  applicable for pass rules (see [below for nested schema](#nestedatt--pass_rules--vlan_tag))

<a id="nestedatt--pass_rules--ip_rewrite"></a>
### Nested Schema for `pass_rules.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)

<a id="nestedatt--pass_rules--rewrite"></a>
### Nested Schema for `pass_rules.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)

<a id="nestedatt--pass_rules--vlan_tag"></a>
### Nested Schema for `pass_rules.vlan_tag`

Required:

* `vlan_action` (String)

Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)

