---
page_title: "gigavuecore_list_env Data Source - gigavuecore"
subcategory: ""
description: |-
  List all unified resource environments
---

# gigavuecore_list_env Data Source

List all unified resource environments

## Example Usage

```terraform
data "gigavuecore_list_env" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `env` (Attributes Set) (see [below for nested schema](#nestedatt--items--env))

<a id="nestedatt--items--env"></a>
### Nested Schema for `items.env`

Read-Only:

* `name` (String)
* `type` (String) - Unified Environment platform types

