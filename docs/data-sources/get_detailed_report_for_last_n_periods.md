---
page_title: "gigavuecore_get_detailed_report_for_last_n_periods Data Source - gigavuecore"
subcategory: ""
description: |-
  Gives a list of detailed volume for 'n' periods
---

# gigavuecore_get_detailed_report_for_last_n_periods Data Source

Gives a list of detailed volume for 'n' periods

## Example Usage

```terraform
data "gigavuecore_get_detailed_report_for_last_n_periods" "example" {
  encode = null
  n = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `encode` (Bool, optional) - encode the response iff true
* `n` (List(String), required) - Number of periods starting from the latest period under collection

### Attributes

In addition to all arguments above, the following attributes are exported:

* `fm_sw_version` (String, computed)
* `period_volumes` (List(Object({app_volume, best_unit, bundle_volume, max_daily_allowance, max_daily_usage, period_volume})), computed)
* `time_of_report_creation` (String, computed)
* `virtual_mac` (String, computed)

