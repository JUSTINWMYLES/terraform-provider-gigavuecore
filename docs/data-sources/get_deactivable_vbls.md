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

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `aid` (String) - activation id as encoded in the license
* `end_date` (Number) - end date of the license (format YYYYMMDD)
* `license_status` (String) - status of the license
* `sku` (String) - the license SKU (Stock Keeping Unit) or product code that uniquely identifies the license type in Gigamon catalog
* `start_date` (Number) - start date of the license (format YYYYMMDD)
* `total_volume` (String) - ingress traffic volume licensed per day (in GigaBytes)

