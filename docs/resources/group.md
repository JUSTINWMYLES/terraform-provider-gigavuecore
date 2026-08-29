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
  description = "example"
  name        = "example"
  roles       = [ "example" ]
  tags = [{
    multi_valued  = true
    override_user = true
    tag_key       = "example"
    tag_values    = [ "example" ]
  }]
}
```

## Schema

### Arguments

The following arguments are supported:

* `description` (String, optional) - description
* `name` (String, required) - name
* `roles` (List of String, optional)
* `tags` (Attributes List, optional) (see [below for nested schema](#nestedatt--tags))

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `description` (String, computed) - description
* `roles` (List of String, computed)
* `tags` (Attributes List, computed) (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

* `tag_key` (String)
* `tag_values` (List of String)
Optional:

* `multi_valued` (Boolean)
* `override_user` (Boolean)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_group.example {name}
```
