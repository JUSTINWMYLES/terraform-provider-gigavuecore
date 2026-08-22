---
page_title: "gigavuecore_get_last_n_periods_summary_with_unit Data Source - gigavuecore"
subcategory: ""
description: |-
  Gives the summary of total allowance, total usage, total overage for each app and bundle for each period, for the last 'n' periods, along with lowest of the best units for each period summary; for bundles, the maximum daily allowance and usage and the number of days of overage are also provided
---

# gigavuecore_get_last_n_periods_summary_with_unit Data Source

Gives the summary of total allowance, total usage, total overage for each app and bundle for each period, for the last 'n' periods, along with lowest of the best units for each period summary; for bundles, the maximum daily allowance and usage and the number of days of overage are also provided

## Example Usage

```terraform
data "gigavuecore_get_last_n_periods_summary_with_unit" "example" {
  n = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `n` (Number, required) - Number of periods

### Attributes

In addition to all arguments above, the following attributes are exported:

* `lowest_unit` (String, computed)
* `summaries` (List(Object({app_volume, best_unit, bundle_volume, period_volume})), computed)

