---
page_title: "gigavuecore_get_period_volume_by_date Data Source - gigavuecore"
subcategory: ""
description: |-
  Gives volume allowed and overuse for the period within which the specified date lies
---

# gigavuecore_get_period_volume_by_date Data Source

Gives volume allowed and overuse for the period within which the specified date lies

## Example Usage

```terraform
data "gigavuecore_get_period_volume_by_date" "example" {
  date = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `date` (Number, required) - Date in YYYYMMDD format, can be any day within the desired period

### Attributes

In addition to all arguments above, the following attributes are exported:

* `apps` (List(String), computed)
* `bundles` (List(String), computed)
* `date_codes` (List(Number), computed)
* `days_completed` (Number, computed) - Number of days completed in the current period
* `days_over95_p` (Number, computed) - Days where 95th percentile usage exceeded allowance, during the period
* `days_overage` (Number, computed) - Number of days with usage exceeding allowance during the current period
* `period` (String, computed) - Period from start to end dates formatted
* `pos_ids` (List(String), computed)

