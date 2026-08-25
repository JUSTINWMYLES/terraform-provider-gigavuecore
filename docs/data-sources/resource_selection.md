---
page_title: "gigavuecore_resource_selection Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all resource details of given resource type that don't have any alert policies associated with.
---

# gigavuecore_resource_selection Data Source

Get all resource details of given resource type that don't have any alert policies associated with.

## Example Usage

```terraform
data "gigavuecore_resource_selection" "example" {
  parent_resource_id = null
  resource_type      = null
  skip_used          = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `parent_resource_id` (String, optional) - Parent resource ID. If the rsource type is port, then providing the clusterId as the parent, fetches the ports of given clusterId.
* `resource_type` (String, required) - Type of the resource
* `skip_used` (Boolean, optional) - Flag to control whether all resources to be returned or only resources that aren't paty of any alert policies. The default value is true.

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `resources` (Attributes List, computed) - List of resources of given resource type (see [below for nested schema](#nestedatt--resources))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--resources"></a>
### Nested Schema for `resources`

Read-Only:

* `parent_resource_id` (String) - Parent resourceId
* `resource_ids` (List of String) - Collection of resource entities

