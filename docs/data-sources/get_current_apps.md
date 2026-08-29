---
page_title: "gigavuecore_get_current_apps Data Source - gigavuecore"
subcategory: ""
description: |-
  Gives the currently licensed application set
---

# gigavuecore_get_current_apps Data Source

Gives the currently licensed application set

## Example Usage

```terraform
data "gigavuecore_get_current_apps" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List of String, computed)


