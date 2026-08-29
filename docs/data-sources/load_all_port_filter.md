---
page_title: "gigavuecore_load_all_port_filter Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all port filters
---

# gigavuecore_load_all_port_filter Data Source

Load all port filters

## Example Usage

```terraform
data "gigavuecore_load_all_port_filter" "example" {
  cluster_id = "example"
  page       = "example"
  sort       = "example"
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

* `port` (String)
* `rules` (Attributes) - Port Filter Rules Container. Private class (see [below for nested schema](#nestedatt--items--rules))
<a id="nestedatt--items--rules"></a>
### Nested Schema for `items.rules`

Read-Only:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--items--rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--items--rules--pass_rules))
<a id="nestedatt--items--rules--drop_rules"></a>
### Nested Schema for `items.rules.drop_rules`

Read-Only:

* `comment` (String)
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)
<a id="nestedatt--items--rules--pass_rules"></a>
### Nested Schema for `items.rules.pass_rules`

Read-Only:

* `comment` (String)
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)

