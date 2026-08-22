---
page_title: "gigavuecore_get_period_volume_detailed_by_date Data Source - gigavuecore"
subcategory: ""
description: |-
  Gives a detailed period volume by date
---

# gigavuecore_get_period_volume_detailed_by_date Data Source

Gives a detailed period volume by date

## Example Usage

```terraform
data "gigavuecore_get_period_volume_detailed_by_date" "example" {
  date = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `date` (Number, required) - Date of detailed period volume to retrieve in format YYYYMMDD

### Attributes

In addition to all arguments above, the following attributes are exported:

* `app_volume` (List(Object({app, day_volumes, total_allowance, total_overage, total_usage})), computed)
* `best_unit` (String, computed) - Best unit to represent the volume (TeraMegaBytes, GigaBytes,..)
* `bundle_volume` (List(Object({bundle, day_volumes, days_overage, max_daily_allowance, max_daily_usage, total_allowance, total_overage, total_usage})), computed)
* `max_daily_allowance` (Number, computed) - Max daily allowance for the period in bytes
* `max_daily_usage` (Number, computed) - Max daily usage for the period in bytes
* `period_volume` (Object({apps, bundles, date_codes, days_completed, days_over95_p, days_overage, period, pos_ids}), computed)
  * `apps` (List(String), computed)
  * `bundles` (List(String), computed)
  * `date_codes` (List(Number), computed)
  * `days_completed` (Number, computed) - Number of days completed in the current period
  * `days_over95_p` (Number, computed) - Days where 95th percentile usage exceeded allowance, during the period
  * `days_overage` (Number, computed) - Number of days with usage exceeding allowance during the current period
  * `period` (String, computed) - Period from start to end dates formatted
  * `pos_ids` (List(String), computed)

