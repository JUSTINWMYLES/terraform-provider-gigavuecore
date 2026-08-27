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

* `app_volume` (Attributes List, computed) (see [below for nested schema](#nestedatt--app_volume))
* `best_unit` (String, computed) - Best unit to represent the volume (TeraMegaBytes, GigaBytes,..)
* `bundle_volume` (Attributes List, computed) (see [below for nested schema](#nestedatt--bundle_volume))
* `max_daily_allowance` (Number, computed) - Max daily allowance for the period in bytes
* `max_daily_usage` (Number, computed) - Max daily usage for the period in bytes
* `period_volume` (Attributes, computed) (see [below for nested schema](#nestedatt--period_volume))

<a id="nestedatt--app_volume"></a>
### Nested Schema for `app_volume`

Read-Only:

* `app` (String) - Name of the app of this app volume
* `day_volumes` (Attributes List) (see [below for nested schema](#nestedatt--app_volume--day_volumes))
* `total_allowance` (Number) - Total allowance of this app volume in bytes
* `total_overage` (Number) - Total overage of this app volume in bytes
* `total_usage` (Number) - Total usage of this app volume in bytes
<a id="nestedatt--app_volume--day_volumes"></a>
### Nested Schema for `app_volume.day_volumes`

Read-Only:

* `allowance` (Number) - Allowance of this day volume in bytes
* `overage` (Number) - Overage on this day in bytes, usually difference between allowance and usage if positive, else 0
* `usage` (Number) - Usage of this day volume in bytes
* `usage95_p` (Number) - 95th percentile usage on this day calculated with a lookback period of 90 days
<a id="nestedatt--bundle_volume"></a>
### Nested Schema for `bundle_volume`

Read-Only:

* `bundle` (String) - Name of the bundle of this bundle volume
* `day_volumes` (Attributes List) (see [below for nested schema](#nestedatt--bundle_volume--day_volumes))
* `days_overage` (Number) - Number of days in the period where licensed volume limit was exceeded
* `max_daily_allowance` (Number) - Max daily allowance of this bundle volume in bytes
* `max_daily_usage` (Number) - Maximum daily usage of this bundle volume in bytes
* `total_allowance` (Number) - Total bundle allowance in bytes
* `total_overage` (Number) - Total bundle overage in bytes
* `total_usage` (Number) - Total bundle usage in bytes
<a id="nestedatt--bundle_volume--day_volumes"></a>
### Nested Schema for `bundle_volume.day_volumes`

Read-Only:

* `allowance` (Number) - Allowance of this day volume in bytes
* `overage` (Number) - Overage on this day in bytes, usually difference between allowance and usage if positive, else 0
* `usage` (Number) - Usage of this day volume in bytes
* `usage95_p` (Number) - 95th percentile usage on this day calculated with a lookback period of 90 days
<a id="nestedatt--period_volume"></a>
### Nested Schema for `period_volume`

Read-Only:

* `apps` (List of String)
* `bundles` (List of String)
* `date_codes` (List of Number)
* `days_completed` (Number) - Number of days completed in the current period
* `days_over95_p` (Number) - Days where 95th percentile usage exceeded allowance, during the period
* `days_overage` (Number) - Number of days with usage exceeding allowance during the current period
* `period` (String) - Period from start to end dates formatted
* `pos_ids` (List of String)

