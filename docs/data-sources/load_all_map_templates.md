---
page_title: "gigavuecore_load_all_map_templates Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all map templates
---

# gigavuecore_load_all_map_templates Data Source

Load all map templates

## Example Usage

```terraform
data "gigavuecore_load_all_map_templates" "example" {
  cluster_id = null
  page       = null
  sort       = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - unique map alias
* `cluster_id` (String) - id of the defining cluster
* `comment` (String)
* `roles` (Attributes) - Map Roles Container. Private class (see [below for nested schema](#nestedatt--items--roles))
* `rules` (Attributes) - Map Rules Container. Private class (see [below for nested schema](#nestedatt--items--rules))
<a id="nestedatt--items--roles"></a>
### Nested Schema for `items.roles`

Read-Only:

* `editors` (List of String)
* `owners` (List of String)
* `viewers` (List of String)
<a id="nestedatt--items--rules"></a>
### Nested Schema for `items.rules`

Read-Only:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--items--rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--items--rules--pass_rules))
<a id="nestedatt--items--rules--drop_rules"></a>
### Nested Schema for `items.rules.drop_rules`

Read-Only:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--items--rules--drop_rules--ip_rewrite))
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--items--rules--drop_rules--rewrite))
* `rule_id` (Number)
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--items--rules--drop_rules--vlan_tag))
<a id="nestedatt--items--rules--drop_rules--ip_rewrite"></a>
### Nested Schema for `items.rules.drop_rules.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--items--rules--drop_rules--rewrite"></a>
### Nested Schema for `items.rules.drop_rules.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--items--rules--drop_rules--vlan_tag"></a>
### Nested Schema for `items.rules.drop_rules.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)
<a id="nestedatt--items--rules--pass_rules"></a>
### Nested Schema for `items.rules.pass_rules`

Read-Only:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--items--rules--pass_rules--ip_rewrite))
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--items--rules--pass_rules--rewrite))
* `rule_id` (Number)
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--items--rules--pass_rules--vlan_tag))
<a id="nestedatt--items--rules--pass_rules--ip_rewrite"></a>
### Nested Schema for `items.rules.pass_rules.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--items--rules--pass_rules--rewrite"></a>
### Nested Schema for `items.rules.pass_rules.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--items--rules--pass_rules--vlan_tag"></a>
### Nested Schema for `items.rules.pass_rules.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)

