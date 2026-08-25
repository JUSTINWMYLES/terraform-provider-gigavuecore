---
page_title: "gigavuecore_get_latest_period_volume_detailed_and_periods Data Source - gigavuecore"
subcategory: ""
description: |-
  Gives a list of all period volumes
---

# gigavuecore_get_latest_period_volume_detailed_and_periods Data Source

Gives a list of all period volumes

## Example Usage

```terraform
data "gigavuecore_get_latest_period_volume_detailed_and_periods" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `latest_volume` (Attributes, computed) (see [below for nested schema](#nestedatt--latest_volume))
* `periods` (List of String, computed)

<a id="nestedatt--latest_volume"></a>
### Nested Schema for `latest_volume`

Read-Only:

* `app_volume` (Attributes List) (see [below for nested schema](#nestedatt--latest_volume--app_volume))
* `best_unit` (String) - Best unit to represent the volume (TeraMegaBytes, GigaBytes,..)
* `bundle_volume` (Attributes List) (see [below for nested schema](#nestedatt--latest_volume--bundle_volume))
* `max_daily_allowance` (Number) - Max daily allowance for the period in bytes
* `max_daily_usage` (Number) - Max daily usage for the period in bytes
* `period_volume` (Attributes) (see [below for nested schema](#nestedatt--latest_volume--period_volume))
<a id="nestedatt--latest_volume--app_volume"></a>
### Nested Schema for `latest_volume.app_volume`

Read-Only:

* `app` (String) - Name of the app of this app volume
* `day_volumes` (Attributes List) (see [below for nested schema](#nestedatt--latest_volume--app_volume--day_volumes))
* `total_allowance` (Number) - Total allowance of this app volume in bytes
* `total_overage` (Number) - Total overage of this app volume in bytes
* `total_usage` (Number) - Total usage of this app volume in bytes
<a id="nestedatt--latest_volume--app_volume--day_volumes"></a>
### Nested Schema for `latest_volume.app_volume.day_volumes`

Read-Only:

* `allowance` (Number) - Allowance of this day volume in bytes
* `overage` (Number) - Overage on this day in bytes, usually difference between allowance and usage if positive, else 0
* `usage` (Number) - Usage of this day volume in bytes
* `usage95_p` (Number) - 95th percentile usage on this day calculated with a lookback period of 90 days
<a id="nestedatt--latest_volume--bundle_volume"></a>
### Nested Schema for `latest_volume.bundle_volume`

Read-Only:

* `bundle` (String) - Name of the bundle of this bundle volume
* `day_volumes` (Attributes List) (see [below for nested schema](#nestedatt--latest_volume--bundle_volume--day_volumes))
* `days_overage` (Number) - Number of days in the period where licensed volume limit was exceeded
* `max_daily_allowance` (Number) - Max daily allowance of this bundle volume in bytes
* `max_daily_usage` (Number) - Maximum daily usage of this bundle volume in bytes
* `total_allowance` (Number) - Total bundle allowance in bytes
* `total_overage` (Number) - Total bundle overage in bytes
* `total_usage` (Number) - Total bundle usage in bytes
<a id="nestedatt--latest_volume--bundle_volume--day_volumes"></a>
### Nested Schema for `latest_volume.bundle_volume.day_volumes`

Read-Only:

* `allowance` (Number) - Allowance of this day volume in bytes
* `overage` (Number) - Overage on this day in bytes, usually difference between allowance and usage if positive, else 0
* `usage` (Number) - Usage of this day volume in bytes
* `usage95_p` (Number) - 95th percentile usage on this day calculated with a lookback period of 90 days
<a id="nestedatt--latest_volume--period_volume"></a>
### Nested Schema for `latest_volume.period_volume`

Read-Only:

* `apps` (List of String)
* `bundles` (List of String)
* `date_codes` (List of Number)
* `days_completed` (Number) - Number of days completed in the current period
* `days_over95_p` (Number) - Days where 95th percentile usage exceeded allowance, during the period
* `days_overage` (Number) - Number of days with usage exceeding allowance during the current period
* `period` (String) - Period from start to end dates formatted
* `pos_ids` (List of String)

