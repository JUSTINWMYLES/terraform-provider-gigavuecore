---
page_title: "gigavuecore_group Resource - gigavuecore"
subcategory: ""
description: |-
  Load a particular group
---

# gigavuecore_group Resource

Load a particular group

## Example Usage

```terraform
resource "gigavuecore_group" "example" {
  description = null
  name = null
  roles = []
  tags = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `description` (String, optional) - description
* `name` (String, required) - name
* `roles` (List(String), optional)
* `tags` (List(Object({multi_valued, override_user, tag_key, tag_values})), optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `description` (String, computed) - description
* `roles` (List(String), computed)
* `tags` (List(Object({multi_valued, override_user, tag_key, tag_values})), computed)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_group.example {name}
```
