---
page_title: "gigavuecore_get_bundles_in_period Data Source - gigavuecore"
subcategory: ""
description: |-
  Returns a list of license bundle names active in a VBL report period, where the input date falls within the period
---

# gigavuecore_get_bundles_in_period Data Source

Returns a list of license bundle names active in a VBL report period, where the input date falls within the period

## Example Usage

```terraform
data "gigavuecore_get_bundles_in_period" "example" {
  date = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `date` (Number, required) - Date falling within a VBL report period, format YYYYMMDD

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(String), computed)

