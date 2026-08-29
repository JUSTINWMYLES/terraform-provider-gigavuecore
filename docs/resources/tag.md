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
  description  = "example"
  hierarchical = true
  multi_valued = true
  tag_key      = "example"
  tag_type     = "example"
  tag_values   = [ "example" ]
}
```

## Schema

### Arguments

The following arguments are supported:

* `description` (String, optional)
* `hierarchical` (Boolean, optional)
* `multi_valued` (Boolean, optional)
* `tag_key` (String, required) - tag key
* `tag_type` (String, required)
* `tag_values` (List of String, required)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `description` (String, computed)
* `hierarchical` (Boolean, computed)
* `multi_valued` (Boolean, computed)


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_tag.example {tag_key}
```
