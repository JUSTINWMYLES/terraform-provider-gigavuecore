---
page_title: "gigavuecore_get_current_app_tiers Data Source - gigavuecore"
subcategory: ""
description: |-
  Gives the set of currently licensed app tiers based on currently licensed application set
---

# gigavuecore_get_current_app_tiers Data Source

Gives the set of currently licensed app tiers based on currently licensed application set

## Example Usage

```terraform
data "gigavuecore_get_current_app_tiers" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Number), computed)

