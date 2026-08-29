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
  n = 0
}
```

## Schema

### Arguments

The following arguments are supported:

* `n` (Number, required) - Number of periods

### Attributes

In addition to all arguments above, the following attributes are exported:

* `lowest_unit` (String, computed)
* `summaries` (Attributes List, computed) (see [below for nested schema](#nestedatt--summaries))

<a id="nestedatt--summaries"></a>
### Nested Schema for `summaries`

Read-Only:

* `app_volume` (Attributes List) (see [below for nested schema](#nestedatt--summaries--app_volume))
* `best_unit` (String) - Best unit to represent the volume (TeraBytes, GigaBytes,..)
* `bundle_volume` (Attributes List) (see [below for nested schema](#nestedatt--summaries--bundle_volume))
* `period_volume` (Attributes) (see [below for nested schema](#nestedatt--summaries--period_volume))
<a id="nestedatt--summaries--app_volume"></a>
### Nested Schema for `summaries.app_volume`

Read-Only:

* `app` (String) - Name of the app of this app volume
* `total_allowance` (Number) - Total allowance of this app volume in bytes
* `total_overage` (Number) - Total overage of this app volume in bytes
* `total_usage` (Number) - Total usage of this app volume in bytes
<a id="nestedatt--summaries--bundle_volume"></a>
### Nested Schema for `summaries.bundle_volume`

Read-Only:

* `bundle` (String) - Name of the license bundle
* `days_overage` (Number) - Number of days in the period where licensed volume allowance was exceeded
* `max_daily_allowance` (Number) - Max daily allowance of this bundle in bytes
* `max_daily_usage` (Number) - Maximum daily usage of this bundle in bytes
* `total_allowance` (Number) - Total bundle allowance for the period in bytes
* `total_overage` (Number) - Total bundle overage for the period in bytes
* `total_usage` (Number) - Total bundle usage for the period in bytes
<a id="nestedatt--summaries--period_volume"></a>
### Nested Schema for `summaries.period_volume`

Read-Only:

* `apps` (List of String)
* `bundles` (List of String)
* `date_codes` (List of Number)
* `days_completed` (Number) - Number of days completed in the current period
* `days_over95_p` (Number) - Days where 95th percentile usage exceeded allowance, during the period
* `days_overage` (Number) - Number of days with usage exceeding allowance during the current period
* `period` (String) - Period from start to end dates formatted
* `pos_ids` (List of String)

