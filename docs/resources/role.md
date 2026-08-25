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
  description = null
  name        = null
  scope       = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `description` (String, optional) - description
* `name` (String, required) - name
* `scope` (Attributes List, required) (see [below for nested schema](#nestedatt--scope))

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `description` (String, computed) - description

<a id="nestedatt--scope"></a>
### Nested Schema for `scope`

Required:

* `actions` (List of String)
* `type` (String)
Optional:

* `hierarchy` (Boolean)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_role.example {name}
```
