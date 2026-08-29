---
page_title: "gigavuecore_get_hierarchial_data Action - gigavuecore"
subcategory: ""
description: |-
  Response for bubble diagram representation
---

# gigavuecore_get_hierarchial_data Action

Response for bubble diagram representation

## Example Usage

```terraform
action "gigavuecore_get_hierarchial_data" "example" {
  config {
    tags_filter = {
      parent_tags = [{
        tag_key   = "example"
        tag_value = "example"
      }]
      tag_key = "example"
    }
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `tags_filter` (Attributes, required) (see [below for nested schema](#nestedatt--tags_filter))

<a id="nestedatt--tags_filter"></a>
### Nested Schema for `tags_filter`

Required:

* `tag_key` (String)
Optional:

* `parent_tags` (Attributes List) (see [below for nested schema](#nestedatt--tags_filter--parent_tags))
<a id="nestedatt--tags_filter--parent_tags"></a>
### Nested Schema for `tags_filter.parent_tags`

Optional:

* `tag_key` (String)
* `tag_value` (String)

