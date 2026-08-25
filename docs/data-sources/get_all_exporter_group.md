---
page_title: "gigavuecore_get_all_exporter_group Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all Apps Exporter Group
---

# gigavuecore_get_all_exporter_group Data Source

Get all Apps Exporter Group

## Example Usage

```terraform
data "gigavuecore_get_all_exporter_group" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `apps_exporter_groups` (Attributes List, computed) (see [below for nested schema](#nestedatt--apps_exporter_groups))
* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))

<a id="nestedatt--apps_exporter_groups"></a>
### Nested Schema for `apps_exporter_groups`

Read-Only:

* `alias` (String)
* `description` (String)
* `exporters` (List of String)
<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type

