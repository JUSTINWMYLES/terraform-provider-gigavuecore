---
page_title: "gigavuecore_display_periods Data Source - gigavuecore"
subcategory: ""
description: |-
  Gives the number of months per period, and mappings from number of periods to the corresponding displayed string
---

# gigavuecore_display_periods Data Source

Gives the number of months per period, and mappings from number of periods to the corresponding displayed string

## Example Usage

```terraform
data "gigavuecore_display_periods" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `months_per_period` (Number, computed)
* `n_to_display_string_map` (Map(String), computed)

