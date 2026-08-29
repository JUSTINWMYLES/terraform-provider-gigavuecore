---
page_title: "gigavuecore_get_period_volumes Data Source - gigavuecore"
subcategory: ""
description: |-
  Gives a list of all period volumes
---

# gigavuecore_get_period_volumes Data Source

Gives a list of all period volumes

## Example Usage

```terraform
data "gigavuecore_get_period_volumes" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `apps` (List of String)
* `bundles` (List of String)
* `date_codes` (List of Number)
* `days_completed` (Number) - Number of days completed in the current period
* `days_over95_p` (Number) - Days where 95th percentile usage exceeded allowance, during the period
* `days_overage` (Number) - Number of days with usage exceeding allowance during the current period
* `period` (String) - Period from start to end dates formatted
* `pos_ids` (List of String)

