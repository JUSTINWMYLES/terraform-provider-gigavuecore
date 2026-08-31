---
page_title: "gigavuecore_get_apps_in_last_n_periods Data Source - gigavuecore"
subcategory: ""
description: |-
  Gives the set of apps for which usage has been recorded in at least one of the last 'n' periods. If 'n' is unspecified or not positive, all periods are considered
---

# gigavuecore_get_apps_in_last_n_periods Data Source

Gives the set of apps for which usage has been recorded in at least one of the last 'n' periods. If 'n' is unspecified or not positive, all periods are considered

## Example Usage

```terraform
data "gigavuecore_get_apps_in_last_n_periods" "example" {
  n = 0
}
```

## Schema

### Arguments

The following arguments are supported:

* `n` (Number, optional) - Number of periods

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List of String, computed)


