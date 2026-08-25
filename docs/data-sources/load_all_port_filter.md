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

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `port_filters` (Attributes List, computed) (see [below for nested schema](#nestedatt--port_filters))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--port_filters"></a>
### Nested Schema for `port_filters`

Read-Only:

* `port` (String)
* `rules` (Attributes) - Port Filter Rules Container. Private class (see [below for nested schema](#nestedatt--port_filters--rules))
<a id="nestedatt--port_filters--rules"></a>
### Nested Schema for `port_filters.rules`

Read-Only:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--port_filters--rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--port_filters--rules--pass_rules))
<a id="nestedatt--port_filters--rules--drop_rules"></a>
### Nested Schema for `port_filters.rules.drop_rules`

Read-Only:

* `comment` (String)
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)
<a id="nestedatt--port_filters--rules--pass_rules"></a>
### Nested Schema for `port_filters.rules.pass_rules`

Read-Only:

* `comment` (String)
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)

