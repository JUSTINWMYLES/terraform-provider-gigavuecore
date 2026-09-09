---
page_title: "gigavuecore_tag Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new tag
---

# gigavuecore_tag Resource

Create a new tag

## Example Usage

```terraform
resource "gigavuecore_tag" "example" {
  description  = "example-value"
  hierarchical = true
  multi_valued = true
  tag_key      = "example"
  tag_type     = "Rbac"
  tag_values   = ["example"]
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

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

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
terraform import gigavuecore_tag.example {tag_key}
```
