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

* `items` (Attributes List, computed) - list of the image files (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `build_id` (String) - Image build id
* `date` (String) - Image build date
* `file_name` (String) - Image filename
* `model` (String) - Image model
* `size` (String) - Image build size
* `version` (String) - Image version

