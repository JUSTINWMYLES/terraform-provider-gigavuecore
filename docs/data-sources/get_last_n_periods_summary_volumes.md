---
page_title: "gigavuecore_get_last_n_periods_summary_volumes Data Source - gigavuecore"
subcategory: ""
description: |-
  Gives the summary of total Allowance, totalUsage, totalOverage for all apps for each period in the set of periods determined by num
---

# gigavuecore_get_last_n_periods_summary_volumes Data Source

Gives the summary of total Allowance, totalUsage, totalOverage for all apps for each period in the set of periods determined by num

## Example Usage

```terraform
data "gigavuecore_get_last_n_periods_summary_volumes" "example" {
  num = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `num` (Number, required) - Number of periods

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({app_volume, best_unit, bundle_volume, period_volume})), computed)

