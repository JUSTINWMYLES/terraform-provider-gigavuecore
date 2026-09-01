---
page_title: "gigavuecore_get_image_server_repo_by_version Data Source - gigavuecore"
subcategory: ""
description: |-
  Find Images By Version
---

# gigavuecore_get_image_server_repo_by_version Data Source

Find Images By Version

## Example Usage

```terraform
data "gigavuecore_get_image_server_repo_by_version" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) - System Image File by version (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `images` (Attributes List) (see [below for nested schema](#nestedatt--items--images))
* `version` (String) - Image version

<a id="nestedatt--items--images"></a>
### Nested Schema for `items.images`

Read-Only:

* `build_id` (String) - Image build id
* `date` (String) - Image build date
* `file_name` (String) - Image filename
* `model` (String) - Image model
* `size` (String) - Image build size
* `version` (String) - Image version

