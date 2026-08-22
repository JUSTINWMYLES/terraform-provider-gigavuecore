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
  resource_type = null
  skip_used = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `parent_resource_id` (String, optional) - Parent resource ID. If the rsource type is port, then providing the clusterId as the parent, fetches the ports of given clusterId.
* `resource_type` (String, required) - Type of the resource
* `skip_used` (Bool, optional) - Flag to control whether all resources to be returned or only resources that aren't paty of any alert policies. The default value is true.

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `resources` (List(Object({parent_resource_id, resource_ids})), computed) - List of resources of given resource type

