---
page_title: "gigavuecore_get_license_skus Data Source - gigavuecore"
subcategory: ""
description: |-
  Load FM Licensing SKUs
---

# gigavuecore_get_license_skus Data Source

Load FM Licensing SKUs

## Example Usage

```terraform
data "gigavuecore_get_license_skus" "example" {
  device_model = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `device_model` (String, optional) - The device model to query

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `device_model` (String) - Gigamon physical device models
* `features` (String)
* `sku` (String)
* `target_type` (String)

