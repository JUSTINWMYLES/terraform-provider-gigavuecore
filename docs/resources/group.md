---
page_title: "gigavuecore_group Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new user group
---

# gigavuecore_group Resource

Create a new user group

## Example Usage

```terraform
resource "gigavuecore_group" "example" {
  description = "example-value"
  name        = "example"
  roles       = ["example"]
  tags = [{
    multi_valued  = true
    override_user = true
    tag_key       = "example"
    tag_values    = ["example"]
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

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

* `tag_key` (String)
* `tag_values` (List of String)

Optional:

* `multi_valued` (Boolean)
* `override_user` (Boolean)
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_group.example {name}
```
