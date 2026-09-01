---
page_title: "gigavuecore_source_rule Resource - gigavuecore"
subcategory: ""
description: |-
  Replace source rules for a policy and sourceRulesAlias
---

# gigavuecore_source_rule Resource

Replace source rules for a policy and sourceRulesAlias

~> **Note:** This resource is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
resource "gigavuecore_source_rule" "example" {
  drop_rules = [{
    bidi    = true
    comment = "example"
    ip_rewrite = {
      dst_ip = "example"
      src_ip = "example"
    }
    matches = ["example"]
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
    matches = ["example"]
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
  rule_ids = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `drop_rules` (Attributes Set, optional) (see [below for nested schema](#nestedatt--drop_rules))
* `pass_rules` (Attributes Set, optional) (see [below for nested schema](#nestedatt--pass_rules))
* `rule_ids` (String, required) - Rule IDs to match

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `id` (String, computed)

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

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
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

