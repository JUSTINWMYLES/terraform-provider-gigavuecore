---
page_title: "gigavuecore_get_deactivable_vbls Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all VBL activations that can be deactivated (and are not already deactivated)
---

# gigavuecore_get_deactivable_vbls Data Source

Get all VBL activations that can be deactivated (and are not already deactivated)

## Example Usage

```terraform
data "gigavuecore_get_deactivable_vbls" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({aid, end_date, license_status, sku, start_date, total_volume})), computed)

