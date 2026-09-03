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
  num = 0
}
```

## Schema

### Arguments

The following arguments are supported:

* `num` (Number, required) - Number of periods

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `app_volume` (Attributes List) (see [below for nested schema](#nestedatt--items--app_volume))
* `best_unit` (String) - Best unit to represent the volume (TeraBytes, GigaBytes,..)
* `bundle_volume` (Attributes List) (see [below for nested schema](#nestedatt--items--bundle_volume))
* `period_volume` (Attributes) (see [below for nested schema](#nestedatt--items--period_volume))

<a id="nestedatt--items--app_volume"></a>
### Nested Schema for `items.app_volume`

Read-Only:

* `app` (String) - Name of the app of this app volume
* `total_allowance` (Number) - Total allowance of this app volume in bytes
* `total_overage` (Number) - Total overage of this app volume in bytes
* `total_usage` (Number) - Total usage of this app volume in bytes

<a id="nestedatt--items--bundle_volume"></a>
### Nested Schema for `items.bundle_volume`

Read-Only:

* `bundle` (String) - Name of the license bundle
* `days_overage` (Number) - Number of days in the period where licensed volume allowance was exceeded
* `max_daily_allowance` (Number) - Max daily allowance of this bundle in bytes
* `max_daily_usage` (Number) - Maximum daily usage of this bundle in bytes
* `total_allowance` (Number) - Total bundle allowance for the period in bytes
* `total_overage` (Number) - Total bundle overage for the period in bytes
* `total_usage` (Number) - Total bundle usage for the period in bytes

<a id="nestedatt--items--period_volume"></a>
### Nested Schema for `items.period_volume`

Read-Only:

* `apps` (List of String)
* `bundles` (List of String)
* `date_codes` (List of Number)
* `days_completed` (Number) - Number of days completed in the current period
* `days_over95_p` (Number) - Days where 95th percentile usage exceeded allowance, during the period
* `days_overage` (Number) - Number of days with usage exceeding allowance during the current period
* `period` (String) - Period from start to end dates formatted
* `pos_ids` (List of String)

