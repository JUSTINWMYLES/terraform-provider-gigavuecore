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

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({apps, bundles, date_codes, days_completed, days_over95_p, days_overage, period, pos_ids})), computed)

