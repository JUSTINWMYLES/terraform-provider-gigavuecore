---
page_title: "gigavuecore_load_all_filter_resources Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Filter Resources
---

# gigavuecore_load_all_filter_resources Data Source

Load all Filter Resources

## Example Usage

```terraform
data "gigavuecore_load_all_filter_resources" "example" {
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

* `cluster_id` (String) - id of the defining cluster
* `filter_template` (String) - alias of filter template
* `lookup_resource_limit` (Number)
* `lookup_resource_used` (Number)
* `map_rules_limit` (Number)
* `map_rules_used` (Number)
* `qualifiers` (List of String) - in use qualifiers
* `slot_id` (String) - Device card slot id
* `tool_port_filter_limit` (Number)
* `tool_port_filter_used` (Number)

