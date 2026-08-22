---
page_title: "gigavuecore_alertable_floating_license_expiries Data Source - gigavuecore"
subcategory: ""
description: |-
  Returns list of expiry info per category for pertinent floating licenses, count and a list of details of the expiry, when expiring (expired), SKU and activation ID
---

# gigavuecore_alertable_floating_license_expiries Data Source

Returns list of expiry info per category for pertinent floating licenses, count and a list of details of the expiry, when expiring (expired), SKU and activation ID

## Example Usage

```terraform
data "gigavuecore_alertable_floating_license_expiries" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({count_, expiry})), computed)

