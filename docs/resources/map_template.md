---
page_title: "gigavuecore_map_template Resource - gigavuecore"
subcategory: ""
description: |-
  Find map template by alias
---

# gigavuecore_map_template Resource

Find map template by alias

## Example Usage

```terraform
resource "gigavuecore_map_template" "example" {
  alias      = null
  cluster_id = null
  comment    = null
  roles      = {}
  rules      = {}
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - unique map alias
* `cluster_id` (String, optional) - id of the defining cluster
* `comment` (String, optional)
* `roles` (Attributes, optional) - Map Roles Container. Private class (see [below for nested schema](#nestedatt--roles))
* `rules` (Attributes, required) - Map Rules Container. Private class (see [below for nested schema](#nestedatt--rules))

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `cluster_id` (String, computed) - id of the defining cluster
* `comment` (String, computed)
* `roles` (Attributes, computed) - Map Roles Container. Private class (see [below for nested schema](#nestedatt--roles))

<a id="nestedatt--roles"></a>
### Nested Schema for `roles`

Optional:

* `editors` (List of String)
* `owners` (List of String)
* `viewers` (List of String)
<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Optional:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--rules--pass_rules))
<a id="nestedatt--rules--drop_rules"></a>
### Nested Schema for `rules.drop_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)
Optional:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--rules--drop_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--rules--drop_rules--rewrite))
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--rules--drop_rules--vlan_tag))
<a id="nestedatt--rules--drop_rules--ip_rewrite"></a>
### Nested Schema for `rules.drop_rules.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--rules--drop_rules--rewrite"></a>
### Nested Schema for `rules.drop_rules.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--rules--drop_rules--vlan_tag"></a>
### Nested Schema for `rules.drop_rules.vlan_tag`

Required:

* `vlan_action` (String)
Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)
<a id="nestedatt--rules--pass_rules"></a>
### Nested Schema for `rules.pass_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)
Optional:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--rules--pass_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--rules--pass_rules--rewrite))
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--rules--pass_rules--vlan_tag))
<a id="nestedatt--rules--pass_rules--ip_rewrite"></a>
### Nested Schema for `rules.pass_rules.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--rules--pass_rules--rewrite"></a>
### Nested Schema for `rules.pass_rules.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--rules--pass_rules--vlan_tag"></a>
### Nested Schema for `rules.pass_rules.vlan_tag`

Required:

* `vlan_action` (String)
Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_map_template.example {alias}
```
