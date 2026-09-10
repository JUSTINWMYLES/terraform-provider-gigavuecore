---
page_title: "gigavuecore_role Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new user role
---

# gigavuecore_role Resource

Create a new user role

## Example Usage

```terraform
resource "gigavuecore_role" "example" {
  description = "example-value"
  name        = "example"
  scope = [{
    actions   = ["all"]
    hierarchy = true
    type      = "example"
  }]
}
```

## Schema

### Arguments

The following arguments are supported:

* `description` (String, optional) - description
* `name` (String, required) - name
* `scope` (Attributes List, required) (see [below for nested schema](#nestedatt--scope))

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--scope"></a>
### Nested Schema for `scope`

Required:

* `actions` (List of String)
* `type` (String)

Optional:

* `hierarchy` (Boolean)
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
terraform import gigavuecore_role.example {name}
```
