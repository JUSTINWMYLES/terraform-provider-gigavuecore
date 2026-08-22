---
page_title: "gigavuecore_source_rule Resource - gigavuecore"
subcategory: ""
description: |-
  Get source rules for a policy and sourceRulesAlias
---

# gigavuecore_source_rule Resource

Get source rules for a policy and sourceRulesAlias

## Example Usage

```terraform
resource "gigavuecore_source_rule" "example" {
  drop_rules = []
  pass_rules = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `drop_rules` (Set(Object({bidi, comment, ip_rewrite, matches, rewrite, rule_id, vlan_tag})), optional)
* `pass_rules` (Set(Object({bidi, comment, ip_rewrite, matches, rewrite, rule_id, vlan_tag})), optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `drop_rules` (Set(Object({bidi, comment, ip_rewrite, matches, rewrite, rule_id, vlan_tag})), computed)
* `id` (String, computed)
* `pass_rules` (Set(Object({bidi, comment, ip_rewrite, matches, rewrite, rule_id, vlan_tag})), computed)

