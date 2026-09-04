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

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `count_` (Number)
* `expiry` (List of String)

