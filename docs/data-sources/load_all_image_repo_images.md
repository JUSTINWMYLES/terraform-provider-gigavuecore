---
page_title: "gigavuecore_load_all_image_repo_images Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Images
---

# gigavuecore_load_all_image_repo_images Data Source

Load all Images

## Example Usage

```terraform
data "gigavuecore_load_all_image_repo_images" "example" {
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

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `image_repo_images_query_response` (Attributes List, computed) - list of the image files (see [below for nested schema](#nestedatt--image_repo_images_query_response))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--image_repo_images_query_response"></a>
### Nested Schema for `image_repo_images_query_response`

Read-Only:

* `build_id` (String) - Image build id
* `date` (String) - Image build date
* `file_name` (String) - Image filename
* `model` (String) - Image model
* `size` (String) - Image build size
* `version` (String) - Image version

