---
page_title: "gigavuecore_load_all_metadata_application_templates Data Source - gigavuecore"
subcategory: ""
description: |-
  Load All Metadata Application Templates
---

# gigavuecore_load_all_metadata_application_templates Data Source

Load All Metadata Application Templates

## Example Usage

```terraform
data "gigavuecore_load_all_metadata_application_templates" "example" {
  page = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `metadata_application_templates` (List(Object({attributes, description, family, name})), computed)

