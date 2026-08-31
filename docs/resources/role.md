---
page_title: "gigavuecore_role Resource - gigavuecore"
subcategory: ""
description: |-
  Load a particular role
---

# gigavuecore_role Resource

Load a particular role

## Example Usage

```terraform
resource "gigavuecore_role" "example" {
  description = "examplee"
  name        = "example"
  scope = [{
    actions   = [ "all" ]
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

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_role.example {name}
```
