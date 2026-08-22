---
page_title: "gigavuecore_tag Resource - gigavuecore"
subcategory: ""
description: |-
  Load a particular tag
---

# gigavuecore_tag Resource

Load a particular tag

## Example Usage

```terraform
resource "gigavuecore_tag" "example" {
  description = null
  hierarchical = null
  multi_valued = null
  tag_key = null
  tag_type = null
  tag_values = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `description` (String, optional)
* `hierarchical` (Bool, optional)
* `multi_valued` (Bool, optional)
* `tag_key` (String, required) - tag key
* `tag_type` (String, required)
* `tag_values` (List(String), required)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `description` (String, computed)
* `hierarchical` (Bool, computed)
* `multi_valued` (Bool, computed)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_tag.example {tag_key}
```
